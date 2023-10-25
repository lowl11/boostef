package repo

import "strings"

func filterQuery(q string) string {
	if strings.Contains(q, "'") {
		return strings.ReplaceAll(q, "'", "''")
	}

	return q
}
