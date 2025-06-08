
export const fetchMemories = async () => {
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

export const addMemory = async (memory: any) => {
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

export const deleteMemory = async (id: string) => {
  try {
    const res = await fetch(`/api/memory/${id}`, {
      method: 'DELETE',
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('メモリ削除失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};