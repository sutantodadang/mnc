package transaction

type TransactionReportResponse struct {
	TransferId      string  `json:"transfer_id,omitempty"`
	PaymentId       string  `json:"payment_id,omitempty"`
	TopUpId         string  `json:"topup_id,omitempty"`
	Status          string  `json:"status"`
	UserId          string  `json:"user_id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
	Remarks         string  `json:"remarks"`
	BalanceBefore   float64 `json:"balance_before"`
	BalanceAfter    float64 `json:"balance_after"`
	CreatedAt       string  `json:"created_at"`
}
