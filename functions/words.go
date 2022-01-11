package functions

import "strings"

func Up(s string) string {
	return strings.ToUpper(s)
}

func Low(s string) string {
	return strings.ToLower(s)
}

func Cap(s string) string {
	return strings.Title(s)
}