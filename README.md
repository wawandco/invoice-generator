# Invoice Generator

Invoice Generator is an API service built with Encore to create PDF invoices with Maroto V2 based on provided data.

## Setup

### Prerequisites

- [Encore CLI](https://encore.dev/docs/getting-started/installation) installed.
- [Go](https://golang.org/doc/install) installed.
- [Maroto](https://github.com/johnfercher/maroto) for PDF generation.

### Installation

**Clone the Repository**

```sh
git clone https://github.com/wawandco/invoice-generator.git
cd invoice-generator
```

**Install Dependencies**

```sh
go mod tidy
```

**Run the Application**

Start the application with Encore:

```sh
encore run
```

And call http://localhost:4000 to use the app's endpoints.

## API Endpoints

### Generate Invoice

- **Method**: POST
- **Path**: `/generate-invoice`
- **Description**: Generates a PDF invoice based on the [provided data](https://github.com/wawandco/invoice-generator/blob/main/model/request.go) and returns the PDF bytes.


## Related Documentation

- [Encore](https://encore.dev/docs/go)
- [Maroto PDF Generation](https://maroto.io/)

## Copyright
This Repo is Copyright © 2024 Wawandco SAS.

Made with ❤️ at [wawand.co](https://wawand.co)