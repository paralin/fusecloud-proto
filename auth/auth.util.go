package auth

import (
	"github.com/synrobo/proto/permissions"
)

type UserPermissionTree struct {
	Global   permissions.SystemPermissions
	ByRegion map[string]permissions.SystemPermissions
}

func BuildUserPermissionTree(user *User) *UserPermissionTree {
	res := &UserPermissionTree{}

	return res
}
