package management

const (
	contentJSON   = "application/json"
	listUserRoles = "/api/management/%s/users/%s/roles"
	listUsers     = "/api/management/%s/users?page=%d&limit=%d&includeRoles=%s"
)
