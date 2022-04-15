package utils

import "strings"

func CapitalizeFirstLetter(s string) string {
	return strings.Title(strings.ToLower(s))
}
