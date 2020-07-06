// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// RuneSupplier represents a supplier of results or an error.
type RuneSupplier func() (rune, error)

// ToSupplier transforms RuneSupplier into Supplier
func (s RuneSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// SilentRuneSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentRuneSupplier func() rune

// ToSilentRuneSupplier transforms RuneSupplier into SilentRuneSupplier
func (s RuneSupplier) ToSilentRuneSupplier() SilentRuneSupplier {
	return func() rune {
		v, _ := s()
		return v
	}
}

// ToSilentSupplier transforms SilentRuneSupplier into SilentSupplier
func (ss SilentRuneSupplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		return ss()
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

// ToMustSupplier transforms MustRuneSupplier into MustSupplier
func (ms MustRuneSupplier) ToMustSupplier() MustSupplier {
	return func() interface{} {
		return ms()
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
