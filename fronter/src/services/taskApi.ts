import { AddTask, Task } from "../model/task";
import { ApplicationError, ErrorCode, parseErrorResponse } from "../errors/errorCodes";
import { fetchApiJsonCore } from "@/utils/fetchApi";

interface ErrorResponse {
  code: ErrorCode;
  message: string;
  detail?: string;
}

export type SuccessResponse<T = any, K extends string = "data"> = {
  success: boolean;
  message?: string;
} & {
  [key in K]?: T;
};

// Todo どこかで参考に使う
const handleApiError = async (response: Response): Promise<never> => {
  let errorData: ErrorResponse;

  try {
    errorData = await response.json();
  } catch {
    // JSONパースできない場合はHTTPステータスコードベースでエラーを作成
    switch (response.status) {
      case 400:
        throw new ApplicationError(ErrorCode.VAL_INVALID_INPUT, 'リクエストが無効です');
      case 401:
        throw new ApplicationError(ErrorCode.AUTH_INVALID_CREDENTIALS, '認証が必要です');
      case 403:
        throw new ApplicationError(ErrorCode.AUTH_UNAUTHORIZED, 'アクセス権限がありません');
      case 404:
        throw new ApplicationError(ErrorCode.RES_NOT_FOUND, 'リソースが見つかりません');
      case 500:
        throw new ApplicationError(ErrorCode.SYS_INTERNAL_ERROR, 'サーバーエラーが発生しました');
      default:
        throw new ApplicationError(ErrorCode.SYS_INTERNAL_ERROR, `HTTPエラー: ${response.status}`);
    }
  }

  // サーバーからエラーコードが返ってきた場合
  if (errorData.code) {
    throw new ApplicationError(errorData.code, errorData.message, errorData.detail);
  }

  throw new ApplicationError(ErrorCode.SYS_INTERNAL_ERROR, '予期しないエラーが発生しました');
};

export const fetchTasksService = async () => {
  const data = await fetchApiJsonCore<undefined,Task[]>({
    endpoint: '/api/task',
    method: 'GET',
    errorMessage: 'error fetchTasksService タスク一覧取得失敗',
    getKey: 'tasks',
  });

  return data;
};

export const addTaskService = async (task: AddTask) => {
  const data = await fetchApiJsonCore<AddTask,Task>({
    endpoint: '/api/task',
    method: 'POST',
    body: task,
    errorMessage: 'error addTaskService タスク追加失敗',
    getKey: 'task',
  });
  return data;
};

export const updateTaskService = async (task: Task) => {
  const data = await fetchApiJsonCore<Task,Task>({
    endpoint: '/api/task', 
    method: 'PUT',
    body: task,
    errorMessage: 'error updateTaskService タスク更新失敗',
    getKey: 'task',
  });
  return data;
};

export const deleteTaskService = async (id: string) => {
  const data = await fetchApiJsonCore<{id:string},Task>({
    endpoint: `/api/task`,
    method: 'DELETE',
    body: { id },
    errorMessage: 'error deleteTaskService タスク削除失敗',
    getKey: 'task',
  });
  
  return data;
};
