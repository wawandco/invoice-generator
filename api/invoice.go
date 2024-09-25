package api

import (
	"context"
	"strings"

	"encore.app/invoice"
	"encore.app/model"
	"encore.dev/beta/auth"
	"encore.dev/beta/errs"
)

// Generates a PDF invoice based on the provided data and returns the PDF's byte data.
//
//encore:api auth method=POST path=/generate-invoice
func GenerateInvoice(ctx context.Context, data *model.Request) (model.Response, error) {
	response, err := invoice.Generate(data)
	if err != nil {
		return model.Response{}, err
	}

	return response, nil
}

// Data can be named whatever you prefer (but must be exported).
// type Data struct {
// 	Username string
// 	// ...
// }

// AuthHandler can be named whatever you prefer (but must be exported).
//
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
