// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// StringSliceSupplier represents a supplier of results or an error.
type StringSliceSupplier func() ([]string, error)

// ToSupplier transforms StringSliceSupplier into Supplier
func (s StringSliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// SilentStringSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentStringSliceSupplier func() []string

// ToSilentStringSliceSupplier transforms StringSliceSupplier into SilentStringSliceSupplier
func (s StringSliceSupplier) ToSilentStringSliceSupplier() SilentStringSliceSupplier {
	return func() []string {
		v, _ := s()
		return v
	}
}

// ToSilentSupplier transforms SilentStringSliceSupplier into SilentSupplier
func (ss SilentStringSliceSupplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		return ss()
	}
}

// MustStringSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustStringSliceSupplier func() []string

// ToMustStringSliceSupplier transforms StringSliceSupplier into MustStringSliceSupplier
func (s StringSliceSupplier) ToMustStringSliceSupplier() MustStringSliceSupplier {
	return func() []string {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToMustSupplier transforms MustStringSliceSupplier into MustSupplier
func (ms MustStringSliceSupplier) ToMustSupplier() MustSupplier {
	return func() interface{} {
		return ms()
	}
}

// ToStringSliceSupplier transforms MustStringSliceSupplier into StringSliceSupplier
func (ms MustStringSliceSupplier) ToStringSliceSupplier() StringSliceSupplier {
	return func() (v []string, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentStringSliceSupplier transforms MustStringSliceSupplier into SilentStringSliceSupplier
func (ms MustStringSliceSupplier) ToSilentStringSliceSupplier() SilentStringSliceSupplier {
	s := ms.ToStringSliceSupplier()
	return s.ToSilentStringSliceSupplier()
}
