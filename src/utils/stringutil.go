package utils

func Reverse(s string) string {
	result := []rune(s)
	for i, j := 0, len(result) - 1; i < len(result) / 2; i, j = i + 1, j - 1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}