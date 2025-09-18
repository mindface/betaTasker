import { AddKnowledgePattern, KnowledgePattern } from "../model/knowledgePattern";

export const fetchKnowledgePatternsService = async () => {
  try {
    const res = await fetch('/api/knowledgePatterns', {
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

export const addKnowledgePatternService = async (knowledgePattern: AddKnowledgePattern) => {
  try {
    const res = await fetch('/api/knowledgePatterns', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(knowledgePattern),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('プロセス最適化追加失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const updateKnowledgePatternService = async (knowledgePattern: KnowledgePattern) => {
  try {
    const res = await fetch('/api/knowledgePatterns', {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(knowledgePattern),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('プロセス最適化更新失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const deleteKnowledgePatternService = async (id: string) => {
  try {
    const res = await fetch(`/api/knowledgePatterns`, {
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
