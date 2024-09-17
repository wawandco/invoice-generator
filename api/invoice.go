package api

import (
	"context"

	"encore.app/invoice"
	"encore.app/model"
)

// Generates a PDF invoice based on the provided data and returns the PDF's byte data.
//
//encore:api public method=POST path=/generate-invoice
func GenerateInvoice(ctx context.Context, data *model.Request) (model.Response, error) {
	response, err := invoice.Generate(data)
	if err != nil {
		return model.Response{}, err
	}

	return response, nil
}
