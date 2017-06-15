package globalidentity

import (
	"testing"
	"github.com/levigross/grequests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ResponseProcessorSuite struct {
	suite.Suite
	badStatusResponse *grequests.Response
}

func TestUserRoles(t *testing.T) {
	suite.Run(t, new(ResponseProcessorSuite))
}

func (suite *ResponseProcessorSuite) SetupTest() {
	suite.badStatusResponse = &grequests.Response{StatusCode: 400}
}

func (suite *ResponseProcessorSuite) TestBadStatus() {

	response := new(Response)

	err := ResponseProcessor.Process(suite.badStatusResponse, response)

	assert.NotNil(suite.T(), err)
}
