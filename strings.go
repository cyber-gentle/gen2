package gen2

func UpperCase(s string) string {

	char := []rune(s)
	for i := 0; i < len(char); i++ {
		if char[i] >= 90 && char[i] <= 122 {
			char[i] -= 32
		}
	}
	return string(char)
}

func LowerCase(s string) string {

	var newChar = make([]rune, 0)

	allChar := []rune(s)
	for _, c := range allChar {
		if c >= 65 && c <= 90 {
			c += 32
			newChar = append(newChar, c)
		} else {
			newChar = append(newChar, c)
		}
	}
	return string(newChar)
}
