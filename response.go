package globalidentity

// Response is the base response of Global Identity.
type Response struct {
	Success         bool     `json:"Success"`
	OperationReport []string `json:"OperationReport"`
}

// Validate checks success of response.
func (r *Response) Validate() error {
	if r.Success != true {
		return GlobalIdentityError(r.OperationReport)
	}

	return nil
}