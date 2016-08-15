package auth_cql

import (
	"fmt"

	"github.com/paralin/cqlpb/marshal"
	"github.com/relops/cqlr"
	"github.com/fuserobotics/proto/auth"

	cup "github.com/fuserobotics/cassandra"
)

func FetchUserByUsername(ctx *cup.CassandraContext, username string) (*auth.User, error) {
	user := &auth.User{}
	m := make(map[string]interface{})
	if err := ctx.Session.Query(fmt.Sprintf("SELECT * FROM %s WHERE username = ? LIMIT 1", auth.AuthTable), username).MapScan(m); err != nil {
		return nil, err
	}
	if err := marshal.Unmarshal(user, m); err != nil {
		return nil, err
	}
	return user, nil
}

func FetchUserAuthPartialByUsername(ctx *cup.CassandraContext, username string) (*auth.User, error) {
	user := &auth.User{}
	m := make(map[string]interface{})
	if err := ctx.Session.Query(fmt.Sprintf("SELECT username, password FROM %s WHERE username = ? LIMIT 1", auth.AuthTable), username).MapScan(m); err != nil {
		return nil, err
	}
	// m["proto"] = nil
	if err := marshal.Unmarshal(user, m); err != nil {
		return nil, err
	}
	return user, nil
}

type insertUserStruct struct {
	Username string
	Password string
	Proto    []byte
}

func CreateUser(ctx *cup.CassandraContext, user *auth.User) error {
	toInsert, err := marshal.Marshal(user, auth.AuthTableTmpl)
	if err != nil {
		return err
	}

	// make this DYNAMICALLY
	toIns := &insertUserStruct{
		Username: toInsert["username"].(string),
		Password: toInsert["password"].(string),
		Proto:    toInsert["proto"].([]byte),
	}

	q := cqlr.Bind("INSERT INTO users (username, password, proto) VALUES (?, ?, ?)", toIns)
	return q.Exec(ctx.Session)
}
