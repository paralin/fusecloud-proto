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

var ReflectGenPath string = "./resolvers/reflectgen.go"
var TypeRegex = regexp.MustCompile("(type)(.+)(struct{)(.+)(})")
var EnumRegex = regexp.MustCompile("(type)(.+)(int32)$")

func main() {
	importer := imp.Default()

	// We are building a map of graphql types -> go type paths.
	typeMapImports := []string{}
	typeMapImportMap := map[string]bool{}
	typeMapStructs := []string{}
	typeMapStructMap := map[string]string{}

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

			objsParts := strings.Split(strings.Split(objs, " ")[1], "/")
			// auth.UserPrivateData
			objsQualified := objsParts[len(objsParts)-1]
			objsQualifiedPts := strings.Split(objsQualified, ".")

			if !isEnum {
				typeMapStructs = append(typeMapStructs, objsQualified)
				typeMapStructMap[objsQualified] = objsQualified // baseName

				objsImportPath := strings.Join(objsParts[:len(objsParts)-1], "/") + "/" + objsQualifiedPts[0]
				if _, ok := typeMapImportMap[objsImportPath]; !ok {
					typeMapImportMap[objsImportPath] = true
					typeMapImports = append(typeMapImports, objsImportPath)
				}
			}

			ptm[nam] = baseName
			packageTypeMapKeys = append(packageTypeMapKeys, pak)
			packageTypeMapPkgKeys[pak] = append(packageTypeMapPkgKeys[pak], nam)
		}
	}

	// Build Go format map with this information.

	// Output needs to be deterministic.
	sort.Strings(typeMapStructs)
	sort.Strings(typeMapImports)
	typeMapGo := bytes.Buffer{}
	typeMapImportsGo := bytes.Buffer{}
	for _, tmi := range typeMapImports {
		typeMapImportsGo.WriteString("\n\t\"")
		typeMapImportsGo.WriteString(tmi)
		typeMapImportsGo.WriteString("\"")
	}
	for _, tmk := range typeMapStructs {
		typeMapGo.WriteString("\n\t\"")
		typeMapGo.WriteString(typeMapStructMap[tmk])
		typeMapGo.WriteString("\": &")
		typeMapGo.WriteString(tmk)
		typeMapGo.WriteString("{},")
	}

	ioutil.WriteFile(ReflectGenPath, []byte(
		fmt.Sprintf(
			`// +build codegen

package main

import (%s
)

type GraphQLProto interface {
	GraphQLTypeName() string
}

// Map of GraphQL simplified names to proto names.
var AllGraphQLTypes = map[string]GraphQLProto{%s
}`, typeMapImportsGo.String(), typeMapGo.String(),
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
