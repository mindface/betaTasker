export enum ErrorCode {
  // 認証・認可エラー (AUTH_xxx)
  AUTH_INVALID_CREDENTIALS = 'AUTH_001',
  AUTH_TOKEN_EXPIRED = 'AUTH_002',
  AUTH_TOKEN_INVALID = 'AUTH_003',
  AUTH_UNAUTHORIZED = 'AUTH_004',
  AUTH_ACCOUNT_DISABLED = 'AUTH_005',

  // バリデーションエラー (VAL_xxx)
  VAL_INVALID_INPUT = 'VAL_001',
  VAL_MISSING_FIELD = 'VAL_002',
  VAL_INVALID_FORMAT = 'VAL_003',
  VAL_DUPLICATE_ENTRY = 'VAL_004',
  VAL_CONSTRAINT_FAILED = 'VAL_005',

  // リソースエラー (RES_xxx)
  RES_NOT_FOUND = 'RES_001',
  RES_ALREADY_EXISTS = 'RES_002',
  RES_ACCESS_DENIED = 'RES_003',
  RES_LOCKED = 'RES_004',

  // データベースエラー (DB_xxx)
  DB_CONNECTION_FAILED = 'DB_001',
  DB_QUERY_FAILED = 'DB_002',
  DB_TRANSACTION_FAILED = 'DB_003',
  DB_TIMEOUT = 'DB_004',

  // ビジネスロジックエラー (BIZ_xxx)
  BIZ_OPERATION_NOT_ALLOWED = 'BIZ_001',
  BIZ_LIMIT_EXCEEDED = 'BIZ_002',
  BIZ_INVALID_STATE = 'BIZ_003',
  BIZ_DEPENDENCY_ERROR = 'BIZ_004',

  // システムエラー (SYS_xxx)
  SYS_INTERNAL_ERROR = 'SYS_001',
  SYS_SERVICE_UNAVAILABLE = 'SYS_002',
  SYS_TIMEOUT = 'SYS_003',
  SYS_RATE_LIMIT_EXCEEDED = 'SYS_004',

  // ネットワークエラー (NET_xxx)
  NET_CONNECTION_FAILED = 'NET_001',
  NET_REQUEST_TIMEOUT = 'NET_002',
  NET_REQUEST_ABORTED = 'NET_003',
}

export interface AppError {
  code: ErrorCode;
  message: string;
  detail?: string;
  timestamp?: string;
  path?: string;
}

export class ApplicationError extends Error {
  public code: ErrorCode;
  public detail?: string;
  public timestamp: string;
  public path?: string;

  constructor(code: ErrorCode, message?: string, detail?: string) {
    const errorMessage = message || getErrorMessage(code);
    super(errorMessage);
    this.name = 'ApplicationError';
    this.code = code;
    this.detail = detail;
    this.timestamp = new Date().toISOString();
    Object.setPrototypeOf(this, ApplicationError.prototype);
  }

  toJSON(): AppError {
    return {
      code: this.code,
      message: this.message,
      detail: this.detail,
      timestamp: this.timestamp,
      path: this.path,
    };
  }
}

const errorMessages: Record<ErrorCode, string> = {
  [ErrorCode.AUTH_INVALID_CREDENTIALS]: '認証情報が無効です',
  [ErrorCode.AUTH_TOKEN_EXPIRED]: 'トークンの有効期限が切れています',
  [ErrorCode.AUTH_TOKEN_INVALID]: '無効なトークンです',
  [ErrorCode.AUTH_UNAUTHORIZED]: '権限がありません',
  [ErrorCode.AUTH_ACCOUNT_DISABLED]: 'アカウントが無効化されています',
  [ErrorCode.VAL_INVALID_INPUT]: '入力値が無効です',
  [ErrorCode.VAL_MISSING_FIELD]: '必須項目が入力されていません',
  [ErrorCode.VAL_INVALID_FORMAT]: '入力形式が正しくありません',
  [ErrorCode.VAL_DUPLICATE_ENTRY]: '既に登録されています',
  [ErrorCode.VAL_CONSTRAINT_FAILED]: '制約条件を満たしていません',
  [ErrorCode.RES_NOT_FOUND]: 'リソースが見つかりません',
  [ErrorCode.RES_ALREADY_EXISTS]: 'リソースが既に存在します',
  [ErrorCode.RES_ACCESS_DENIED]: 'アクセスが拒否されました',
  [ErrorCode.RES_LOCKED]: 'リソースがロックされています',
  [ErrorCode.DB_CONNECTION_FAILED]: 'データベース接続に失敗しました',
  [ErrorCode.DB_QUERY_FAILED]: 'クエリの実行に失敗しました',
  [ErrorCode.DB_TRANSACTION_FAILED]: 'トランザクションに失敗しました',
  [ErrorCode.DB_TIMEOUT]: 'データベースタイムアウト',
  [ErrorCode.BIZ_OPERATION_NOT_ALLOWED]: 'この操作は許可されていません',
  [ErrorCode.BIZ_LIMIT_EXCEEDED]: '制限を超えました',
  [ErrorCode.BIZ_INVALID_STATE]: '無効な状態です',
  [ErrorCode.BIZ_DEPENDENCY_ERROR]: '依存関係エラー',
  [ErrorCode.SYS_INTERNAL_ERROR]: '内部エラーが発生しました',
  [ErrorCode.SYS_SERVICE_UNAVAILABLE]: 'サービスが利用できません',
  [ErrorCode.SYS_TIMEOUT]: 'タイムアウトしました',
  [ErrorCode.SYS_RATE_LIMIT_EXCEEDED]: 'レート制限を超えました',
  [ErrorCode.NET_CONNECTION_FAILED]: 'ネットワーク接続に失敗しました',
  [ErrorCode.NET_REQUEST_TIMEOUT]: 'リクエストがタイムアウトしました',
  [ErrorCode.NET_REQUEST_ABORTED]: 'リクエストが中断されました',
};

export function getErrorMessage(code: ErrorCode): string {
  return errorMessages[code] || 'エラーが発生しました';
}

export function parseErrorResponse(error: any): ApplicationError {
  if (error instanceof ApplicationError) {
    return error;
  }

  if (error.response?.data?.code) {
    return new ApplicationError(
      error.response.data.code,
      error.response.data.message,
      error.response.data.detail
    );
  }

  if (error.code === 'ECONNABORTED') {
    return new ApplicationError(ErrorCode.NET_REQUEST_ABORTED);
  }

  if (error.code === 'ENOTFOUND' || error.code === 'ECONNREFUSED') {
    return new ApplicationError(ErrorCode.NET_CONNECTION_FAILED);
  }

  if (error.message?.includes('timeout')) {
    return new ApplicationError(ErrorCode.NET_REQUEST_TIMEOUT);
  }

  return new ApplicationError(
    ErrorCode.SYS_INTERNAL_ERROR,
    error.message || '予期しないエラーが発生しました'
  );
}