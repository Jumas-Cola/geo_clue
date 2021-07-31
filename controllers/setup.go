package controllers

import (
	"strings"
	"unicode"
)

type headers struct {
	ContentType string `header:"Content-Type"`
}

var err error

func CheckValid(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) && !strings.ContainsRune("- ", r) {
			return false
		}
	}
	return true
}
