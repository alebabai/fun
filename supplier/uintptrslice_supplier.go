// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// UintPtrSliceSupplier represents a supplier of results or an error.
type UintPtrSliceSupplier func() ([]*uint, error)

// ToSupplier transforms UintPtrSliceSupplier into Supplier
func (s UintPtrSliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// SilentUintPtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentUintPtrSliceSupplier func() []*uint

// ToSilentUintPtrSliceSupplier transforms UintPtrSliceSupplier into SilentUintPtrSliceSupplier
func (s UintPtrSliceSupplier) ToSilentUintPtrSliceSupplier() SilentUintPtrSliceSupplier {
	return func() []*uint {
		v, _ := s()
		return v
	}
}

// ToSilentSupplier transforms SilentUintPtrSliceSupplier into SilentSupplier
func (ss SilentUintPtrSliceSupplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		return ss()
	}
}

// MustUintPtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustUintPtrSliceSupplier func() []*uint

// ToMustUintPtrSliceSupplier transforms UintPtrSliceSupplier into MustUintPtrSliceSupplier
func (s UintPtrSliceSupplier) ToMustUintPtrSliceSupplier() MustUintPtrSliceSupplier {
	return func() []*uint {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToMustSupplier transforms MustUintPtrSliceSupplier into MustSupplier
func (ms MustUintPtrSliceSupplier) ToMustSupplier() MustSupplier {
	return func() interface{} {
		return ms()
	}
}

// ToUintPtrSliceSupplier transforms MustUintPtrSliceSupplier into UintPtrSliceSupplier
func (ms MustUintPtrSliceSupplier) ToUintPtrSliceSupplier() UintPtrSliceSupplier {
	return func() (v []*uint, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentUintPtrSliceSupplier transforms MustUintPtrSliceSupplier into SilentUintPtrSliceSupplier
func (ms MustUintPtrSliceSupplier) ToSilentUintPtrSliceSupplier() SilentUintPtrSliceSupplier {
	s := ms.ToUintPtrSliceSupplier()
	return s.ToSilentUintPtrSliceSupplier()
}
