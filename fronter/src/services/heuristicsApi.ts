import {
  HeuristicsAnalysis,
  HeuristicsAnalysisRequest,
  HeuristicsTracking,
  HeuristicsTrackingData,
  HeuristicsPattern,
  HeuristicsModel,
  HeuristicsTrainRequest
} from '../model/heuristics';
import { ApplicationError, ErrorCode, parseErrorResponse } from '../errors/errorCodes';
import { SuccessResponse } from './taskApi';

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
    const res = await fetch(`${API_BASE}/analyze`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(request),
      credentials: 'include',
    });

    if (!res.ok) {
      await handleApiError(res);
    }
    
    const data: SuccessResponse<HeuristicsAnalysis, 'analysis'> = await res.json();
    console.log("analyzeData HeuristicsAnalysis response");
    console.log(data);
    return data.analysis || data;
  } catch (err) {
    const appError = parseErrorResponse(err);
    return { error: appError.message, code: appError.code };
  }
};

export const getAnalysisById = async (id: string) => {
  try {
    const res = await fetch(`${API_BASE}/analyze/${id}`, {
      method: 'GET',
      credentials: 'include',
    });
    
    if (!res.ok) {
      await handleApiError(res);
    }

    const data: SuccessResponse<HeuristicsAnalysis, 'analysis'> = await res.json();
    return data.analysis || data;
  } catch (err) {
    const appError = parseErrorResponse(err);
    return { error: appError.message, code: appError.code };
  }
};

// トラッキング関連
export const trackBehavior = async (trackData: HeuristicsTrackingData) => {
  try {
    const res = await fetch(`${API_BASE}/track`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(trackData),
      credentials: 'include',
    });

    if (!res.ok) {
      await handleApiError(res);
    }

    const data = await res.json();
    console.log("trackBehavior HeuristicsTracking response");
    console.log(data);
    
    return data;
  } catch (err) {
    const appError = parseErrorResponse(err);
    return { error: appError.message, code: appError.code };
  }
};

export const getTrackingData = async (userId: string) => {
  try {
    const res = await fetch(`${API_BASE}/track/${userId}`, {
      method: 'GET',
      credentials: 'include',
    });
    
    if (!res.ok) {
      await handleApiError(res);
    }

    const response: { success: boolean; data: { tracking_data: HeuristicsTracking[] } } = await res.json();
    return response.data?.tracking_data || [];
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
    
    const res = await fetch(url, {
      method: 'GET',
      credentials: 'include',
    });
    
    if (!res.ok) {
      await handleApiError(res);
    }

    const response = await res.json();
    // レスポンス構造: { success: true, data: { insights: [...], total: number, limit: number, offset: number } }
    console.log("fetchInsights response");
    console.log(response);
    return response.data || response;
  } catch (err) {
    const appError = parseErrorResponse(err);
    return { error: appError.message, code: appError.code };
  }
};

export const getInsightById = async (id: string) => {
  try {
    const res = await fetch(`${API_BASE}/insights/${id}`, {
      method: 'GET',
      credentials: 'include',
    });

    if (!res.ok) {
      await handleApiError(res);
    }
    
    const response = await res.json();
    // レスポンス構造: { success: true, data: { insight: {...} } }
    console.log("getInsightById response");
    console.log(response);
    return response.data?.insight || response;
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
 
  const res = await fetch(url, {
    method: 'GET',
    credentials: 'include',
  });

  if (!res.ok) {
    await handleApiError(res);
  }

  const data: SuccessResponse<{ metadata: any; patterns:
  HeuristicsPattern[] }, 'data'> = await res.json();
  return data.data?.patterns || data;
};

// モデルトレーニング関連
export const trainModel = async (request: HeuristicsTrainRequest) => {
  try {
    const res = await fetch(`${API_BASE}/patterns/train`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(request),
      credentials: 'include',
    });
    
    if (!res.ok) {
      await handleApiError(res);
    }

    const data: SuccessResponse<HeuristicsModel, 'model'> = await res.json();
    return data.model || data;
  } catch (err) {
    const appError = parseErrorResponse(err);
    return { error: appError.message, code: appError.code };
  }
};