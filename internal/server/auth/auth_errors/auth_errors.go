package auth_errors

import "errors"

var (
	ErrorInvalidAccessToken        error = errors.New("invalid access token")
	ErrorInvalidRefreshToken       error = errors.New("invalid access token")
	ErrorMissingAccessToken        error = errors.New("missing access token")
	ErrorMissingRefreshToken       error = errors.New("missing refresh token")
	ErrorFailToParseNewAccessToken error = errors.New("failed to parse new access token")
)
