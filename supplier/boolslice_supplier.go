// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// BoolSliceSupplier represents a supplier of results or an error.
type BoolSliceSupplier func() ([]bool, error)

// SilentBoolSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentBoolSliceSupplier func() []bool

// ToSupplier transforms BoolSliceSupplier into Supplier
func (s BoolSliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentBoolSliceSupplier transforms BoolSliceSupplier into SilentBoolSliceSupplier
func (s BoolSliceSupplier) ToSilentBoolSliceSupplier() SilentBoolSliceSupplier {
	return func() []bool {
		v, _ := s()
		return v
	}
}

// MustBoolSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustBoolSliceSupplier func() []bool

// ToMustBoolSliceSupplier transforms BoolSliceSupplier into MustBoolSliceSupplier
func (s BoolSliceSupplier) ToMustBoolSliceSupplier() MustBoolSliceSupplier {
	return func() []bool {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToBoolSliceSupplier transforms MustBoolSliceSupplier into BoolSliceSupplier
func (ms MustBoolSliceSupplier) ToBoolSliceSupplier() BoolSliceSupplier {
	return func() (v []bool, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentBoolSliceSupplier transforms MustBoolSliceSupplier into SilentBoolSliceSupplier
func (ms MustBoolSliceSupplier) ToSilentBoolSliceSupplier() SilentBoolSliceSupplier {
	s := ms.ToBoolSliceSupplier()
	return s.ToSilentBoolSliceSupplier()
}
