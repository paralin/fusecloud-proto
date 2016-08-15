package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"sort"
	"strings"

	"github.com/fuserobotics/proto/auth"
	"github.com/fuserobotics/proto/permissions"
)

var HeaderTemplatePath string = "./roles/roles.hdr.tmpl"
var RolesJsonPath string = "./roles/roles.json"

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("No argument for destination.\n")
		os.Exit(1)
	}
	targetPath := os.Args[1]
	fmt.Printf("Will write to %s...\n", targetPath)

	var resultBuf bytes.Buffer
	d, err := ioutil.ReadFile(HeaderTemplatePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	resultBuf.WriteString(string(d))

	d, err = ioutil.ReadFile(RolesJsonPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	var roleMap map[string]permissions.SystemPermissions
	err = json.Unmarshal(d, &roleMap)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	roleIntMap := make(map[int]permissions.SystemPermissions)
	var sortedKeys []int
	for k, v := range roleMap {
		roleName := strings.ToUpper(k)
		// find role for k
		roleId, ok := auth.User_UserRole_value[roleName]
		if !ok {
			fmt.Printf("Couldn't find role ID for %sd\n", roleName)
			continue
		}

		roleIntMap[int(roleId)] = v
		sortedKeys = append(sortedKeys, int(roleId))
	}
	sort.Ints(sortedKeys)

	for roleId := range sortedKeys {
		v := roleIntMap[roleId]
		resultBuf.WriteString(fmt.Sprintf("\tres[%d] = permissions.SystemPermissions{\n", roleId))
		// iterate over boolean properties on the object
		s := reflect.ValueOf(&v).Elem()
		to := s.Type()
		for i := 0; i < s.NumField(); i++ {
			fn := to.Field(i).Name
			fv := s.Field(i).Bool()
			if !fv {
				continue
			}
			resultBuf.WriteString(fmt.Sprintf("\t\t%s: true,\n", fn))
		}
		resultBuf.WriteString("\t}\n\n")
	}
	resultBuf.WriteString("\treturn res\n}")

	ioutil.WriteFile(targetPath, resultBuf.Bytes(), 0644)
}
