import { useState, useCallback, useMemo } from 'react';
import { useAppSelector } from '../../app/hooks';
import { QuantificationLabel, LabelRevision } from '../../model/quantificationLabel';
import { getLabelHistoryService } from '../../services/labelApi';

interface HistoryEntry {
  id: string;
  timestamp: string;
  action: 'created' | 'updated' | 'verified' | 'annotated';
  userId: string;
  changes?: {
    field: string;
    oldValue: any;
    newValue: any;
    reason?: string;
  }[];
  metadata?: {
    confidence: number;
    accuracy: number;
    verification?: 'approved' | 'rejected';
  };
}

interface HistoryAnalytics {
  totalEntries: number;
  averageAccuracyTrend: number[];
  confidenceTrend: number[];
  verificationRate: number;
  mostActiveUsers: Array<{ userId: string; count: number }>;
  popularConcepts: Array<{ concept: string; count: number }>;
  domainDistribution: Record<string, number>;
}

interface UseQuantificationHistoryReturn {
  // データ
  history: HistoryEntry[];
  analytics: HistoryAnalytics;
  
  // フィルタリング
  filteredHistory: HistoryEntry[];
  filters: {
    dateRange: { start: string; end: string } | null;
    userId: string | null;
    action: string | null;
    concept: string | null;
  };
  
  // 状態
  loading: boolean;
  error: string | null;
  
  // アクション
  loadHistory: (labelId?: string, period?: string) => Promise<void>;
  setDateFilter: (start: string, end: string) => void;
  setUserFilter: (userId: string | null) => void;
  setActionFilter: (action: string | null) => void;
  setConceptFilter: (concept: string | null) => void;
  clearFilters: () => void;
  
  // 解析
  getAccuracyEvolution: (labelId: string) => Array<{ date: string; accuracy: number }>;
  getUsagePattern: (concept: string) => Array<{ period: string; count: number }>;
  compareLabels: (labelIds: string[]) => any;
  predictTrend: (metric: string, days: number) => Array<{ date: string; predicted: number }>;
}

