package payment

type MakePaymentRequest struct {
	Amount  float64 `json:"amount" binding:"required,min=1"`
	Remarks string  `json:"remarks" binding:"required"`
	UserID  string  `json:"user_id"`
}

type MakePaymentResponse struct {
	PaymentId     string  `json:"payment_id"`
	Amount        float64 `json:"amount" binding:"required,min=1"`
	Remarks       string  `json:"remarks" binding:"required"`
	BalanceBefore float64 `json:"balance_before"`
	BalanceAfter  float64 `json:"balance_after"`
	CreatedAt     string  `json:"created_at"`
}
