// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// RunePtrSliceSupplier represents a supplier of results or an error.
type RunePtrSliceSupplier func() ([]*rune, error)

// ToSupplier transforms RunePtrSliceSupplier into Supplier
func (s RunePtrSliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// SilentRunePtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentRunePtrSliceSupplier func() []*rune

// ToSilentRunePtrSliceSupplier transforms RunePtrSliceSupplier into SilentRunePtrSliceSupplier
func (s RunePtrSliceSupplier) ToSilentRunePtrSliceSupplier() SilentRunePtrSliceSupplier {
	return func() []*rune {
		v, _ := s()
		return v
	}
}

// ToSilentSupplier transforms SilentRunePtrSliceSupplier into SilentSupplier
func (ss SilentRunePtrSliceSupplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		return ss()
	}
}

// MustRunePtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustRunePtrSliceSupplier func() []*rune

// ToMustRunePtrSliceSupplier transforms RunePtrSliceSupplier into MustRunePtrSliceSupplier
func (s RunePtrSliceSupplier) ToMustRunePtrSliceSupplier() MustRunePtrSliceSupplier {
	return func() []*rune {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToMustSupplier transforms MustRunePtrSliceSupplier into MustSupplier
func (ms MustRunePtrSliceSupplier) ToMustSupplier() MustSupplier {
	return func() interface{} {
		return ms()
	}
}

// ToRunePtrSliceSupplier transforms MustRunePtrSliceSupplier into RunePtrSliceSupplier
func (ms MustRunePtrSliceSupplier) ToRunePtrSliceSupplier() RunePtrSliceSupplier {
	return func() (v []*rune, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentRunePtrSliceSupplier transforms MustRunePtrSliceSupplier into SilentRunePtrSliceSupplier
func (ms MustRunePtrSliceSupplier) ToSilentRunePtrSliceSupplier() SilentRunePtrSliceSupplier {
	s := ms.ToRunePtrSliceSupplier()
	return s.ToSilentRunePtrSliceSupplier()
}
