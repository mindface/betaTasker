import { LearningData } from "../model/learning";
import { fetchApiJsonCore } from "@/utils/fetchApi";

// APIからlearningDataを取得する関数
export async function fetchLearningData() {
  // TODO 調整が必要
  const data = await fetchApiJsonCore<undefined, LearningData>({
    endpoint: "/api/learning",
    method: "GET",
    errorMessage: "error fetchLearningData アセスメント追加失敗",
  });

  return data;
}
