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
	var num = len(chars)
	var ekranSymbol = false
	for i := 0; i < num; i++ {
		curr := string(chars[i])
		if chars[i] == 92 && !ekranSymbol {
			ekranSymbol = true
		} else if i < (num-1) && unicode.IsDigit(chars[i+1]) {
			if IsNextDigital(i, num, chars) {
				return "", ErrInvalidString
			}

			next := string(chars[i+1])
			var repeatNum, _ = strconv.Atoi(next)
			if repeatNum != 0 {
				sb.WriteString(strings.Repeat(curr, repeatNum))
			}

			i++
		} else {
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
