package account

// Account represents the format of the data that will be sent/received
type Account struct {
	Id      string `json:"id"`
	Status  string `json:"status"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Phone   string `json:"phone"`
	Gender  string `json:"gender"`
	Address string `json:"address"`
}

// Login represents the format of the data used for authentication in the API.
type Login struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}
