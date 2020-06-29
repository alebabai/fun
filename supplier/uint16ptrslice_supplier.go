// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Uint16PtrSliceSupplier represents a supplier of results or an error.
type Uint16PtrSliceSupplier func() ([]*uint16, error)

// SilentUint16PtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentUint16PtrSliceSupplier func() []*uint16

// ToSupplier transforms Uint16PtrSliceSupplier into Supplier
func (s Uint16PtrSliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentUint16PtrSliceSupplier transforms Uint16PtrSliceSupplier into SilentUint16PtrSliceSupplier
func (s Uint16PtrSliceSupplier) ToSilentUint16PtrSliceSupplier() SilentUint16PtrSliceSupplier {
	return func() []*uint16 {
		v, _ := s()
		return v
	}
}

// MustUint16PtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustUint16PtrSliceSupplier func() []*uint16

// ToMustUint16PtrSliceSupplier transforms Uint16PtrSliceSupplier into MustUint16PtrSliceSupplier
func (s Uint16PtrSliceSupplier) ToMustUint16PtrSliceSupplier() MustUint16PtrSliceSupplier {
	return func() []*uint16 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToUint16PtrSliceSupplier transforms MustUint16PtrSliceSupplier into Uint16PtrSliceSupplier
func (ms MustUint16PtrSliceSupplier) ToUint16PtrSliceSupplier() Uint16PtrSliceSupplier {
	return func() (v []*uint16, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentUint16PtrSliceSupplier transforms MustUint16PtrSliceSupplier into SilentUint16PtrSliceSupplier
func (ms MustUint16PtrSliceSupplier) ToSilentUint16PtrSliceSupplier() SilentUint16PtrSliceSupplier {
	s := ms.ToUint16PtrSliceSupplier()
	return s.ToSilentUint16PtrSliceSupplier()
}
