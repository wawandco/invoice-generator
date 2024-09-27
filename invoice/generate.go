package invoice

import (
	"fmt"
	"log"

	"encore.app/model"
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func Generate(data *model.Request) ([]byte, error) {
	m := GetMaroto(data)

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	return document.GetBytes(), nil
}

func GetMaroto(data *model.Request) core.Maroto {
	cfg := config.NewBuilder().
		WithPageSize(pagesize.A4).
		WithDefaultFont(&props.Font{Color: getDarkGrayColor()}).
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	err := m.RegisterHeader(getPageHeader())
	if err != nil {
		log.Fatal(err.Error())
	}

	m.AddRows(getBilledToInfo(data)...)
	m.AddRows(getTransactions(data)...)

	err = m.RegisterFooter(getPageFooter(data))
	if err != nil {
		log.Fatal(err.Error())
	}

	return m
}

func getPageHeader() core.Row {
	return row.New(35).Add(
		col.New(6).Add(
			text.New("INVOICE", props.Text{
				Size:   35,
				Align:  align.Left,
				Color:  getGrayColor(),
				Style:  fontstyle.Bold,
				Family: "courier",
			}),
		),
		col.New(6).Add(
			text.New("Really Great Company", props.Text{
				Top:   4,
				Size:  12,
				Align: align.Right,
				Style: fontstyle.Bold,
			}),
			text.New("Your Business Partner", props.Text{
				Top:   10,
				Style: fontstyle.Normal,
				Size:  10,
				Align: align.Right,
			}),
		),
	)
}

func getBilledToInfo(data *model.Request) []core.Row {
	rows := []core.Row{
		row.New(50).Add(
			col.New(6).Add(
				text.New("BILLED TO:", props.Text{
					Size:  12,
					Style: fontstyle.Bold,
				}),
				text.New(data.CustomerName, props.Text{
					Size: 10,
					Top:  6,
				}),
				text.New(data.CustomerPhoneNumber, props.Text{
					Size: 10,
					Top:  11,
				}),
				text.New(data.CustomerAddress, props.Text{
					Size: 10,
					Top:  16,
				}),
			),
			col.New(6).Add(
				text.New(fmt.Sprintf("Invoice No. %s", data.InvoiceNumber), props.Text{
					Size:  10,
					Align: align.Right,
				}),
				text.New(data.InvoiceDate, props.Text{
					Size:  10,
					Align: align.Right,
					Top:   6,
				}),
			),
		),
	}

	return rows
}

func getTransactions(data *model.Request) []core.Row {
	// Table Header
	rows := []core.Row{
		row.New(5).Add(line.NewCol(12, props.Line{SizePercent: 100})),
		row.New(8).Add(
			text.NewCol(6, "Item", props.Text{Size: 10, Align: align.Left, Style: fontstyle.Bold, Left: 4}),
			text.NewCol(2, "Quantity", props.Text{Size: 10, Align: align.Center, Style: fontstyle.Bold}),
			text.NewCol(2, "Unit Price", props.Text{Size: 10, Align: align.Center, Style: fontstyle.Bold}),
			text.NewCol(2, "Total", props.Text{Size: 10, Align: align.Center, Style: fontstyle.Bold}),
		),
		row.New(5).Add(line.NewCol(12, props.Line{SizePercent: 100})),
	}

	// Transactions
	var contentsRow []core.Row
	for _, transaction := range data.Transactions {
		r := row.New(6).Add(
			text.NewCol(6, transaction.Name, props.Text{Size: 10, Align: align.Left, Left: 4}),
			text.NewCol(2, transaction.Quantity, props.Text{Size: 10, Align: align.Center}),
			text.NewCol(2, transaction.UnitPrice, props.Text{Size: 10, Align: align.Center}),
			text.NewCol(2, transaction.Total, props.Text{Size: 10, Align: align.Center}),
		)

		divider := row.New(7).Add(line.NewCol(12, props.Line{SizePercent: 100, OffsetPercent: 50}))
		contentsRow = append(contentsRow, r, divider)
	}

	rows = append(rows, contentsRow...)

	// Totals
	rows = append(rows, row.New(10), row.New(10).Add(
		col.New(8),
		text.NewCol(2, "Subtotal", props.Text{
			Style: fontstyle.Bold,
			Size:  10,
			Align: align.Center,
		}),
		text.NewCol(2, data.Subtotal, props.Text{
			Size:  10,
			Align: align.Center,
		}),
	), row.New(10).Add(
		col.New(8),
		text.NewCol(2, fmt.Sprintf("Tax (%v)", data.TaxPercentage), props.Text{
			Style: fontstyle.Bold,
			Size:  10,
			Align: align.Center,
		}),
		text.NewCol(2, data.TaxAmount, props.Text{
			Size:  10,
			Align: align.Center,
		}),
	), row.New(3).Add(
		col.New(8),
		line.NewCol(4, props.Line{
			SizePercent:   100,
			OffsetPercent: 50,
			Color:         &props.BlackColor,
			Thickness:     0.5,
		}),
	), row.New(10).Add(
		col.New(8),
		text.NewCol(2, "Total", props.Text{
			Style: fontstyle.Bold,
			Size:  16,
			Align: align.Center,
			Top:   2,
		}),
		text.NewCol(2, data.Total, props.Text{
			Style: fontstyle.Bold,
			Size:  16,
			Align: align.Center,
			Top:   2,
		}),
	))

	return rows
}

func getPageFooter(data *model.Request) core.Row {
	return row.New(25).Add(
		col.New(6).Add(
			text.New("PAYMENT INFORMATION:", props.Text{
				Size:  12,
				Style: fontstyle.Bold,
			}),
			text.New(data.BankName, props.Text{
				Size: 10,
				Top:  6,
			}),
			text.New(fmt.Sprintf("Account Name: %s", data.OwnerAccountName), props.Text{
				Size: 10,
				Top:  11,
			}),
			text.New(fmt.Sprintf("Account No.: %s", data.OwnerAccountNumber), props.Text{
				Size: 10,
				Top:  16,
			}),
			text.New(fmt.Sprintf("Pay by: %s", data.PaymentDate), props.Text{
				Size: 10,
				Top:  21,
			}),
		),
		col.New(6).Add(
			text.New(data.OwnerName, props.Text{
				Size:  15,
				Align: align.Right,
			}),
			text.New(data.OwnerAddress, props.Text{
				Size:  10,
				Align: align.Right,
				Top:   7,
			}),
		),
	)
}

func getDarkGrayColor() *props.Color {
	return &props.Color{
		Red:   55,
		Green: 55,
		Blue:  55,
	}
}

func getGrayColor() *props.Color {
	return &props.Color{
		Red:   180,
		Green: 180,
		Blue:  180,
	}
}
