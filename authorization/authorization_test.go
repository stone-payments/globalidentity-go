package authorization

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const (
	globalApplicationUrl   = "https://dlpgi.dlp-payments.com"
	validateApplicationUrl = "https://dlpgi.dlp-payments.com/api/authorization/validateapplication"
	authenticateUserUrl    = "https://dlpgi.dlp-payments.com/api/authorization/authenticate"
	isUserInRolesUrl       = "https://dlpgi.dlp-payments.com/api/authorization/isuserinroles"
	validateTokenUrl       = "https://dlpgi.dlp-payments.com/api/authorization/validatetokenresponse"
	renewTokenUrl          = "https://dlpgi.dlp-payments.com/api/authorization/renewtoken"
	recoverPasswordUrl     = "https://dlpgi.dlp-payments.com/api/authorization/recoverPassword"
)

func TestRecoverPasswordWrongStatusCode(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", recoverPasswordUrl, httpmock.NewStringResponder(http.StatusInternalServerError, ""))

	gim := New("test", globalApplicationUrl)
	_, err := gim.RecoverPassword("test@test.com.br")
	assert.NotNil(t, err)
}

func TestRecoverPasswordWrongResponse(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", recoverPasswordUrl, httpmock.NewStringResponder(http.StatusOK, "mock"))

	gim := New("test", globalApplicationUrl)
	_, err := gim.RecoverPassword("test@test.com.br")
	assert.NotNil(t, err)
}

func TestRecoverPasswordWithOperationReport(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", recoverPasswordUrl, httpmock.NewStringResponder(http.StatusOK, `{"Success": false, "OperationReport": ["test", "mock"]}`))

	gim := New("test", globalApplicationUrl)
	_, err := gim.RecoverPassword("test@test.com.br")
	assert.NotNil(t, err)
}

func TestRecoverPasswordOk(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", recoverPasswordUrl, httpmock.NewStringResponder(http.StatusOK, `{"Success": true, "OperationReport": []}`))

	gim := New("test", globalApplicationUrl)
	ok, err := gim.RecoverPassword("test@test.com.br")
	assert.True(t, ok)
	assert.Nil(t, err)
}

func TestGlobalIdentityManager_ValidateApplication(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", validateApplicationUrl, httpmock.NewStringResponder(http.StatusInternalServerError, ""))

	gim := New("test", globalApplicationUrl)
	_, err := gim.ValidateApplication("", "", "")
	if err == nil {
		t.FailNow()
	}

	okResponse, _ := json.Marshal(&validateApplicationResponse{
		Success:         true,
		OperationReport: make([]string, 0),
	})

	httpmock.RegisterResponder("POST", validateApplicationUrl, httpmock.NewStringResponder(http.StatusOK, string(okResponse)))

	gim = New("test", globalApplicationUrl)
	ok, err := gim.ValidateApplication("", "", "")
	if !ok || err != nil {
		t.FailNow()
	}

	notOkResponse, _ := json.Marshal(&validateApplicationResponse{
		Success:         false,
		OperationReport: []string{"error"},
	})

	httpmock.RegisterResponder("POST", validateApplicationUrl, httpmock.NewStringResponder(http.StatusOK, string(notOkResponse)))

	ok, err = gim.ValidateApplication("", "", "")
	if ok {
		t.FailNow()
	}
	giErr := err.(GlobalIdentityError)
	if len(giErr) != 1 || giErr[0] != "error" {
		t.FailNow()
	}

	httpmock.RegisterResponder("POST", validateApplicationUrl, httpmock.NewStringResponder(http.StatusOK, "{\"saa}"))

	_, err = gim.ValidateApplication("", "", "")

	if err == nil {
		t.FailNow()
	}
}

func TestGlobalIdentityManager_AuthenticateUser(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", authenticateUserUrl, httpmock.NewStringResponder(http.StatusInternalServerError, ""))

	gim := New("test", globalApplicationUrl)
	_, err := gim.AuthenticateUser("", "", 1)
	if err == nil {
		t.FailNow()
	}

	okResponse, _ := json.Marshal(&authenticateUserResponse{
		Success:                  true,
		AuthenticationToken:      "banana",
		TokenExpirationInMinutes: 1,
		UserKey:                  "user",
		Name:                     "user",
	})

	httpmock.RegisterResponder("POST", authenticateUserUrl, httpmock.NewStringResponder(http.StatusOK, string(okResponse)))

	gim = New("test", globalApplicationUrl)
	_, err = gim.AuthenticateUser("", "", 1)
	if err != nil {
		t.FailNow()
	}

	oprep := []loginOperationReport{
		{Message: "error1", Field: "login"},
		{Message: "error2", Field: "login"},
	}
	notOkResponse, _ := json.Marshal(&authenticateUserResponse{
		Success:                  false,
		AuthenticationToken:      "banana",
		TokenExpirationInMinutes: 1,
		UserKey:                  "user",
		Name:                     "user",
		OperationReport:          oprep,
	})

	httpmock.RegisterResponder("POST", authenticateUserUrl, httpmock.NewStringResponder(http.StatusOK, string(notOkResponse)))

	_, err = gim.AuthenticateUser("", "")
	if err.Error() != `[]string{"error1", "error2"}` {
		t.FailNow()
	}

	httpmock.RegisterResponder("POST", authenticateUserUrl, httpmock.NewStringResponder(http.StatusOK, "{\"saa}"))

	_, err = gim.AuthenticateUser("", "")

	if err == nil {
		t.FailNow()
	}
}

