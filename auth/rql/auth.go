package rql

import (
	"github.com/fuserobotics/proto/auth"
	r "gopkg.in/dancannon/gorethink.v2"
)

func FetchUserByUsername(rctx *r.Session, username string) (*auth.User, error) {
	user := &auth.User{}
	cursor, err := r.DB("global").Table("users").Get(username).Run(rctx)
	defer cursor.Close()
	if err != nil {
		return nil, err
	}
	return user, cursor.One(user)
}

// Gets just username and password
func FetchUserAuthPartialByUsername(rctx *r.Session, username string) (*auth.User, error) {
	user := &auth.User{}
	cursor, err := r.DB("global").Table("users").Get(username).Pluck("username", "password").Run(rctx)
	defer cursor.Close()
	if err != nil {
		return nil, err
	}
	return user, cursor.One(user)
}

func CreateUser(rctx *r.Session, user *auth.User) error {
	_, err := r.DB("global").Table("users").Insert(user).RunWrite(rctx)
	return err
}
