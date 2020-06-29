// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// RunePtrSupplier represents a supplier of results or an error.
type RunePtrSupplier func() (*rune, error)

// SilentRunePtrSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentRunePtrSupplier func() *rune

// ToSupplier transforms RunePtrSupplier into Supplier
func (s RunePtrSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentRunePtrSupplier transforms RunePtrSupplier into SilentRunePtrSupplier
func (s RunePtrSupplier) ToSilentRunePtrSupplier() SilentRunePtrSupplier {
	return func() *rune {
		v, _ := s()
		return v
	}
}

// MustRunePtrSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustRunePtrSupplier func() *rune

// ToMustRunePtrSupplier transforms RunePtrSupplier into MustRunePtrSupplier
func (s RunePtrSupplier) ToMustRunePtrSupplier() MustRunePtrSupplier {
	return func() *rune {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToRunePtrSupplier transforms MustRunePtrSupplier into RunePtrSupplier
func (ms MustRunePtrSupplier) ToRunePtrSupplier() RunePtrSupplier {
	return func() (v *rune, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentRunePtrSupplier transforms MustRunePtrSupplier into SilentRunePtrSupplier
func (ms MustRunePtrSupplier) ToSilentRunePtrSupplier() SilentRunePtrSupplier {
	s := ms.ToRunePtrSupplier()
	return s.ToSilentRunePtrSupplier()
}
