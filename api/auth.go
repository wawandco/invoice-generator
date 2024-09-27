package api

import (
	"context"
	"strings"

	"encore.dev/beta/auth"
	"encore.dev/beta/errs"
)

//encore:authhandler
func AuthHandler(ctx context.Context, token string) (auth.UID, error) {
	if strings.EqualFold(token, secrets.APIToken) {
		return "accepted", nil
	}

	return "", &errs.Error{
		Code:    errs.Unauthenticated,
		Message: "invalid token",
	}
}
