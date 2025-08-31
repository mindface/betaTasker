import { AddTask, Task } from "../model/task";
import { ApplicationError, ErrorCode, parseErrorResponse } from "../errors/errorCodes";

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
  try {
    const res = await fetch('/api/task', {
      method: 'GET',
      credentials: 'include',
    });
    
    if (!res.ok) {
      await handleApiError(res);
    }

    const data: SuccessResponse<Task[], 'tasks'> = await res.json();
    return data.tasks || data;
  } catch (err) {
    const appError = parseErrorResponse(err);
    return { error: appError.message, code: appError.code };
  }
};

export const addTaskService = async (task: AddTask) => {
  try {
    // クライアントサイドバリデーション
    if (!task.title || task.title.trim() === '') {
      throw new ApplicationError(
        ErrorCode.VAL_MISSING_FIELD,
        'タイトルは必須項目です'
      );
    }

    const res = await fetch('/api/task', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(task),
      credentials: 'include',
    });

    if (!res.ok) {
      await handleApiError(res);
    }

    const data: SuccessResponse<Task, 'task'> = await res.json();
    return data.task || data;
  } catch (err) {
    const appError = parseErrorResponse(err);
    return { error: appError.message, code: appError.code };
  }
};

export const updateTaskService = async (task: Task) => {
  try {
    if (!task.id) {
      throw new ApplicationError(
        ErrorCode.VAL_MISSING_FIELD,
        'タスクIDが指定されていません'
      );
    }

    const res = await fetch('/api/task', {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(task),
      credentials: 'include',
    });
    
    if (!res.ok) {
      await handleApiError(res);
    }

    const data: SuccessResponse<Task, 'task'> = await res.json();
    return data.task || data;
  } catch (err) {
    const appError = parseErrorResponse(err);
    return { error: appError.message, code: appError.code };
  }
};

export const deleteTaskService = async (id: string) => {
  try {
    if (!id) {
      throw new ApplicationError(
        ErrorCode.VAL_MISSING_FIELD,
        '削除するタスクのIDが指定されていません'
      );
    }

    const res = await fetch(`/api/task`, {
      method: 'DELETE',
      credentials: 'include',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ id }),
    });
    
    if (!res.ok) {
      await handleApiError(res);
    }
    
    const data: SuccessResponse<void> = await res.json();
    return data;
  } catch (err) {
    const appError = parseErrorResponse(err);
    return { error: appError.message, code: appError.code };
  }
};
