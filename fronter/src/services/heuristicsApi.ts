import {
  HeuristicsAnalysis,
  HeuristicsAnalysisRequest,
  HeuristicsTracking,
  HeuristicsTrackingData,
  HeuristicsPattern,
  HeuristicsModel,
  HeuristicsTrainRequest,
  HeuristicsInsight,
} from '../model/heuristics';
import { ApplicationError, ErrorCode, parseErrorResponse } from '../errors/errorCodes';
import { SuccessResponse } from './taskApi';
import { fetchApiJsonCore } from "@/utils/fetchApi";

const API_BASE = '/api/heuristics';

const handleApiError = async (response: Response): Promise<never> => {
  let errorData: any;

  try {
    errorData = await response.json();
  } catch {
    switch (response.status) {
      case 400:
        throw new ApplicationError(ErrorCode.VAL_INVALID_INPUT, 'リクエストが無効です');
      case 401:
        throw new ApplicationError(ErrorCode.AUTH_INVALID_CREDENTIALS, '認証が必要です');
      case 403:
        throw new ApplicationError(ErrorCode.AUTH_UNAUTHORIZED, 'アクセス権限がありません');
      case 404:
        throw new ApplicationError(ErrorCode.RES_NOT_FOUND, 'リソースが見つかりません');
      case 500:
        throw new ApplicationError(ErrorCode.SYS_INTERNAL_ERROR, 'サーバーエラーが発生しました');
      default:
        throw new ApplicationError(ErrorCode.SYS_INTERNAL_ERROR, `HTTPエラー: ${response.status}`);
    }
  }

  if (errorData.code) {
    throw new ApplicationError(errorData.code, errorData.message, errorData.detail);
  }
  
  throw new ApplicationError(ErrorCode.SYS_INTERNAL_ERROR, '予期しないエラーが発生しました');
};

// 分析関連
export const analyzeData = async (request: HeuristicsAnalysisRequest) => {
  try {
    const data = await fetchApiJsonCore<HeuristicsAnalysisRequest,HeuristicsAnalysis[]>({
      endpoint: `${API_BASE}/analyze`,
      method: 'POST',
      body: request,
      errorMessage: 'error analyzeData アナリシス一覧取得失敗',
    });

    return data;
  } catch (err) {
    const appError = parseErrorResponse(err);
    return { error: appError.message, code: appError.code };
  }
};

export const getAnalysisById = async (id: string) => {
  try {
    const data = await fetchApiJsonCore<HeuristicsAnalysisRequest,HeuristicsAnalysis>({
      endpoint: `${API_BASE}/analyze/${id}`,
      method: 'GET',
      errorMessage: 'error getAnalysisById アナリシス情報取得失敗',
    });

    return data;
  } catch (err) {
    const appError = parseErrorResponse(err);
    return { error: appError.message, code: appError.code };
  }
};

// トラッキング関連
export const trackBehavior = async (trackData: HeuristicsTrackingData) => {
  try {
    const data = await fetchApiJsonCore<HeuristicsAnalysisRequest,HeuristicsTracking>({
      endpoint: `${API_BASE}/track`,
      method: 'POST',
      body: trackData,
      errorMessage: 'error trackBehavior トラッキング情報一覧の情報取得失敗',
    });

    return data;
  } catch (err) {
    const appError = parseErrorResponse(err);
    return { error: appError.message, code: appError.code };
  }
};

export const getTrackingData = async (userId: string) => {
  try {
    const data = await fetchApiJsonCore<undefined,HeuristicsTracking>({
      endpoint: `${API_BASE}/track/${userId}`,
      method: 'GET',
      errorMessage: 'error getTrackingData トラッキング一覧取得失敗',
    });
    return data;
  } catch (err) {
    const appError = parseErrorResponse(err);
    return { error: appError.message, code: appError.code };
  }
};

// インサイト関連
export const fetchInsights = async (params?: { limit?: number; offset?: number; user_id?: string }) => {
  try {
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
  } catch (err) {
    const appError = parseErrorResponse(err);
    return { error: appError.message, code: appError.code };
  }
};

export const getInsightById = async (id: string) => {
  try {

    const data = await fetchApiJsonCore<undefined,HeuristicsInsight>({
      endpoint: `${API_BASE}/insights/${id}`,
      method: 'GET',
      errorMessage: 'error getInsightById インサイト情報取得失敗',
    });

    return data;
  } catch (err) {
    const appError = parseErrorResponse(err);
    return { error: appError.message, code: appError.code };
  }
};

// パターン検出関連
export const detectPatterns = async (params?: { user_id?: string; data_type?: string; period?: string }) => {
  const queryParams = new URLSearchParams();
  if (params?.user_id) queryParams.append('user_id', params.user_id);
  if (params?.data_type) queryParams.append('data_type', params.data_type);
  if (params?.period) queryParams.append('period', params.period);

  const url = `${API_BASE}/patterns${queryParams.toString() ? `?${queryParams.toString()}` : ''}`;

  const data = await fetchApiJsonCore<undefined,HeuristicsInsight>({
    endpoint: url,
    method: 'GET',
    errorMessage: 'error detectPatterns パターンでの情報取得失敗',
  });
  if('error' in data) {
    return data.error
  }
  return data
};

// モデルトレーニング関連
export const trainModel = async (request: HeuristicsTrainRequest) => {
  try {
    const data = await fetchApiJsonCore<HeuristicsTrainRequest,HeuristicsInsight>({
      endpoint: `${API_BASE}/patterns/train`,
      method: 'POST',
      body: request,
      errorMessage: 'error trainModel trainModelでの情報取得失敗',
    });

    if('error' in data) {
      return data.error
    }
    return data
  } catch (err) {
    const appError = parseErrorResponse(err);
    return { error: appError.message, code: appError.code };
  }
};