// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Uint32PtrSliceSupplier represents a supplier of results or an error.
type Uint32PtrSliceSupplier func() ([]*uint32, error)

// SilentUint32PtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentUint32PtrSliceSupplier func() []*uint32

// ToSupplier transforms Uint32PtrSliceSupplier into Supplier
func (s Uint32PtrSliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentUint32PtrSliceSupplier transforms Uint32PtrSliceSupplier into SilentUint32PtrSliceSupplier
func (s Uint32PtrSliceSupplier) ToSilentUint32PtrSliceSupplier() SilentUint32PtrSliceSupplier {
	return func() []*uint32 {
		v, _ := s()
		return v
	}
}

// MustUint32PtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustUint32PtrSliceSupplier func() []*uint32

// ToMustUint32PtrSliceSupplier transforms Uint32PtrSliceSupplier into MustUint32PtrSliceSupplier
func (s Uint32PtrSliceSupplier) ToMustUint32PtrSliceSupplier() MustUint32PtrSliceSupplier {
	return func() []*uint32 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToUint32PtrSliceSupplier transforms MustUint32PtrSliceSupplier into Uint32PtrSliceSupplier
func (ms MustUint32PtrSliceSupplier) ToUint32PtrSliceSupplier() Uint32PtrSliceSupplier {
	return func() (v []*uint32, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentUint32PtrSliceSupplier transforms MustUint32PtrSliceSupplier into SilentUint32PtrSliceSupplier
func (ms MustUint32PtrSliceSupplier) ToSilentUint32PtrSliceSupplier() SilentUint32PtrSliceSupplier {
	s := ms.ToUint32PtrSliceSupplier()
	return s.ToSilentUint32PtrSliceSupplier()
}
