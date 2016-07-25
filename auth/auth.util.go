package auth

import (
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"time"

	"github.com/synrobo/proto/permissions"
	"github.com/synrobo/proto/roles"
)

const AuthDomain string = "users.synrobo.com"

var roleMap map[int32]permissions.SystemPermissions = roles.BuildRoleMap()

type UserPermissionTree struct {
	Global   *permissions.SystemPermissions
	ByRegion map[string]*permissions.SystemPermissions
}

func (user *User) GenerateFqdn(authDomain string) string {
	return fmt.Sprintf("%s.%s", user.Username, authDomain)
}

func ParseFqdn(fqdn string) (username string, err error) {
	n, err := fmt.Sscanf(fqdn, fmt.Sprintf("%%s.%s", AuthDomain), &username)
	if err != nil {
		return "", err
	}
	if n < 1 {
		return "", fmt.Errorf("FQDN %s failed to parse.", fqdn)
	}
	return username, nil
}

func (usr *User_UserCert) ParsePublicKey() (*rsa.PublicKey, error) {
	if usr.PublicKey == "" {
		return nil, errors.New("User has no public key.")
	}
	blk, _ := pem.Decode([]byte(usr.PublicKey))
	if blk == nil {
		return nil, errors.New("Private key pem failed to parse")
	}
	if blk.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("Public key pem type %s is not PUBLIC KEY.", blk.Type)
	}
	pkr, err := x509.ParsePKIXPublicKey(blk.Bytes)
	if err != nil {
		return nil, err
	}
	return pkr.(*rsa.PublicKey), nil
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

func (user *User) BuildSubject(authDomain string) pkix.Name {
	return pkix.Name{
		Country:            []string{"United States"},
		Province:           []string{"California"},
		Locality:           []string{"Cloud"},
		Organization:       []string{"Synergy Robotics"},
		OrganizationalUnit: []string{"users:" + user.Username},
		CommonName:         user.GenerateFqdn(authDomain),
	}
}

func UserCertExpiryTime(from time.Time) time.Time {
	return from.Add(time.Duration(375) * 24 * time.Hour)
}