func TestGlobalIdentityManager_IsUserInRoles(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", isUserInRolesUrl, httpmock.NewStringResponder(http.StatusInternalServerError, ""))

	gim := New("test", globalApplicationUrl)
	_, err := gim.IsUserInRoles("", "", "")
	if err == nil {
		t.FailNow()
	}

	okResponse, _ := json.Marshal(&isUserInRoleResponse{
		Success:         true,
		OperationReport: make([]string, 0),
	})

	httpmock.RegisterResponder("POST", isUserInRolesUrl, httpmock.NewStringResponder(http.StatusOK, string(okResponse)))

	gim = New("test", globalApplicationUrl)
	ok, err := gim.IsUserInRoles("", "")
	if !ok || err != nil {
		t.FailNow()
	}

	notOkResponse, _ := json.Marshal(&isUserInRoleResponse{
		Success:         false,
		OperationReport: []string{"error"},
	})

	httpmock.RegisterResponder("POST", isUserInRolesUrl, httpmock.NewStringResponder(http.StatusOK, string(notOkResponse)))

	ok, err = gim.IsUserInRoles("", "")
	if ok {
		t.FailNow()
	}
	giErr := err.(GlobalIdentityError)
	if len(giErr) != 1 || giErr[0] != "error" {
		t.FailNow()
	}

	httpmock.RegisterResponder("POST", isUserInRolesUrl, httpmock.NewStringResponder(http.StatusOK, "{\"saa}"))

	_, err = gim.IsUserInRoles("", "")

	if err == nil {
		t.FailNow()
	}
}

func TestGlobalIdentityManager_ValidateToken(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", validateTokenUrl, httpmock.NewStringResponder(http.StatusInternalServerError, ""))

	gim := New("test", globalApplicationUrl)
	_, err := gim.ValidateToken("")
	if err == nil {
		t.FailNow()
	}

	okResponse, _ := json.Marshal(&validateTokenResponse{
		Success:         true,
		OperationReport: make([]string, 0),
	})

	httpmock.RegisterResponder("POST", validateTokenUrl, httpmock.NewStringResponder(http.StatusOK, string(okResponse)))

	gim = New("test", globalApplicationUrl)
	ok, err := gim.ValidateToken("")
	if !ok || err != nil {
		t.FailNow()
	}

	notOkResponse, _ := json.Marshal(&validateTokenResponse{
		Success:         false,
		OperationReport: []string{"error"},
	})

	httpmock.RegisterResponder("POST", validateTokenUrl, httpmock.NewStringResponder(http.StatusOK, string(notOkResponse)))

	ok, err = gim.ValidateToken("")
	if ok {
		t.FailNow()
	}
	giErr := err.(GlobalIdentityError)
	if len(giErr) != 1 || giErr[0] != "error" {
		t.FailNow()
	}

	httpmock.RegisterResponder("POST", validateTokenUrl, httpmock.NewStringResponder(http.StatusOK, "{\"saa}"))

	_, err = gim.ValidateToken("")

	if err == nil {
		t.FailNow()
	}
}

func TestGlobalIdentityManager_RenewToken(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", renewTokenUrl, httpmock.NewStringResponder(http.StatusInternalServerError, ""))

	gim := New("test", globalApplicationUrl)
	_, err := gim.RenewToken("")
	if err == nil {
		t.FailNow()
	}

	okResponse, _ := json.Marshal(&validateTokenResponse{
		Success:         true,
		OperationReport: make([]string, 0),
	})

	httpmock.RegisterResponder("POST", renewTokenUrl, httpmock.NewStringResponder(http.StatusOK, string(okResponse)))

	gim = New("test", globalApplicationUrl)
	_, err = gim.RenewToken("")
	if err != nil {
		t.FailNow()
	}

	notOkResponse, _ := json.Marshal(&renewTokenResponse{
		Success:         false,
		NewToken:        "token",
		OperationReport: []string{"error"},
	})

	httpmock.RegisterResponder("POST", renewTokenUrl, httpmock.NewStringResponder(http.StatusOK, string(notOkResponse)))

	_, err = gim.RenewToken("")

	if err == nil {
		t.FailNow()
	}

	giErr := err.(GlobalIdentityError)
	if len(giErr) != 1 || giErr[0] != "error" {
		t.FailNow()
	}

	httpmock.RegisterResponder("POST", renewTokenUrl, httpmock.NewStringResponder(http.StatusOK, "{\"saa}"))

	_, err = gim.RenewToken("")

	if err == nil {
		t.FailNow()
	}
}

func TestGlobalIdentityError_Error(t *testing.T) {
	err := GlobalIdentityError([]string{"error01", "error01"})
	if err.Error() != `[]string{"error01", "error01"}` {
		t.FailNow()
	}
}
