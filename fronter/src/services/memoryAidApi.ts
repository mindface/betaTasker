// /src/services/memoryAidApi.ts
import { MemoryContext } from "../model/memoryAid";

export const fetchMemoryAidsByCode = async (code: string) => {
  try {
    const res = await fetch(`/api/memoryAid?code=${code}`, {
      method: 'GET',
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('メモリー支援データ取得失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};
