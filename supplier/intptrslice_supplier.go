// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// IntPtrSliceSupplier represents a supplier of results or an error.
type IntPtrSliceSupplier func() ([]*int, error)

// SilentIntPtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentIntPtrSliceSupplier func() []*int

// ToSupplier transforms IntPtrSliceSupplier into Supplier
func (s IntPtrSliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentIntPtrSliceSupplier transforms IntPtrSliceSupplier into SilentIntPtrSliceSupplier
func (s IntPtrSliceSupplier) ToSilentIntPtrSliceSupplier() SilentIntPtrSliceSupplier {
	return func() []*int {
		v, _ := s()
		return v
	}
}

// MustIntPtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustIntPtrSliceSupplier func() []*int

// ToMustIntPtrSliceSupplier transforms IntPtrSliceSupplier into MustIntPtrSliceSupplier
func (s IntPtrSliceSupplier) ToMustIntPtrSliceSupplier() MustIntPtrSliceSupplier {
	return func() []*int {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToIntPtrSliceSupplier transforms MustIntPtrSliceSupplier into IntPtrSliceSupplier
func (ms MustIntPtrSliceSupplier) ToIntPtrSliceSupplier() IntPtrSliceSupplier {
	return func() (v []*int, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentIntPtrSliceSupplier transforms MustIntPtrSliceSupplier into SilentIntPtrSliceSupplier
func (ms MustIntPtrSliceSupplier) ToSilentIntPtrSliceSupplier() SilentIntPtrSliceSupplier {
	s := ms.ToIntPtrSliceSupplier()
	return s.ToSilentIntPtrSliceSupplier()
}
