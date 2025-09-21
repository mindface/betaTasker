import { AddLanguageOptimization, LanguageOptimization } from "../model/languageOptimization";

export const fetchLanguageOptimizationsService = async () => {
  try {
    const res = await fetch('/api/languageOptimization', {
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

export const addLanguageOptimizationService = async (languageOptimization: AddLanguageOptimization) => {
  try {
    const res = await fetch('/api/languageOptimization', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(languageOptimization),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('プロセス最適化追加失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const updateLanguageOptimizationService = async (languageOptimization: LanguageOptimization) => {
  try {
    const res = await fetch('/api/languageOptimization', {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(languageOptimization),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('プロセス最適化更新失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const deleteLanguageOptimizationService = async (id: string) => {
  try {
    const res = await fetch(`/api/languageOptimization`, {
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
