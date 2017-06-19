package management

import (
	"fmt"
	core "github.com/stone-pagamentos/globalidentity-go"
)

type GlobalIdentityManager interface {
	UserRoles(email string) ([]core.Role, error)
}

type globalIdentityManager struct {
	applicationKey     string
	apiKey             string
	globalIdentityHost string
	requester 		   core.Requester
}

func New(applicationKey string, apiKey string, globalIdentityHost string) GlobalIdentityManager {
	return &globalIdentityManager{
		applicationKey,
		apiKey,
		globalIdentityHost,
		core.NewRequester(),
	}
}

func (gim *globalIdentityManager) UserRoles(email string) ([]core.Role, error) {

	url := fmt.Sprintf(gim.globalIdentityHost+listUserRoles, gim.applicationKey, email)

	resp, err := gim.requester.Get(url, gim.requestOptions())

	if err != nil {
		return nil, err
	}

	response := new(rolesResponse)
	if err = resp.JSON(&response); err != nil {
		return nil, err
	}

	if err = response.Validate(); err != nil {
		return nil, err
	}

	roles := make([]core.Role, len(response.Roles))

	for i, role := range response.Roles {
		roles[i] = core.Role{
			Name:        role.RoleName,
			Description: role.Description,
			Active:      role.Active,
		}
	}

	return roles, nil
}

func (gim *globalIdentityManager) requestOptions() *core.RequestOptions {
	ro := new(core.RequestOptions) 
	ro.Headers = map[string]string{
			"Accept": contentJSON,
			"Authorization": "bearer " + gim.apiKey,
			"Content-Type":  contentJSON,
		}
	return ro
}
