package model

const (
	UserRoleCode    = "8001"
	CashierRoleCode = "8002"
	AdminRoleCode   = "8003"
)

type UserRole struct {
	RoleCode string
	RoleName string
}

func IsAdmin(roleCode string) bool {
	return roleCode == AdminRoleCode
}
