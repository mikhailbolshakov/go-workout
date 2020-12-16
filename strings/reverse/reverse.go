package reverse

func reverse(input string) string {
	runes := []rune(input)
	revRunes := make([]rune, len(input))

	l := len(runes)

	for i := 0; i < l; i = i + 1 {
		revRunes[i] = runes[l - i - 1]
	}
	return string(revRunes)

}

