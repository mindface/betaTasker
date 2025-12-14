import { AddProcessOptimization, ProcessOptimization } from "../model/processOptimization";
import { fetchApiJsonCore } from "@/utils/fetchApi";

export const fetchProcessOptimizationsClient = async () => {
  const data = await fetchApiJsonCore<undefined,ProcessOptimization[]>({
    endpoint: '/api/processOptimization',
    method: 'GET',
    errorMessage: 'error fetchProcessOptimizationsClient プロセス最適化一覧取得失敗',
    getKey: 'process_optimizations',
  });
  return data;
};

export const addProcessOptimizationClient = async (processOptimization: AddProcessOptimization) => {
  const data = await fetchApiJsonCore<AddProcessOptimization,ProcessOptimization>({
    endpoint: '/api/processOptimization',
    method: 'POST',
    body: processOptimization,
    errorMessage: 'error addProcessOptimizationClient プロセス最適化追加失敗',
    getKey: 'process_optimization',
  });
  return data;
};

export const updateProcessOptimizationClient = async (processOptimization: ProcessOptimization) => {
  const data = await fetchApiJsonCore<ProcessOptimization,ProcessOptimization>({
    endpoint: '/api/processOptimization',
    method: 'PUT',
    body: processOptimization,
    errorMessage: 'error updateProcessOptimizationClient プロセス最適化更新失敗',
    getKey: 'process_optimization',
  });
  return data;
};

export const deleteProcessOptimizationClient = async (id: string) => {
  const data = await fetchApiJsonCore<{id: string},ProcessOptimization>({
    endpoint: `/api/processOptimization`,
    method: 'DELETE',
    body: { id },
    errorMessage: 'error deleteProcessOptimizationClient プロセス最適化削除失敗',
    getKey: 'process_optimization',
  });
  return data;
};
