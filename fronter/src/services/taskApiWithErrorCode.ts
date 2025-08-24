import { Task, AddTask } from '../model/task';
import { ApplicationError, ErrorCode, parseErrorResponse } from '../errors/errorCodes';

const API_BASE_URL = process.env.REACT_APP_API_URL || 'http://localhost:8081';

interface ApiResponse<T> {
  success?: boolean;
  data?: T;
  message?: string;
  code?: string;
  detail?: string;
}

interface ErrorResponse {
  code: string;
  message: string;
  detail?: string;
}

class TaskApiClient {
  private baseURL: string;
  private timeout: number;

  constructor(baseURL: string = API_BASE_URL, timeout: number = 10000) {
    this.baseURL = baseURL;
    this.timeout = timeout;
  }

  private async handleError(response: Response): Promise<never> {
    let errorData: ErrorResponse | null = null;

    try {
      errorData = await response.json();
    } catch {
      // JSONパースできない場合
    }

    if (errorData?.code) {
      throw new ApplicationError(
        errorData.code as ErrorCode,
        errorData.message,
        errorData.detail
      );
    }

    // エラーコードがない場合はHTTPステータスコードから判定
    switch (response.status) {
      case 400:
        throw new ApplicationError(
          ErrorCode.VAL_INVALID_INPUT,
          errorData?.message || '不正なリクエストです'
        );
      case 401:
        throw new ApplicationError(
          ErrorCode.AUTH_INVALID_CREDENTIALS,
          errorData?.message || '認証が必要です'
        );
      case 403:
        throw new ApplicationError(
          ErrorCode.AUTH_UNAUTHORIZED,
          errorData?.message || 'アクセス権限がありません'
        );
      case 404:
        throw new ApplicationError(
          ErrorCode.RES_NOT_FOUND,
          errorData?.message || 'リソースが見つかりません'
        );
      case 409:
        throw new ApplicationError(
          ErrorCode.RES_ALREADY_EXISTS,
          errorData?.message || 'リソースが既に存在します'
        );
      case 429:
        throw new ApplicationError(
          ErrorCode.SYS_RATE_LIMIT_EXCEEDED,
          errorData?.message || 'リクエスト数が制限を超えました'
        );
      case 500:
        throw new ApplicationError(
          ErrorCode.SYS_INTERNAL_ERROR,
          errorData?.message || 'サーバーエラーが発生しました'
        );
      case 503:
        throw new ApplicationError(
          ErrorCode.SYS_SERVICE_UNAVAILABLE,
          errorData?.message || 'サービスが一時的に利用できません'
        );
      default:
        throw new ApplicationError(
          ErrorCode.SYS_INTERNAL_ERROR,
          `HTTPエラー: ${response.status}`
        );
    }
  }

  private async fetchWithTimeout(url: string, options: RequestInit = {}): Promise<Response> {
    const controller = new AbortController();
    const timeoutId = setTimeout(() => controller.abort(), this.timeout);

    try {
      const response = await fetch(url, {
        ...options,
        signal: controller.signal,
        credentials: 'include',
      });
      clearTimeout(timeoutId);
      return response;
    } catch (error: any) {
      clearTimeout(timeoutId);
      
      if (error.name === 'AbortError') {
        throw new ApplicationError(
          ErrorCode.NET_REQUEST_TIMEOUT,
          'リクエストがタイムアウトしました'
        );
      }
      
      throw new ApplicationError(
        ErrorCode.NET_CONNECTION_FAILED,
        'サーバーに接続できません'
      );
    }
  }

  async getTasks(): Promise<Task[]> {
    try {
      const response = await this.fetchWithTimeout(
        `${this.baseURL}/api/tasks`
      );
      
      if (!response.ok) {
        await this.handleError(response);
      }

      const data: ApiResponse<Task[]> = await response.json();
      return data.data || [];
    } catch (error) {
      if (error instanceof ApplicationError) {
        throw error;
      }
      throw new ApplicationError(
        ErrorCode.SYS_INTERNAL_ERROR,
        '予期しないエラーが発生しました'
      );
    }
  }

