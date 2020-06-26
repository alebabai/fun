// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// RuneSupplier represents a supplier of results or an error.
type RuneSupplier func() (rune, error)

// SilentRuneSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentRuneSupplier func() rune

// ToSupplier transforms RuneSupplier into Supplier
func (s RuneSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentRuneSupplier transforms RuneSupplier into SilentRuneSupplier
func (s RuneSupplier) ToSilentRuneSupplier() SilentRuneSupplier {
	return func() rune {
		v, _ := s()
		return v
	}
}

// MustRuneSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustRuneSupplier func() rune

// ToMustRuneSupplier transforms RuneSupplier into MustRuneSupplier
func (s RuneSupplier) ToMustRuneSupplier() MustRuneSupplier {
	return func() rune {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToRuneSupplier transforms MustRuneSupplier into RuneSupplier
func (ms MustRuneSupplier) ToRuneSupplier() RuneSupplier {
	return func() (v rune, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentRuneSupplier transforms MustRuneSupplier into SilentRuneSupplier
func (ms MustRuneSupplier) ToSilentRuneSupplier() SilentRuneSupplier {
	s := ms.ToRuneSupplier()
	return s.ToSilentRuneSupplier()
}
