import { AddProcessOptimization, ProcessOptimization } from "../model/processOptimization";
import { fetchApiJsonCore } from "@/utils/fetchApi";

export const fetchProcessOptimizationsService = async () => {
  try {
    const data = await fetchApiJsonCore<undefined,ProcessOptimization[]>({
      endpoint: '/api/processOptimization',
      method: 'GET',
      errorMessage: 'error fetchProcessOptimizationsService プロセス最適化一覧取得失敗',
    });
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const addProcessOptimizationService = async (processOptimization: AddProcessOptimization) => {
  try {
    const data = await fetchApiJsonCore<AddProcessOptimization,ProcessOptimization>({
      endpoint: '/api/processOptimization',
      method: 'POST',
      body: processOptimization,
      errorMessage: 'error addProcessOptimizationService プロセス最適化追加失敗',
    });
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const updateProcessOptimizationService = async (processOptimization: ProcessOptimization) => {
  try {
    const data = await fetchApiJsonCore<ProcessOptimization,ProcessOptimization>({
      endpoint: '/api/processOptimization',
      method: 'PUT',
      body: processOptimization,
      errorMessage: 'error updateProcessOptimizationService プロセス最適化更新失敗',
    });
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const deleteProcessOptimizationService = async (id: string) => {
  try {
    const data = await fetchApiJsonCore<{id: string},ProcessOptimization>({
      endpoint: `/api/processOptimization`,
      method: 'DELETE',
      body: { id },
      errorMessage: 'error deleteProcessOptimizationService プロセス最適化削除失敗',
    });
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};
