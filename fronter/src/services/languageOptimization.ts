import { fetchApiJsonCore } from "@/utils/fetchApi";
import { AddLanguageOptimization, LanguageOptimization } from "../model/languageOptimization";

export const fetchLanguageOptimizationsService = async () => {
  const data = await fetchApiJsonCore<undefined,LanguageOptimization[]>({
    endpoint: '/api/languageOptimization',
    method: 'GET',
    errorMessage: 'error fetchLanguageOptimizationsService 言語最適化データ一覧取得失敗',
  });
  return data;
};

export const addLanguageOptimizationService = async (languageOptimization: AddLanguageOptimization) => {
  const data = await fetchApiJsonCore<AddLanguageOptimization,LanguageOptimization>({
    endpoint: '/api/languageOptimization',
    method: 'POST',
    body: languageOptimization,
    errorMessage: 'error addLanguageOptimizationService 言語最適化データ追加失敗',
  });
  return data;
};

export const updateLanguageOptimizationService = async (languageOptimization: LanguageOptimization) => {
  const data = await fetchApiJsonCore<LanguageOptimization,LanguageOptimization>({
    endpoint: '/api/languageOptimization',
    method: 'PUT',
    body: languageOptimization,
    errorMessage: 'error updateLanguageOptimizationService 言語最適化データ更新失敗',
  });
  return data;
};

export const deleteLanguageOptimizationService = async (id: string) => {
  const data = await fetchApiJsonCore<{id:string},undefined>({
    endpoint: `/api/languageOptimization`,
    method: 'DELETE',
    body: ({ id }),
    errorMessage: 'error deleteLanguageOptimizationService 言語最適化データ削除失敗',
  });
  return data;
};
