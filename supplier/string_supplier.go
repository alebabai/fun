// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// StringSupplier represents a supplier of results or an error.
type StringSupplier func() (string, error)

// SilentStringSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentStringSupplier func() string

// ToSupplier transforms StringSupplier into Supplier
func (s StringSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentStringSupplier transforms StringSupplier into SilentStringSupplier
func (s StringSupplier) ToSilentStringSupplier() SilentStringSupplier {
	return func() string {
		v, _ := s()
		return v
	}
}

// MustStringSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustStringSupplier func() string

// ToMustStringSupplier transforms StringSupplier into MustStringSupplier
func (s StringSupplier) ToMustStringSupplier() MustStringSupplier {
	return func() string {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToStringSupplier transforms MustStringSupplier into StringSupplier
func (ms MustStringSupplier) ToStringSupplier() StringSupplier {
	return func() (v string, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentStringSupplier transforms MustStringSupplier into SilentStringSupplier
func (ms MustStringSupplier) ToSilentStringSupplier() SilentStringSupplier {
	s := ms.ToStringSupplier()
	return s.ToSilentStringSupplier()
}
