// +build codegen

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"reflect"
	"sort"
	"strings"
)

func resolverName(typeName string) string {
	return fmt.Sprintf("%sResolver", typeName)
}

func stripLeadingQualifiers(returnTypeStr string) string {
	if strings.HasPrefix(returnTypeStr, "*") {
		returnTypeStr = returnTypeStr[1:]
	}
	if strings.HasPrefix(returnTypeStr, "[]") {
		returnTypeStr = returnTypeStr[2:]
	}
	if strings.HasPrefix(returnTypeStr, "*") {
		returnTypeStr = returnTypeStr[1:]
	}
	return returnTypeStr
}

func main() {
	var outpBuf bytes.Buffer

	outpBuf.WriteString("// +build !codegen\n\npackage resolvers\n")

	// First, to enable sorting order, iterate over types and sort them.
	allTypes := make([]string, len(AllGraphQLTypes))
	var allPaks []string
	i := 0
	for typeStr, typeInstance := range AllGraphQLTypes {
		allTypes[i] = typeStr

		typePkg := reflect.TypeOf(typeInstance).Elem().PkgPath()
		if typePkg == "" {
			continue
		}
		found := false
		for _, pak := range allPaks {
			if pak == typePkg {
				found = true
				break
			}
		}
		if !found {
			allPaks = append(allPaks, typePkg)
		}
		i++
	}

	sort.Strings(allPaks)
	sort.Strings(allTypes)

	outpBuf.WriteString("\nimport (")
	for _, pak := range allPaks {
		outpBuf.WriteString("\n\t\"")
		outpBuf.WriteString(pak)
		outpBuf.WriteString("\"")
	}
	outpBuf.WriteString("\n)\n")

	for _, typeStr := range allTypes {
		typeInstance := AllGraphQLTypes[typeStr]
		typePts := strings.Split(typeStr, ".")

		// Package name (github.com/fuserobotics/proto/device)
		// typePkg := typePts[0]
		// Type name (Device)
		typeName := typePts[1]
		// Type name (DeviceResolver)
		resolverTypeName := resolverName(typeName)

		typ := reflect.TypeOf(typeInstance)
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}
		if typ.Kind() != reflect.Struct {
			continue
		}

		fmt.Printf("%s:\n", resolverTypeName)
		outpBuf.WriteString(
			fmt.Sprintf("\ntype %s struct {\n\tBase *%s\n}\n", resolverTypeName, typeStr),
		)
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)

			isPointer := false
			isSlice := false
			isPrimitive := false
			requiresCast := false
			// uint32, string, or blank if pointer / slice / other
			returnType := field.Type
			// uint32, float64, or ptr or slice
			returnTypeKind := returnType.Kind()

			if returnTypeKind == reflect.Slice {
				isSlice = true
				returnType = returnType.Elem()
				returnTypeKind = returnType.Kind()
			}

			if returnTypeKind == reflect.Ptr {
				isPointer = true
				returnType = returnType.Elem()
				returnTypeKind = returnType.Kind()
			}

			fmt.Printf("  - %s: %v\n", returnType.Name(), returnTypeKind)

			// If is primitive, Kind will be anything but struct.
			// Could have isPointer and isSlice if we had []*Device
			// returnType: type by name (IPAddress, PrivateKeyStrategy) (includes enums)
			// returnTypeKind: base kind, struct or primitive.

			// First, determine our return type - *DeviceResolver, []*DeviceResolver, etc.
			var returnTypeStr string
			// We have a resolver if it's a struct.
			if returnTypeKind == reflect.Struct {
				// We would need to return []*DeviceResolver
				returnTypeStr = returnType.Name() + "Resolver"
			} else {
				returnTypeStr = returnTypeKind.String()
				isPrimitive = true
			}
			if isPointer || (isPrimitive && isSlice) {
				returnTypeStr = "*" + returnTypeStr
			}
			if isSlice {
				returnTypeStr = "[]" + returnTypeStr
			}
			if returnTypeKind != reflect.Struct && returnTypeStr != field.Type.String() {
				requiresCast = true
			}
			fmt.Printf("    -> %s\n", returnTypeStr)

			actReturnTypeStr := returnTypeStr
			if returnTypeKind != reflect.Struct || isSlice {
				actReturnTypeStr = "*" + returnTypeStr
			}

			outpBuf.WriteString(
				fmt.Sprintf(
					"\nfunc (r *%s) %s() %s {\n",
					resolverTypeName,
					field.Name,
					actReturnTypeStr,
				),
			)

			if isSlice {
				outpBuf.WriteString(
					fmt.Sprintf(
						"\tres := make(%s, len(r.Base.%s))\n",
						returnTypeStr,
						field.Name,
					),
				)
			}

			// If we are just returning a primitive, we can be simple.
			if isPrimitive {
				if isSlice {
					outpBuf.WriteString(
						fmt.Sprintf(
							"\tfor idx, val := range r.Base.%s {\n\t\tres[idx] = &val\n\t}\n\treturn &res\n}\n",
							field.Name,
						),
					)
				} else if requiresCast {
					outpBuf.WriteString(
						fmt.Sprintf(
							"\tres := (*%s)(&r.Base.%s)\n\treturn res\n}\n",
							returnTypeStr,
							field.Name,
						),
					)
				} else {
					outpBuf.WriteString(
						fmt.Sprintf(
							"\treturn &r.Base.%s\n}\n",
							field.Name,
						),
					)
				}
			} else {
				// We are returning a resolver.
				if isSlice {
					outpBuf.WriteString(
						fmt.Sprintf(
							"\tfor idx, val := range r.Base.%s {\n\t\tres[idx] = &%s{Base: val}\n\t}\n\treturn &res\n}\n",
							field.Name,
							stripLeadingQualifiers(returnTypeStr),
						),
					)
				} else {
					outpBuf.WriteString(
						fmt.Sprintf(
							"\tif r.Base.%s == nil { return nil }\n",
							field.Name,
						),
					)
					outpBuf.WriteString(
						fmt.Sprintf(
							"\treturn &%s{Base: r.Base.%s}\n}\n",
							stripLeadingQualifiers(returnTypeStr),
							field.Name,
						),
					)
				}
			}
		}
	}

	// fmt.Printf("\n%s\n", outpBuf.String())
	ioutil.WriteFile("./resolvers.go", []byte(outpBuf.String()), 0644)
}
