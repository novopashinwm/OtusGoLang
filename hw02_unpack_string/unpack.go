package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inStr string) (string, error) {
	// Place your code here.
	var sb strings.Builder
	chars := []rune(inStr)
	var err error
	if len(chars) == 0 {
		return "", nil
	}
	if unicode.IsDigit(chars[0]) {
		return "", ErrInvalidString
	}
	num := len(chars)
	ekranSymbol := false
	for i := 0; i < num; i++ {
		curr := string(chars[i])
		switch {
		case chars[i] == 92 && !ekranSymbol:
			ekranSymbol = true
		case i < (num-1) && unicode.IsDigit(chars[i+1]):
			if IsNextDigital(i, num, chars) {
				return "", ErrInvalidString
			}

			next := string(chars[i+1])
			repeatNum, _ := strconv.Atoi(next)
			if repeatNum != 0 {
				sb.WriteString(strings.Repeat(curr, repeatNum))
			}

			i++
		default:

			sb.WriteString(curr)
			ekranSymbol = false
		}
	}
	return sb.String(), err
}

func IsNextDigital(i int, num int, chars []rune) bool {
	if i < num-2 {
		return unicode.IsDigit(chars[i+2])
	}
	return false
}
