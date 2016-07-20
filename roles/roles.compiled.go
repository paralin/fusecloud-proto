package roles

/* Compiled roles to golang */

import (
	"github.com/synrobo/proto/permissions"
)

func BuildRoleMap() map[int32]permissions.SystemPermissions {
	res := make(map[int32]permissions.SystemPermissions)

	res[0] = permissions.SystemPermissions{
	}

	res[1] = permissions.SystemPermissions{
		AllPermissions: true,
	}

	return res
}