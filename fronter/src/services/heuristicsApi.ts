import { AddHeuristicsAnalysis, HeuristicsAnalysis, AddHeuristicsTracking, HeuristicsTracking, AddHeuristicsInsight, HeuristicsInsight, AddHeuristicsPattern, HeuristicsPattern } from "../model/heuristics";

export const fetchAnalysesService = async () => {
  try {
    const res = await fetch('/api/heuristics/analysis', {
      method: 'GET',
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('分析結果一覧取得失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const getAnalysisForTaskUserService = async (userId: number, taskId: number) => {
  try {
    const res = await fetch('/api/heuristics/analysisForTaskUser', {
      method: 'POST',
      credentials: 'include',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ userId, taskId }),
    });
    const data = await res.json();
    if (!res.ok) throw new Error('分析結果一覧取得失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const addAnalysisService = async (analysis: AddHeuristicsAnalysis) => {
  try {
    const res = await fetch('/api/heuristics/analysis', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(analysis),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('分析結果追加失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const updateAnalysisService = async (analysis: HeuristicsAnalysis) => {
  try {
    const res = await fetch('/api/heuristics/analysis', {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(analysis),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('分析結果更新失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const deleteAnalysisService = async (id: string) => {
  try {
    const res = await fetch('/api/heuristics/analysis', {
      method: 'DELETE',
      credentials: 'include',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ id }),
    });
    const data = await res.json();
    if (!res.ok) throw new Error('分析結果削除失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const fetchTrackingService = async () => {
  try {
    const res = await fetch('/api/heuristics/tracking', {
      method: 'GET',
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('トラッキング一覧取得失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const getTrackingForUserService = async (userId: number) => {
  try {
    const res = await fetch('/api/heuristics/trackingForUser', {
      method: 'POST',
      credentials: 'include',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ userId }),
    });
    const data = await res.json();
    if (!res.ok) throw new Error('トラッキング一覧取得失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const addTrackingService = async (tracking: AddHeuristicsTracking) => {
  try {
    const res = await fetch('/api/heuristics/tracking', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(tracking),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('トラッキング追加失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const updateTrackingService = async (tracking: HeuristicsTracking) => {
  try {
    const res = await fetch('/api/heuristics/tracking', {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(tracking),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('トラッキング更新失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const deleteTrackingService = async (id: string) => {
  try {
    const res = await fetch('/api/heuristics/tracking', {
      method: 'DELETE',
      credentials: 'include',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ id }),
    });
    const data = await res.json();
    if (!res.ok) throw new Error('トラッキング削除失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const fetchInsightsService = async () => {
  try {
    const res = await fetch('/api/heuristics/insights', {
      method: 'GET',
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('インサイト一覧取得失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const getInsightsForUserService = async (userId: number) => {
  try {
    const res = await fetch('/api/heuristics/insightsForUser', {
      method: 'POST',
      credentials: 'include',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ userId }),
    });
    const data = await res.json();
    if (!res.ok) throw new Error('インサイト一覧取得失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const addInsightService = async (insight: AddHeuristicsInsight) => {
  try {
    const res = await fetch('/api/heuristics/insights', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(insight),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('インサイト追加失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const updateInsightService = async (insight: HeuristicsInsight) => {
  try {
    const res = await fetch('/api/heuristics/insights', {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(insight),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('インサイト更新失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const deleteInsightService = async (id: string) => {
  try {
    const res = await fetch('/api/heuristics/insights', {
      method: 'DELETE',
      credentials: 'include',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ id }),
    });
    const data = await res.json();
    if (!res.ok) throw new Error('インサイト削除失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const fetchPatternsService = async () => {
  try {
    const res = await fetch('/api/heuristics/patterns', {
      method: 'GET',
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('パターン一覧取得失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const addPatternService = async (pattern: AddHeuristicsPattern) => {
  try {
    const res = await fetch('/api/heuristics/patterns', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(pattern),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('パターン追加失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const updatePatternService = async (pattern: HeuristicsPattern) => {
  try {
    const res = await fetch('/api/heuristics/patterns', {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(pattern),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('パターン更新失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const deletePatternService = async (id: string) => {
  try {
    const res = await fetch('/api/heuristics/patterns', {
      method: 'DELETE',
      credentials: 'include',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ id }),
    });
    const data = await res.json();
    if (!res.ok) throw new Error('パターン削除失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};