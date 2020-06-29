// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Uint8PtrSliceSupplier represents a supplier of results or an error.
type Uint8PtrSliceSupplier func() ([]*uint8, error)

// SilentUint8PtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentUint8PtrSliceSupplier func() []*uint8

// ToSupplier transforms Uint8PtrSliceSupplier into Supplier
func (s Uint8PtrSliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentUint8PtrSliceSupplier transforms Uint8PtrSliceSupplier into SilentUint8PtrSliceSupplier
func (s Uint8PtrSliceSupplier) ToSilentUint8PtrSliceSupplier() SilentUint8PtrSliceSupplier {
	return func() []*uint8 {
		v, _ := s()
		return v
	}
}

// MustUint8PtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustUint8PtrSliceSupplier func() []*uint8

// ToMustUint8PtrSliceSupplier transforms Uint8PtrSliceSupplier into MustUint8PtrSliceSupplier
func (s Uint8PtrSliceSupplier) ToMustUint8PtrSliceSupplier() MustUint8PtrSliceSupplier {
	return func() []*uint8 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToUint8PtrSliceSupplier transforms MustUint8PtrSliceSupplier into Uint8PtrSliceSupplier
func (ms MustUint8PtrSliceSupplier) ToUint8PtrSliceSupplier() Uint8PtrSliceSupplier {
	return func() (v []*uint8, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentUint8PtrSliceSupplier transforms MustUint8PtrSliceSupplier into SilentUint8PtrSliceSupplier
func (ms MustUint8PtrSliceSupplier) ToSilentUint8PtrSliceSupplier() SilentUint8PtrSliceSupplier {
	s := ms.ToUint8PtrSliceSupplier()
	return s.ToSilentUint8PtrSliceSupplier()
}
