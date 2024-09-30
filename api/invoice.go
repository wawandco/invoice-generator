package api

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"encore.app/invoice"
	"encore.app/model"
	"encore.dev/beta/errs"
)

// Generates a PDF invoice based on the provided data and returns the PDF's byte data.
//
//encore:api auth method=POST path=/generate-invoice
func GenerateInvoice(ctx context.Context, data *model.Request) (model.Response, error) {
	invoiceData, err := invoice.Generate(data)
	if err != nil {
		errorResponse := &errs.Error{
			Code:    errs.Internal,
			Message: err.Error(),
		}

		return model.Response{}, errorResponse
	}

	// Storing the invoice
	invoice, err := createInvoice(ctx, invoiceData)
	if err != nil {
		errorResponse := &errs.Error{
			Code:    errs.Internal,
			Message: err.Error(),
		}

		return model.Response{}, errorResponse
	}

	response := model.Response{
		Status:  http.StatusCreated,
		Invoice: &invoice,
	}

	return response, nil
}

// GetInvoice returns the PDF's byte data based on the given id.
//
//encore:api auth method=GET path=/get-invoice/:id
func GetInvoice(ctx context.Context, id int) (model.Response, error) {
	invoice, err := findInvoice(ctx, id)

	if err != nil {
		apiResponse := model.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}

		if errors.Is(err, sql.ErrNoRows) {
			apiResponse.Status = http.StatusNotFound
			apiResponse.Message = "invoice not found"
		}

		return apiResponse, err
	}

	response := model.Response{
		Status:  http.StatusOK,
		Invoice: &invoice,
	}

	return response, nil
}
