// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// BoolSupplier represents a supplier of results or an error.
type BoolSupplier func() (bool, error)

// ToSupplier transforms BoolSupplier into Supplier
func (s BoolSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// SilentBoolSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentBoolSupplier func() bool

// ToSilentBoolSupplier transforms BoolSupplier into SilentBoolSupplier
func (s BoolSupplier) ToSilentBoolSupplier() SilentBoolSupplier {
	return func() bool {
		v, _ := s()
		return v
	}
}

// ToSilentSupplier transforms SilentBoolSupplier into SilentSupplier
func (ss SilentBoolSupplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		return ss()
	}
}

// MustBoolSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustBoolSupplier func() bool

// ToMustBoolSupplier transforms BoolSupplier into MustBoolSupplier
func (s BoolSupplier) ToMustBoolSupplier() MustBoolSupplier {
	return func() bool {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToMustSupplier transforms MustBoolSupplier into MustSupplier
func (ms MustBoolSupplier) ToMustSupplier() MustSupplier {
	return func() interface{} {
		return ms()
	}
}

// ToBoolSupplier transforms MustBoolSupplier into BoolSupplier
func (ms MustBoolSupplier) ToBoolSupplier() BoolSupplier {
	return func() (v bool, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentBoolSupplier transforms MustBoolSupplier into SilentBoolSupplier
func (ms MustBoolSupplier) ToSilentBoolSupplier() SilentBoolSupplier {
	s := ms.ToBoolSupplier()
	return s.ToSilentBoolSupplier()
}
