package vo

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidKodeProvinsi = errors.New("kode provinsi tidak valid")

	rxKodeProvinsi = regexp.MustCompile(`^\d{2}$`)
)

// KodeProvinsi structure
type KodeProvinsi struct {
	value string
}

// NewKodeProvinsi create new KodeProvinsi
func NewKodeProvinsi(value string) (KodeProvinsi, error) {
	var e = KodeProvinsi{value: value}

	if !e.validate() {
		return KodeProvinsi{}, ErrInvalidKodeProvinsi
	}

	return e, nil
}

func (e KodeProvinsi) validate() bool {
	return rxKodeProvinsi.MatchString(e.value)
}

// Value return value KodeProvinsi
func (e KodeProvinsi) Value() string {
	return e.value
}

// String returns string representation of the KodeProvinsi
func (e KodeProvinsi) String() string {
	return e.value
}

// Equals checks that two KodeProvinsi are the same
func (e KodeProvinsi) Equals(value Value) bool {
	o, ok := value.(KodeProvinsi)
	return ok && e.value == o.value
}

// NewKodeProvinsiTest create new KodeProvinsi for test
func NewKodeProvinsiTest(e string) KodeProvinsi {
	return KodeProvinsi{value: e}
}
