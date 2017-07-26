package management

import core "github.com/stone-payments/globalidentity-go"

type rolesResponse struct {
	Roles []role `json:"roles"`
	*core.Response
}

type role struct {
	RoleName    string `json:"roleName"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}

type listUsersResponse struct {
	Users     []core.User `json:"users"`
	FirstPage int `json:"FirstPage"`
	NextPage  int `json:"NextPage"`
	LastPage  int `json:"LastPage"`
	TotalRows int `json:"TotalRows"`
	*core.Response
}
