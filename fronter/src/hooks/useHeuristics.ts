import { useCallback } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { AppDispatch, RootState } from '../store';
import {
  analyzeData,
  fetchAnalysisById,
  trackUserBehavior,
  fetchTrackingData,
  loadInsights,
  fetchInsightById,
  loadPatterns,
  trainHeuristicsModel,
  clearAnalysisError,
  clearTrackingError,
  clearInsightsError,
  clearPatternsError,
  clearModelError,
} from '../features/heuristics/heuristicsSlice';
import {
  HeuristicsAnalysisRequest,
  HeuristicsTrackingData,
  HeuristicsTrainRequest,
} from '../model/heuristics';

export const useHeuristics = () => {
  const dispatch = useDispatch<AppDispatch>();
  const heuristicsState = useSelector((state: RootState) => state.heuristics);

  // 分析関連
  const analyze = useCallback(
    (request: HeuristicsAnalysisRequest) => {
      return dispatch(analyzeData(request));
    },
    [dispatch]
  );

  const getAnalysis = useCallback(
    (id: string) => {
      return dispatch(fetchAnalysisById(id));
    },
    [dispatch]
  );

  // トラッキング関連
  const trackBehavior = useCallback(
    (data: HeuristicsTrackingData) => {
      return dispatch(trackUserBehavior(data));
    },
    [dispatch]
  );

  const getTracking = useCallback(
    (userId: string) => {
      return dispatch(fetchTrackingData(userId));
    },
    [dispatch]
  );

  // インサイト関連
  const getInsights = useCallback(
    (params?: { limit?: number; offset?: number; user_id?: string }) => {
      return dispatch(loadInsights(params));
    },
    [dispatch]
  );

  const getInsight = useCallback(
    (id: string) => {
      return dispatch(fetchInsightById(id));
    },
    [dispatch]
  );

  // パターン検出関連
  const getPatterns = useCallback(
    (params?: { user_id?: string; data_type?: string; period?: string }) => {
      return dispatch(loadPatterns(params));
    },
    [dispatch]
  );

  // モデルトレーニング関連
  const trainModel = useCallback(
    (request: HeuristicsTrainRequest) => {
      return dispatch(trainHeuristicsModel(request));
    },
    [dispatch]
  );

  // エラークリア関連
  const clearErrors = useCallback(() => {
    dispatch(clearAnalysisError());
    dispatch(clearTrackingError());
    dispatch(clearInsightsError());
    dispatch(clearPatternsError());
    dispatch(clearModelError());
  }, [dispatch]);

  return {
    // State
    ...heuristicsState,
    
    // Actions
    analyze,
    getAnalysis,
    trackBehavior,
    getTracking,
    getInsights,
    getInsight,
    getPatterns,
    trainModel,
    clearErrors,
  };
};

// 個別のフックも提供
export const useHeuristicsAnalysis = () => {
  const dispatch = useDispatch<AppDispatch>();
  const { analyses, currentAnalysis, analysisLoading, analysisError } = useSelector(
    (state: RootState) => state.heuristics
  );

  const analyze = useCallback(
    (request: HeuristicsAnalysisRequest) => dispatch(analyzeData(request)),
    [dispatch]
  );

  const getAnalysis = useCallback(
    (id: string) => dispatch(fetchAnalysisById(id)),
    [dispatch]
  );

  const clearError = useCallback(
    () => dispatch(clearAnalysisError()),
    [dispatch]
  );

  return {
    analyses,
    currentAnalysis,
    loading: analysisLoading,
    error: analysisError,
    analyze,
    getAnalysis,
    clearError,
  };
};

export const useHeuristicsInsights = () => {
  const dispatch = useDispatch<AppDispatch>();
  const { insights, currentInsight, insightsTotal, insightsLoading, insightsError } = useSelector(
    (state: RootState) => state.heuristics
  );

  const getInsights = useCallback(
    (params?: { limit?: number; offset?: number; user_id?: string }) => 
      dispatch(loadInsights(params)),
    [dispatch]
  );

  const getInsight = useCallback(
    (id: string) => dispatch(fetchInsightById(id)),
    [dispatch]
  );

  const clearError = useCallback(
    () => dispatch(clearInsightsError()),
    [dispatch]
  );

  return {
    insights,
    currentInsight,
    total: insightsTotal,
    loading: insightsLoading,
    error: insightsError,
    getInsights,
    getInsight,
    clearError,
  };
};

export const useHeuristicsTracking = () => {
  const dispatch = useDispatch<AppDispatch>();
  const { trackingData, trackingLoading, trackingError } = useSelector(
    (state: RootState) => state.heuristics
  );

  const track = useCallback(
    (data: HeuristicsTrackingData) => dispatch(trackUserBehavior(data)),
    [dispatch]
  );

  const getTracking = useCallback(
    (userId: string) => dispatch(fetchTrackingData(userId)),
    [dispatch]
  );

  const clearError = useCallback(
    () => dispatch(clearTrackingError()),
    [dispatch]
  );

  return {
    trackingData,
    loading: trackingLoading,
    error: trackingError,
    track,
    getTracking,
    clearError,
  };
};

export const useHeuristicsPatterns = () => {
  const dispatch = useDispatch<AppDispatch>();
  const { patterns, patternsLoading, patternsError } = useSelector(
    (state: RootState) => state.heuristics
  );

  const getPatterns = useCallback(
    (params?: { user_id?: string; data_type?: string; period?: string }) =>
      dispatch(loadPatterns(params)),
    [dispatch]
  );

  const clearError = useCallback(
    () => dispatch(clearPatternsError()),
    [dispatch]
  );

  return {
    patterns,
    loading: patternsLoading,
    error: patternsError,
    getPatterns,
    clearError,
  };
};