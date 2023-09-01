package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func MarshalNonZeroFields(v interface{}) ([]byte, error) {
	// Create a map to store non-zero fields
	nonZeroFields := make(map[string]interface{})

	// Use reflection to iterate through the struct fields
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input is not a struct")
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldVal := field.Interface()

		// Check if the field is zero-valued
		if !reflect.DeepEqual(fieldVal, reflect.Zero(typ.Field(i).Type).Interface()) {
			// Field is non-zero, include it in the map
			fieldName := typ.Field(i).Name
			jsonTag := typ.Field(i).Tag.Get("json")
			if jsonTag != "" {
				jsonTagParts := strings.Split(jsonTag, ",")
				if jsonTagParts[0] != "-" {
					fieldName = jsonTagParts[0]
				}
			}
			nonZeroFields[fieldName] = fieldVal
		}
	}

	// Marshal the non-zero fields to JSON
	return json.Marshal(nonZeroFields)
}
