package stringutil

func UpperCase(s string) string {
	chars := []rune(s)

	for i := range chars {
		if chars[i] >= 'a' && chars[i] <= 'z' {
			chars[i] -= 'a' - 'A'
		}
	}

	return string(chars)
}

func LowerCase(s string) string {
	chars := []rune(s)

	for i := range chars {
		if chars[i] >= 'A' && chars[i] <= 'Z' {
			chars[i] += 'a' - 'A'
		}
	}

	return string(chars)
}

func Capitalize(s string) string {
	chars := []rune(s)

	if len(chars) == 0 {
		return s
	}

	if chars[0] >= 'a' && chars[0] <= 'z' {
		chars[0] -= 'a' - 'A'
	}

	for i := 0; i <= len(chars)-1; i++ {
		if chars[i] == ' ' &&
			chars[i+1] >= 'a' &&
			chars[i+1] <= 'z' {
			chars[i+1] -= 'a' - 'A'
		}
	}

	return string(chars)
}
