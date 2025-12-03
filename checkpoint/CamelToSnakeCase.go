package checkpoint

func CamelToSnakeCase(arg string) string {
	IsCamelcaseDisqualifier := func(r rune) bool {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' {
			return false
		}
		return true
	}
	IsValidCamelCase := func(s string) bool {
		for i, r := range s {
			// checking for numbers or non letters
			if IsCamelcaseDisqualifier(r) {
				return false
			}
			// checking if the character is a Capital letter
			if r >= 'A' && r <= 'Z' {
				// checking if the capital letter is at the end of the string
				if len(s) == i+1 {
					return false
				} else if i > 0 && (s[i-1] >= 'A' && s[i-1] <= 'Z') { // checking if the previouse letter was uppercase
					return false
				}

			}
		}
		return true
	}
	finalString := ""
	if IsValidCamelCase(arg) {
		for i, r := range arg {
			if i != 0 && (r >= 'A' && r <= 'Z') {
				finalString += string('_')
				finalString += string(r)
			} else {
				finalString += string(r)
			}
		}

	}
	return finalString
}
