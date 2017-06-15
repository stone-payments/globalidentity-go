package globalidentity

import (
	"github.com/levigross/grequests"
	"fmt"
)

type ResponseProcessable interface {
	Process(resp *grequests.Response, data interface{}) error
}

type responseProcessor struct{}

func (r *responseProcessor) Process(resp *grequests.Response, data interface{}) error {

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return GlobalIdentityError([]string{fmt.Sprintf("%v", resp.StatusCode)})
	}

	resp.JSON(data)
	return nil
}
