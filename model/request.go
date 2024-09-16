package model

// Transaction is the details for each invoice item.
type Transaction struct {
	Name      string
	Quantity  string
	UnitPrice string
	Total     string
}

// Request is the information needed to generate the PDF invoice
type Request struct {
	// Billed To
	CustomerName        string
	CustomerPhoneNumber string
	CustomerAddress     string
	InvoiceNumber       string
	InvoiceDate         string

	// Transactions
	Transactions  []Transaction
	Subtotal      string
	TaxPercentage string
	TaxAmount     string
	Total         string

	// Payment
	BankName           string
	OwnerAccountName   string
	OwnerAccountNumber string
	PaymentDate        string
	OwnerName          string
	OwnerAddress       string
}
