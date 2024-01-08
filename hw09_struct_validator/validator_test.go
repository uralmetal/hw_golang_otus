package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require" //nolint:depguard
)

type UserRole string

// Test the function on different structures and other types.
type (
	NonExistedTag struct {
		Field string `validate:"nonexisted:tag"`
	}

	InvalidStringLenTag struct {
		Field string `validate:"len:tag"`
	}

	InvalidStringRegexpTag struct {
		Field string `validate:"regexp:["`
	}

	InvalidIntMinTag struct {
		Field int `validate:"min:tag"`
	}

	InvalidIntMaxTag struct {
		Field int `validate:"max:tag"`
	}

	InvalidIntInTag struct {
		Field int `validate:"in:tag1,tag2"`
	}

	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in:          "invalid argument",
			expectedErr: ErrInvalidArgument,
		},
		{
			in:          NonExistedTag{Field: "test"},
			expectedErr: ErrNonExistedValidationFunction,
		},
		{
			in:          InvalidStringLenTag{Field: "test"},
			expectedErr: ErrInvalidTag,
		},
		{
			in:          InvalidStringRegexpTag{Field: "test"},
			expectedErr: ErrInvalidTag,
		},
		{
			in:          InvalidIntMinTag{Field: 123},
			expectedErr: ErrInvalidTag,
		},
		{
			in:          InvalidIntMaxTag{Field: 123},
			expectedErr: ErrInvalidTag,
		},
		{
			in:          InvalidIntInTag{Field: 123},
			expectedErr: ErrInvalidTag,
		},
		{
			in: User{
				ID:     "012345678901234567890123456789012345",
				Name:   "test",
				Age:    30,
				Email:  "test@test.com",
				Role:   "stuff",
				Phones: []string{"89012345671"},
				meta:   json.RawMessage("{}"),
			},
			expectedErr: nil,
		},
		{
			in: User{
				ID:     "0123456789012345678",
				Name:   "test",
				Age:    3000,
				Email:  "test-test.com",
				Role:   "kiddy",
				Phones: []string{"invalid"},
				meta:   json.RawMessage("{}"),
			},
			expectedErr: ValidationErrors{
				ValidationError{Field: "ID", Err: ErrValidationStringLen},
				ValidationError{Field: "Age", Err: ErrValidationIntMax},
				ValidationError{Field: "Email", Err: ErrValidationStringRegexp},
				ValidationError{Field: "Role", Err: ErrValidationStringIn},
				ValidationError{Field: "Phones", Err: ErrValidationStringLen},
			},
		},
		{
			in: User{
				ID:     "012345678901234567890123456789012345",
				Name:   "test",
				Age:    2,
				Email:  "test@test.com",
				Role:   "stuff",
				Phones: []string{"89012345671"},
				meta:   json.RawMessage("{}"),
			},
			expectedErr: ValidationErrors{
				ValidationError{Field: "Age", Err: ErrValidationIntMin},
			},
		},
		{
			in:          App{Version: "12345"},
			expectedErr: nil,
		},
		{
			in: App{Version: "123456"},
			expectedErr: ValidationErrors{
				ValidationError{Field: "Version", Err: ErrValidationStringLen},
			},
		},
		{
			in: Token{
				Header:    []byte("Header"),
				Payload:   []byte("Payload"),
				Signature: []byte("Signature"),
			},
			expectedErr: nil,
		},
		{
			in:          Response{Body: "Body", Code: 200},
			expectedErr: nil,
		},
		{
			in: Response{Body: "Body", Code: 900},
			expectedErr: ValidationErrors{
				ValidationError{Field: "Code", Err: ErrValidationIntIn},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			validateErrors := Validate(tt.in)
			require.Equal(t, tt.expectedErr, validateErrors)
		})
	}
}
