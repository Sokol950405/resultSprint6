package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

const lettersAndDigits = "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ" +
	"абвгдеёжзийклмнопрстуфхцчшщъыьэюя" +
	"0123456789"

// AutoDetect - func conversion.
func AutoDetect(data []byte) (string, error) {
	// convert bytes to a string
	input := string(data)
	if input == "" {
		return "", errors.New("empty input")
	}

	// determining the data type and return conversion
	if isMorseCode(input) {
		return morse.ToText(input), nil
	}
	return morse.ToMorse(input), nil
}

// isMorseCode - check function.
func isMorseCode(s string) bool {
	if strings.ContainsAny(s, lettersAndDigits) {
		return false
	}

	return true
}
