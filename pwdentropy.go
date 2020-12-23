// Package pwdentropy responsible to measure password strength using entropy value.
// https://generatepasswords.org/how-to-calculate-entropy/
//
// Only supports basic latin unicode
// https://en.wikipedia.org/wiki/List_of_Unicode_characters#Latin_script
//
package pwdentropy

import "math"

// Calculate measures password strength using entropy value.
// Value is rounded to the nearest.
//
func Calculate(password string) float64 {
	s := getSizeOfThePoolOfCharacterSet(password)

	// TODO: need to check dummy password sequence
	// Like: aaaaaaaaa
	l := len(password)

	possibleCombination := math.Pow(float64(s), float64(l))
	entropy := math.Log2(possibleCombination)
	return math.Round(entropy*100) / 100
}

func getSizeOfThePoolOfCharacterSet(password string) int {
	var (
		hasPunctuationAndSymbols bool
		hasDigits                bool
		hasAlphabetUppercase     bool
		hasAlphabetLowercase     bool
	)

	for _, p := range password {
		if isPunctuationAndSymbols(p) {
			hasPunctuationAndSymbols = true
			continue
		}

		if isDigits(p) {
			hasDigits = true
			continue
		}

		if isAlphabetUppercase(p) {
			hasAlphabetUppercase = true
			continue
		}

		if isAlphabetLowercase(p) {
			hasAlphabetLowercase = true
			continue
		}
	}

	var size int

	if hasPunctuationAndSymbols {
		size += poolSizePunctuationAndSymbols
	}

	if hasDigits {
		size += poolSizeDigits
	}

	if hasAlphabetUppercase {
		size += poolSizeAlphabetUppercase
	}

	if hasAlphabetLowercase {
		size += poolSizeAlphabetLowercase
	}

	return size
}

const (
	poolSizePunctuationAndSymbols int = 33
	poolSizeDigits                    = 10
	poolSizeAlphabetUppercase         = 26
	poolSizeAlphabetLowercase         = 26
)

func isPunctuationAndSymbols(r rune) bool {
	return (r >= 32 && r <= 47) || (r >= 58 && r <= 64) || (r >= 91 && r <= 96) || (r >= 123 && r <= 126)
}

func isDigits(r rune) bool {
	return r >= 48 && r <= 57
}

func isAlphabetUppercase(r rune) bool {
	return r >= 65 && r <= 90
}

func isAlphabetLowercase(r rune) bool {
	return r >= 97 && r <= 122
}
