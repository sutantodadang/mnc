package kafka

type TransferMessage struct {
	TransferID    string  `json:"transfer_id"`
	TransactionID string  `json:"transaction_id"`
	Amount        float64 `json:"amount"`
	From          string  `json:"from"`
	To            string  `json:"to"`
	Status        string  `json:"status"`
	Remarks       string  `json:"remarks"`
}
