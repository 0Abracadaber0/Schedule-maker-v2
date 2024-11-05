package service

import "strings"

func endsWithLab(s string) bool {
	return strings.HasSuffix(s, "lab")
}
