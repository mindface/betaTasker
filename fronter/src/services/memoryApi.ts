import { AddMemory, Memory } from "../model/memory";

export const fetchMemoriesService = async () => {
  try {
    const res = await fetch('/api/memory', {
      method: 'GET',
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('メモリ一覧取得失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const fetchMemoryService = async (memoryId: number) => {
  try {
    const res = await fetch(`/api/memory/${memoryId}`, {
      method: 'GET',
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('メモリ一覧取得失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const addMemoryService = async (memory: AddMemory) => {
  try {
    const res = await fetch('/api/memory', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(memory),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('メモリ追加失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const updateMemoryService = async (memory: Memory) => {
  try {
    const res = await fetch('/api/memory', {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(memory),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('メモリ追加失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const deleteMemoryService = async (id: string) => {
  try {
    const res = await fetch(`/api/memory`, {
      method: 'DELETE',
      credentials: 'include',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ id }),
    });
    const data = await res.json();
    if (!res.ok) throw new Error('メモリ削除失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};