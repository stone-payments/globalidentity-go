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
