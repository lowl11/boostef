package keywords

var (
	keywords = [...]string{
		"id", "group", "user",
	}
)

func Is(value string) bool {
	for _, keyword := range keywords {
		if value == keyword {
			return true
		}
	}

	return false
}
