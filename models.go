package globalidentity

type Authorization struct {
	Token string
	Key   string
}

type Role struct {
	Name        string
	Description string
	Active      bool
}

type User struct {
	UserKey   string `json:"userKey"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Comment   string `json:"comment"`
	Active    bool `json:"active"`
	LockedOut bool `json:"lockedOut"`
	Roles     []string `json:"roles"`
}
