package management

const (
	contentJSON   = "application/json"
	listUserRoles = "/api/management/%s/users/%s/roles"
	listUsers     = "/api/management/%s/users?page=%d&limit=%d&includeRoles=%t"
	getUser     = "/api/management/%s/users/%s?includeRoles=%t"
)
