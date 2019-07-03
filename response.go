package globalidentity

// Response is the base response of Global Identity.
type Response struct {
	Success         bool `json:"Success"`
	OperationReport []struct {
		Field     string `json:"Field"`
		Message   string `json:"Message"`
		ErrorCode int    `json:"ErrorCode"`
	} `json:"OperationReport,omitempty"`
}

// Validate checks success of response.
func (r *Response) Validate() error {
	if r.Success != true {
		el := make([]string, 0)
		for _, e := range r.OperationReport {
			el = append(el, e.Message)
		}
		return GlobalIdentityError(el)
	}

	return nil
}
