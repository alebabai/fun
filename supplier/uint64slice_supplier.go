// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Uint64SliceSupplier represents a supplier of results or an error.
type Uint64SliceSupplier func() ([]uint64, error)

// SilentUint64SliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentUint64SliceSupplier func() []uint64

// ToSupplier transforms Uint64SliceSupplier into Supplier
func (s Uint64SliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentUint64SliceSupplier transforms Uint64SliceSupplier into SilentUint64SliceSupplier
func (s Uint64SliceSupplier) ToSilentUint64SliceSupplier() SilentUint64SliceSupplier {
	return func() []uint64 {
		v, _ := s()
		return v
	}
}

// MustUint64SliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustUint64SliceSupplier func() []uint64

// ToMustUint64SliceSupplier transforms Uint64SliceSupplier into MustUint64SliceSupplier
func (s Uint64SliceSupplier) ToMustUint64SliceSupplier() MustUint64SliceSupplier {
	return func() []uint64 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToUint64SliceSupplier transforms MustUint64SliceSupplier into Uint64SliceSupplier
func (ms MustUint64SliceSupplier) ToUint64SliceSupplier() Uint64SliceSupplier {
	return func() (v []uint64, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentUint64SliceSupplier transforms MustUint64SliceSupplier into SilentUint64SliceSupplier
func (ms MustUint64SliceSupplier) ToSilentUint64SliceSupplier() SilentUint64SliceSupplier {
	s := ms.ToUint64SliceSupplier()
	return s.ToSilentUint64SliceSupplier()
}
