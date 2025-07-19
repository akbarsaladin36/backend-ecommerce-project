package inputs

type CreateInvoiceInput struct {
	CartCode     string `json:"cart_code"`
	InvoicePrice string `json:"invoice_price"`
	InvoiceDesc  string `json:"invoice_desc"`
}

type UpdateInvoiceInput struct {
	InvoiceStatusCd string `json:"invoice_status_cd"`
}
