import { useCallback, useMemo, useState } from 'react';
import { useAppDispatch, useAppSelector } from '../../app/hooks';
import {
  processMultimodal,
  calibrateUser,
  verifyQuantification,
  startTextProcessing,
  addImage,
  addDirectMapping,
  addVisualMetaphor,
  resetProcessing,
  clearError,
} from '../../features/heuristics/multimodalSlice';

interface UseMultimodalReturn {
  // 状態
  multimodalData: any[];
  visualMetaphors: any[];
  userCalibrations: any[];
  currentProcessing: any;
  statistics: any;
  loading: boolean;
  error: string | null;
  
  // アクション
  processTextWithImage: (text: string, imageFile?: File, userId?: number, taskId?: number) => Promise<any>;
  calibrate: (userId: number, referenceObject: string, imageFile: File) => Promise<any>;
  verify: (dataId: string, feedback: string) => Promise<any>;
  findDirectMapping: (text: string) => any;
  matchVisualMetaphor: (text: string) => any;
  generateConfirmationImages: (text: string) => string[];
  reset: () => void;
}

export const useMultimodal = (): UseMultimodalReturn => {
  const dispatch = useAppDispatch();
  const multimodalState = useAppSelector((state) => state.multimodal);
  const [processedImages, setProcessedImages] = useState<Map<string, string>>(new Map());
  
  // テキストと画像を処理
  const processTextWithImage = useCallback(async (
    text: string,
    imageFile?: File,
    userId: number = 1,
    taskId: number = 1
  ) => {
    // テキスト処理開始
    dispatch(startTextProcessing(text));
    
    // 画像がある場合はプレビュー生成
    if (imageFile) {
      const imageUrl = URL.createObjectURL(imageFile);
      dispatch(addImage(imageUrl));
      setProcessedImages(prev => new Map(prev).set(text, imageUrl));
    }
    
    // マルチモーダル処理実行
    const result = await dispatch(processMultimodal({
      text,
      imageFile,
      userId,
      taskId,
    }));
    
    return result.payload;
  }, [dispatch]);
  
  // ユーザーキャリブレーション
  const calibrate = useCallback(async (
    userId: number,
    referenceObject: string,
    imageFile: File
  ) => {
    const result = await dispatch(calibrateUser({
      userId,
      referenceObject,
      imageFile,
    }));
    
    return result.payload;
  }, [dispatch]);
  
  // 定量化結果の検証
  const verify = useCallback(async (
    dataId: string,
    feedback: string
  ) => {
    const result = await dispatch(verifyQuantification({
      dataId,
      feedback,
    }));
    
    return result.payload;
  }, [dispatch]);
  
  // 直接マッピングを検索
  const findDirectMapping = useCallback((text: string) => {
    // 完全一致を検索
    if (multimodalState.directMappings[text]) {
      return multimodalState.directMappings[text];
    }
    
    // 部分一致を検索
    for (const [pattern, value] of Object.entries(multimodalState.directMappings)) {
      if (text.includes(pattern)) {
        return value;
      }
    }
    
    return null;
  }, [multimodalState.directMappings]);
  
  // ビジュアルメタファーをマッチング
  const matchVisualMetaphor = useCallback((text: string) => {
    const lowerText = text.toLowerCase();
    
    return multimodalState.visualMetaphors.find(metaphor => {
      const metaphorLower = metaphor.metaphor.toLowerCase();
      return lowerText.includes(metaphorLower) || 
             lowerText.includes(metaphor.referenceObject.toLowerCase());
    });
  }, [multimodalState.visualMetaphors]);
  
  // 確認用画像の生成
  const generateConfirmationImages = useCallback((text: string): string[] => {
    const images: string[] = [];
    
    // 直接マッピングの画像
    const directMapping = findDirectMapping(text);
    if (directMapping) {
      // プレースホルダー画像URL（実際の実装では画像生成APIを使用）
      images.push(`/api/images/generate?value=${directMapping.value}&unit=${directMapping.unit}`);
    }
    
    // メタファーの画像
    const metaphor = matchVisualMetaphor(text);
    if (metaphor && metaphor.imageUrl) {
      images.push(metaphor.imageUrl);
    }
    
    // 過去の類似処理の画像
    const similarData = multimodalState.multimodalData.find(
      data => data.linguistic.text.includes(text.substring(0, 5))
    );
    if (similarData && similarData.visual?.imageUrl) {
      images.push(similarData.visual.imageUrl);
    }
    
    return images;
  }, [findDirectMapping, matchVisualMetaphor, multimodalState.multimodalData]);
  
  // リセット
  const reset = useCallback(() => {
    dispatch(resetProcessing());
    dispatch(clearError());
    processedImages.forEach(url => URL.revokeObjectURL(url));
    setProcessedImages(new Map());
  }, [dispatch, processedImages]);
  
  // 統計情報の計算
  const enhancedStatistics = useMemo(() => {
    const stats = multimodalState.statistics;
    const hasCalibration = multimodalState.userCalibrations.length > 0;
    
    return {
      ...stats,
      hasUserCalibration: hasCalibration,
      calibrationQuality: hasCalibration 
        ? multimodalState.userCalibrations.reduce((sum, c) => sum + c.confidence, 0) / multimodalState.userCalibrations.length
        : 0,
      metaphorCoverage: multimodalState.visualMetaphors.length,
      directMappingCount: Object.keys(multimodalState.directMappings).length,
    };
  }, [multimodalState]);
  
  return {
    // 状態
    multimodalData: multimodalState.multimodalData,
    visualMetaphors: multimodalState.visualMetaphors,
    userCalibrations: multimodalState.userCalibrations,
    currentProcessing: multimodalState.currentProcessing,
    statistics: enhancedStatistics,
    loading: multimodalState.loading,
    error: multimodalState.error,
    
    // アクション
    processTextWithImage,
    calibrate,
    verify,
    findDirectMapping,
    matchVisualMetaphor,
    generateConfirmationImages,
    reset,
  };
};

