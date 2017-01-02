package graphql

import (
	"fmt"
	"reflect"
	"strings"
)

type GraphQLProto interface {
	GraphQLTypeName() string
}

type GraphQLProtoEnum interface {
	GraphQLTypeName() string
	String() string
}

// Recursively marshal to a map result, including __typename
func MarshalToMap(obj interface{}) interface{} {
	if obj == nil {
		return nil
	}

	// Try as map[string]interface{}
	if smap, ok := obj.(map[string]interface{}); ok {
		return smap
	}

	sourceVal := reflect.ValueOf(obj)
	sourceValType := sourceVal.Type()
	sourceValKind := sourceVal.Kind()

	// Follow pointer
	if sourceVal.IsValid() && sourceValKind == reflect.Ptr {
		sourceVal = sourceVal.Elem()
		if !sourceVal.IsValid() {
			return nil
		}
		sourceValType = sourceVal.Type()
		sourceValKind = sourceVal.Kind()
	}
	if !sourceVal.IsValid() {
		return nil
	}

	// If struct, convert to map.
	if sourceValKind == reflect.Struct {
		rmap := make(map[string]interface{})
		if ptn, ok := obj.(GraphQLProto); ok {
			rmap["__typename"] = ptn.GraphQLTypeName()
		}
		for i := 0; i < sourceVal.NumField(); i++ {
			valueField := sourceVal.Field(i)
			typeField := sourceValType.Field(i)

			tag := typeField.Tag
			jsonTag := tag.Get("json")
			jsonOptions := strings.Split(jsonTag, ",")
			rmapName := typeField.Name
			if len(jsonOptions) > 0 {
				rmapName = jsonOptions[0]
			}

			if valueField.Kind() == reflect.Ptr && valueField.IsNil() {
				rmap[rmapName] = nil
				continue
			}

			if valueField.IsValid() {
				elem := valueField
				if elem.Kind() == reflect.Ptr {
					elem = valueField.Elem()
				}
				rmap[rmapName] = MarshalToMap(valueField.Interface())
				continue
			}

			rmap[rmapName] = valueField.Interface()
		}

		return rmap
	}

	if sourceValKind == reflect.String {
		return obj.(string)
	}

	if sourceValKind == reflect.Array || sourceValKind == reflect.Slice {
		rarr := make([]interface{}, sourceVal.Len())
		for i := 0; i < sourceVal.Len(); i++ {
			e := sourceVal.Index(i)
			rarr[i] = MarshalToMap(e.Interface())
		}
		return rarr
	}

	if sourceValKind == reflect.Int32 {
		// Detect enum value, to see if it's a graphql enum
		if objes, ok := obj.(GraphQLProtoEnum); ok {
			// Use it as a string
			return objes.String()
		}
		return int32(sourceVal.Int())
	}

	if sourceValKind == reflect.Uint32 {
		return uint32(sourceVal.Uint())
	}

	fmt.Printf("Unhandled type (%v): %v\n", sourceValKind, obj)

	return obj
}
