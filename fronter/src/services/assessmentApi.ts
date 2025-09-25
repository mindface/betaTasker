import { AddAssessment, Assessment } from "../model/assessment";
import { fetchApiJsonCore } from "@/utils/fetchApi";

export const fetchAssessmentsService = async () => {
  try {
    const data = await fetchApiJsonCore<undefined,Assessment[]>({
      endpoint: '/api/assessment',
      method: 'GET',
      errorMessage: 'error fetchAssessmentsService アセスメント一覧取得失敗',
    });
    if ('error' in data) {
      return data;
    }
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

// Todoデータ内容を確認
export const getAssessmentsForTaskUserService = async (userId: number,taskId:number) => {
  try {
    const data = await fetchApiJsonCore<{userId: number, taskId: number},Assessment[]>({
      endpoint: '/api/assessmentsForTaskUser',
      method: 'POST',
      body: ({ userId, taskId }),
      errorMessage: 'error getAssessmentsForTaskUserService アセスメントIdでの情報取得失敗',
    });
    if ('error' in data) {
      return data;
    }
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const addAssessmentService = async (assessment: AddAssessment) => {
  try {
    const data = await fetchApiJsonCore<AddAssessment,Assessment>({
      endpoint: '/api/assessment',
      method: 'POST',
      body: assessment,
      errorMessage: 'error addAssessmentService アセスメント追加失敗',
    });
    return data
  } catch (err: any) {
    return { error: err.message };
  }
};

export const updateAssessmentService = async (assessment: Assessment) => {
  try {
    const data = await fetchApiJsonCore<Assessment,Assessment>({
      endpoint: '/api/assessment',
      method: 'PUT',
      body: assessment,
      errorMessage: 'error updateAssessmentService アセスメント更新失敗',
    });
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const deleteAssessmentService = async (id: string) => {
  try {
    const data = await fetchApiJsonCore<{id: string},undefined>({
      endpoint: '/api/assessment',
      method: 'PUT',
      body: { id },
      errorMessage: 'error updateAssessmentService アセスメント更新失敗',
    });
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};
