package utils

import (
	"fmt"
	"strings"

	"github.com/rs/xid"
)

// generate a unique slug for a post
// @TODO
// do more validations and checking
func GenerateSlug(title string, addrand bool) string {
	title = strings.ToLower(title)

	title = FirstN(title, 80)
	title = strings.ToLower(title)
	title = strings.ReplaceAll(title, " ", "-")
	title = strings.ReplaceAll(title, ",", "-")
	title = strings.ReplaceAll(title, ".", "-")

	if addrand {
		guid := xid.New()
		title = fmt.Sprintf("%s-%s", title, guid.String())
	}

	return title
}

//@utils
// Return first n chars of a string
// https://stackoverflow.com/a/41604514/17126147
func FirstN(s string, n int) string {
	i := 0
	for j := range s {
		if i == n {
			return s[:j]
		}
		i++
	}
	return s
}
