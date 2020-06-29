// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// UintPtrSliceSupplier represents a supplier of results or an error.
type UintPtrSliceSupplier func() ([]*uint, error)

// SilentUintPtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentUintPtrSliceSupplier func() []*uint

// ToSupplier transforms UintPtrSliceSupplier into Supplier
func (s UintPtrSliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentUintPtrSliceSupplier transforms UintPtrSliceSupplier into SilentUintPtrSliceSupplier
func (s UintPtrSliceSupplier) ToSilentUintPtrSliceSupplier() SilentUintPtrSliceSupplier {
	return func() []*uint {
		v, _ := s()
		return v
	}
}

// MustUintPtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustUintPtrSliceSupplier func() []*uint

// ToMustUintPtrSliceSupplier transforms UintPtrSliceSupplier into MustUintPtrSliceSupplier
func (s UintPtrSliceSupplier) ToMustUintPtrSliceSupplier() MustUintPtrSliceSupplier {
	return func() []*uint {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToUintPtrSliceSupplier transforms MustUintPtrSliceSupplier into UintPtrSliceSupplier
func (ms MustUintPtrSliceSupplier) ToUintPtrSliceSupplier() UintPtrSliceSupplier {
	return func() (v []*uint, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentUintPtrSliceSupplier transforms MustUintPtrSliceSupplier into SilentUintPtrSliceSupplier
func (ms MustUintPtrSliceSupplier) ToSilentUintPtrSliceSupplier() SilentUintPtrSliceSupplier {
	s := ms.ToUintPtrSliceSupplier()
	return s.ToSilentUintPtrSliceSupplier()
}
