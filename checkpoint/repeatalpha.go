package checkpoint

import "strings"

func RepeatAlpha(s string) string {
	var res strings.Builder

	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			res.WriteString(strings.Repeat(string(r), int(r-'a')+1))
		} else if r >= 'A' && r <= 'Z' {
			res.WriteString(strings.Repeat(string(r), int(r-'A')+1))
		} else {
			res.WriteString(string(r))
		}
	}
	return res.String()
}
