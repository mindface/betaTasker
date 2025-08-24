package errors

import (
	"fmt"
	"net/http"
)

type ErrorCode string

const (
	// 認証・認可エラー (AUTH_xxx)
	AUTH_INVALID_CREDENTIALS ErrorCode = "AUTH_001"
	AUTH_TOKEN_EXPIRED       ErrorCode = "AUTH_002"
	AUTH_TOKEN_INVALID       ErrorCode = "AUTH_003"
	AUTH_UNAUTHORIZED        ErrorCode = "AUTH_004"
	AUTH_ACCOUNT_DISABLED    ErrorCode = "AUTH_005"

	// バリデーションエラー (VAL_xxx)
	VAL_INVALID_INPUT     ErrorCode = "VAL_001"
	VAL_MISSING_FIELD     ErrorCode = "VAL_002"
	VAL_INVALID_FORMAT    ErrorCode = "VAL_003"
	VAL_DUPLICATE_ENTRY   ErrorCode = "VAL_004"
	VAL_CONSTRAINT_FAILED ErrorCode = "VAL_005"

	// リソースエラー (RES_xxx)
	RES_NOT_FOUND      ErrorCode = "RES_001"
	RES_ALREADY_EXISTS ErrorCode = "RES_002"
	RES_ACCESS_DENIED  ErrorCode = "RES_003"
	RES_LOCKED         ErrorCode = "RES_004"

	// データベースエラー (DB_xxx)
	DB_CONNECTION_FAILED ErrorCode = "DB_001"
	DB_QUERY_FAILED      ErrorCode = "DB_002"
	DB_TRANSACTION_FAILED ErrorCode = "DB_003"
	DB_TIMEOUT           ErrorCode = "DB_004"

	// ビジネスロジックエラー (BIZ_xxx)
	BIZ_OPERATION_NOT_ALLOWED ErrorCode = "BIZ_001"
	BIZ_LIMIT_EXCEEDED        ErrorCode = "BIZ_002"
	BIZ_INVALID_STATE         ErrorCode = "BIZ_003"
	BIZ_DEPENDENCY_ERROR      ErrorCode = "BIZ_004"

	// システムエラー (SYS_xxx)
	SYS_INTERNAL_ERROR ErrorCode = "SYS_001"
	SYS_SERVICE_UNAVAILABLE ErrorCode = "SYS_002"
	SYS_TIMEOUT ErrorCode = "SYS_003"
	SYS_RATE_LIMIT_EXCEEDED ErrorCode = "SYS_004"
)

type AppError struct {
	Code       ErrorCode `json:"code"`
	Message    string    `json:"message"`
	Detail     string    `json:"detail,omitempty"`
	HTTPStatus int       `json:"-"`
}

func (e *AppError) Error() string {
	if e.Detail != "" {
		return fmt.Sprintf("[%s] %s: %s", e.Code, e.Message, e.Detail)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func NewAppError(code ErrorCode, message string, detail string) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		Detail:     detail,
		HTTPStatus: getHTTPStatus(code),
	}
}

func getHTTPStatus(code ErrorCode) int {
	switch code {
	case AUTH_INVALID_CREDENTIALS, AUTH_TOKEN_EXPIRED, AUTH_TOKEN_INVALID:
		return http.StatusUnauthorized
	case AUTH_UNAUTHORIZED, AUTH_ACCOUNT_DISABLED:
		return http.StatusForbidden
	case VAL_INVALID_INPUT, VAL_MISSING_FIELD, VAL_INVALID_FORMAT, VAL_DUPLICATE_ENTRY, VAL_CONSTRAINT_FAILED:
		return http.StatusBadRequest
	case RES_NOT_FOUND:
		return http.StatusNotFound
	case RES_ALREADY_EXISTS:
		return http.StatusConflict
	case RES_ACCESS_DENIED:
		return http.StatusForbidden
	case RES_LOCKED:
		return http.StatusLocked
	case DB_CONNECTION_FAILED, DB_QUERY_FAILED, DB_TRANSACTION_FAILED, DB_TIMEOUT:
		return http.StatusInternalServerError
	case BIZ_OPERATION_NOT_ALLOWED:
		return http.StatusForbidden
	case BIZ_LIMIT_EXCEEDED:
		return http.StatusTooManyRequests
	case BIZ_INVALID_STATE, BIZ_DEPENDENCY_ERROR:
		return http.StatusBadRequest
	case SYS_INTERNAL_ERROR:
		return http.StatusInternalServerError
	case SYS_SERVICE_UNAVAILABLE:
		return http.StatusServiceUnavailable
	case SYS_TIMEOUT:
		return http.StatusRequestTimeout
	case SYS_RATE_LIMIT_EXCEEDED:
		return http.StatusTooManyRequests
	default:
		return http.StatusInternalServerError
	}
}

var errorMessages = map[ErrorCode]string{
	AUTH_INVALID_CREDENTIALS:  "認証情報が無効です",
	AUTH_TOKEN_EXPIRED:        "トークンの有効期限が切れています",
	AUTH_TOKEN_INVALID:        "無効なトークンです",
	AUTH_UNAUTHORIZED:         "権限がありません",
	AUTH_ACCOUNT_DISABLED:     "アカウントが無効化されています",
	VAL_INVALID_INPUT:         "入力値が無効です",
	VAL_MISSING_FIELD:         "必須項目が入力されていません",
	VAL_INVALID_FORMAT:        "入力形式が正しくありません",
	VAL_DUPLICATE_ENTRY:       "既に登録されています",
	VAL_CONSTRAINT_FAILED:     "制約条件を満たしていません",
	RES_NOT_FOUND:             "リソースが見つかりません",
	RES_ALREADY_EXISTS:        "リソースが既に存在します",
	RES_ACCESS_DENIED:         "アクセスが拒否されました",
	RES_LOCKED:                "リソースがロックされています",
	DB_CONNECTION_FAILED:      "データベース接続に失敗しました",
	DB_QUERY_FAILED:           "クエリの実行に失敗しました",
	DB_TRANSACTION_FAILED:     "トランザクションに失敗しました",
	DB_TIMEOUT:                "データベースタイムアウト",
	BIZ_OPERATION_NOT_ALLOWED: "この操作は許可されていません",
	BIZ_LIMIT_EXCEEDED:        "制限を超えました",
	BIZ_INVALID_STATE:         "無効な状態です",
	BIZ_DEPENDENCY_ERROR:      "依存関係エラー",
	SYS_INTERNAL_ERROR:        "内部エラーが発生しました",
	SYS_SERVICE_UNAVAILABLE:   "サービスが利用できません",
	SYS_TIMEOUT:               "タイムアウトしました",
	SYS_RATE_LIMIT_EXCEEDED:   "レート制限を超えました",
}

func GetErrorMessage(code ErrorCode) string {
	if msg, ok := errorMessages[code]; ok {
		return msg
	}
	return "エラーが発生しました"
}