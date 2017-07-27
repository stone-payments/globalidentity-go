package authorization

import (
	core "github.com/stone-payments/globalidentity-go"
)

type GlobalIdentityManager interface {
	AuthenticateUser(email string, password string, expirationInMinutes ...int) (*core.Authorization, error)
	ValidateToken(token string) (bool, error)
	IsUserInRoles(userKey string, roles ...string) (bool, error)
	RenewToken(token string) (string, error)
	ValidateApplication(clientApplicationKey string, rawData string, encryptedData string) (bool, error)
	RecoverPassword(email string) (bool, error)
}

type globalIdentityManager struct {
	applicationKey     string
	globalIdentityHost string
	requester          core.Requester
}

func New(applicationKey string, globalIdentityHost string) GlobalIdentityManager {
	return &globalIdentityManager{
		applicationKey,
		globalIdentityHost,
		core.NewRequester(),
	}
}

func (gim *globalIdentityManager) AuthenticateUser(email string, password string, expirationInMinutes ...int) (*core.Authorization, error) {
	expirationInMinutes = append(expirationInMinutes, 15)
	request := &authenticateUserRequest{
		ApplicationKey:           gim.applicationKey,
		TokenExpirationInMinutes: expirationInMinutes[0],
		Email:                    email,
		Password:                 password,
	}

	requestOptions := gim.requestOptions()
	requestOptions.JSON = request

	resp, err := gim.requester.Post(gim.globalIdentityHost+authenticateUserSuffix, requestOptions)

	if err != nil {
		return nil, err
	}

	var response authenticateUserResponse
	err = resp.JSON(&response)

	if err != nil {
		return nil, err
	}

	var globalIdentityUser *core.Authorization
	if response.Success {
		globalIdentityUser = &core.Authorization{
			Token: response.AuthenticationToken,
			Key:   response.UserKey,
		}
	} else {
		var messages []string
		for _, operationReport := range response.OperationReport {
			messages = append(messages, operationReport.Message)
		}
		err = core.GlobalIdentityError(messages)
	}

	return globalIdentityUser, err
}

func (gim *globalIdentityManager) RecoverPassword(email string) (bool, error) {
	request := recoverPasswordRequest{
		ApplicationKey: gim.applicationKey,
		Email:          email,
	}

	requestOptions := gim.requestOptions()
	requestOptions.JSON = request

	resp, err := gim.requester.Post(gim.globalIdentityHost+recoverPasswordSuffix, requestOptions)

	if err != nil {
		return false, err
	}

	var response core.Response
	if err = resp.JSON(&response); err != nil {
		return false, err
	}

	if err = response.Validate(); err != nil {
		return false, err
	}

	return response.Success, err
}

func (gim *globalIdentityManager) ValidateToken(token string) (bool, error) {
	request := &validateTokenRequest{
		ApplicationKey: gim.applicationKey,
		Token:          token,
	}

	requestOptions := gim.requestOptions()
	requestOptions.JSON = request

	resp, err := gim.requester.Post(gim.globalIdentityHost+validateTokenSuffix, requestOptions)
	if err != nil {
		return false, err
	}

	var response core.Response
	if err = resp.JSON(&response); err != nil {
		return false, err
	}

	if err = response.Validate(); err != nil {
		return false, err
	}

	return response.Success, err
}

func (gim *globalIdentityManager) IsUserInRoles(userKey string, roles ...string) (bool, error) {
	request := &isUserInHolesRequest{
		ApplicationKey: gim.applicationKey,
		UserKey:        userKey,
		RoleCollection: roles,
	}

	requestOptions := gim.requestOptions()
	requestOptions.JSON = request

	resp, err := gim.requester.Post(gim.globalIdentityHost+isUserInRolesSuffix, requestOptions)
	if err != nil {
		return false, err
	}

	var response core.Response
	if err = resp.JSON(&response); err != nil {
		return false, err
	}

	if err = response.Validate(); err != nil {
		return false, err
	}

	return response.Success, err
}

func (gim *globalIdentityManager) RenewToken(token string) (string, error) {
	request := &renewTokenRequest{
		ApplicationKey: gim.applicationKey,
		Token:          token,
	}
	requestOptions := gim.requestOptions()
	requestOptions.JSON = request

	resp, err := gim.requester.Post(gim.globalIdentityHost+renewTokenSuffix, requestOptions)
	if err != nil {
		return "", err
	}

	var response renewTokenResponse
	if err = resp.JSON(&response); err != nil {
		return "", err
	}

	if err = response.Validate(); err != nil {
		return "", err
	}

	return response.NewToken, err
}

func (gim *globalIdentityManager) ValidateApplication(clientApplicationKey string, rawData string, encryptedData string) (bool, error) {

	request := &validateApplicationRequest{
		ApplicationKey:       gim.applicationKey,
		ClientApplicationKey: clientApplicationKey,
		RawData:              rawData,
		EncryptedData:        encryptedData,
	}
	requestOptions := gim.requestOptions()
	requestOptions.JSON = request

	resp, err := gim.requester.Post(gim.globalIdentityHost+validateApplicationSuffix, requestOptions)
	if err != nil {
		return false, err
	}

	var response core.Response
	if err = resp.JSON(&response); err != nil {
		return false, err
	}

	if err = response.Validate(); err != nil {
		return false, err
	}

	return response.Success, err
}

func (gim *globalIdentityManager) requestOptions() *core.RequestOptions {
	ro := new(core.RequestOptions)
	ro.Headers = map[string]string{
		"Accept":       contentJson,
		"Content-Type": contentJson,
	}
	return ro
}
