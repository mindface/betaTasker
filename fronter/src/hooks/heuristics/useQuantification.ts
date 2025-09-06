import { useCallback, useMemo } from 'react';
import { useAppDispatch, useAppSelector } from '../../app/hooks';
import {
  collectQuantificationData,
  analyzePattern,
  calculateMetrics,
  addLocalQuantificationData,
  updateQuantificationLevel,
  recordPatternEvolution,
} from '../../features/heuristics/quantificationSlice';

interface QuantificationInput {
  rawValue: string;
  context: Record<string, any>;
  domain: string;
}

interface QuantificationResult {
  value: number;
  unit: string;
  confidence: number;
  level: 1 | 2 | 3 | 4;
}

export const useQuantification = () => {
  const dispatch = useAppDispatch();
  const quantificationState = useAppSelector((state) => state.quantification);
  
  // 感覚的入力を定量化
  const quantifySensoryInput = useCallback((input: QuantificationInput): QuantificationResult => {
    // 感覚的表現のマッピング例
    const sensoryMappings: Record<string, { value: number; unit: string; variance: number }> = {
      '少し': { value: 0.2, unit: 'ratio', variance: 0.1 },
      '小さじ1杯': { value: 5, unit: 'ml', variance: 0.5 },
      '中程度': { value: 0.5, unit: 'ratio', variance: 0.15 },
      '多め': { value: 0.8, unit: 'ratio', variance: 0.1 },
      'かなり': { value: 0.9, unit: 'ratio', variance: 0.05 },
    };
    
    // パターンマッチング
    let bestMatch = null;
    let confidence = 0;
    
    for (const [pattern, mapping] of Object.entries(sensoryMappings)) {
      if (input.rawValue.includes(pattern)) {
        bestMatch = mapping;
        confidence = 0.8; // 基本信頼度
        break;
      }
    }
    
    // コンテキストによる調整
    if (bestMatch && input.context.historicalData) {
      confidence += 0.1; // 履歴データがある場合は信頼度向上
    }
    
    // デフォルト値
    if (!bestMatch) {
      bestMatch = { value: 0, unit: 'unknown', variance: 1 };
      confidence = 0.1;
    }
    
    // 定量化レベルの判定
    let level: 1 | 2 | 3 | 4 = 1;
    if (confidence >= 0.9 && bestMatch.variance < 0.1) {
      level = 4; // 体系的定量化
    } else if (confidence >= 0.7 && bestMatch.variance < 0.3) {
      level = 3; // 構造化定量化
    } else if (confidence >= 0.5) {
      level = 2; // 部分的定量化
    }
    
    return {
      value: bestMatch.value,
      unit: bestMatch.unit,
      confidence,
      level,
    };
  }, []);
  
  // パターン認識と進化
  const recognizePattern = useCallback((dataPoints: any[]) => {
    // 頻出パターンの検出
    const patterns: Record<string, number> = {};
    
    dataPoints.forEach(point => {
      const key = `${point.action}_${point.result}`;
      patterns[key] = (patterns[key] || 0) + 1;
    });
    
    // 最頻出パターンの特定
    const topPattern = Object.entries(patterns)
      .sort(([, a], [, b]) => b - a)[0];
    
    if (topPattern) {
      const [patternKey, frequency] = topPattern;
      const confidence = frequency / dataPoints.length;
      
      return {
        pattern: patternKey,
        frequency,
        confidence,
        type: classifyPatternType(confidence, dataPoints),
      };
    }
    
    return null;
  }, []);
  
  // パターンタイプの分類
  const classifyPatternType = (confidence: number, dataPoints: any[]): string => {
    if (confidence > 0.9) return '同一事象';
    if (confidence > 0.7) return '類似事象';
    if (confidence > 0.5) return '類推可能';
    return '組み合わせ';
  };
  
  // データ収集と定量化
  const collectAndQuantify = useCallback(async (
    userId: number,
    taskId: number,
    rawData: any
  ) => {
    const quantificationResult = quantifySensoryInput({
      rawValue: rawData.value || '',
      context: rawData.context || {},
      domain: rawData.domain || 'general',
    });
    
    const quantificationData = {
      id: `${userId}_${taskId}_${Date.now()}`,
      userId,
      taskId,
      level: quantificationResult.level,
      levelDescription: getLevelDescription(quantificationResult.level),
      rawValue: rawData.value || '',
      quantifiedValue: quantificationResult.value,
      unit: quantificationResult.unit,
      variance: 0.1, // デフォルト分散
      patternType: '類似事象' as const,
      patternConfidence: quantificationResult.confidence,
      metrics: {
        reproducibility: quantificationResult.confidence,
        shareability: quantificationResult.level >= 3 ? 0.8 : 0.4,
        standardization: quantificationResult.level >= 3 ? 0.7 : 0.3,
        confidence: quantificationResult.confidence,
      },
      domain: rawData.domain || 'general',
      relatedDomains: [],
      transferability: 0.5,
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString(),
    };
    
    // ローカルに追加
    dispatch(addLocalQuantificationData(quantificationData));
    
    // サーバーに送信
    return dispatch(collectQuantificationData({ userId, taskId, rawData }));
  }, [dispatch, quantifySensoryInput]);
  
  // メトリクス計算
  const refreshMetrics = useCallback(async (userId?: number, period?: string) => {
    return dispatch(calculateMetrics({ userId, period }));
  }, [dispatch]);
  
  // 定量化レベルの評価
  const evaluateQuantificationLevel = useCallback((data: any): 1 | 2 | 3 | 4 => {
    const { reproducibility, standardization } = data.metrics || {};
    
    if (reproducibility >= 0.9 && standardization >= 0.8) return 4;
    if (reproducibility >= 0.7 && standardization >= 0.6) return 3;
    if (reproducibility >= 0.5) return 2;
    return 1;
  }, []);
  
  // 集計メトリクスの取得
  const aggregateMetrics = useMemo(() => {
    return quantificationState.aggregateMetrics;
  }, [quantificationState.aggregateMetrics]);
  
  // KPI達成状況
  const kpiStatus = useMemo(() => {
    const metrics = quantificationState.aggregateMetrics;
    return {
      reproducibility: {
        current: metrics.averageReproducibility,
        target: 0.8,
        achieved: metrics.averageReproducibility >= 0.8,
      },
      level3Rate: {
        current: metrics.level3OrHigherRate,
        target: 0.7,
        achieved: metrics.level3OrHigherRate >= 0.7,
      },
      standardization: {
        current: metrics.averageShareability,
        target: 0.6,
        achieved: metrics.averageShareability >= 0.6,
      },
    };
  }, [quantificationState.aggregateMetrics]);
  
  return {
    // 状態
    quantificationData: quantificationState.quantificationData,
    patternEvolutions: quantificationState.patternEvolutions,
    loading: quantificationState.loading,
    error: quantificationState.error,
    
    // メトリクス
    aggregateMetrics,
    kpiStatus,
    
    // アクション
    quantifySensoryInput,
    recognizePattern,
    collectAndQuantify,
    refreshMetrics,
    evaluateQuantificationLevel,
    
    // Redux アクション
    updateLevel: (id: string, level: 1 | 2 | 3 | 4) => 
      dispatch(updateQuantificationLevel({ id, level })),
    recordEvolution: (evolution: any) => 
      dispatch(recordPatternEvolution(evolution)),
  };
};

// ヘルパー関数
function getLevelDescription(level: 1 | 2 | 3 | 4): string {
  const descriptions = {
    1: '感覚的定量化',
    2: '部分的定量化',
    3: '構造化定量化',
    4: '体系的定量化',
  };
  return descriptions[level];
}