import {
  HeuristicsAnalysis,
  HeuristicsAnalysisRequest,
  HeuristicsTracking,
  HeuristicsTrackingData,
  HeuristicsTrainRequest,
  HeuristicsInsight,
  HeuristicsPattern,
  HeuristicsModel
} from '../model/heuristics';
import { ApplicationError, ErrorCode, parseErrorResponse } from '../errors/errorCodes';
import { fetchApiJsonCore } from "@/utils/fetchApi";

const API_BASE = '/api/heuristics';

// 分析関連
export const analyzeData = async (request: HeuristicsAnalysisRequest) => {
  const data = await fetchApiJsonCore<HeuristicsAnalysisRequest,HeuristicsAnalysis[]>({
    endpoint: `${API_BASE}/analyze`,
    method: 'POST',
    body: request,
    errorMessage: 'error analyzeData アナリシス一覧取得失敗',
  });
  return data;
};

export const getAnalysisById = async (id: string) => {
  const data = await fetchApiJsonCore<HeuristicsAnalysisRequest,HeuristicsAnalysis>({
    endpoint: `${API_BASE}/analyze/${id}`,
    method: 'GET',
    errorMessage: 'error getAnalysisById アナリシス情報取得失敗',
  });
  if ('error' in data) {
    return data;
  }
  return data.value;
};

// トラッキング関連
export const trackBehavior = async (trackData: HeuristicsTrackingData) => {
  const data = await fetchApiJsonCore<HeuristicsAnalysisRequest,HeuristicsTracking>({
    endpoint: `${API_BASE}/track`,
    method: 'POST',
    body: trackData,
    errorMessage: 'error trackBehavior トラッキング情報一覧の情報取得失敗',
  });

  return data;
};

export const getTrackingData = async () => {
  const data = await fetchApiJsonCore<undefined,HeuristicsTracking[]>({
    endpoint: `${API_BASE}/track`,
    method: 'GET',
    errorMessage: 'error getTrackingData トラッキング一覧取得失敗',
  });
  return data;
};

// インサイト関連
export const fetchInsights = async (params?: { limit?: number; offset?: number; user_id?: string }) => {
  const queryParams = new URLSearchParams();
  if (params?.limit) queryParams.append('limit', params.limit.toString());
  if (params?.offset) queryParams.append('offset', params.offset.toString());
  if (params?.user_id) queryParams.append('user_id', params.user_id);

  const url = `${API_BASE}/insights${queryParams.toString() ? `?${queryParams.toString()}` : ''}`;

  const data = await fetchApiJsonCore<undefined,HeuristicsInsight[]>({
    endpoint: url,
    method: 'GET',
    errorMessage: 'error getTrackingData トラッキング一覧取得失敗',
  });
  return data;
};

export const getInsightById = async (id: string) => {
  const data = await fetchApiJsonCore<undefined,HeuristicsInsight>({
    endpoint: `${API_BASE}/insights/${id}`,
    method: 'GET',
    errorMessage: 'error getInsightById インサイト情報取得失敗',
  });
  return data;
};

// パターン検出関連
export const detectPatterns = async (params?: { user_id?: string; data_type?: string; period?: string }) => {
  const queryParams = new URLSearchParams();
  if (params?.user_id) queryParams.append('user_id', params.user_id);
  if (params?.data_type) queryParams.append('data_type', params.data_type);
  if (params?.period) queryParams.append('period', params.period);

  const url = `${API_BASE}/patterns${queryParams.toString() ? `?${queryParams.toString()}` : ''}`;

  const data = await fetchApiJsonCore<undefined,HeuristicsPattern[]>({
    endpoint: url,
    method: 'GET',
    errorMessage: 'error detectPatterns パターンでの情報取得失敗',
  });
  return data
};

// モデルトレーニング関連
export const trainModel = async (request: HeuristicsTrainRequest) => {
  const data = await fetchApiJsonCore<HeuristicsTrainRequest,HeuristicsModel>({
    endpoint: `${API_BASE}/patterns/train`,
    method: 'POST',
    body: request,
    errorMessage: 'error trainModel trainModelでの情報取得失敗',
  });

  return data
};