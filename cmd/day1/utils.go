package day1

import (
	"strings"
)

func isDigit(d uint8) bool {
	if d >= '0' && d <= '9' {
		return true
	}

	return false
}

func getDigitFromLetters(s string) (string, bool) {
	digits := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for letters, digit := range digits {
		if strings.Contains(s, letters) {
			return digit, true
		}
	}

	return "", false
}
