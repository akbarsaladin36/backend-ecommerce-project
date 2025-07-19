package inputs

type CreateCartInput struct {
	ProductCode     string `json:"product_code"`
	CartDescription string `json:"cart_description"`
	CartQuantity    string `json:"cart_quantity"`
}

type UpdateCartInput struct {
	ProductCode     string `json:"product_code"`
	CartDescription string `json:"cart_description"`
	CartQuantity    string `json:"cart_quantity"`
}
