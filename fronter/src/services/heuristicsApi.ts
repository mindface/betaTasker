import {
  HeuristicsAnalysis,
  HeuristicsAnalysisRequest,
  HeuristicsTracking,
  HeuristicsTrackingData,
  HeuristicsPattern,
  HeuristicsModel,
  HeuristicsTrainRequest,
  HeuristicsInsight,
  AnalysisFilters,
  PatternFilters,
  InsightFilters,
  PaginationState,
  validateAnalysisRequest,
  validateTrackingData,
  validateTrainRequest,
  VALID_ANALYSIS_TYPES,
  DEFAULT_PAGINATION
} from '../model/heuristics';
import { 
  ApplicationError, 
  ErrorCode, 
  parseErrorResponse,
  HeuristicsErrorCode
} from '../errors/errorCodes';

const API_BASE = '/api/heuristics';

// 統一されたエラーハンドリング
const handleApiError = async (response: Response): Promise<never> => {
  let errorData: any;

  try {
    errorData = await response.json();
  } catch {
    errorData = { message: 'Unknown error' };
  }

  // HTTPステータスコードに基づくエラーコードのマッピング
  const errorCode = mapHttpStatusToErrorCode(response.status);
  throw new ApplicationError(errorCode, errorData.message || 'API呼び出しに失敗しました');
};

// HTTPステータスコードをエラーコードにマッピング
const mapHttpStatusToErrorCode = (status: number): ErrorCode => {
  switch (status) {
    case 400:
      return ErrorCode.VAL_INVALID_INPUT;
    case 401:
      return ErrorCode.AUTH_INVALID_CREDENTIALS;
    case 403:
      return ErrorCode.AUTH_UNAUTHORIZED;
    case 404:
      return ErrorCode.RES_NOT_FOUND;
    case 408:
      return HeuristicsErrorCode.ANALYSIS_TIMEOUT;
    case 409:
      return ErrorCode.RES_CONFLICT;
    case 500:
      return ErrorCode.SYS_INTERNAL_ERROR;
    default:
      return ErrorCode.SYS_INTERNAL_ERROR;
  }
};

// 統一されたレスポンス処理
const handleResponse = async <T>(response: Response): Promise<T> => {
  if (!response.ok) {
    await handleApiError(response);
  }
  
  const data = await response.json();
  
  // レスポンス構造の正規化
  if (data.success === false) {
    throw new ApplicationError(
      ErrorCode.API_ERROR,
      data.message || 'API呼び出しに失敗しました'
    );
  }
  
  return data.data || data;
};

// 分析関連のAPI
export const analyzeData = async (request: HeuristicsAnalysisRequest): Promise<HeuristicsAnalysis> => {
  try {
    // バリデーション
    const validation = validateAnalysisRequest(request);
    if (!validation.isValid) {
      throw new ApplicationError(
        HeuristicsErrorCode.ANALYSIS_INVALID_PARAMETERS,
        validation.errors.join(', ')
      );
    }

    const response = await fetch(`${API_BASE}/analyze`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(request),
      credentials: 'include',
    });

    return await handleResponse<HeuristicsAnalysis>(response);
  } catch (err) {
    if (err instanceof ApplicationError) {
      throw err;
    }
    const appError = parseErrorResponse(err);
    throw new ApplicationError(
      HeuristicsErrorCode.ANALYSIS_FAILED,
      '分析処理に失敗しました',
      appError.message
    );
  }
};

export const getAnalysisById = async (id: string): Promise<HeuristicsAnalysis> => {
  try {
    const response = await fetch(`${API_BASE}/analyze/${id}`, {
      method: 'GET',
      credentials: 'include',
    });
    
    return await handleResponse<HeuristicsAnalysis>(response);
  } catch (err) {
    if (err instanceof ApplicationError) {
      throw err;
    }
    const appError = parseErrorResponse(err);
    throw new ApplicationError(
      HeuristicsErrorCode.ANALYSIS_FAILED,
      '分析結果の取得に失敗しました',
      appError.message
    );
  }
};

