import { AddMemory, Memory } from "../model/memory";
import { fetchApiJsonCore } from "@/utils/fetchApi";

export const fetchMemoriesService = async () => {
  try {
    const data = await fetchApiJsonCore<undefined,Memory[]>({
      endpoint: '/api/memory',
      method: 'GET',
      errorMessage: 'error fetchMemoriesService メモリ一覧取得失敗',
    });
    if('error' in data) {
      return data.error
    }
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const fetchMemoryService = async (memoryId: number) => {
  try {
    const data = await fetchApiJsonCore<undefined,Memory>({
      endpoint: `/api/memory/${memoryId}`,
      method: 'GET',
      errorMessage: 'error fetchMemoriesService メモリ情報取得失敗',
    });
    if('error' in data) {
      return data.error
    }
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const addMemoryService = async (memory: AddMemory) => {
  try {
    const data = await fetchApiJsonCore<AddMemory,Memory>({
      endpoint: '/api/memory',
      method: 'POST',
      body: memory,
      errorMessage: 'error addMemoryService メモリ情報追加失敗',
    });
    if('error' in data) {
      return data.error
    }
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const updateMemoryService = async (memory: Memory) => {
  try {
    const data = await fetchApiJsonCore<Memory,Memory>({
      endpoint: '/api/memory',
      method: 'PUT',
      body: memory,
      errorMessage: 'error updateMemoryService メモリ追加失敗',
    });
    if('error' in data) {
      return data.error
    }
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const deleteMemoryService = async (id: string) => {
  try {
    const data = await fetchApiJsonCore<{ id: string },Memory>({
      endpoint: '/api/memory',
      method: 'DELETE',
      body: {id},
      errorMessage: 'error deleteMemoryService メモリ削除失敗',
    });
    if('error' in data) {
      return data.error
    }
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};