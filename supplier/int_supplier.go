// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// IntSupplier represents a supplier of results or an error.
type IntSupplier func() (int, error)

// ToSupplier transforms IntSupplier into Supplier
func (s IntSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// SilentIntSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentIntSupplier func() int

// ToSilentIntSupplier transforms IntSupplier into SilentIntSupplier
func (s IntSupplier) ToSilentIntSupplier() SilentIntSupplier {
	return func() int {
		v, _ := s()
		return v
	}
}

// ToSilentSupplier transforms SilentIntSupplier into SilentSupplier
func (ss SilentIntSupplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		return ss()
	}
}

// MustIntSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustIntSupplier func() int

// ToMustIntSupplier transforms IntSupplier into MustIntSupplier
func (s IntSupplier) ToMustIntSupplier() MustIntSupplier {
	return func() int {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToMustSupplier transforms MustIntSupplier into MustSupplier
func (ms MustIntSupplier) ToMustSupplier() MustSupplier {
	return func() interface{} {
		return ms()
	}
}

// ToIntSupplier transforms MustIntSupplier into IntSupplier
func (ms MustIntSupplier) ToIntSupplier() IntSupplier {
	return func() (v int, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentIntSupplier transforms MustIntSupplier into SilentIntSupplier
func (ms MustIntSupplier) ToSilentIntSupplier() SilentIntSupplier {
	s := ms.ToIntSupplier()
	return s.ToSilentIntSupplier()
}
