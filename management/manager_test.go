package management

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserRolesSuite struct {
	suite.Suite
	url             string
	manager         GlobalIdentityManager
	errorResponder  httpmock.Responder
	okResponder     httpmock.Responder
	failedResponder httpmock.Responder
	wrongResponder  httpmock.Responder
}

func TestUserRoles(t *testing.T) {
	suite.Run(t, new(UserRolesSuite))
}

func (suite *UserRolesSuite) SetupTest() {
	suite.manager = New("key", "key", "http://url")
	suite.url = "http://url/api/management/key/users/user/roles"

	suite.okResponder = httpmock.NewStringResponder(http.StatusOK, `{"Success": true, "OperationReport": [], "roles":[{"statusName":"mock"}]}`)
	suite.errorResponder = httpmock.NewStringResponder(http.StatusInternalServerError, "")
	suite.wrongResponder = httpmock.NewStringResponder(http.StatusOK, "mock")
	suite.failedResponder = httpmock.NewStringResponder(http.StatusOK, `{"Success": false, "OperationReport": []}`)
}

func (suite *UserRolesSuite) TestOk() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", suite.url, suite.okResponder)

	roles, err := suite.manager.UserRoles("user")

	assert.True(suite.T(), len(roles) > 0)
	assert.Nil(suite.T(), err)
}

func (suite *UserRolesSuite) TestWrongResponse() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", suite.url, suite.wrongResponder)

	roles, err := suite.manager.UserRoles("user")

	assert.Nil(suite.T(), roles)
	assert.NotNil(suite.T(), err)
}

func (suite *UserRolesSuite) TestErrorResponse() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", suite.url, suite.errorResponder)

	roles, err := suite.manager.UserRoles("user")

	assert.Nil(suite.T(), roles)
	assert.NotNil(suite.T(), err)
}

func (suite *UserRolesSuite) TestFailedResponse() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", suite.url, suite.failedResponder)

	roles, err := suite.manager.UserRoles("user")

	assert.Nil(suite.T(), roles)
	assert.NotNil(suite.T(), err)
}
