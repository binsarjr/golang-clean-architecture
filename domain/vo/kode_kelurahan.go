package vo

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidKodeKelurahan = errors.New("kode kelurahan tidak valid")

	rxKodeKelurahan = regexp.MustCompile(`^\d{4}$`)
)

// KodeKelurahan structure
type KodeKelurahan struct {
	value string
}

// NewKodeKelurahan create new KodeKelurahan
func NewKodeKelurahan(value string) (KodeKelurahan, error) {
	var e = KodeKelurahan{value: value}

	if !e.validate() {
		return KodeKelurahan{}, ErrInvalidKodeKelurahan
	}

	return e, nil
}

func (e KodeKelurahan) validate() bool {
	return rxKodeKelurahan.MatchString(e.value)
}

// Value return value KodeKelurahan
func (e KodeKelurahan) Value() string {
	return e.value
}

// String returns string representation of the KodeKelurahan
func (e KodeKelurahan) String() string {
	return e.value
}

// Equals checks that two KodeKelurahan are the same
func (e KodeKelurahan) Equals(value Value) bool {
	o, ok := value.(KodeKelurahan)
	return ok && e.value == o.value
}

// NewKodeKelurahanTest create new KodeKelurahan for test
func NewKodeKelurahanTest(e string) KodeKelurahan {
	return KodeKelurahan{value: e}
}
