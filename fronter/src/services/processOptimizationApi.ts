import { AddProcessOptimization, ProcessOptimization } from "../model/processOptimization";

export const fetchProcessOptimizationsService = async () => {
  try {
    const res = await fetch('/api/processOptimization', {
      method: 'GET',
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('プロセス最適化一覧取得失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const addProcessOptimizationService = async (processOptimization: AddProcessOptimization) => {
  try {
    const res = await fetch('/api/processOptimization', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(processOptimization),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('プロセス最適化追加失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const updateProcessOptimizationService = async (processOptimization: ProcessOptimization) => {
  try {
    const res = await fetch('/api/processOptimization', {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(processOptimization),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('プロセス最適化更新失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const deleteProcessOptimizationService = async (id: string) => {
  try {
    const res = await fetch(`/api/processOptimization`, {
      method: 'DELETE',
      credentials: 'include',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ id }),
    });
    const data = await res.json();
    if (!res.ok) throw new Error('プロセス最適化削除失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};
