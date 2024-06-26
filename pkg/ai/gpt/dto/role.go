package dto

type Role = string

const (
	// RoleUser represents a user role.
	RoleUser Role = "user"
	// RoleSystem represents a system role.
	RoleSystem    Role = "system"
	RoleFunction  Role = "function"
	RoleAssistant Role = "assistant"
	RoleTool      Role = "tool"
)
