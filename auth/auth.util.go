package auth

import (
	"github.com/synrobo/proto/permissions"
	"github.com/synrobo/proto/roles"
)

var roleMap map[int32]permissions.SystemPermissions = roles.BuildRoleMap()

type UserPermissionTree struct {
	Global   *permissions.SystemPermissions
	ByRegion map[string]*permissions.SystemPermissions
}

/* Builds a permission tree from a user object */
func BuildUserPermissionTree(user *User, includeGlobal bool) *UserPermissionTree {
	if user == nil {
		return nil
	}

	res := &UserPermissionTree{
		Global:   &permissions.SystemPermissions{},
		ByRegion: make(map[string]*permissions.SystemPermissions),
	}

	// process global roles
	if user.GlobalRole == nil {
		user.GlobalRole = &User_UserRoleSet{}
	}

	for _, globalRole := range user.GlobalRole.Role {
		globalRolePerms, ok := roleMap[int32(globalRole)]
		if ok {
			res.Global.Extend(&globalRolePerms)
		}
	}

	if user.GlobalExtraPermission != nil {
		res.Global.Extend(user.GlobalExtraPermission)
	}

	// Process each region
	for regionId, regionRoleSet := range user.RegionRole {
		for _, regionRole := range regionRoleSet.Role {
			regionRoleV, ok := roleMap[int32(regionRole)]
			if !ok {
				continue
			}
			rperm := &permissions.SystemPermissions{}
			res.ByRegion[regionId] = rperm
			rperm.Extend(&regionRoleV)
			if includeGlobal {
				rperm.Extend(res.Global)
			}
		}
	}

	for regionId, regionPermission := range user.RegionExtraPermission {
		rperm, ok := res.ByRegion[regionId]
		if !ok {
			rperm = &permissions.SystemPermissions{}
			res.ByRegion[regionId] = rperm
		}
		rperm.Extend(regionPermission)
		if includeGlobal {
			rperm.Extend(res.Global)
		}
	}

	return res
}
