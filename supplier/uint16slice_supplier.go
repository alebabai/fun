// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Uint16SliceSupplier represents a supplier of results or an error.
type Uint16SliceSupplier func() ([]uint16, error)

// SilentUint16SliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentUint16SliceSupplier func() []uint16

// ToSupplier transforms Uint16SliceSupplier into Supplier
func (s Uint16SliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentUint16SliceSupplier transforms Uint16SliceSupplier into SilentUint16SliceSupplier
func (s Uint16SliceSupplier) ToSilentUint16SliceSupplier() SilentUint16SliceSupplier {
	return func() []uint16 {
		v, _ := s()
		return v
	}
}

// MustUint16SliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustUint16SliceSupplier func() []uint16

// ToMustUint16SliceSupplier transforms Uint16SliceSupplier into MustUint16SliceSupplier
func (s Uint16SliceSupplier) ToMustUint16SliceSupplier() MustUint16SliceSupplier {
	return func() []uint16 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToUint16SliceSupplier transforms MustUint16SliceSupplier into Uint16SliceSupplier
func (ms MustUint16SliceSupplier) ToUint16SliceSupplier() Uint16SliceSupplier {
	return func() (v []uint16, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentUint16SliceSupplier transforms MustUint16SliceSupplier into SilentUint16SliceSupplier
func (ms MustUint16SliceSupplier) ToSilentUint16SliceSupplier() SilentUint16SliceSupplier {
	s := ms.ToUint16SliceSupplier()
	return s.ToSilentUint16SliceSupplier()
}
