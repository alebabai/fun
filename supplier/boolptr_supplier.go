// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// BoolPtrSupplier represents a supplier of results or an error.
type BoolPtrSupplier func() (*bool, error)

// SilentBoolPtrSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentBoolPtrSupplier func() *bool

// ToSupplier transforms BoolPtrSupplier into Supplier
func (s BoolPtrSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentBoolPtrSupplier transforms BoolPtrSupplier into SilentBoolPtrSupplier
func (s BoolPtrSupplier) ToSilentBoolPtrSupplier() SilentBoolPtrSupplier {
	return func() *bool {
		v, _ := s()
		return v
	}
}

// MustBoolPtrSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustBoolPtrSupplier func() *bool

// ToMustBoolPtrSupplier transforms BoolPtrSupplier into MustBoolPtrSupplier
func (s BoolPtrSupplier) ToMustBoolPtrSupplier() MustBoolPtrSupplier {
	return func() *bool {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToBoolPtrSupplier transforms MustBoolPtrSupplier into BoolPtrSupplier
func (ms MustBoolPtrSupplier) ToBoolPtrSupplier() BoolPtrSupplier {
	return func() (v *bool, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentBoolPtrSupplier transforms MustBoolPtrSupplier into SilentBoolPtrSupplier
func (ms MustBoolPtrSupplier) ToSilentBoolPtrSupplier() SilentBoolPtrSupplier {
	s := ms.ToBoolPtrSupplier()
	return s.ToSilentBoolPtrSupplier()
}
