import { AddMemory, Memory } from "../model/memory";
import { fetchApiJsonCore } from "@/utils/fetchApi";

export const fetchMemoriesService = async () => {
  const data = await fetchApiJsonCore<undefined,Memory[]>({
    endpoint: '/api/memory',
    method: 'GET',
    errorMessage: 'error fetchMemoriesService メモリ一覧取得失敗',
  });
  return data;
};

export const fetchMemoryService = async (memoryId: number) => {
  const data = await fetchApiJsonCore<undefined,Memory>({
    endpoint: `/api/memory/${memoryId}`,
    method: 'GET',
    errorMessage: 'error fetchMemoryService メモリ情報取得失敗',
  });
  return data;
};

export const addMemoryService = async (memory: AddMemory) => {
  const data = await fetchApiJsonCore<AddMemory,Memory>({
    endpoint: '/api/memory',
    method: 'POST',
    body: memory,
    errorMessage: 'error addMemoryService メモリ情報追加失敗',
  });
  return data;
};

export const updateMemoryService = async (memory: Memory) => {
  const data = await fetchApiJsonCore<Memory,Memory>({
    endpoint: '/api/memory',
    method: 'PUT',
    body: memory,
    errorMessage: 'error updateMemoryService メモリ追加失敗',
  });
  return data;
};

export const deleteMemoryService = async (id: string) => {
  const data = await fetchApiJsonCore<{ id: string },Memory>({
    endpoint: '/api/memory',
    method: 'DELETE',
    body: {id},
    errorMessage: 'error deleteMemoryService メモリ削除失敗',
  });
  return data;
};