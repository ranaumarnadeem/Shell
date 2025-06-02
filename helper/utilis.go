package utils

import (
	"strings"
)


func Contains(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}


func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