// ヘルパーフック: 画像アノテーション
export const useImageAnnotation = () => {
  const [annotations, setAnnotations] = useState<Map<string, any[]>>(new Map());
  
  const annotateImage = useCallback((
    imageId: string,
    region: { x: number; y: number; width: number; height: number },
    quantification: { value: number; unit: string },
    description: string
  ) => {
    setAnnotations(prev => {
      const newMap = new Map(prev);
      const existing = newMap.get(imageId) || [];
      existing.push({
        id: Date.now().toString(),
        region,
        quantification,
        description,
        timestamp: new Date().toISOString(),
      });
      newMap.set(imageId, existing);
      return newMap;
    });
  }, []);
  
  const getAnnotations = useCallback((imageId: string) => {
    return annotations.get(imageId) || [];
  }, [annotations]);
  
  const clearAnnotations = useCallback((imageId?: string) => {
    if (imageId) {
      setAnnotations(prev => {
        const newMap = new Map(prev);
        newMap.delete(imageId);
        return newMap;
      });
    } else {
      setAnnotations(new Map());
    }
  }, []);
  
  return {
    annotations,
    annotateImage,
    getAnnotations,
    clearAnnotations,
  };
};

// ヘルパーフック: 視覚的比較
export const useVisualComparison = () => {
  const [comparisonPairs, setComparisonPairs] = useState<Array<{
    id: string;
    imageA: string;
    imageB: string;
    similarity: number;
  }>>([]);
  
  const compareImages = useCallback(async (
    imageA: string,
    imageB: string
  ): Promise<number> => {
    // 実際の実装では画像比較APIを呼び出す
    // ここではダミーの類似度を返す
    const similarity = Math.random() * 0.5 + 0.5; // 0.5-1.0の範囲
    
    setComparisonPairs(prev => [...prev, {
      id: Date.now().toString(),
      imageA,
      imageB,
      similarity,
    }]);
    
    return similarity;
  }, []);
  
  const findSimilarImages = useCallback((
    targetImage: string,
    threshold: number = 0.7
  ) => {
    return comparisonPairs
      .filter(pair => 
        (pair.imageA === targetImage || pair.imageB === targetImage) &&
        pair.similarity >= threshold
      )
      .map(pair => ({
        image: pair.imageA === targetImage ? pair.imageB : pair.imageA,
        similarity: pair.similarity,
      }));
  }, [comparisonPairs]);
  
  return {
    comparisonPairs,
    compareImages,
    findSimilarImages,
  };
};