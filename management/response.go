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
