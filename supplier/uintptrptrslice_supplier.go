// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// UintptrPtrSliceSupplier represents a supplier of results or an error.
type UintptrPtrSliceSupplier func() ([]*uintptr, error)

// SilentUintptrPtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentUintptrPtrSliceSupplier func() []*uintptr

// ToSupplier transforms UintptrPtrSliceSupplier into Supplier
func (s UintptrPtrSliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentUintptrPtrSliceSupplier transforms UintptrPtrSliceSupplier into SilentUintptrPtrSliceSupplier
func (s UintptrPtrSliceSupplier) ToSilentUintptrPtrSliceSupplier() SilentUintptrPtrSliceSupplier {
	return func() []*uintptr {
		v, _ := s()
		return v
	}
}

// MustUintptrPtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustUintptrPtrSliceSupplier func() []*uintptr

// ToMustUintptrPtrSliceSupplier transforms UintptrPtrSliceSupplier into MustUintptrPtrSliceSupplier
func (s UintptrPtrSliceSupplier) ToMustUintptrPtrSliceSupplier() MustUintptrPtrSliceSupplier {
	return func() []*uintptr {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToUintptrPtrSliceSupplier transforms MustUintptrPtrSliceSupplier into UintptrPtrSliceSupplier
func (ms MustUintptrPtrSliceSupplier) ToUintptrPtrSliceSupplier() UintptrPtrSliceSupplier {
	return func() (v []*uintptr, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentUintptrPtrSliceSupplier transforms MustUintptrPtrSliceSupplier into SilentUintptrPtrSliceSupplier
func (ms MustUintptrPtrSliceSupplier) ToSilentUintptrPtrSliceSupplier() SilentUintptrPtrSliceSupplier {
	s := ms.ToUintptrPtrSliceSupplier()
	return s.ToSilentUintptrPtrSliceSupplier()
}
