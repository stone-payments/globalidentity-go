package authorization

import (
	"fmt"
	core "github.com/stone-pagamentos/globalidentity-go"
	"net/http"
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
}

func New(applicationKey string, globalIdentityHost string) GlobalIdentityManager {
	return &globalIdentityManager{
		applicationKey,
		globalIdentityHost,
	}
}

func (gim *globalIdentityManager) AuthenticateUser(email string, password string, expirationInMinutes ...int) (*core.Authorization, error) {
	expirationInMinutes = append(expirationInMinutes, 15)
	request := &authenticateUserRequest{
		ApplicationKey:           gim.applicationKey,
		TokenExpirationInMinutes: expirationInMinutes[0],
		Email:    email,
		Password: password,
	}
	json, err := core.ToJson(request)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(gim.globalIdentityHost+authenticateUserSuffix, contentJson, json)
	if err != nil {
		return nil, err
	}

	var response authenticateUserResponse

	err = core.FromJson(&response, resp.Body)
	if err != nil {
		return nil, err
	}
	if !response.Success {
		var messages []string
		for _, operationReport := range response.OperationReport {
			messages = append(messages, operationReport.Message)
		}
		err = core.GlobalIdentityError(messages)
	}

	var globalIdentityUser core.Authorization
	globalIdentityUser.Token = response.AuthenticationToken
	globalIdentityUser.Key = response.UserKey
	return &globalIdentityUser, err
}

func (gim *globalIdentityManager) RecoverPassword(email string) (bool, error) {
	request := recoverPasswordRequest{
		ApplicationKey: gim.applicationKey,
		Email:          email,
	}

	json, err := core.ToJson(request)
	if err != nil {
		return false, err
	}

	resp, err := http.Post(gim.globalIdentityHost+recoverPasswordSuffix, contentJson, json)
	if err != nil {
		return false, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return false, core.GlobalIdentityError([]string{fmt.Sprintf("%v", resp.StatusCode)})
	}

	var response recoverPasswordResponse

	err = core.FromJson(&response, resp.Body)
	if err != nil {
		return false, err
	}

	if len(response.OperationReport) > 0 {
		err = core.GlobalIdentityError(response.OperationReport)
	}

	return response.Success, err
}

func (gim *globalIdentityManager) ValidateToken(token string) (bool, error) {
	request := &validateTokenRequest{
		ApplicationKey: gim.applicationKey,
		Token:          token,
	}
	json, err := core.ToJson(request)
	if err != nil {
		return false, err
	}

	resp, err := http.Post(gim.globalIdentityHost+validateTokenSuffix, contentJson, json)
	if err != nil {
		return false, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return false, core.GlobalIdentityError([]string{fmt.Sprintf("%v", resp.StatusCode)})
	}

	var response validateTokenResponse

	err = core.FromJson(&response, resp.Body)
	if err != nil {
		return false, err
	}
	if !response.Success {
		err = core.GlobalIdentityError(response.OperationReport)
	}
	return response.Success, err
}

func (gim *globalIdentityManager) IsUserInRoles(userKey string, roles ...string) (bool, error) {
	request := &isUserInHolesRequest{
		ApplicationKey: gim.applicationKey,
		UserKey:        userKey,
		RoleCollection: roles,
	}
	json, err := core.ToJson(request)
	if err != nil {
		return false, err
	}

	resp, err := http.Post(gim.globalIdentityHost+isUserInRolesSuffix, contentJson, json)
	if err != nil {
		return false, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return false, core.GlobalIdentityError([]string{fmt.Sprintf("%v", resp.StatusCode)})
	}

	var response isUserInRoleResponse

	err = core.FromJson(&response, resp.Body)
	if err != nil {
		return false, err
	}
	if !response.Success {
		err = core.GlobalIdentityError(response.OperationReport)
	}
	return response.Success, err
}

func (gim *globalIdentityManager) RenewToken(token string) (string, error) {
	request := &renewTokenRequest{
		ApplicationKey: gim.applicationKey,
		Token:          token,
	}
	json, err := core.ToJson(request)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(gim.globalIdentityHost+renewTokenSuffix, contentJson, json)
	if err != nil {
		return "", err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", core.GlobalIdentityError([]string{fmt.Sprintf("%v", resp.StatusCode)})
	}

	var response renewTokenResponse

	err = core.FromJson(&response, resp.Body)
	if err != nil {
		return "", err
	}
	if !response.Success {
		err = core.GlobalIdentityError(response.OperationReport)
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
	json, err := core.ToJson(request)
	if err != nil {
		return false, err
	}

	resp, err := http.Post(gim.globalIdentityHost+validateApplicationSuffix, contentJson, json)
	if err != nil {
		return false, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return false, core.GlobalIdentityError([]string{fmt.Sprintf("%v", resp.StatusCode)})
	}

	var response validateApplicationResponse

	err = core.FromJson(&response, resp.Body)
	if err != nil {
		return false, err
	}
	if !response.Success {
		err = core.GlobalIdentityError(response.OperationReport)
	}
	return response.Success, err
}
