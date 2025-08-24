package errors

import (
	"net/http"
	"testing"
)

func TestNewAppError(t *testing.T) {
	tests := []struct {
		name           string
		code           ErrorCode
		message        string
		detail         string
		expectedStatus int
	}{
		{
			name:           "認証エラー",
			code:           AUTH_INVALID_CREDENTIALS,
			message:        "認証情報が無効です",
			detail:         "ユーザー名またはパスワードが間違っています",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "バリデーションエラー",
			code:           VAL_MISSING_FIELD,
			message:        "必須項目が入力されていません",
			detail:         "タイトルフィールドが空です",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "リソース未発見エラー",
			code:           RES_NOT_FOUND,
			message:        "タスクが見つかりません",
			detail:         "ID: 123のタスクは存在しません",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "データベースエラー",
			code:           DB_CONNECTION_FAILED,
			message:        "データベース接続エラー",
			detail:         "",
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "レート制限エラー",
			code:           SYS_RATE_LIMIT_EXCEEDED,
			message:        "API呼び出し制限を超えました",
			detail:         "1分間に60回までです",
			expectedStatus: http.StatusTooManyRequests,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewAppError(tt.code, tt.message, tt.detail)

			// エラーコードの確認
			if err.Code != tt.code {
				t.Errorf("Expected code %s, got %s", tt.code, err.Code)
			}

			// メッセージの確認
			if err.Message != tt.message {
				t.Errorf("Expected message %s, got %s", tt.message, err.Message)
			}

			// 詳細の確認
			if err.Detail != tt.detail {
				t.Errorf("Expected detail %s, got %s", tt.detail, err.Detail)
			}

			// HTTPステータスコードの確認
			if err.HTTPStatus != tt.expectedStatus {
				t.Errorf("Expected HTTP status %d, got %d", tt.expectedStatus, err.HTTPStatus)
			}
		})
	}
}

func TestGetErrorMessage(t *testing.T) {
	tests := []struct {
		name     string
		code     ErrorCode
		expected string
	}{
		{
			name:     "認証エラーメッセージ",
			code:     AUTH_INVALID_CREDENTIALS,
			expected: "認証情報が無効です",
		},
		{
			name:     "バリデーションエラーメッセージ",
			code:     VAL_MISSING_FIELD,
			expected: "必須項目が入力されていません",
		},
		{
			name:     "リソースエラーメッセージ",
			code:     RES_NOT_FOUND,
			expected: "リソースが見つかりません",
		},
		{
			name:     "存在しないエラーコード",
			code:     ErrorCode("UNKNOWN_999"),
			expected: "エラーが発生しました",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := GetErrorMessage(tt.code)
			if msg != tt.expected {
				t.Errorf("Expected message %s, got %s", tt.expected, msg)
			}
		})
	}
}

func TestAppErrorString(t *testing.T) {
	tests := []struct {
		name     string
		err      *AppError
		expected string
	}{
		{
			name: "詳細付きエラー",
			err: &AppError{
				Code:    VAL_INVALID_INPUT,
				Message: "入力エラー",
				Detail:  "メールアドレスの形式が不正です",
			},
			expected: "[VAL_001] 入力エラー: メールアドレスの形式が不正です",
		},
		{
			name: "詳細なしエラー",
			err: &AppError{
				Code:    AUTH_TOKEN_EXPIRED,
				Message: "トークン期限切れ",
				Detail:  "",
			},
			expected: "[AUTH_002] トークン期限切れ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.err.Error()
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestHTTPStatusMapping(t *testing.T) {
	statusTests := []struct {
		code         ErrorCode
		expectedCode int
	}{
		// 認証系
		{AUTH_INVALID_CREDENTIALS, http.StatusUnauthorized},
		{AUTH_TOKEN_EXPIRED, http.StatusUnauthorized},
		{AUTH_UNAUTHORIZED, http.StatusForbidden},
		
		// バリデーション系
		{VAL_INVALID_INPUT, http.StatusBadRequest},
		{VAL_MISSING_FIELD, http.StatusBadRequest},
		{VAL_DUPLICATE_ENTRY, http.StatusBadRequest},
		
		// リソース系
		{RES_NOT_FOUND, http.StatusNotFound},
		{RES_ALREADY_EXISTS, http.StatusConflict},
		{RES_ACCESS_DENIED, http.StatusForbidden},
		{RES_LOCKED, http.StatusLocked},
		
		// データベース系
		{DB_CONNECTION_FAILED, http.StatusInternalServerError},
		{DB_QUERY_FAILED, http.StatusInternalServerError},
		
		// ビジネスロジック系
		{BIZ_OPERATION_NOT_ALLOWED, http.StatusForbidden},
		{BIZ_LIMIT_EXCEEDED, http.StatusTooManyRequests},
		
		// システム系
		{SYS_INTERNAL_ERROR, http.StatusInternalServerError},
		{SYS_SERVICE_UNAVAILABLE, http.StatusServiceUnavailable},
		{SYS_TIMEOUT, http.StatusRequestTimeout},
		{SYS_RATE_LIMIT_EXCEEDED, http.StatusTooManyRequests},
	}

	for _, tt := range statusTests {
		t.Run(string(tt.code), func(t *testing.T) {
			status := getHTTPStatus(tt.code)
			if status != tt.expectedCode {
				t.Errorf("Code %s: expected status %d, got %d", 
					tt.code, tt.expectedCode, status)
			}
		})
	}
}