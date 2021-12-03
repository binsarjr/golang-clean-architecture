package vo

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidKodeKecamatan = errors.New("kode kecamatan tidak valid")

	rxKodeKecamatan = regexp.MustCompile(`^\d{2}$`)
)

// KodeKecamatan structure
type KodeKecamatan struct {
	value string
}

// NewKodeKecamatan create new KodeKecamatan
func NewKodeKecamatan(value string) (KodeKecamatan, error) {
	var e = KodeKecamatan{value: value}

	if !e.validate() {
		return KodeKecamatan{}, ErrInvalidKodeKecamatan
	}

	return e, nil
}

func (e KodeKecamatan) validate() bool {
	return rxKodeKecamatan.MatchString(e.value)
}

// Value return value KodeKecamatan
func (e KodeKecamatan) Value() string {
	return e.value
}

// String returns string representation of the KodeKecamatan
func (e KodeKecamatan) String() string {
	return e.value
}

// Equals checks that two KodeKecamatan are the same
func (e KodeKecamatan) Equals(value Value) bool {
	o, ok := value.(KodeKecamatan)
	return ok && e.value == o.value
}

// NewKodeKecamatanTest create new KodeKecamatan for test
func NewKodeKecamatanTest(e string) KodeKecamatan {
	return KodeKecamatan{value: e}
}
