import { AddTeachingFreeControl, TeachingFreeControl } from "../model/teachingFreeControl";

export const fetchTeachingFreeControlService = async () => {
  try {
    const res = await fetch('/api/teachingFreeControl', {
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

export const addTeachingFreeControlService = async (teachingFreeControl: AddTeachingFreeControl) => {
  try {
    const res = await fetch('/api/teachingFreeControl', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(teachingFreeControl),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('プロセス最適化追加失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const updateTeachingFreeControlService = async (teachingFreeControl: TeachingFreeControl) => {
  try {
    const res = await fetch('/api/teachingFreeControl', {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(teachingFreeControl),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('プロセス最適化更新失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const deleteTeachingFreeControlService = async (id: string) => {
  try {
    const res = await fetch(`/api/teachingFreeControl`, {
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
