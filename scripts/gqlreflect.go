package main

import (
	"bytes"
	"fmt"
	imp "go/importer"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/fuserobotics/proto"
)

var TypeRegex = regexp.MustCompile("(type)(.+)(struct{)(.+)(})")

func main() {
	importer := imp.Default()

	// We are building a map of graphql types -> go type paths.
	typeMap := make(map[string]string)

	// Also, append a function to each type that returns the graphql name.
	packageTypeMap := make(map[string]map[string]string)

	for _, pak := range proto.ProtoPackages {
		ppkg := fmt.Sprintf("github.com/fuserobotics/proto/%s", pak)

		ptm := make(map[string]string)
		packageTypeMap[pak] = ptm

		ipak, err := importer.Import(ppkg)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s:\n", ppkg)
		scope := ipak.Scope()
		for _, nam := range scope.Names() {
			obj := scope.Lookup(nam)
			objs := obj.String()
			if !obj.Exported() ||
				!TypeRegex.MatchString(objs) {
				continue
			}
			baseNamePts := strings.Split(nam, "_")
			baseName := baseNamePts[len(baseNamePts)-1]
			fmt.Printf(" -> %s: %s\n", nam, baseName)
			if _, ok := typeMap[baseName]; ok {
				fmt.Printf("  ! warn: %s redefined\n", baseName)
			}

			typeMap[baseName] = strings.Split(objs, " ")[1]
			ptm[nam] = baseName
		}
	}

	// Build Go format map with this information.
	typeMapGo := fmt.Sprintf("%#v\n", typeMap)
	typeMapGo = strings.Replace(typeMapGo, ", ", ",\n\t", -1)
	typeMapGo = strings.Replace(typeMapGo, "{", "{\n\t", -1)
	typeMapGo = strings.Replace(typeMapGo, "}", ",\n}", -1)
	typeMapGo = strings.Replace(typeMapGo, ":", ": ", -1)
	ioutil.WriteFile("root_gql.go", []byte(
		fmt.Sprintf(
			`package proto

// Map of GraphQL simplified names to proto names.
var GraphQLTypes = %s
`, typeMapGo,
		),
	),
		0644,
	)

	// For each type, write a file with functions for gql type names.
	for packageName, pakTypeMap := range packageTypeMap {
		if len(pakTypeMap) == 0 {
			continue
		}

		outPath := fmt.Sprintf("./%s/graphql.types.go", packageName)
		outp := bytes.Buffer{}

		outp.WriteString("package ")
		outp.WriteString(packageName)
		outp.WriteString("\n\n// WARNING: This file is auto-generated.\n")

		for typeName, gqlTypeName := range pakTypeMap {
			outp.WriteString(fmt.Sprintf(
				"\nfunc (*%s) GraphQLTypeName() string {\n\treturn \"%s\"\n}\n",
				typeName,
				gqlTypeName,
			))
		}

		ioutil.WriteFile(outPath, []byte(outp.String()), 0644)
	}
}
