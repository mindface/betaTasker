import { fetchApiJsonCore } from "@/utils/fetchApi";
import { AddLanguageOptimization, LanguageOptimization } from "../model/languageOptimization";

export const fetchLanguageOptimizationsService = async () => {
  try {
    const data = await fetchApiJsonCore<undefined,LanguageOptimization>({
      endpoint: '/api/languageOptimization',
      method: 'GET',
      errorMessage: 'error fetchLanguageOptimizationsService 言語最適化データ一覧取得失敗',
    });
    if ('error' in data) {
      return data;
    }
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const addLanguageOptimizationService = async (languageOptimization: AddLanguageOptimization) => {
  try {
    const data = await fetchApiJsonCore<AddLanguageOptimization,LanguageOptimization>({
      endpoint: '/api/languageOptimization',
      method: 'POST',
      body: languageOptimization,
      errorMessage: 'error addLanguageOptimizationService 言語最適化データ追加失敗',
    });
    if ('error' in data) {
      return data;
    }
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const updateLanguageOptimizationService = async (languageOptimization: LanguageOptimization) => {
  try {
    const data = await fetchApiJsonCore<LanguageOptimization,LanguageOptimization>({
      endpoint: '/api/languageOptimization',
      method: 'PUT',
      body: languageOptimization,
      errorMessage: 'error updateLanguageOptimizationService 言語最適化データ更新失敗',
    });
    if ('error' in data) {
      return data;
    }
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const deleteLanguageOptimizationService = async (id: string) => {
  try {
    const data = await fetchApiJsonCore<{id:string},undefined>({
      endpoint: `/api/languageOptimization`,
      method: 'DELETE',
      body: ({ id }),
      errorMessage: 'error deleteLanguageOptimizationService 言語最適化データ削除失敗',
    });
    if ('error' in data) {
      return data;
    }
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};
