package management

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	core "github.com/stone-payments/globalidentity-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/fortytw2/leaktest"
)

type ManagementSuite struct {
	suite.Suite
	userRolesUrl         string
	listUsersUrl         string
	manager              GlobalIdentityManager
	errorResponder       httpmock.Responder
	okUserRolesResponder httpmock.Responder
	okListUsersResponder httpmock.Responder
	failedResponder      httpmock.Responder
	wrongResponder       httpmock.Responder
}

func TestManager(t *testing.T) {
	suite.Run(t, new(ManagementSuite))
}

func (suite *ManagementSuite) SetupTest() {
	suite.manager = New("key", "key", "http://userRolesUrl")
	suite.userRolesUrl = "http://userRolesUrl/api/management/key/users/user/roles"
	suite.listUsersUrl = "http://userRolesUrl/api/management/key/users?page=1&limit=1&includeRoles=true"

	suite.okUserRolesResponder = httpmock.NewStringResponder(http.StatusOK, `{"Success": true, "OperationReport": [], "roles":[{"statusName":"mock"}]}`)
	suite.okListUsersResponder = httpmock.NewStringResponder(http.StatusOK, `{ "OperationReport": [], "Success": true, "users": [ { "active": true, "comment": "Comments about the user in the application", "email": "user1@email.com", "lockedOut": false, "name": "User's name", "roles": [ "ADMIN" ], "userKey": "00000000-0000-0000-0000-000000000000" } ] }`)
	suite.errorResponder = httpmock.NewStringResponder(http.StatusInternalServerError, "")
	suite.wrongResponder = httpmock.NewStringResponder(http.StatusOK, "mock")
	suite.failedResponder = httpmock.NewStringResponder(http.StatusOK, `{"Success": false, "OperationReport": []}`)
}

func (suite *ManagementSuite) TestUserRolesOk() {
	defer leaktest.Check(suite.T())()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()


	httpmock.RegisterResponder("GET", suite.userRolesUrl, suite.okUserRolesResponder)

	roles, err := suite.manager.UserRoles("user")

	assert.True(suite.T(), len(roles) > 0)
	assert.Nil(suite.T(), err)
}

func (suite *ManagementSuite) TestUserRolesWrongResponse() {
	defer leaktest.Check(suite.T())()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", suite.userRolesUrl, suite.wrongResponder)

	roles, err := suite.manager.UserRoles("user")

	assert.Nil(suite.T(), roles)
	assert.NotNil(suite.T(), err)
}

func (suite *ManagementSuite) TestUserRolesErrorResponse() {
	defer leaktest.Check(suite.T())()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", suite.userRolesUrl, suite.errorResponder)

	roles, err := suite.manager.UserRoles("user")

	_, ok := err.(core.GlobalIdentityError)

	assert.Nil(suite.T(), roles)
	assert.True(suite.T(), ok)
}

func (suite *ManagementSuite) TestUserRolesFailedResponse() {
	defer leaktest.Check(suite.T())()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", suite.userRolesUrl, suite.failedResponder)

	roles, err := suite.manager.UserRoles("user")

	_, ok := err.(core.GlobalIdentityError)

	assert.Nil(suite.T(), roles)
	assert.True(suite.T(), ok)
}

func (suite *ManagementSuite) TestListUsersOk() {
	defer leaktest.Check(suite.T())()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", suite.listUsersUrl, suite.okListUsersResponder)

	users, err := suite.manager.ListUsers(1, 1, true)

	assert.True(suite.T(), len(users.Users) > 0)
	assert.Nil(suite.T(), err)
}

func (suite *ManagementSuite) TestListUsersWrongResponse() {
	defer leaktest.Check(suite.T())()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", suite.listUsersUrl, suite.wrongResponder)

	users, err := suite.manager.ListUsers(1, 1, true)

	assert.Nil(suite.T(), users)
	assert.NotNil(suite.T(), err)
}

func (suite *ManagementSuite) TestListUsersErrorResponse() {
	defer leaktest.Check(suite.T())()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", suite.listUsersUrl, suite.errorResponder)

	users, err := suite.manager.ListUsers(1, 1, true)

	_, ok := err.(core.GlobalIdentityError)

	assert.Nil(suite.T(), users)
	assert.True(suite.T(), ok)
}

func (suite *ManagementSuite) TestListUsersFailedResponse() {
	defer leaktest.Check(suite.T())()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", suite.listUsersUrl, suite.failedResponder)

	users, err := suite.manager.ListUsers(1, 1, true)

	_, ok := err.(core.GlobalIdentityError)

	assert.Nil(suite.T(), users)
	assert.True(suite.T(), ok)
}
