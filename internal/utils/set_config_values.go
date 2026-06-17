package utils

import (
	"reflect"
)

// DefaultAPIVersion is used when no API version is configured.
const DefaultAPIVersion = "v3"

// Configuration holds the common configuration values.
type Configuration struct {
	Username   string
	Password   string
	BaseURL    string
	APIVersion string
}

// SetConfigValues sets common configuration values for a given config struct.
func SetConfigValues(config interface{}, commonConfig Configuration) {
	v := reflect.ValueOf(config).Elem()
	t := v.Type()

	apiVersion := commonConfig.APIVersion
	if apiVersion == "" {
		apiVersion = DefaultAPIVersion
	}

	for i := 0; i < t.NumField(); i++ {
		field := v.Field(i)
		fieldName := t.Field(i).Name

		switch fieldName {
		case "Username":
			field.SetString(commonConfig.Username)
		case "Password":
			field.SetString(commonConfig.Password)
		case "BaseURL":
			field.SetString(commonConfig.BaseURL)
		case "APIVersion":
			field.SetString(apiVersion)
		}
	}
}
