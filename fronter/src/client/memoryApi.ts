import { AddMemory, Memory } from "@/model/memory";
import { LimitResponse } from "@/model/respose";
import { fetchApiJsonCore } from "@/utils/fetchApi";

export const fetchMemoriesClient = async () => {
  const data = await fetchApiJsonCore<undefined, LimitResponse<Memory,"memories">>({
    endpoint: "/api/memory",
    method: "GET",
    errorMessage: "error fetchMemoriesClient メモリ一覧取得失敗",
  });
  if ("error" in data) {
    return data;
  }
  return data.value;
};

export const getMemoriesLimitClient = async (
  page: number,
  limit: number,
) => {
  const data = await fetchApiJsonCore<
    undefined,
     LimitResponse<Memory,"memories">
  >({
    endpoint: `/api/memory/pager?page=${page}&limit=${limit}`,
    method: "GET",
    errorMessage:
      "error getMemoriesLimitClient メモリLimiterでの情報取得失敗",
  });
  if ("error" in data) {
    return data;
  }
  return data.value;
};

export const getMemoryClient = async (memoryId: number) => {
  const data = await fetchApiJsonCore<undefined, Memory>({
    endpoint: `/api/memory/${memoryId}`,
    method: "GET",
    errorMessage: "error getMemoryClient メモリ情報取得失敗",
  });
  return data;
};

export const addMemoryClient = async (memory: AddMemory) => {
  const data = await fetchApiJsonCore<AddMemory, Memory>({
    endpoint: "/api/memory",
    method: "POST",
    body: memory,
    errorMessage: "error addMemoryClient メモリ情報追加失敗",
  });
  return data;
};

export const updateMemoryClient = async (memory: Memory) => {
  const data = await fetchApiJsonCore<Memory, Memory>({
    endpoint: "/api/memory",
    method: "PUT",
    body: memory,
    errorMessage: "error updateMemoryClient メモリ追加失敗",
  });
  return data;
};

export const deleteMemoryClient = async (id: string) => {
  const data = await fetchApiJsonCore<{ id: string }, Memory>({
    endpoint: "/api/memory",
    method: "DELETE",
    body: { id },
    errorMessage: "error deleteMemoryClient メモリ削除失敗",
  });
  return data;
};
