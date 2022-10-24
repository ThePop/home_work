package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(source string) (string, error) {
	var builder strings.Builder

	runes := []rune(source)
	for i := 0; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) {
			return "", ErrInvalidString
		}

		// проверка на последний символ строки
		if i == len(runes)-1 {
			if !unicode.IsDigit(runes[i]) {
				builder.WriteString(string(runes[i]))
				break
			} else {
				return "", ErrInvalidString
			}
		}

		if unicode.IsDigit(runes[i+1]) {
			digit, err := strconv.Atoi(string(runes[i+1]))
			if err != nil {
				return "", err
			}

			builder.WriteString(strings.Repeat(string(runes[i]), digit))
			i++
		} else {
			builder.WriteString(string(runes[i]))
		}
	}

	return builder.String(), nil
}
