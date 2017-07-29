package globalidentity

import (
	"github.com/levigross/grequests"
	"fmt"
)

type HttpResponse struct {
	*grequests.Response
}

type Requester interface {
	Post(url string, requestOptions *RequestOptions) (*HttpResponse, error)
	Get(url string, requestOptions *RequestOptions) (*HttpResponse, error)
}

type requester struct{}

type RequestOptions struct {
	grequests.RequestOptions
}

func (r requester) Post(url string, ro *RequestOptions) (*HttpResponse, error) {
	response, err := grequests.Post(url, &ro.RequestOptions)

	if err == nil {
		err = r.processResponse(&HttpResponse{Response: response})
	}

	return &HttpResponse{Response: response}, err
}
func (r requester) Get(url string, ro *RequestOptions) (*HttpResponse, error) {
	response, err := grequests.Get(url, &ro.RequestOptions)

	if err == nil {
		err = r.processResponse(&HttpResponse{Response: response})
	}

	return &HttpResponse{Response: response}, err
}

func (r *requester) processResponse(resp *HttpResponse) error {

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return GlobalIdentityError([]string{fmt.Sprintf("%d", resp.StatusCode)})
	}

	return nil
}

func NewRequester() Requester {
	return requester{}
}
