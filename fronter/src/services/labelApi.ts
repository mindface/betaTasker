import { 
  QuantificationLabel, 
  CreateLabelRequest, 
  UpdateLabelRequest,
  VerifyLabelRequest,
  LabelSearchQuery,
  LabelStatistics
} from '../model/quantificationLabel';

// ラベル一覧取得
export const fetchLabelsService = async () => {
  try {
    const res = await fetch('/api/heuristics/labels', {
      method: 'GET',
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('ラベル一覧取得失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

// ラベル作成
export const createLabelService = async (formData: FormData) => {
  try {
    const res = await fetch('/api/heuristics/labels', {
      method: 'POST',
      body: formData,
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('ラベル作成失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

// ラベル更新
export const updateLabelService = async (request: UpdateLabelRequest) => {
  try {
    const res = await fetch(`/api/heuristics/labels/${request.id}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(request),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('ラベル更新失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

// ラベル削除
export const deleteLabelService = async (id: string) => {
  try {
    const res = await fetch(`/api/heuristics/labels/${id}`, {
      method: 'DELETE',
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('ラベル削除失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

// ラベル検索
export const searchLabelsService = async (query: LabelSearchQuery) => {
  try {
    const queryParams = new URLSearchParams();
    
    if (query.text) queryParams.append('text', query.text);
    if (query.domain) queryParams.append('domain', query.domain);
    if (query.category) queryParams.append('category', query.category);
    if (query.minConfidence) queryParams.append('minConfidence', query.minConfidence.toString());
    if (query.verified !== undefined) queryParams.append('verified', query.verified.toString());
    if (query.limit) queryParams.append('limit', query.limit.toString());
    if (query.offset) queryParams.append('offset', query.offset.toString());
    if (query.sortBy) queryParams.append('sortBy', query.sortBy);
    if (query.sortOrder) queryParams.append('sortOrder', query.sortOrder);
    
    if (query.valueRange) {
      queryParams.append('minValue', query.valueRange.min.toString());
      queryParams.append('maxValue', query.valueRange.max.toString());
      queryParams.append('unit', query.valueRange.unit);
    }
    
    if (query.concepts) {
      queryParams.append('concepts', JSON.stringify(query.concepts));
    }
    
    if (query.dateRange) {
      queryParams.append('from', query.dateRange.from);
      queryParams.append('to', query.dateRange.to);
    }
    
    const res = await fetch(`/api/heuristics/labels/search?${queryParams}`, {
      method: 'GET',
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('ラベル検索失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

// ラベル検証
export const verifyLabelService = async (request: VerifyLabelRequest) => {
  try {
    const res = await fetch(`/api/heuristics/labels/${request.labelId}/verify`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(request),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('ラベル検証失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

// 統計取得
export const getStatisticsService = async () => {
  try {
    const res = await fetch('/api/heuristics/labels/statistics', {
      method: 'GET',
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('統計取得失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

// 定量化提案
export const suggestQuantificationService = async (params: { 
  text: string; 
  imageUrl?: string; 
  domain?: string 
}) => {
  try {
    const res = await fetch('/api/heuristics/labels/suggest', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(params),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('定量化提案失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

// 類似ラベル検索
export const findSimilarLabelsService = async (params: {
  text?: string;
  imageUrl?: string;
  value?: number;
  unit?: string;
  limit?: number;
}) => {
  try {
    const queryParams = new URLSearchParams();
    if (params.text) queryParams.append('text', params.text);
    if (params.imageUrl) queryParams.append('imageUrl', params.imageUrl);
    if (params.value) queryParams.append('value', params.value.toString());
    if (params.unit) queryParams.append('unit', params.unit);
    if (params.limit) queryParams.append('limit', params.limit.toString());

    const res = await fetch(`/api/heuristics/labels/similar?${queryParams}`, {
      method: 'GET',
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('類似ラベル検索失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

// バルク操作
export const bulkOperationService = async (params: {
  operation: 'verify' | 'delete' | 'export';
  labelIds: string[];
  options?: any;
}) => {
  try {
    const res = await fetch('/api/heuristics/labels/bulk', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(params),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('バルク操作失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

// ラベル履歴取得
export const getLabelHistoryService = async (id: string) => {
  try {
    const res = await fetch(`/api/heuristics/labels/${id}/history`, {
      method: 'GET',
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('履歴取得失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

// 概念関係マップ取得
export const getConceptRelationsService = async (conceptId: string) => {
  try {
    const res = await fetch(`/api/heuristics/concepts/${conceptId}/relations`, {
      method: 'GET',
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('概念関係取得失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

// データセット作成
export const createDatasetService = async (params: {
  name: string;
  description: string;
  labelIds: string[];
  domain: string;
}) => {
  try {
    const res = await fetch('/api/heuristics/datasets', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(params),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('データセット作成失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

// データセットエクスポート
export const exportDatasetService = async (datasetId: string, format: 'json' | 'csv' | 'xml') => {
  try {
    const res = await fetch(`/api/heuristics/datasets/${datasetId}/export?format=${format}`, {
      method: 'GET',
      credentials: 'include',
    });
    
    if (!res.ok) {
      const data = await res.json();
      throw new Error(data.error || 'エクスポート失敗');
    }
    
    // バイナリデータとして返す
    const blob = await res.blob();
    const url = URL.createObjectURL(blob);
    
    // ダウンロードリンクを作成
    const a = document.createElement('a');
    a.href = url;
    a.download = `dataset_${datasetId}.${format}`;
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
    URL.revokeObjectURL(url);
    
    return { success: true };
  } catch (err: any) {
    return { error: err.message };
  }
};

// 画像アノテーション保存
export const saveImageAnnotationsService = async (params: {
  labelId: string;
  annotations: any[];
}) => {
  try {
    const res = await fetch(`/api/heuristics/labels/${params.labelId}/annotations`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ annotations: params.annotations }),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('アノテーション保存失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

// 概念階層取得
export const getConceptHierarchyService = async (domain?: string) => {
  try {
    const queryParams = domain ? `?domain=${domain}` : '';
    const res = await fetch(`/api/heuristics/concepts/hierarchy${queryParams}`, {
      method: 'GET',
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('概念階層取得失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};