export const getAnalyses = async (filters?: AnalysisFilters, pagination?: Partial<PaginationState>): Promise<{
  analyses: HeuristicsAnalysis[];
  pagination: PaginationState;
}> => {
  try {
    const queryParams = new URLSearchParams();
    
    if (filters) {
      Object.entries(filters).forEach(([key, value]) => {
        if (value !== undefined && value !== null) {
          if (key === 'date_range' && value.start && value.end) {
            queryParams.append('date_start', value.start);
            queryParams.append('date_end', value.end);
          } else {
            queryParams.append(key, String(value));
          }
        }
      });
    }
    
    if (pagination) {
      if (pagination.page) queryParams.append('page', pagination.page.toString());
      if (pagination.limit) queryParams.append('limit', pagination.limit.toString());
    }
    
    const url = `${API_BASE}/analyze${queryParams.toString() ? `?${queryParams.toString()}` : ''}`;
    const response = await fetch(url, {
      method: 'GET',
      credentials: 'include',
    });
    
    const data = await handleResponse<{
      analyses: HeuristicsAnalysis[];
      pagination: PaginationState;
    }>(response);
    
    return {
      analyses: data.analyses || [],
      pagination: data.pagination || DEFAULT_PAGINATION
    };
  } catch (err) {
    if (err instanceof ApplicationError) {
      throw err;
    }
    const appError = parseErrorResponse(err);
    throw new ApplicationError(
      HeuristicsErrorCode.ANALYSIS_FAILED,
      '分析一覧の取得に失敗しました',
      appError.message
    );
  }
};

// トラッキング関連のAPI
export const trackBehavior = async (trackData: HeuristicsTrackingData): Promise<{ tracking_id: number }> => {
  try {
    // バリデーション
    const validation = validateTrackingData(trackData);
    if (!validation.isValid) {
      throw new ApplicationError(
        HeuristicsErrorCode.TRACKING_INVALID_ACTION,
        validation.errors.join(', ')
      );
    }

    const response = await fetch(`${API_BASE}/track`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(trackData),
      credentials: 'include',
    });

    return await handleResponse<{ tracking_id: number }>(response);
  } catch (err) {
    if (err instanceof ApplicationError) {
      throw err;
    }
    const appError = parseErrorResponse(err);
    throw new ApplicationError(
      HeuristicsErrorCode.TRACKING_FAILED,
      '行動追跡に失敗しました',
      appError.message
    );
  }
};

export const getTrackingData = async (userId: string): Promise<HeuristicsTracking[]> => {
  try {
    const response = await fetch(`${API_BASE}/track/${userId}`, {
      method: 'GET',
      credentials: 'include',
    });
    
    const data = await handleResponse<{ tracking_data: HeuristicsTracking[] }>(response);
    return data.tracking_data || [];
  } catch (err) {
    if (err instanceof ApplicationError) {
      throw err;
    }
    const appError = parseErrorResponse(err);
    throw new ApplicationError(
      HeuristicsErrorCode.TRACKING_FAILED,
      'トラッキングデータの取得に失敗しました',
      appError.message
    );
  }
};

// インサイト関連のAPI
export const fetchInsights = async (params?: {
  limit?: number;
  offset?: number;
  user_id?: string;
  filters?: InsightFilters;
}): Promise<{
  insights: HeuristicsInsight[];
  total: number;
  limit: number;
  offset: number;
}> => {
  try {
    const queryParams = new URLSearchParams();
    
    if (params?.limit) queryParams.append('limit', params.limit.toString());
    if (params?.offset) queryParams.append('offset', params.offset.toString());
    if (params?.user_id) queryParams.append('user_id', params.user_id);
    
    if (params?.filters) {
      Object.entries(params.filters).forEach(([key, value]) => {
        if (value !== undefined && value !== null) {
          if (key === 'date_range' && value.start && value.end) {
            queryParams.append('date_start', value.start);
            queryParams.append('date_end', value.end);
          } else {
            queryParams.append(key, String(value));
          }
        }
      });
    }

    const url = `${API_BASE}/insights${queryParams.toString() ? `?${queryParams.toString()}` : ''}`;
    
    const response = await fetch(url, {
      method: 'GET',
      credentials: 'include',
    });
    
    const data = await handleResponse<{
      insights: HeuristicsInsight[];
      total: number;
      limit: number;
      offset: number;
    }>(response);
    
    return {
      insights: data.insights || [],
      total: data.total || 0,
      limit: data.limit || 20,
      offset: data.offset || 0
    };
  } catch (err) {
    if (err instanceof ApplicationError) {
      throw err;
    }
    const appError = parseErrorResponse(err);
    throw new ApplicationError(
      HeuristicsErrorCode.INSIGHT_GENERATION_FAILED,
      'インサイトの取得に失敗しました',
      appError.message
    );
  }
};

