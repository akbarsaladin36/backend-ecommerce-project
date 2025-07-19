package inputs

type CreateProductInput struct {
	ProductName        string `json:"product_name"`
	ProductDescription string `json:"product_description"`
	ProductPrice       string `json:"product_price"`
	ProductQuantity    string `json:"product_quantity"`
}

type UpdateProductInput struct {
	ProductName        string `json:"product_name"`
	ProductDescription string `json:"product_description"`
	ProductPrice       string `json:"product_price"`
	ProductQuantity    string `json:"product_quantity"`
}
