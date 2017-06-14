package management

type response struct {
	Success         bool          `json:"Success"`
	OperationReport []interface{} `json:"OperationReport"`
}

type rolesResponse struct {
	Roles []role `json:"roles"`
	response
}

type role struct {
	RoleName    string `json:"roleName"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}
