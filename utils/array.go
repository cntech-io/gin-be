package utils

import "strings"

func FindInStringArray(arr []string, value string) bool {
	if len(arr) > 0 && value == "" {
		return false
	}
	for _, item := range arr {
		if strings.EqualFold(item, value) {
			return true
		}
	}
	return false
}
