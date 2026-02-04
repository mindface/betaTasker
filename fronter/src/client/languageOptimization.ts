import { fetchApiJsonCore } from "@/utils/fetchApi";
import {
  AddLanguageOptimization,
  LanguageOptimization,
} from "../model/languageOptimization";

export const fetchLanguageOptimizationsClient = async () => {
  const data = await fetchApiJsonCore<undefined, LanguageOptimization[]>({
    endpoint: "/api/languageOptimization",
    method: "GET",
    errorMessage:
      "error fetchLanguageOptimizationsClient 言語最適化データ一覧取得失敗",
    getKey: "language_optimizations",
  });
  return data;
};

export const addLanguageOptimizationClient = async (
  languageOptimization: AddLanguageOptimization,
) => {
  const data = await fetchApiJsonCore<
    AddLanguageOptimization,
    LanguageOptimization
  >({
    endpoint: "/api/languageOptimization",
    method: "POST",
    body: languageOptimization,
    errorMessage:
      "error addLanguageOptimizationClient 言語最適化データ追加失敗",
    getKey: "language_optimization",
  });
  return data;
};

export const updateLanguageOptimizationClient = async (
  languageOptimization: LanguageOptimization,
) => {
  const data = await fetchApiJsonCore<
    LanguageOptimization,
    LanguageOptimization
  >({
    endpoint: "/api/languageOptimization",
    method: "PUT",
    body: languageOptimization,
    errorMessage:
      "error updateLanguageOptimizationClient 言語最適化データ更新失敗",
    getKey: "language_optimization",
  });
  return data;
};

export const deleteLanguageOptimizationClient = async (id: string) => {
  const data = await fetchApiJsonCore<{ id: string }, undefined>({
    endpoint: `/api/languageOptimization`,
    method: "DELETE",
    body: { id },
    errorMessage:
      "error deleteLanguageOptimizationClient 言語最適化データ削除失敗",
    getKey: "language_optimization",
  });
  return data;
};
