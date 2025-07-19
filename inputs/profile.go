package inputs

type UpdateProfileInput struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type UpdateProfileBalanceTransactionInput struct {
	BalanceTransaction string `json:"balance_transaction"`
}