export const getInsightById = async (id: string): Promise<HeuristicsInsight> => {
  try {
    const response = await fetch(`${API_BASE}/insights/${id}`, {
      method: 'GET',
      credentials: 'include',
    });
    
    const data = await handleResponse<{ insight: HeuristicsInsight }>(response);
    return data.insight;
  } catch (err) {
    if (err instanceof ApplicationError) {
      throw err;
    }
    const appError = parseErrorResponse(err);
    throw new ApplicationError(
      HeuristicsErrorCode.INSIGHT_GENERATION_FAILED,
      'インサイトの取得に失敗しました',
      appError.message
    );
  }
};

// パターン検出関連のAPI
export const detectPatterns = async (params?: {
  user_id?: string;
  data_type?: string;
  period?: string;
  filters?: PatternFilters;
}): Promise<{
  patterns: HeuristicsPattern[];
  metadata: {
    user_id: string;
    data_type: string;
    period: string;
  };
}> => {
  try {
    const queryParams = new URLSearchParams();
    
    if (params?.user_id) queryParams.append('user_id', params.user_id);
    if (params?.data_type) queryParams.append('data_type', params.data_type);
    if (params?.period) queryParams.append('period', params.period);
    
    if (params?.filters) {
      Object.entries(params.filters).forEach(([key, value]) => {
        if (value !== undefined && value !== null) {
          queryParams.append(key, String(value));
        }
      });
    }

    const url = `${API_BASE}/patterns${queryParams.toString() ? `?${queryParams.toString()}` : ''}`;
 
    const response = await fetch(url, {
      method: 'GET',
      credentials: 'include',
    });

    const data = await handleResponse<{
      patterns: HeuristicsPattern[];
      metadata: {
        user_id: string;
        data_type: string;
        period: string;
      };
    }>(response);
    
    return {
      patterns: data.patterns || [],
      metadata: data.metadata || {
        user_id: params?.user_id || '',
        data_type: params?.data_type || 'all',
        period: params?.period || 'week'
      }
    };
  } catch (err) {
    if (err instanceof ApplicationError) {
      throw err;
    }
    const appError = parseErrorResponse(err);
    throw new ApplicationError(
      HeuristicsErrorCode.PATTERN_DETECTION_FAILED,
      'パターン検出に失敗しました',
      appError.message
    );
  }
};

// モデルトレーニング関連のAPI
export const trainModel = async (request: HeuristicsTrainRequest): Promise<HeuristicsModel> => {
  try {
    // バリデーション
    const validation = validateTrainRequest(request);
    if (!validation.isValid) {
      throw new ApplicationError(
        HeuristicsErrorCode.MODEL_INVALID_TYPE,
        validation.errors.join(', ')
      );
    }

    const response = await fetch(`${API_BASE}/patterns/train`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(request),
      credentials: 'include',
    });
    
    const data = await handleResponse<{ model: HeuristicsModel }>(response);
    return data.model;
  } catch (err) {
    if (err instanceof ApplicationError) {
      throw err;
    }
    const appError = parseErrorResponse(err);
    throw new ApplicationError(
      HeuristicsErrorCode.MODEL_TRAINING_FAILED,
      'モデル学習に失敗しました',
      appError.message
    );
  }
};

// ユーティリティ関数
export const getAnalysisTypeLabel = (type: string): string => {
  const labels: Record<string, string> = {
    performance: 'パフォーマンス',
    behavior: '行動',
    pattern: 'パターン',
    cognitive: '認知的',
    efficiency: '効率性'
  };
  return labels[type] || type;
};

export const getStatusColor = (status: string): string => {
  const colors: Record<string, string> = {
    pending: '#FF9800',
    completed: '#4CAF50',
    failed: '#F44336',
    training: '#2196F3',
    ready: '#4CAF50',
    deprecated: '#9E9E9E'
  };
  return colors[status] || '#9E9E9E';
};