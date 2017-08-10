package authorization

import core "github.com/stone-payments/globalidentity-go"

type authenticateUserResponse struct {
	AuthenticationToken      string                 `json:"AuthenticationToken"`
	TokenExpirationInMinutes int                    `json:"TokenExpirationInMinutes"`
	UserKey                  string                 `json:"UserKey"`
	Name                     string                 `json:"Name"`
	Success                  bool                   `json:"Success"`
	OperationReport          []loginOperationReport `json:"OperationReport"`
}

type loginOperationReport struct {
	Field   string `json:"Field"`
	Message string `json:"Message"`
}

type renewTokenResponse struct {
	NewToken            string `json:"NewToken"`
	ExpirationInMinutes int    `json:"ExpirationInMinutes"`
	core.Response
}
