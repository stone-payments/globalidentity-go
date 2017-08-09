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
	UserKey   string   `json:"userKey"`
	Email     string   `json:"email"`
	Name      string   `json:"name"`
	Comment   string   `json:"comment"`
	Active    bool     `json:"active"`
	LockedOut bool     `json:"lockedOut"`
	Roles     []string `json:"roles,omitempty"`
}

type ListUsersResponse struct {
	Users     []User `json:"users"`
	FirstPage int    `json:"FirstPage"`
	NextPage  int    `json:"NextPage"`
	LastPage  int    `json:"LastPage"`
	TotalRows int    `json:"TotalRows"`
	*Response
}
