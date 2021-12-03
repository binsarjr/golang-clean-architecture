package vo

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidKodeKota = errors.New("kode kota tidak valid")

	rxKodeKota = regexp.MustCompile(`^\d{2}$`)
)

// KodeKota structure
type KodeKota struct {
	value string
}

// NewKodeKota create new KodeKota
func NewKodeKota(value string) (KodeKota, error) {
	var e = KodeKota{value: value}

	if !e.validate() {
		return KodeKota{}, ErrInvalidKodeKota
	}

	return e, nil
}

func (e KodeKota) validate() bool {
	return rxKodeKota.MatchString(e.value)
}

// Value return value KodeKota
func (e KodeKota) Value() string {
	return e.value
}

// String returns string representation of the KodeKota
func (e KodeKota) String() string {
	return e.value
}

// Equals checks that two KodeKota are the same
func (e KodeKota) Equals(value Value) bool {
	o, ok := value.(KodeKota)
	return ok && e.value == o.value
}

// NewKodeKotaTest create new KodeKota for test
func NewKodeKotaTest(e string) KodeKota {
	return KodeKota{value: e}
}
