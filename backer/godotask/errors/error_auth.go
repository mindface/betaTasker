package errors

var (
	ErrAuthTokenExpired = NewAppError(
		AUTH_TOKEN_EXPIRED,
		"トークンの有効期限が切れています",
		"",
	)

	ErrAuthTokenInvalid = NewAppError(
		AUTH_TOKEN_INVALID,
		"無効なトークンです",
		"",
	)
)