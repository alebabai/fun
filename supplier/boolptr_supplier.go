// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// BoolPtrSupplier represents a supplier of results or an error.
type BoolPtrSupplier func() (*bool, error)

// ToSupplier transforms BoolPtrSupplier into Supplier
func (s BoolPtrSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// SilentBoolPtrSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentBoolPtrSupplier func() *bool

// ToSilentBoolPtrSupplier transforms BoolPtrSupplier into SilentBoolPtrSupplier
func (s BoolPtrSupplier) ToSilentBoolPtrSupplier() SilentBoolPtrSupplier {
	return func() *bool {
		v, _ := s()
		return v
	}
}

// ToSilentSupplier transforms SilentBoolPtrSupplier into SilentSupplier
func (ss SilentBoolPtrSupplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		return ss()
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

// ToMustSupplier transforms MustBoolPtrSupplier into MustSupplier
func (ms MustBoolPtrSupplier) ToMustSupplier() MustSupplier {
	return func() interface{} {
		return ms()
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
