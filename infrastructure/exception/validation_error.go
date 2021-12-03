package exception

import (
	"errors"
	"fmt"
)

func NewMissingRequired(field string) error {
	msg := fmt.Sprintf("%s wajib diisi.", field)
	return errors.New(msg)
}

func NewValidationNotMatch(field string, targetField string) error {
	msg := fmt.Sprintf("%s dan %s harus sama.", field, targetField)
	return errors.New(msg)
}
