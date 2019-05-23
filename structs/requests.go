package structs

// LoginRequest is a structure for login requests
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UpdatePasswordRequest is a structure for change password requests
type UpdatePasswordRequest struct {
	Password string `json:"password"`
	UserID   int
}

// UpdateUsernameRequest is a structure for change username requests
type UpdateUsernameRequest struct {
	Username string `json:"username"`
	UserID   int
}

// InvoiceRequest is a structure for creating an invoice
type InvoiceRequest struct {
	Amount int64  `json:"amount"`
	Memo   string `json:"memo"`
}
