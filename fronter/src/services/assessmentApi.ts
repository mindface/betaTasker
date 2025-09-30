import { error } from "console";
import { AddAssessment, Assessment } from "../model/assessment";
import { fetchApiJsonCore } from "@/utils/fetchApi";

export const fetchAssessmentsService = async () => {
  const data = await fetchApiJsonCore<undefined,Assessment[]>({
    endpoint: '/api/assessment',
    method: 'GET',
    errorMessage: 'error fetchAssessmentsService アセスメント一覧取得失敗',
  });
  return data;
};

// Todoデータ内容を確認
export const getAssessmentsForTaskUserService = async (userId: number,taskId:number) => {
  const data = await fetchApiJsonCore<{userId: number, taskId: number},Assessment[]>({
    endpoint: '/api/assessmentsForTaskUser',
    method: 'POST',
    body: ({ userId, taskId }),
    errorMessage: 'error getAssessmentsForTaskUserService アセスメントIdでの情報取得失敗',
  });
  if ('error' in data) {
    return data;
  }
  return data.value;
};

export const addAssessmentService = async (assessment: AddAssessment) => {
  const data = await fetchApiJsonCore<AddAssessment,Assessment>({
    endpoint: '/api/assessment',
    method: 'POST',
    body: assessment,
    errorMessage: 'error addAssessmentService アセスメント追加失敗',
  });
  if ('error' in data) {
    return data;
  }
  return data.value;
};

export const updateAssessmentService = async (assessment: Assessment) => {
  const data = await fetchApiJsonCore<Assessment,Assessment>({
    endpoint: '/api/assessment',
    method: 'PUT',
    body: assessment,
    errorMessage: 'error updateAssessmentService アセスメント更新失敗',
  });
  if ('error' in data) {
    return data;
  }
  return data.value;
};

export const deleteAssessmentService = async (id: string) => {
  const data = await fetchApiJsonCore<{id: string},{ id: string }>({
    endpoint: '/api/assessment',
    method: 'PUT',
    body: { id },
    errorMessage: 'error updateAssessmentService アセスメント更新失敗',
  });
  if ('error' in data) {
    return data;
  }
  return data.value;
};
