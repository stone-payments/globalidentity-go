package globalidentity

import "fmt"

// Response is the base response of Global Identity.
type Response struct {
	Success         bool     `json:"Success"`
	OperationReport []string `json:"OperationReport"`
}

// Validate checks success of response.
func (r *Response) Validate() error {
	if r.Success != true {
		fmt.Printf("%#v \n", r)
		return GlobalIdentityError(r.OperationReport)
	}

	return nil
}
