// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Uint8SliceSupplier represents a supplier of results or an error.
type Uint8SliceSupplier func() ([]uint8, error)

// ToSupplier transforms Uint8SliceSupplier into Supplier
func (s Uint8SliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// SilentUint8SliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentUint8SliceSupplier func() []uint8

// ToSilentUint8SliceSupplier transforms Uint8SliceSupplier into SilentUint8SliceSupplier
func (s Uint8SliceSupplier) ToSilentUint8SliceSupplier() SilentUint8SliceSupplier {
	return func() []uint8 {
		v, _ := s()
		return v
	}
}

// ToSilentSupplier transforms SilentUint8SliceSupplier into SilentSupplier
func (ss SilentUint8SliceSupplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		return ss()
	}
}

// MustUint8SliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustUint8SliceSupplier func() []uint8

// ToMustUint8SliceSupplier transforms Uint8SliceSupplier into MustUint8SliceSupplier
func (s Uint8SliceSupplier) ToMustUint8SliceSupplier() MustUint8SliceSupplier {
	return func() []uint8 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToMustSupplier transforms MustUint8SliceSupplier into MustSupplier
func (ms MustUint8SliceSupplier) ToMustSupplier() MustSupplier {
	return func() interface{} {
		return ms()
	}
}

// ToUint8SliceSupplier transforms MustUint8SliceSupplier into Uint8SliceSupplier
func (ms MustUint8SliceSupplier) ToUint8SliceSupplier() Uint8SliceSupplier {
	return func() (v []uint8, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentUint8SliceSupplier transforms MustUint8SliceSupplier into SilentUint8SliceSupplier
func (ms MustUint8SliceSupplier) ToSilentUint8SliceSupplier() SilentUint8SliceSupplier {
	s := ms.ToUint8SliceSupplier()
	return s.ToSilentUint8SliceSupplier()
}
