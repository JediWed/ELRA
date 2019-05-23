package structs

// LoginResponse is a structure for response to a successful login
type LoginResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Role      int    `json:"role"`
	Token     string `json:"token"`
}

type InvoiceResponse struct {
	PaymentRequest string `json:"payment_request"`
}
