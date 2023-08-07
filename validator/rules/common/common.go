package common

import (
	"fmt"
	"unicode/utf8"
	"validator-refactor/config"
)

type CommonValidator struct{}

func (v *CommonValidator) Validate(config config.Config) error {
	return ValidateMaxLength(config.GetName())
}

const maxLength = 63

func ValidateMaxLength(field string) error {
	if utf8.RuneCountInString(field) > maxLength {
		return fmt.Errorf("%s must be less than or equal to %d characters", field, maxLength)
	} else {
		fmt.Printf("%s is of valid length\n", field)
	}
	return nil
}
