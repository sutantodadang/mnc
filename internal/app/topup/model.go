package topup

type AddTopup struct {
	Amount float64 `json:"amount" binding:"required,min=5000"`
	Userid string  `json:"userid"`
}

type AddTopupResponse struct {
	TopUpId       string  `json:"top_up_id"`
	AmountTopup   float64 `json:"amount_topup"`
	BalanceBefore float64 `json:"balance_before"`
	BalanceAfter  float64 `json:"balance_after"`
	CreatedAt     string  `json:"created_at"`
}