  async getTask(id: number): Promise<Task> {
    try {
      const response = await this.fetchWithTimeout(
        `${this.baseURL}/api/task/${id}`
      );
      
      if (!response.ok) {
        await this.handleError(response);
      }
      
      const data: ApiResponse<Task> = await response.json();
      
      if (!data.data) {
        throw new ApplicationError(
          ErrorCode.RES_NOT_FOUND,
          `タスクID: ${id} が見つかりません`
        );
      }
      
      return data.data;
    } catch (error) {
      if (error instanceof ApplicationError) {
        throw error;
      }
      throw new ApplicationError(
        ErrorCode.SYS_INTERNAL_ERROR,
        '予期しないエラーが発生しました'
      );
    }
  }

  async addTask(task: AddTask): Promise<Task> {
    try {
      // バリデーション
      if (!task.title || task.title.trim() === '') {
        throw new ApplicationError(
          ErrorCode.VAL_MISSING_FIELD,
          'タイトルは必須項目です'
        );
      }

      const response = await this.fetchWithTimeout(
        `${this.baseURL}/api/task`,
        {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(task),
        }
      );
      
      if (!response.ok) {
        await this.handleError(response);
      }
      
      const data: ApiResponse<Task> = await response.json();
      
      if (!data.data) {
        throw new ApplicationError(
          ErrorCode.SYS_INTERNAL_ERROR,
          'タスクの作成に失敗しました'
        );
      }
      
      return data.data;
    } catch (error) {
      if (error instanceof ApplicationError) {
        throw error;
      }
      throw new ApplicationError(
        ErrorCode.SYS_INTERNAL_ERROR,
        '予期しないエラーが発生しました'
      );
    }
  }

  async updateTask(task: Task): Promise<Task> {
    try {
      if (!task.id) {
        throw new ApplicationError(
          ErrorCode.VAL_MISSING_FIELD,
          'タスクIDが指定されていません'
        );
      }

      const response = await this.fetchWithTimeout(
        `${this.baseURL}/api/task/${task.id}`,
        {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(task),
        }
      );
      
      if (!response.ok) {
        await this.handleError(response);
      }
      
      const data: ApiResponse<Task> = await response.json();
      
      if (!data.data) {
        throw new ApplicationError(
          ErrorCode.SYS_INTERNAL_ERROR,
          'タスクの更新に失敗しました'
        );
      }
      
      return data.data;
    } catch (error) {
      if (error instanceof ApplicationError) {
        throw error;
      }
      throw new ApplicationError(
        ErrorCode.SYS_INTERNAL_ERROR,
        '予期しないエラーが発生しました'
      );
    }
  }

  async deleteTask(id: number): Promise<void> {
    try {
      const response = await this.fetchWithTimeout(
        `${this.baseURL}/api/task/${id}`,
        {
          method: 'DELETE',
        }
      );
      
      if (!response.ok) {
        await this.handleError(response);
      }
    } catch (error) {
      if (error instanceof ApplicationError) {
        throw error;
      }
      throw new ApplicationError(
        ErrorCode.SYS_INTERNAL_ERROR,
        '予期しないエラーが発生しました'
      );
    }
  }
}

export const taskApiClient = new TaskApiClient();

// 使用例を示すヘルパー関数
export async function handleTaskOperation<T>(
  operation: () => Promise<T>,
  onSuccess?: (result: T) => void,
  onError?: (error: ApplicationError) => void
): Promise<T | null> {
  try {
    const result = await operation();
    onSuccess?.(result);
    return result;
  } catch (error) {
    const appError = error instanceof ApplicationError 
      ? error 
      : parseErrorResponse(error);
    
    console.error(`[${appError.code}] ${appError.message}`, appError.detail);
    onError?.(appError);
    return null;
  }
}