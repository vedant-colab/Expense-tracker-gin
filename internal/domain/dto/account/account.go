package account

type CreateAccountRequest struct {
	Name    string  `json:"name" binding:"required"`
	Type    string  `json:"type" binding:"required"`
	Balance float64 `json:"balance"`
}

type UpdateAccountRequest struct {
	Name    string  `json:"name"`
	Type    string  `json:"type"`
	Balance float64 `json:"balance"`
}
