package authorization

type authenticateUserRequest struct {
	ApplicationKey           string `json:"ApplicationKey"`
	Email                    string `json:"Email"`
	Password                 string `json:"Password"`
	TokenExpirationInMinutes int    `json:"TokenExpirationInMinutes"`
}

type recoverPasswordRequest struct {
	ApplicationKey string `json:"ApplicationKey"`
	Email          string `json:"Email"`
}

type renewTokenRequest struct {
	ApplicationKey string `json:"ApplicationKey"`
	Token          string `json:"Token"`
}

type validateTokenRequest struct {
	ApplicationKey string `json:"ApplicationKey"`
	Token          string `json:"Token"`
}

type isUserInHolesRequest struct {
	ApplicationKey string   `json:"ApplicationKey"`
	UserKey        string   `json:"UserKey"`
	RoleCollection []string `json:"RoleCollection"`
}

type validateApplicationRequest struct {
	ApplicationKey       string `json:"ApplicationKey"`
	ClientApplicationKey string `json:"ClientApplicationKey"`
	RawData              string `json:"RawData"`
	EncryptedData        string `json:"EncryptedData"`
}
