package model

// Account Modelo de dados de conta
type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	ServedBy string `json:"serverdBy"`
}
