'use client';

import React, { useState, useRef, useCallback } from 'react';
import { useMultimodal, useImageAnnotation } from '../../hooks/heuristics/useMultimodal';
import { useQuantification } from '../../hooks/heuristics/useQuantification';

interface AnnotationRegion {
  x: number;
  y: number;
  width: number;
  height: number;
}

const MultimodalQuantification: React.FC = () => {
  const {
    multimodalData,
    visualMetaphors,
    currentProcessing,
    statistics,
    loading,
    error,
    processTextWithImage,
    calibrate,
    verify,
    findDirectMapping,
    matchVisualMetaphor,
    generateConfirmationImages,
    reset,
  } = useMultimodal();
  
  const {
    annotations,
    annotateImage,
    getAnnotations,
    clearAnnotations,
  } = useImageAnnotation();
  
  const { quantifySensoryInput } = useQuantification();
  
  // ローカル状態
  const [textInput, setTextInput] = useState('');
  const [selectedImage, setSelectedImage] = useState<File | null>(null);
  const [previewUrl, setPreviewUrl] = useState<string>('');
  const [showCalibration, setShowCalibration] = useState(false);
  const [annotationMode, setAnnotationMode] = useState(false);
  const [currentRegion, setCurrentRegion] = useState<AnnotationRegion | null>(null);
  
  const imageInputRef = useRef<HTMLInputElement>(null);
  const canvasRef = useRef<HTMLCanvasElement>(null);
  
  // 画像選択処理
  const handleImageSelect = useCallback((e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      setSelectedImage(file);
      const url = URL.createObjectURL(file);
      setPreviewUrl(url);
    }
  }, []);
  
  // テキストと画像の処理
  const handleProcess = async () => {
    if (!textInput.trim()) return;
    
    try {
      const result = await processTextWithImage(
        textInput,
        selectedImage || undefined,
        1, // userId
        1  // taskId
      );
      
      console.log('処理結果:', result);
    } catch (err) {
      console.error('処理エラー:', err);
    }
  };
  
  // キャリブレーション処理
  const handleCalibration = async (referenceObject: string) => {
    if (!selectedImage) return;
    
    try {
      const result = await calibrate(1, referenceObject, selectedImage);
      console.log('キャリブレーション結果:', result);
      setShowCalibration(false);
    } catch (err) {
      console.error('キャリブレーションエラー:', err);
    }
  };
  
  // フィードバック処理
  const handleFeedback = async (dataId: string, feedback: string) => {
    try {
      await verify(dataId, feedback);
    } catch (err) {
      console.error('検証エラー:', err);
    }
  };
  
  // アノテーション描画
  const drawAnnotations = useCallback(() => {
    if (!canvasRef.current || !previewUrl) return;
    
    const canvas = canvasRef.current;
    const ctx = canvas.getContext('2d');
    if (!ctx) return;
    
    const img = new Image();
    img.onload = () => {
      canvas.width = img.width;
      canvas.height = img.height;
      ctx.drawImage(img, 0, 0);
      
      // アノテーションを描画
      const imageAnnotations = getAnnotations(previewUrl);
      ctx.strokeStyle = '#00ff00';
      ctx.lineWidth = 2;
      ctx.font = '14px Arial';
      ctx.fillStyle = '#00ff00';
      
      imageAnnotations.forEach((annotation: any) => {
        ctx.strokeRect(
          annotation.region.x,
          annotation.region.y,
          annotation.region.width,
          annotation.region.height
        );
        ctx.fillText(
          `${annotation.quantification.value}${annotation.quantification.unit}`,
          annotation.region.x,
          annotation.region.y - 5
        );
      });
    };
    img.src = previewUrl;
  }, [previewUrl, getAnnotations]);
  
  // クリーンアップ
  React.useEffect(() => {
    return () => {
      if (previewUrl) {
        URL.revokeObjectURL(previewUrl);
      }
    };
  }, [previewUrl]);
  
  return (
    <div className="p-6 max-w-6xl mx-auto">
      <h2 className="text-2xl font-bold mb-6">マルチモーダル定量化</h2>
      
      {/* 入力セクション */}
      <div className="bg-white rounded-lg shadow p-6 mb-6">
        <h3 className="text-lg font-semibold mb-4">入力</h3>
        
        <div className="space-y-4">
          {/* テキスト入力 */}
          <div>
            <label className="block text-sm font-medium mb-2">
              言語表現
            </label>
            <input
              type="text"
              value={textInput}
              onChange={(e) => setTextInput(e.target.value)}
              placeholder="例: 小さじ1杯、手のひらサイズ、少し多め"
              className="w-full px-3 py-2 border rounded-lg"
            />
          </div>
          
          {/* 画像入力 */}
          <div>
            <label className="block text-sm font-medium mb-2">
              参照画像（オプション）
            </label>
            <input
              ref={imageInputRef}
              type="file"
              accept="image/*"
              onChange={handleImageSelect}
              className="hidden"
            />
            <button
              onClick={() => imageInputRef.current?.click()}
              className="px-4 py-2 bg-gray-200 rounded hover:bg-gray-300"
            >
              画像を選択
            </button>
            
            {previewUrl && (
              <div className="mt-4 relative">
                <img
                  src={previewUrl}
                  alt="Preview"
                  className="max-w-md rounded"
                />
                {annotationMode && (
                  <canvas
                    ref={canvasRef}
                    className="absolute top-0 left-0 cursor-crosshair"
                    onClick={(e) => {
                      // アノテーション追加ロジック
                    }}
                  />
                )}
              </div>
            )}
          </div>
          
          {/* アクションボタン */}
          <div className="flex space-x-4">
            <button
              onClick={handleProcess}
              disabled={loading || !textInput.trim()}
              className="px-6 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 disabled:opacity-50"
            >
              {loading ? '処理中...' : '定量化'}
            </button>
            
            <button
              onClick={() => setAnnotationMode(!annotationMode)}
              disabled={!previewUrl}
              className="px-4 py-2 bg-green-500 text-white rounded hover:bg-green-600 disabled:opacity-50"
            >
              {annotationMode ? 'アノテーション終了' : 'アノテーション'}
            </button>
            
            <button
              onClick={() => setShowCalibration(!showCalibration)}
              className="px-4 py-2 bg-purple-500 text-white rounded hover:bg-purple-600"
            >
              キャリブレーション
            </button>
            
            <button
              onClick={reset}
              className="px-4 py-2 bg-gray-500 text-white rounded hover:bg-gray-600"
            >
              リセット
            </button>
          </div>
        </div>
      </div>
      
      {/* 直接マッピング候補 */}
      {textInput && (
        <div className="bg-white rounded-lg shadow p-6 mb-6">
          <h3 className="text-lg font-semibold mb-4">マッピング候補</h3>
          
          <div className="grid grid-cols-2 gap-4">
            {/* 直接マッピング */}
            <div>
              <h4 className="font-medium mb-2">直接マッピング</h4>
              {(() => {
                const mapping = findDirectMapping(textInput);
                return mapping ? (
                  <div className="p-3 bg-blue-50 rounded">
                    <p className="font-semibold">
                      {mapping.value} {mapping.unit}
                    </p>
                    <p className="text-sm text-gray-600">
                      範囲: {mapping.range[0]}-{mapping.range[1]}
                    </p>
                    <p className="text-sm text-gray-600">
                      信頼度: {(mapping.confidence * 100).toFixed(0)}%
                    </p>
                  </div>
                ) : (
                  <p className="text-gray-500">該当なし</p>
                );
              })()}
            </div>
            
            {/* ビジュアルメタファー */}
            <div>
              <h4 className="font-medium mb-2">ビジュアルメタファー</h4>
              {(() => {
                const metaphor = matchVisualMetaphor(textInput);
                return metaphor ? (
                  <div className="p-3 bg-green-50 rounded">
                    <p className="font-semibold">{metaphor.metaphor}</p>
                    <p className="text-sm text-gray-600">
                      サイズ: {metaphor.dimensions.width} × {metaphor.dimensions.height} cm
                    </p>
                    <p className="text-sm text-gray-600">
                      変動幅: {metaphor.variability.min}-{metaphor.variability.max}倍
                    </p>
                  </div>
                ) : (
                  <p className="text-gray-500">該当なし</p>
                );
              })()}
            </div>
          </div>
          
          {/* 確認画像 */}
          <div className="mt-4">
            <h4 className="font-medium mb-2">確認用画像</h4>
            <div className="flex space-x-4">
              {generateConfirmationImages(textInput).map((url, index) => (
                <img
                  key={index}
                  src={url}
                  alt={`確認画像${index + 1}`}
                  className="w-32 h-32 object-cover rounded"
                  onError={(e) => {
                    (e.target as HTMLImageElement).src = '/placeholder.png';
                  }}
                />
              ))}
            </div>
          </div>
        </div>
      )}
      
      {/* 処理結果 */}
      {currentProcessing.status === 'completed' && currentProcessing.result && (
        <div className="bg-white rounded-lg shadow p-6 mb-6">
          <h3 className="text-lg font-semibold mb-4">定量化結果</h3>
          
          <div className="p-4 bg-yellow-50 rounded">
            <p className="text-2xl font-bold">
              {currentProcessing.result.value} {currentProcessing.result.unit}
            </p>
            <p className="text-gray-600">
              範囲: {currentProcessing.result.range[0]}-{currentProcessing.result.range[1]}
            </p>
            <p className="text-gray-600">
              信頼度: {(currentProcessing.result.confidence * 100).toFixed(0)}%
            </p>
          </div>
          
          {/* フィードバックボタン */}
          <div className="mt-4 flex space-x-2">
            <button
              onClick={() => handleFeedback('current', 'correct')}
              className="px-3 py-1 bg-green-100 text-green-700 rounded hover:bg-green-200"
            >
              正確
            </button>
            <button
              onClick={() => handleFeedback('current', 'too_high')}
              className="px-3 py-1 bg-orange-100 text-orange-700 rounded hover:bg-orange-200"
            >
              多すぎる
            </button>
            <button
              onClick={() => handleFeedback('current', 'too_low')}
              className="px-3 py-1 bg-orange-100 text-orange-700 rounded hover:bg-orange-200"
            >
              少なすぎる
            </button>
            <button
              onClick={() => handleFeedback('current', 'incorrect')}
              className="px-3 py-1 bg-red-100 text-red-700 rounded hover:bg-red-200"
            >
              不正確
            </button>
          </div>
        </div>
      )}
      
      {/* 統計情報 */}
      <div className="bg-white rounded-lg shadow p-6">
        <h3 className="text-lg font-semibold mb-4">統計情報</h3>
        
        <div className="grid grid-cols-4 gap-4">
          <div className="text-center">
            <p className="text-2xl font-bold">{statistics.totalProcessed}</p>
            <p className="text-sm text-gray-600">処理済み</p>
          </div>
          <div className="text-center">
            <p className="text-2xl font-bold">
              {(statistics.accuracyRate * 100).toFixed(0)}%
            </p>
            <p className="text-sm text-gray-600">精度</p>
          </div>
          <div className="text-center">
            <p className="text-2xl font-bold">
              {(statistics.averageConfidence * 100).toFixed(0)}%
            </p>
            <p className="text-sm text-gray-600">平均信頼度</p>
          </div>
          <div className="text-center">
            <p className="text-2xl font-bold">
              {(statistics.userConfirmationRate * 100).toFixed(0)}%
            </p>
            <p className="text-sm text-gray-600">確認率</p>
          </div>
        </div>
      </div>
      
      {/* キャリブレーションモーダル */}
      {showCalibration && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white rounded-lg p-6 max-w-md">
            <h3 className="text-lg font-semibold mb-4">
              ユーザーキャリブレーション
            </h3>
            <p className="mb-4">
              参照オブジェクトの画像を選択してください
            </p>
            <div className="space-y-2">
              <button
                onClick={() => handleCalibration('hand')}
                className="w-full px-4 py-2 bg-blue-100 rounded hover:bg-blue-200"
              >
                手のひら
              </button>
              <button
                onClick={() => handleCalibration('finger')}
                className="w-full px-4 py-2 bg-blue-100 rounded hover:bg-blue-200"
              >
                指
              </button>
              <button
                onClick={() => handleCalibration('common_object')}
                className="w-full px-4 py-2 bg-blue-100 rounded hover:bg-blue-200"
              >
                身近な物
              </button>
            </div>
            <button
              onClick={() => setShowCalibration(false)}
              className="mt-4 w-full px-4 py-2 bg-gray-200 rounded hover:bg-gray-300"
            >
              キャンセル
            </button>
          </div>
        </div>
      )}
      
      {/* エラー表示 */}
      {error && (
        <div className="mt-4 p-4 bg-red-100 text-red-700 rounded">
          {error}
        </div>
      )}
    </div>
  );
};

export default MultimodalQuantification;