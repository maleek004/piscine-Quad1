package checkpoint

import "strings"

// func CamelToSnakeCase(arg string) string {
// 	IsCamelcaseDisqualifier := func(r rune) bool {
// 		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' {
// 			return false
// 		}
// 		return true
// 	}
// 	IsValidCamelCase := func(s string) bool {
// 		for i, r := range s {
// 			// checking for numbers or non letters
// 			if IsCamelcaseDisqualifier(r) {
// 				return false
// 			}
// 			// checking if the character is a Capital letter
// 			if r >= 'A' && r <= 'Z' {
// 				// checking if the capital letter is at the end of the string
// 				if len(s) == i+1 {
// 					return false
// 				} else if i > 0 && (s[i-1] >= 'A' && s[i-1] <= 'Z') { // checking if the previouse letter was uppercase
// 					return false
// 				}

// 			}
// 		}
// 		return true
// 	}
// 	finalString := ""
// 	if IsValidCamelCase(arg) {
// 		for i, r := range arg {
// 			if i != 0 && (r >= 'A' && r <= 'Z') {
// 				finalString += string('_')
// 				finalString += string(r)
// 			} else {
// 				finalString += string(r)
// 			}
// 		}

// 	}
// 	return finalString
// }

func IsCamel(s string) bool {
	IsAlphaOnly := func(str string) bool {
		for _, r := range str {
			if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
				continue
			} else {
				return false
			}
		}
		return true
	}

	if !IsAlphaOnly(s) {
		return false
	}
	sSize := len(s)
	for i, r := range s {
		if i > 0 {
			if (r >= 'A' && r <= 'Z') && (rune(s[i-1]) >= 'A' && rune(s[i-1]) <= 'Z') {
				return false
			}
			if (r >= 'A' && r <= 'Z') && i == sSize-1 {
				return false
			}
		}
	}
	return true
}

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	if !IsCamel(s) {
		//return "error: not camel case"
		return s
	}

	var res strings.Builder
	for i, r := range s {
		if r >= 'A' && r <= 'Z' {
			if i == 0 {
				res.WriteRune(r)
			} else {
				res.WriteRune('_')
				res.WriteRune(r)
			}
		} else {
			res.WriteRune(r)
		}
	}
	return res.String()
}
