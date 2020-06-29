// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// UintSliceSupplier represents a supplier of results or an error.
type UintSliceSupplier func() ([]uint, error)

// SilentUintSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentUintSliceSupplier func() []uint

// ToSupplier transforms UintSliceSupplier into Supplier
func (s UintSliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentUintSliceSupplier transforms UintSliceSupplier into SilentUintSliceSupplier
func (s UintSliceSupplier) ToSilentUintSliceSupplier() SilentUintSliceSupplier {
	return func() []uint {
		v, _ := s()
		return v
	}
}

// MustUintSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustUintSliceSupplier func() []uint

// ToMustUintSliceSupplier transforms UintSliceSupplier into MustUintSliceSupplier
func (s UintSliceSupplier) ToMustUintSliceSupplier() MustUintSliceSupplier {
	return func() []uint {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToUintSliceSupplier transforms MustUintSliceSupplier into UintSliceSupplier
func (ms MustUintSliceSupplier) ToUintSliceSupplier() UintSliceSupplier {
	return func() (v []uint, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentUintSliceSupplier transforms MustUintSliceSupplier into SilentUintSliceSupplier
func (ms MustUintSliceSupplier) ToSilentUintSliceSupplier() SilentUintSliceSupplier {
	s := ms.ToUintSliceSupplier()
	return s.ToSilentUintSliceSupplier()
}
