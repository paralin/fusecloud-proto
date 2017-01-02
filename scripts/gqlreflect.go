package main

import (
	"bytes"
	"fmt"
	imp "go/importer"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/fuserobotics/proto"
)

var TypeRegex = regexp.MustCompile("(type)(.+)(struct{)(.+)(})")
var EnumRegex = regexp.MustCompile("(type)(.+)(int32)$")

func main() {
	importer := imp.Default()

	// We are building a map of graphql types -> go type paths.
	typeMap := make(map[string]string)
	typeMapKeys := make([]string, 0)

	// Also, append a function to each type that returns the graphql name.
	packageTypeMap := make(map[string]map[string]string)

	// Sort the package type map keys
	packageTypeMapPkgKeys := make(map[string][]string)

	// Sort the keys of that map too
	packageTypeMapKeys := make([]string, 0)

	// List of enums
	enumTypeMap := make(map[string]bool)

	for _, pak := range proto.ProtoPackages {
		ppkg := fmt.Sprintf("github.com/fuserobotics/proto/%s", pak)

		ptm := make(map[string]string)
		packageTypeMap[pak] = ptm
		packageTypeMapKeys = append(packageTypeMapKeys, pak)

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
			if !obj.Exported() {
				continue
			}

			isEnum := false
			if EnumRegex.MatchString(objs) {
				isEnum = true
			} else if !TypeRegex.MatchString(objs) {
				continue
			}

			baseNamePts := strings.Split(nam, "_")
			baseName := baseNamePts[len(baseNamePts)-1]
			fmt.Printf(" -> %s: %s\n", nam, baseName)
			if isEnum {
				enumTypeMap[baseName] = true
			}
			if _, ok := typeMap[baseName]; ok {
				fmt.Printf("  ! warn: %s redefined\n", baseName)
			} else {
				typeMapKeys = append(typeMapKeys, baseName)
			}

			typeMap[baseName] = strings.Split(objs, " ")[1]
			ptm[nam] = baseName
			packageTypeMapKeys = append(packageTypeMapKeys, pak)
			packageTypeMapPkgKeys[pak] = append(packageTypeMapPkgKeys[pak], nam)
		}
	}

	// Build Go format map with this information.

	// Output needs to be deterministic.
	sort.Strings(typeMapKeys)
	typeMapGo := bytes.Buffer{}
	for _, tmk := range typeMapKeys {
		typeMapGo.WriteString("\n\t\"")
		typeMapGo.WriteString(tmk)
		typeMapGo.WriteString("\": \"")
		typeMapGo.WriteString(typeMap[tmk])
		typeMapGo.WriteString("\",")
	}

	ioutil.WriteFile("root_gql.go", []byte(
		fmt.Sprintf(
			`package proto

// Map of GraphQL simplified names to proto names.
var GraphQLTypes = map[string]string{%s
}`, typeMapGo.String(),
		),
	),
		0644,
	)

	// For each type, write a file with functions for gql type names.
	sort.Strings(packageTypeMapKeys)
	for _, packageName := range packageTypeMapKeys {
		pakTypeMap := packageTypeMap[packageName]
		pakTypeMapKeys := packageTypeMapPkgKeys[packageName]

		if len(pakTypeMap) == 0 {
			continue
		}

		outPath := fmt.Sprintf("./%s/graphql.types.go", packageName)
		outp := bytes.Buffer{}

		outp.WriteString("package ")
		outp.WriteString(packageName)
		outp.WriteString("\n\n// WARNING: This file is auto-generated.\n")

		sort.Strings(pakTypeMapKeys)
		for _, typeName := range pakTypeMapKeys {
			gqlTypeName := pakTypeMap[typeName]
			if _, ok := enumTypeMap[gqlTypeName]; !ok {
				typeName = "*" + typeName
			}

			outp.WriteString(fmt.Sprintf(
				"\nfunc (%s) GraphQLTypeName() string {\n\treturn \"%s\"\n}\n",
				typeName,
				gqlTypeName,
			))
		}

		ioutil.WriteFile(outPath, []byte(outp.String()), 0644)
	}
}
