package api

import (
	"context"

	"encore.app/model"
	"encore.dev/storage/sqldb"
)

// Creates the invoice_generator database and assign it to the "db" variable
var DB = sqldb.NewDatabase("invoice_generator", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})

// createInvoice inserts an invoice item into the database.
func createInvoice(ctx context.Context, data []byte) (model.Invoice, error) {
	invoice := model.Invoice{}

	err := DB.QueryRow(ctx, `
		INSERT INTO invoices (data, created_at)
		VALUES ($1, CURRENT_TIMESTAMP)
		RETURNING id, data, created_at
	`, data).Scan(&invoice.ID, &invoice.Data, &invoice.CreatedAt)

	return invoice, err
}

// find gets an invoice from the database based on the given id.
func findInvoice(ctx context.Context, id int) (model.Invoice, error) {
	invoice := model.Invoice{}

	err := DB.QueryRow(ctx, `
	    SELECT id, data, created_at
	    FROM invoices
	    WHERE id = $1
	`, id).Scan(&invoice.ID, &invoice.Data, &invoice.CreatedAt)

	return invoice, err
}
