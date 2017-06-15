package globalidentity

type Response struct {
	Success         bool     `json:"Success"`
	OperationReport []string `json:"OperationReport"`
}
