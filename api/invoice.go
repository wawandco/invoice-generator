package api

import (
	"context"

	"encore.app/invoice"
	"encore.app/model"
)

//encore:api public method=POST path=/generate-invoice
func GenerateInvoice(ctx context.Context, data *model.Request) (model.Response, error) {
	response, err := invoice.Generate(data)
	if err != nil {
		return model.Response{}, err
	}

	return response, nil
}
