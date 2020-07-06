// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// RuneSliceSupplier represents a supplier of results or an error.
type RuneSliceSupplier func() ([]rune, error)

// ToSupplier transforms RuneSliceSupplier into Supplier
func (s RuneSliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// SilentRuneSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentRuneSliceSupplier func() []rune

// ToSilentRuneSliceSupplier transforms RuneSliceSupplier into SilentRuneSliceSupplier
func (s RuneSliceSupplier) ToSilentRuneSliceSupplier() SilentRuneSliceSupplier {
	return func() []rune {
		v, _ := s()
		return v
	}
}

// ToSilentSupplier transforms SilentRuneSliceSupplier into SilentSupplier
func (ss SilentRuneSliceSupplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		return ss()
	}
}

// MustRuneSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustRuneSliceSupplier func() []rune

// ToMustRuneSliceSupplier transforms RuneSliceSupplier into MustRuneSliceSupplier
func (s RuneSliceSupplier) ToMustRuneSliceSupplier() MustRuneSliceSupplier {
	return func() []rune {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToMustSupplier transforms MustRuneSliceSupplier into MustSupplier
func (ms MustRuneSliceSupplier) ToMustSupplier() MustSupplier {
	return func() interface{} {
		return ms()
	}
}

// ToRuneSliceSupplier transforms MustRuneSliceSupplier into RuneSliceSupplier
func (ms MustRuneSliceSupplier) ToRuneSliceSupplier() RuneSliceSupplier {
	return func() (v []rune, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentRuneSliceSupplier transforms MustRuneSliceSupplier into SilentRuneSliceSupplier
func (ms MustRuneSliceSupplier) ToSilentRuneSliceSupplier() SilentRuneSliceSupplier {
	s := ms.ToRuneSliceSupplier()
	return s.ToSilentRuneSliceSupplier()
}
