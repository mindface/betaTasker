import { AddKnowledgePattern, KnowledgePattern } from "../model/knowledgePattern";
import { fetchApiJsonCore } from "@/utils/fetchApi";

export const fetchKnowledgePatternsService = async () => {
  const data = await fetchApiJsonCore<undefined,KnowledgePattern[]>({
    endpoint: '/api/knowledgePattern',
    method: 'GET',
    errorMessage: 'error fetchKnowledgePatternsService プロセス最適化一覧取得失敗',
  });
  return data;
};

export const addKnowledgePatternService = async (knowledgePattern: AddKnowledgePattern) => {
  const data = await fetchApiJsonCore<AddKnowledgePattern,KnowledgePattern>({
    endpoint: '/api/knowledgePattern',
    method: 'POST',
    body: knowledgePattern,
    errorMessage: 'error addKnowledgePatternService アセスメント一覧取得失敗',
  });
  return data;
};

export const updateKnowledgePatternService = async (knowledgePattern: KnowledgePattern) => {
  const data = await fetchApiJsonCore<KnowledgePattern,KnowledgePattern>({
    endpoint: '/api/knowledgePattern',
    method: 'PUT',
    body: knowledgePattern,
    errorMessage: 'error updateKnowledgePatternService プロセス最適化更新失敗',
  });
  return data
};

export const deleteKnowledgePatternService = async (id: string) => {
  const data = await fetchApiJsonCore<{id:string},undefined>({
    endpoint: '/api/knowledgePattern',
    method: 'DELETE',
    body: { id },
    errorMessage: 'error deleteKnowledgePatternService プロセス最適化削除失敗',
  });
  return data;
};
