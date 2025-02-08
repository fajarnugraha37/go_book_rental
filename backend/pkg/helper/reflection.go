package helper

import (
	"fmt"
	"reflect"
)

// Recursive function to get all field names, including nested structs
func GetAllFieldNames(withPrefix bool, object any, prefix ...string) []string {
	// Get the type of the struct
	t := reflect.TypeOf(object)

	// Check if it's a struct
	if t.Kind() != reflect.Struct {
		fmt.Println("Not a struct")
		return []string{}
	}

	// Iterate through the fields
	var fields []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		name := field.Name
		if len(prefix) > 0 {
			if withPrefix {
				name = prefix[0] + "." + name // For nested fields
			}
		}

		fields = append(fields, name)

		if field.Type.Kind() == reflect.Struct {
			fields = append(fields, GetAllFieldNames(withPrefix, field.Type, name)...) // Recursive call for nested structs
		}
	}
	return fields
}

// Recursive function to get all field names, including nested structs
func GetAllFieldNamesWithoutPrefix(object any, prefix ...string) []string {
	return GetAllFieldNames(false, object, prefix...)
}

type FieldInfo struct {
	Name  string
	Value any
	Tag   reflect.StructTag
}

func GetAllFields(t reflect.Type, v reflect.Value, prefix ...string) []FieldInfo {
	var fields []FieldInfo
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		name := field.Name
		if len(prefix) > 0 {
			name = prefix[0] + "." + name
		}

		fieldValue := v.Field(i)
		tag := field.Tag

		if field.Type.Kind() == reflect.Struct {
			fields = append(fields, GetAllFields(field.Type, fieldValue, name)...)
		} else {
			if !fieldValue.CanInterface() {
				continue
			}

			var value any
			if (fieldValue.Kind() == reflect.Ptr && fieldValue.IsNil()) || (fieldValue.Kind() == reflect.Interface && fieldValue.IsNil()) {
				value = nil
			} else {
				value = fieldValue.Interface()
			}

			fields = append(fields, FieldInfo{
				Name:  name,
				Value: value,
				Tag:   tag,
			})
		}
	}

	return fields
}
