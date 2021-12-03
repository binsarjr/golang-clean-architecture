package exception

import (
	"errors"
	"fmt"
	"net/http"
)

func ErrValidationIfNeeded(err error) {
	if err != nil {
		panic(ErrResponse{Code: http.StatusUnprocessableEntity, Message: err.Error()})
	}
}

func NewMissingRequired(field string) error {
	msg := fmt.Sprintf("%s wajib diisi.", field)
	return errors.New(msg)
}

func NewValidationNotMatch(field string, targetField string) error {
	msg := fmt.Sprintf("%s dan %s harus sama.", field, targetField)
	return errors.New(msg)
}
