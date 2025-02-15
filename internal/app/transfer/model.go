package transfer

type MakeTransferRequest struct {
	TargetUser string  `json:"target_user" binding:"required"`
	Amount     float64 `json:"amount" binding:"required"`
	Remarks    string  `json:"remarks" binding:"required"`
	UserId     string  `json:"user_id"`
}

type MakeTransferResponse struct {
	TransferId    string  `json:"transfer_id"`
	Amount        float64 `json:"amount"`
	Remarks       string  `json:"remarks"`
	BalanceBefore float64 `json:"balance_before"`
	BalanceAfter  float64 `json:"balance_after"`
	CreatedAt     string  `json:"created_at"`
}
