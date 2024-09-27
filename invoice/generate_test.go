package invoice_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"encore.app/invoice"
	"encore.app/model"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePDF(t *testing.T) {
	request := &model.Request{
		CustomerName:        "John Doe",
		CustomerPhoneNumber: "(+1) 555-123-4567",
		CustomerAddress:     "742 Evergreen Terrace, Springfield, IL 62704",
		InvoiceNumber:       "12345",
		InvoiceDate:         "October 1, 2024",
		Transactions: []model.Transaction{
			{Name: "Blue Denim Jacket", Quantity: "1", UnitPrice: "$150", Total: "$150"},
			{Name: "Striped Polo Shirt", Quantity: "3", UnitPrice: "$80", Total: "$240"},
			{Name: "Patterned Skirt", Quantity: "2", UnitPrice: "$65", Total: "$130"},
		},
		Subtotal:           "$520",
		TaxPercentage:      "5%",
		TaxAmount:          "$26.00",
		Total:              "$546.00",
		BankName:           "First National Bank",
		OwnerAccountName:   "John Smith",
		OwnerAccountNumber: "123-456-789-10",
		PaymentDate:        "October 10, 2024",
		OwnerName:          "John Smith",
		OwnerAddress:       "123 Maple St, Anytown, CA 90210",
	}

	response, err := invoice.Generate(request)

	assert.NoError(t, err)
	assert.NotEmpty(t, response, "PDF data should not be empty")

	outputFolder := "pdf"
	outputFile := filepath.Join(outputFolder, fmt.Sprintf("invoice_%s.pdf", request.InvoiceNumber))

	// Create the folder if it doesn't exist
	err = os.MkdirAll(outputFolder, os.ModePerm)
	assert.NoError(t, err)

	err = os.WriteFile(outputFile, response, 0644)
	assert.NoError(t, err)
}
