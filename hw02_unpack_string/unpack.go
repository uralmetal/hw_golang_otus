package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var builder strings.Builder

	for i := 0; i < len(str); i++ {
		if i > 0 && str[i] >= '0' && str[i] <= '9' && str[i-1] != '\\' {
			num, _ := strconv.Atoi(string(str[i]))
			if num > 0 {
				builder.Write([]byte(strings.Repeat(string(str[i-1]), num)))
			}
			continue
		}
		if i == 0 || i > 0 && str[i-1] != '\\' {
			builder.WriteByte(str[i])
		}
	}
	return builder.String(), nil
}
