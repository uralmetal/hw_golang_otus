package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	ErrInvalidString  = errors.New("invalid string")
	ErrInvalidEscape  = errors.New("incorrect escaping")
	ErrInvalidDigit   = errors.New("invalid digit")
	lastCheckedSym    byte
	preLastCheckedSym byte
	escape            = false
)

func isDigit(sym byte) bool {
	return sym >= '0' && sym <= '9'
}

func validateString(str string) error {
	/*
		Проверяется - начинается ли строка
	*/
	if len(str) == 0 {
		return nil
	}
	if isDigit(str[0]) {
		return ErrInvalidString
	}
	if str[len(str)-1] == '\\' {
		return ErrInvalidEscape
	}
	return nil
}

func checkRunTime(sym byte) error {
	if lastCheckedSym == 0 {
		lastCheckedSym = sym
		return nil
	}
	if escape && !isDigit(sym) && sym != '\\' {
		return ErrInvalidEscape
	}
	if preLastCheckedSym != '\\' && isDigit(sym) && isDigit(lastCheckedSym) {
		return ErrInvalidDigit
	}
	preLastCheckedSym = lastCheckedSym
	lastCheckedSym = sym
	return nil
}

func Unpack(str string) (string, error) {
	/*
		Функция для распаковки строки с поддержкой экранирования.
		Строка не валидная, если:
		1. Экранируется только цифра или слэш
		2. Не содержит чисел, а только цифры
		3. Строка заканчивается экранированием
		4. Строка начинается с числа
	*/
	var unpack strings.Builder

	lastCheckedSym = 0
	escape = false
	errorString := validateString(str)
	if errorString != nil {
		return "", errorString
	}
	for i := 0; i < len(str); i++ {
		errorString = checkRunTime(str[i])
		if errorString != nil {
			return "", errorString
		}

		if str[i] == '\\' && !escape {
			escape = true
			continue
		}
		if isDigit(str[i]) && !escape {
			num, _ := strconv.Atoi(string(str[i]))
			unpackString := unpack.String()
			last, sizeLast := utf8.DecodeLastRuneInString(unpackString)
			// ToDo: need type with pop element for improve performance
			unpack.Reset()
			unpack.WriteString(unpackString[0 : len(unpackString)-sizeLast])
			unpack.WriteString(strings.Repeat(string(last), num))
			continue
		}
		unpack.WriteByte(str[i])
		escape = false
	}
	return unpack.String(), nil
}
