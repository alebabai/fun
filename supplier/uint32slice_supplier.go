// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Uint32SliceSupplier represents a supplier of results or an error.
type Uint32SliceSupplier func() ([]uint32, error)

// ToSupplier transforms Uint32SliceSupplier into Supplier
func (s Uint32SliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// SilentUint32SliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentUint32SliceSupplier func() []uint32

// ToSilentUint32SliceSupplier transforms Uint32SliceSupplier into SilentUint32SliceSupplier
func (s Uint32SliceSupplier) ToSilentUint32SliceSupplier() SilentUint32SliceSupplier {
	return func() []uint32 {
		v, _ := s()
		return v
	}
}

// ToSilentSupplier transforms SilentUint32SliceSupplier into SilentSupplier
func (ss SilentUint32SliceSupplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		return ss()
	}
}

// MustUint32SliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustUint32SliceSupplier func() []uint32

// ToMustUint32SliceSupplier transforms Uint32SliceSupplier into MustUint32SliceSupplier
func (s Uint32SliceSupplier) ToMustUint32SliceSupplier() MustUint32SliceSupplier {
	return func() []uint32 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToMustSupplier transforms MustUint32SliceSupplier into MustSupplier
func (ms MustUint32SliceSupplier) ToMustSupplier() MustSupplier {
	return func() interface{} {
		return ms()
	}
}

// ToUint32SliceSupplier transforms MustUint32SliceSupplier into Uint32SliceSupplier
func (ms MustUint32SliceSupplier) ToUint32SliceSupplier() Uint32SliceSupplier {
	return func() (v []uint32, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentUint32SliceSupplier transforms MustUint32SliceSupplier into SilentUint32SliceSupplier
func (ms MustUint32SliceSupplier) ToSilentUint32SliceSupplier() SilentUint32SliceSupplier {
	s := ms.ToUint32SliceSupplier()
	return s.ToSilentUint32SliceSupplier()
}
