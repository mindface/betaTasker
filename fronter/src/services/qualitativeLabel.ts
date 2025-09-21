import { AddQualitativeLabel, QualitativeLabel } from "../model/qualitativeLabel";

export const fetchQualitativeLabelsService = async () => {
  try {
    const res = await fetch('/api/qualitativeLabel', {
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

export const addQualitativeLabelService = async (qualitativeLabel: AddQualitativeLabel) => {
  try {
    const res = await fetch('/api/qualitativeLabel', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(qualitativeLabel),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('プロセス最適化追加失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const updateQualitativeLabelService = async (qualitativeLabel: QualitativeLabel) => {
  try {
    const res = await fetch('/api/qualitativeLabel', {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(qualitativeLabel),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('プロセス最適化更新失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const deleteQualitativeLabelService = async (id: string) => {
  try {
    const res = await fetch(`/api/qualitativeLabel`, {
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
