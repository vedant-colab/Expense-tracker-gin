package recurring

type CreateRecurringRequest struct {
	Amount   float64 `json:"amount" binding:"required"`
	Interval string  `json:"interval" binding:"required"`
	Category string  `json:"category" binding:"required"`
}

type UpdateRecurringRequest struct {
	Amount   float64 `json:"amount"`
	Interval string  `json:"interval"`
	Category string  `json:"category"`
}
