import {
  AddKnowledgePattern,
  KnowledgePattern,
} from "../model/knowledgePattern";
import { fetchApiJsonCore } from "@/utils/fetchApi";

export const fetchKnowledgePatternsClient = async () => {
  const data = await fetchApiJsonCore<undefined, KnowledgePattern[]>({
    endpoint: "/api/knowledgePattern",
    method: "GET",
    errorMessage:
      "error fetchKnowledgePatternsClient プロセス最適化一覧取得失敗",
    getKey: "knowledge_patterns",
  });
  return data;
};

export const addKnowledgePatternClient = async (
  knowledgePattern: AddKnowledgePattern,
) => {
  const data = await fetchApiJsonCore<AddKnowledgePattern, KnowledgePattern>({
    endpoint: "/api/knowledgePattern",
    method: "POST",
    body: knowledgePattern,
    errorMessage: "error addKnowledgePatternClient アセスメント一覧取得失敗",
    getKey: "knowledge_pattern",
  });
  return data;
};

export const updateKnowledgePatternClient = async (
  knowledgePattern: KnowledgePattern,
) => {
  const data = await fetchApiJsonCore<KnowledgePattern, KnowledgePattern>({
    endpoint: "/api/knowledgePattern",
    method: "PUT",
    body: knowledgePattern,
    errorMessage: "error updateKnowledgePatternClient プロセス最適化更新失敗",
    getKey: "knowledge_pattern",
  });
  return data;
};

export const deleteKnowledgePatternClient = async (id: string) => {
  const data = await fetchApiJsonCore<{ id: string }, undefined>({
    endpoint: "/api/knowledgePattern",
    method: "DELETE",
    body: { id },
    errorMessage: "error deleteKnowledgePatternClient プロセス最適化削除失敗",
    getKey: "knowledge_pattern",
  });
  return data;
};
