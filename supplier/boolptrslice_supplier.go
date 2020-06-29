// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// BoolPtrSliceSupplier represents a supplier of results or an error.
type BoolPtrSliceSupplier func() ([]*bool, error)

// SilentBoolPtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentBoolPtrSliceSupplier func() []*bool

// ToSupplier transforms BoolPtrSliceSupplier into Supplier
func (s BoolPtrSliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentBoolPtrSliceSupplier transforms BoolPtrSliceSupplier into SilentBoolPtrSliceSupplier
func (s BoolPtrSliceSupplier) ToSilentBoolPtrSliceSupplier() SilentBoolPtrSliceSupplier {
	return func() []*bool {
		v, _ := s()
		return v
	}
}

// MustBoolPtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustBoolPtrSliceSupplier func() []*bool

// ToMustBoolPtrSliceSupplier transforms BoolPtrSliceSupplier into MustBoolPtrSliceSupplier
func (s BoolPtrSliceSupplier) ToMustBoolPtrSliceSupplier() MustBoolPtrSliceSupplier {
	return func() []*bool {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToBoolPtrSliceSupplier transforms MustBoolPtrSliceSupplier into BoolPtrSliceSupplier
func (ms MustBoolPtrSliceSupplier) ToBoolPtrSliceSupplier() BoolPtrSliceSupplier {
	return func() (v []*bool, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentBoolPtrSliceSupplier transforms MustBoolPtrSliceSupplier into SilentBoolPtrSliceSupplier
func (ms MustBoolPtrSliceSupplier) ToSilentBoolPtrSliceSupplier() SilentBoolPtrSliceSupplier {
	s := ms.ToBoolPtrSliceSupplier()
	return s.ToSilentBoolPtrSliceSupplier()
}
