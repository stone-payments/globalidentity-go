package globalidentity

type authenticateUserResponse struct {
	AuthenticationToken      string `json:"AuthenticationToken"`
	TokenExpirationInMinutes int    `json:"TokenExpirationInMinutes"`
	UserKey                  string `json:"UserKey"`
	Name                     string `json:"Name"`
	Success                  bool   `json:"Success"`
}

type validateTokenResponse struct {
	ExpirationInMinutes int      `json:"ExpirationInMinutes"`
	Success             bool     `json:"Success"`
	OperationReport     []string `json:"OperationReport"`
}

type recoverPasswordResponse struct {
	Success         bool     `json:"Success"`
	OperationReport []string `json:"OperationReport"`
}

type renewTokenResponse struct {
	NewToken            string   `json:"NewToken"`
	ExpirationInMinutes int      `json:"ExpirationInMinutes"`
	Success             bool     `json:"Success"`
	OperationReport     []string `json:"OperationReport"`
}

type renewToken struct {
	NewToken            string   `json:"NewToken"`
	ExpirationInMinutes int      `json:"ExpirationInMinutes"`
	Success             bool     `json:"Success"`
	OperationReport     []string `json:"OperationReport"`
}

type validateApplicationResponse struct {
	Success         bool     `json:"Success"`
	OperationReport []string `json:"OperationReport"`
}

type isUserInRoleResponse struct {
	Success         bool     `json:"Success"`
	OperationReport []string `json:"OperationReport"`
}
