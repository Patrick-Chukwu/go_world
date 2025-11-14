package piscine

func JumpOver(str string) string {
	if len(str) <= 2 {
		return "\n"

	}
	result := ""
	for i, ch := range str {
		if i%3 == 0 {
			result += string(ch)
		}
	}
	return result
}