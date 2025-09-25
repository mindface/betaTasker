// import { LearningData } from '../model/learning';
import { fetchApiJsonCore } from "@/utils/fetchApi";

// APIからlearningDataを取得する関数
export async function fetchLearningData() {
  try {
    // TODO 調整が必要
    const data = await fetchApiJsonCore<undefined,undefined>({
      endpoint: '/api/learning',
      method: 'GET',
      errorMessage: 'error addAssessmentService アセスメント追加失敗',
    });
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
}
