package util

import (
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