package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrInvalidTag                   = errors.New("invalid tag")
	ErrNonExistedValidationFunction = errors.New("non existed validation function")
	ErrInvalidArgument              = errors.New("invalid argument")
	ErrValidationStringLen          = errors.New("string is not required len")
	ErrValidationStringRegexp       = errors.New("string is not valid by regexp")
	ErrValidationStringIn           = errors.New("string is not contained in valid strings")
	ErrValidationIntMin             = errors.New("number is less than min")
	ErrValidationIntMax             = errors.New("number is more than max")
	ErrValidationIntIn              = errors.New("number is not contained in valid numbers")
)

type (
	ValidatorInt    func(fieldValue int, checkValue string) error
	ValidatorString func(fieldValue, checkValue string) error
)

var validationFunctionsString = map[string]ValidatorString{
	"len":    validateStringLen,
	"regexp": validateStringRegexp,
	"in":     validateStringIn,
}

var validationFunctionsInt = map[string]ValidatorInt{
	"min": validateIntMin,
	"max": validateIntMax,
	"in":  validateIntIn,
}

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	errString := ""
	for _, item := range v {
		errString += fmt.Sprintf("%s: %s\n", item.Field, item.Err)
	}
	return errString
}

func (v ValidationErrors) Add(field string, err error) ValidationErrors {
	if err == nil {
		return v
	}
	return append(v, ValidationError{Field: field, Err: err})
}

func validateIntMin(fieldValue int, checkValue string) error {
	num, err := strconv.Atoi(checkValue)
	if err != nil {
		return ErrInvalidTag
	}
	if fieldValue < num {
		return ErrValidationIntMin
	}
	return nil
}

func validateIntMax(fieldValue int, checkValue string) error {
	num, err := strconv.Atoi(checkValue)
	if err != nil {
		return ErrInvalidTag
	}
	if fieldValue > num {
		return ErrValidationIntMax
	}
	return nil
}

func validateIntIn(fieldValue int, checkValue string) error {
	for _, item := range strings.Split(checkValue, ",") {
		num, err := strconv.Atoi(item)
		if err != nil {
			return ErrInvalidTag
		}
		if fieldValue == num {
			return nil
		}
	}
	return ErrValidationIntIn
}

func validateStringLen(fieldValue string, checkValue string) error {
	num, err := strconv.Atoi(checkValue)
	if err != nil {
		return ErrInvalidTag
	}
	if len(fieldValue) != num {
		return ErrValidationStringLen
	}
	return nil
}

func validateStringRegexp(fieldValue string, checkValue string) error {
	regex, err := regexp.Compile(checkValue)
	if err != nil {
		return ErrInvalidTag
	}
	if !regex.MatchString(fieldValue) {
		return ErrValidationStringRegexp
	}
	return nil
}

func validateStringIn(fieldValue string, checkValue string) error {
	for _, item := range strings.Split(checkValue, ",") {
		if item == fieldValue {
			return nil
		}
	}
	return ErrValidationStringIn
}

func validateValue(
	v interface{}, validationErrors ValidationErrors, validation, fieldName string,
) (ValidationErrors, error) {
	var err error

	valueType := reflect.TypeOf(v).Kind()
	value := reflect.ValueOf(v)
	params := strings.SplitN(validation, ":", 2)
	if len(params) < 2 {
		return validationErrors, ErrInvalidTag
	}
	checkName := params[0]
	checkValue := params[1]
	switch valueType { //nolint:exhaustive
	case reflect.Array, reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			validationErrors, err = validateValue(value.Index(i).Interface(), validationErrors, validation, fieldName)
			if err != nil {
				return validationErrors, err
			}
		}

	case reflect.Int:
		validateFunc := validationFunctionsInt[checkName]
		if validateFunc == nil {
			err = ErrNonExistedValidationFunction
		} else {
			err = validateFunc(int(value.Int()), checkValue)
		}

	case reflect.String:
		validateFunc := validationFunctionsString[checkName]
		if validateFunc == nil {
			err = ErrNonExistedValidationFunction
		} else {
			err = validateFunc(value.String(), checkValue)
		}
	}
	if !errors.Is(ErrInvalidTag, err) && !errors.Is(ErrNonExistedValidationFunction, err) {
		validationErrors = validationErrors.Add(fieldName, err)
		err = nil
	}
	return validationErrors, err
}

func Validate(v interface{}) error {
	var err error

	rv := reflect.ValueOf(v)
	var validationErrors ValidationErrors
	if rv.Kind() != reflect.Struct {
		return ErrInvalidArgument
	}

	for i := 0; i < rv.NumField(); i++ {
		field := rv.Type().Field(i)
		validateTag := field.Tag.Get("validate")
		if validateTag == "" {
			continue
		}
		validations := strings.Split(validateTag, "|")
		for _, validation := range validations {
			validationErrors, err = validateValue(rv.Field(i).Interface(), validationErrors, validation, field.Name)
			if err != nil {
				return err
			}
		}
	}
	if len(validationErrors) == 0 {
		return nil
	}
	return validationErrors
}
