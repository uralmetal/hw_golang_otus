package hw02unpackstring

import (
	"testing"

	"github.com/stretchr/testify/require" //nolint:depguard
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "d\n5abc", expected: "d\n\n\n\n\nabc"},
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
		{input: "Hello, 世4界", expected: "Hello, 世世世世界"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	tests := []struct {
		input string
		err   error
	}{
		{input: "3abc", err: ErrInvalidString},
		{input: "45", err: ErrInvalidString},
		{input: "aaa10b", err: ErrInvalidDigit},
		{input: `qw\ne`, err: ErrInvalidEscape},
		{input: `test\`, err: ErrInvalidEscape},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.Equal(t, "", result)
			require.ErrorIs(t, tc.err, err)
		})
	}
}
