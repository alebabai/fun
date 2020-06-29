// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Uint64PtrSliceSupplier represents a supplier of results or an error.
type Uint64PtrSliceSupplier func() ([]*uint64, error)

// SilentUint64PtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentUint64PtrSliceSupplier func() []*uint64

// ToSupplier transforms Uint64PtrSliceSupplier into Supplier
func (s Uint64PtrSliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentUint64PtrSliceSupplier transforms Uint64PtrSliceSupplier into SilentUint64PtrSliceSupplier
func (s Uint64PtrSliceSupplier) ToSilentUint64PtrSliceSupplier() SilentUint64PtrSliceSupplier {
	return func() []*uint64 {
		v, _ := s()
		return v
	}
}

// MustUint64PtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustUint64PtrSliceSupplier func() []*uint64

// ToMustUint64PtrSliceSupplier transforms Uint64PtrSliceSupplier into MustUint64PtrSliceSupplier
func (s Uint64PtrSliceSupplier) ToMustUint64PtrSliceSupplier() MustUint64PtrSliceSupplier {
	return func() []*uint64 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToUint64PtrSliceSupplier transforms MustUint64PtrSliceSupplier into Uint64PtrSliceSupplier
func (ms MustUint64PtrSliceSupplier) ToUint64PtrSliceSupplier() Uint64PtrSliceSupplier {
	return func() (v []*uint64, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentUint64PtrSliceSupplier transforms MustUint64PtrSliceSupplier into SilentUint64PtrSliceSupplier
func (ms MustUint64PtrSliceSupplier) ToSilentUint64PtrSliceSupplier() SilentUint64PtrSliceSupplier {
	s := ms.ToUint64PtrSliceSupplier()
	return s.ToSilentUint64PtrSliceSupplier()
}
