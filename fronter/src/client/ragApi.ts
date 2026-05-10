import type {
  OntologySearchRequest,
  OntologySearchResponse,
} from "@/model/rag";
import { fetchApiJsonCore } from "@/utils/fetchApi";

export const rankOntologySearchClient = async (
  request: OntologySearchRequest,
) => {
  const data = await fetchApiJsonCore<
    OntologySearchRequest,
    OntologySearchResponse
  >({
    endpoint: "/api/rag/ontology/rank",
    method: "POST",
    body: request,
    errorMessage: "error rankOntologySearchClient RAG検索ランキング取得失敗",
  });

  if ("error" in data) {
    return data;
  }

  return data.value;
};