export const useQuantificationHistory = (): UseQuantificationHistoryReturn => {
  const [history, setHistory] = useState<HistoryEntry[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  
  const [filters, setFilters] = useState({
    dateRange: null as { start: string; end: string } | null,
    userId: null as string | null,
    action: null as string | null,
    concept: null as string | null,
  });
  
  // Redux状態から関連データを取得
  const labels = useAppSelector((state) => state.label?.labels || []);
  const multimodalData = useAppSelector((state) => state.multimodal?.multimodalData || []);
  
  // 履歴データのロード
  const loadHistory = useCallback(async (labelId?: string, period?: string) => {
    setLoading(true);
    setError(null);
    
    try {
      // 複数のソースから履歴データを収集
      const historyPromises: Promise<any>[] = [];
      
      if (labelId) {
        // 特定ラベルの履歴
        historyPromises.push(getLabelHistoryService(labelId));
      } else {
        // 全体の履歴を構築
        historyPromises.push(
          Promise.resolve({
            entries: labels.flatMap(label => 
              label.history.revisions.map(revision => ({
                id: `${label.id}_${revision.revisionId}`,
                timestamp: revision.timestamp,
                action: 'updated' as const,
                userId: revision.userId,
                changes: revision.changes,
                labelId: label.id,
                concept: label.concept.relatedConcepts[0] || 'unknown',
              }))
            )
          })
        );
      }
      
      // マルチモーダル処理履歴
      const multimodalHistory = multimodalData.map(data => ({
        id: data.id,
        timestamp: data.timestamp,
        action: 'created' as const,
        userId: data.userId.toString(),
        metadata: {
          confidence: data.quantification.confidence,
          accuracy: data.association.correlationScore,
        },
        concept: data.linguistic.text,
      }));
      
      const results = await Promise.all(historyPromises);
      const combinedHistory = [
        ...results.flatMap(result => result.entries || []),
        ...multimodalHistory,
      ].sort((a, b) => new Date(b.timestamp).getTime() - new Date(a.timestamp).getTime());
      
      setHistory(combinedHistory);
      
    } catch (err: any) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  }, [labels, multimodalData]);
  
  // フィルタリングされた履歴
  const filteredHistory = useMemo(() => {
    let filtered = [...history];
    
    if (filters.dateRange) {
      const start = new Date(filters.dateRange.start);
      const end = new Date(filters.dateRange.end);
      filtered = filtered.filter(entry => {
        const entryDate = new Date(entry.timestamp);
        return entryDate >= start && entryDate <= end;
      });
    }
    
    if (filters.userId) {
      filtered = filtered.filter(entry => entry.userId === filters.userId);
    }
    
    if (filters.action) {
      filtered = filtered.filter(entry => entry.action === filters.action);
    }
    
    if (filters.concept) {
      filtered = filtered.filter(entry => 
        (entry as any).concept?.includes(filters.concept)
      );
    }
    
    return filtered;
  }, [history, filters]);
  
  // 分析データ
  const analytics = useMemo((): HistoryAnalytics => {
    if (history.length === 0) {
      return {
        totalEntries: 0,
        averageAccuracyTrend: [],
        confidenceTrend: [],
        verificationRate: 0,
        mostActiveUsers: [],
        popularConcepts: [],
        domainDistribution: {},
      };
    }
    
    // ユーザー活動度
    const userActivity: Record<string, number> = {};
    history.forEach(entry => {
      userActivity[entry.userId] = (userActivity[entry.userId] || 0) + 1;
    });
    
    // 概念の人気度
    const conceptPopularity: Record<string, number> = {};
    history.forEach(entry => {
      const concept = (entry as any).concept;
      if (concept) {
        conceptPopularity[concept] = (conceptPopularity[concept] || 0) + 1;
      }
    });
    
    // 精度トレンド（過去30日）
    const last30Days = Array.from({ length: 30 }, (_, i) => {
      const date = new Date();
      date.setDate(date.getDate() - i);
      return date.toISOString().split('T')[0];
    }).reverse();
    
    const accuracyTrend = last30Days.map(date => {
      const dayEntries = history.filter(entry => 
        entry.timestamp.startsWith(date) && entry.metadata?.accuracy
      );
      
      if (dayEntries.length === 0) return 0;
      
      const avgAccuracy = dayEntries.reduce(
        (sum, entry) => sum + (entry.metadata?.accuracy || 0), 0
      ) / dayEntries.length;
      
      return avgAccuracy;
    });
    
    // 信頼度トレンド
    const confidenceTrend = last30Days.map(date => {
      const dayEntries = history.filter(entry => 
        entry.timestamp.startsWith(date) && entry.metadata?.confidence
      );
      
      if (dayEntries.length === 0) return 0;
      
      const avgConfidence = dayEntries.reduce(
        (sum, entry) => sum + (entry.metadata?.confidence || 0), 0
      ) / dayEntries.length;
      
      return avgConfidence;
    });
    
    // 検証率
    const verificationEntries = history.filter(entry => entry.action === 'verified');
    const verificationRate = verificationEntries.length / history.length;
    
    return {
      totalEntries: history.length,
      averageAccuracyTrend: accuracyTrend,
      confidenceTrend,
      verificationRate,
      mostActiveUsers: Object.entries(userActivity)
        .sort(([, a], [, b]) => b - a)
        .slice(0, 5)
        .map(([userId, count]) => ({ userId, count })),
      popularConcepts: Object.entries(conceptPopularity)
        .sort(([, a], [, b]) => b - a)
        .slice(0, 10)
        .map(([concept, count]) => ({ concept, count })),
      domainDistribution: {}, // TODO: ドメイン情報から計算
    };
  }, [history]);
  
  // フィルタ設定関数
  const setDateFilter = useCallback((start: string, end: string) => {
    setFilters(prev => ({ ...prev, dateRange: { start, end } }));
  }, []);
  
  const setUserFilter = useCallback((userId: string | null) => {
    setFilters(prev => ({ ...prev, userId }));
  }, []);
  
  const setActionFilter = useCallback((action: string | null) => {
    setFilters(prev => ({ ...prev, action }));
  }, []);
  
  const setConceptFilter = useCallback((concept: string | null) => {
    setFilters(prev => ({ ...prev, concept }));
  }, []);
  
  const clearFilters = useCallback(() => {
    setFilters({
      dateRange: null,
      userId: null,
      action: null,
      concept: null,
    });
  }, []);
  
  // 精度の進化を取得
  const getAccuracyEvolution = useCallback((labelId: string) => {
    const labelHistory = history.filter(entry => (entry as any).labelId === labelId);
    return labelHistory
      .filter(entry => entry.metadata?.accuracy)
      .map(entry => ({
        date: entry.timestamp.split('T')[0],
        accuracy: entry.metadata!.accuracy,
      }))
      .sort((a, b) => new Date(a.date).getTime() - new Date(b.date).getTime());
  }, [history]);
  
  // 使用パターン取得
  const getUsagePattern = useCallback((concept: string) => {
    const conceptEntries = history.filter(entry => 
      (entry as any).concept?.includes(concept)
    );
    
    // 週別の使用回数を集計
    const weeklyUsage: Record<string, number> = {};
    conceptEntries.forEach(entry => {
      const date = new Date(entry.timestamp);
      const week = getWeekString(date);
      weeklyUsage[week] = (weeklyUsage[week] || 0) + 1;
    });
    
    return Object.entries(weeklyUsage)
      .map(([period, count]) => ({ period, count }))
      .sort((a, b) => a.period.localeCompare(b.period));
  }, [history]);
  
  // ラベル比較
  const compareLabels = useCallback((labelIds: string[]) => {
    const comparison = labelIds.map(labelId => {
      const label = labels.find(l => l.id === labelId);
      const labelHistory = history.filter(entry => (entry as any).labelId === labelId);
      
      return {
        labelId,
        label,
        historyCount: labelHistory.length,
        averageAccuracy: labelHistory
          .filter(entry => entry.metadata?.accuracy)
          .reduce((sum, entry, _, arr) => 
            sum + (entry.metadata!.accuracy / arr.length), 0
          ),
        lastUpdated: labelHistory[0]?.timestamp || null,
        verificationStatus: label?.metadata.validated || false,
      };
    });
    
    return comparison;
  }, [labels, history]);
  
  // トレンド予測（簡単な線形回帰）
  const predictTrend = useCallback((metric: string, days: number) => {
    let data: number[] = [];
    
    if (metric === 'accuracy') {
      data = analytics.averageAccuracyTrend.filter(v => v > 0);
    } else if (metric === 'confidence') {
      data = analytics.confidenceTrend.filter(v => v > 0);
    }
    
    if (data.length < 2) return [];
    
    // 簡単な線形回帰
    const n = data.length;
    const sumX = data.reduce((sum, _, i) => sum + i, 0);
    const sumY = data.reduce((sum, y) => sum + y, 0);
    const sumXY = data.reduce((sum, y, i) => sum + (i * y), 0);
    const sumXX = data.reduce((sum, _, i) => sum + (i * i), 0);
    
    const slope = (n * sumXY - sumX * sumY) / (n * sumXX - sumX * sumX);
    const intercept = (sumY - slope * sumX) / n;
    
    // 未来の予測値を計算
    const predictions = Array.from({ length: days }, (_, i) => {
      const x = data.length + i;
      const predicted = slope * x + intercept;
      const futureDate = new Date();
      futureDate.setDate(futureDate.getDate() + i + 1);
      
      return {
        date: futureDate.toISOString().split('T')[0],
        predicted: Math.max(0, Math.min(1, predicted)), // 0-1の範囲に制限
      };
    });
    
    return predictions;
  }, [analytics]);
  
  return {
    history,
    analytics,
    filteredHistory,
    filters,
    loading,
    error,
    loadHistory,
    setDateFilter,
    setUserFilter,
    setActionFilter,
    setConceptFilter,
    clearFilters,
    getAccuracyEvolution,
    getUsagePattern,
    compareLabels,
    predictTrend,
  };
};

// ヘルパー関数：週の文字列表現を取得
function getWeekString(date: Date): string {
  const year = date.getFullYear();
  const week = getWeekNumber(date);
  return `${year}-W${week.toString().padStart(2, '0')}`;
}

function getWeekNumber(date: Date): number {
  const firstDayOfYear = new Date(date.getFullYear(), 0, 1);
  const pastDaysOfYear = (date.getTime() - firstDayOfYear.getTime()) / 86400000;
  return Math.ceil((pastDaysOfYear + firstDayOfYear.getDay() + 1) / 7);
}