package utilities

import "strings"

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const base = int64(len(alphabet))

func Encode(num uint64) string {
	if num == 0 {
		return string(alphabet[0])
	}

	var result strings.Builder
	for num > 0 {
		rem := num % uint64(base)
		result.WriteByte(alphabet[rem])
		num /= uint64(base)
	}

	encoded := []rune(result.String())
	for i, j := 0, len(encoded)-1; i < j; i, j = i+1, j-1 {
		encoded[i], encoded[j] = encoded[j], encoded[i]
	}

	return string(encoded)
}
