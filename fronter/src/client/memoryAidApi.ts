// /src/services/memoryAidApi.ts
import { MemoryContext } from "../model/memoryAid";
import { fetchApiJsonCore } from "@/utils/fetchApi";

export const fetchMemoryAidsByCode = async (code: string) => {
  return [];
  const data = await fetchApiJsonCore<undefined,MemoryContext[]>({
    endpoint: `/api/memoryAid?code=${code}`,
    method: 'GET',
    errorMessage: 'error fetchMemoryAidsByCode メモリー支援データ取得失敗',
    getKey: 'contexts',
  });
  return data
};
