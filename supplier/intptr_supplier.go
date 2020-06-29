// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// IntPtrSupplier represents a supplier of results or an error.
type IntPtrSupplier func() (*int, error)

// SilentIntPtrSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentIntPtrSupplier func() *int

// ToSupplier transforms IntPtrSupplier into Supplier
func (s IntPtrSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentIntPtrSupplier transforms IntPtrSupplier into SilentIntPtrSupplier
func (s IntPtrSupplier) ToSilentIntPtrSupplier() SilentIntPtrSupplier {
	return func() *int {
		v, _ := s()
		return v
	}
}

// MustIntPtrSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustIntPtrSupplier func() *int

// ToMustIntPtrSupplier transforms IntPtrSupplier into MustIntPtrSupplier
func (s IntPtrSupplier) ToMustIntPtrSupplier() MustIntPtrSupplier {
	return func() *int {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToIntPtrSupplier transforms MustIntPtrSupplier into IntPtrSupplier
func (ms MustIntPtrSupplier) ToIntPtrSupplier() IntPtrSupplier {
	return func() (v *int, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentIntPtrSupplier transforms MustIntPtrSupplier into SilentIntPtrSupplier
func (ms MustIntPtrSupplier) ToSilentIntPtrSupplier() SilentIntPtrSupplier {
	s := ms.ToIntPtrSupplier()
	return s.ToSilentIntPtrSupplier()
}
