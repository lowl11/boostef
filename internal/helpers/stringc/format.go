package stringc

import "strings"

func Format(template, key, value string) string {
	return strings.ReplaceAll(template, "{% "+key+" %}", value)
}
