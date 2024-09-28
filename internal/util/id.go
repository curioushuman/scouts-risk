package util

import (
	"reflect"
	"regexp"
	"strings"
)

func IdFromString(input string) string {
    // Convert to lowercase
    lowercaseStr := strings.ToLower(input)

    // Remove spaces and unacceptable characters using regex
    re := regexp.MustCompile(`[^a-z0-9]+`)
    cleanedStr := re.ReplaceAllString(lowercaseStr, "")

    return cleanedStr
}

func FindByStringId[T any](id string, items []T) T {
	var stringId string
	for _, item := range items {
		v := reflect.ValueOf(item)
		labelField := v.FieldByName("Label")
		if (labelField.IsValid() && labelField.Kind() == reflect.String) {
			stringId = IdFromString(labelField.String())
			if stringId == id {
				return item
			}
		}
	}

	var emptyT T
	return emptyT
}