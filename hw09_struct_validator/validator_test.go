package hw09structvalidator

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

type UserRole string

// Test the function on different structures and other types.
type (
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

	InvalidRuleLen struct {
		ID string `validate:"len:short"`
	}
	InvalidRuleMin struct {
		Age int `validate:"min:youngster"`
	}
	InvalidRuleMax struct {
		Age int `validate:"max:adult"`
	}
	InvalidRuleInteger struct {
		Code int `validate:"in:one,two"`
	}
	InvalidRuleRegExp struct {
		Email string `validate:"regexp:(@"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name        string
		in          interface{}
		expectedErr error
	}{
		{
			name:        "01 - some string - Error",
			in:          "not a struct",
			expectedErr: ErrUnsupportedType,
		},
		{
			name:        "02 - some number - Error",
			in:          1234,
			expectedErr: ErrUnsupportedType,
		},
		{
			name: "03 - User valid",
			in: User{
				Name:   "all fields valid",
				ID:     "testID123456789101112131415161789012",
				Age:    45,
				Email:  "test@mail.test",
				Role:   "admin",
				Phones: []string{"89611111118", "89614111741"},
				meta:   nil,
			},
			expectedErr: nil,
		},
		{
			name: "04 - User invalid content",
			in: User{
				Name:   "all fields invalid",
				ID:     "invalid",
				Age:    450,
				Email:  "testmail.test",
				Role:   "unknown",
				Phones: []string{"8961", "821265"},
				meta:   nil,
			},
			expectedErr: ValidationErrors{
				{Field: "ID", Err: ErrExactLen},
				{Field: "Age", Err: ErrLessOrEqual},
				{Field: "Email", Err: ErrMatchRegExp},
				{Field: "Role", Err: ErrNotInList},
				{Field: "Phones.0", Err: ErrExactLen},
				{Field: "Phones.1", Err: ErrExactLen},
			},
		},
		{
			name: "05 - User - invalid content",
			in: User{
				Name:   "one phone is invalid age is less then needed",
				ID:     "testID123456789101112131415161789012",
				Age:    12,
				Email:  "test@mail.test",
				Role:   "admin",
				Phones: []string{"89611111118", "8961741"},
				meta:   nil,
			},
			expectedErr: ValidationErrors{
				{Field: "Age", Err: ErrGreaterOrEqual},
				{Field: "Phones.1", Err: ErrExactLen},
			},
		},
		{
			name:        "06 - App - valid content",
			in:          App{Version: "1.2.3"},
			expectedErr: nil,
		},
		{
			name: "07 - App - error version",
			in:   App{Version: "1.2"},
			expectedErr: ValidationErrors{
				{Field: "Version", Err: ErrExactLen},
			},
		},
		{
			name: "08 - Token - valid",
			in: Token{
				Header:    []byte("test"),
				Payload:   []byte("without"),
				Signature: []byte("validation"),
			},
			expectedErr: nil,
		},
		{
			name: "09 - Response - valid",
			in: Response{
				Code: 200,
				Body: "OK",
			},
			expectedErr: nil,
		},
		{
			name: "10 - Response invalid",
			in: Response{
				Code: 401,
				Body: "Unauthorized",
			},
			expectedErr: ValidationErrors{
				{Field: "Code", Err: ErrNotInList},
			},
		},
		{
			name:        "11 - validate rule len",
			in:          InvalidRuleLen{ID: "ID123456"},
			expectedErr: ErrInvalidRule,
		},
		{
			name:        "12 - validate rule min",
			in:          InvalidRuleMin{Age: 12},
			expectedErr: ErrInvalidRule,
		},
		{
			name:        "13 - validate rule max",
			in:          InvalidRuleMax{Age: 50},
			expectedErr: ErrInvalidRule,
		},
		{
			name:        "14 - validate rule integer",
			in:          InvalidRuleInteger{Code: 200},
			expectedErr: ErrInvalidRule,
		},
		{
			name:        "15 - validate rule RegExp",
			in:          InvalidRuleRegExp{Email: "test@mail.test"},
			expectedErr: ErrInvalidRule,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("case %s", tt.name), func(t *testing.T) {
			tt := tt
			t.Parallel()

			err := Validate(tt.in)

			if tt.expectedErr == nil {
				require.NoError(t, err)
				return
			}

			require.Error(t, err)

			var wantValidateErrs ValidationErrors
			if errors.As(tt.expectedErr, &wantValidateErrs) {
				var gotValidateErrs ValidationErrors

				require.ErrorAs(t, err, &gotValidateErrs)
				require.Len(t, gotValidateErrs, len(wantValidateErrs))

				for j, gotE := range gotValidateErrs {
					wantE := wantValidateErrs[j]

					require.Equal(t, wantE.Field, gotE.Field)
					require.ErrorIs(t, gotE.Err, wantE.Err)
				}
			} else {
				require.ErrorIs(t, err, tt.expectedErr)
			}
		})
	}
}
