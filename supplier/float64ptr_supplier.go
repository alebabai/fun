// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Float64PtrSupplier represents a supplier of results or an error.
type Float64PtrSupplier func() (*float64, error)

// SilentFloat64PtrSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentFloat64PtrSupplier func() *float64

// ToSupplier transforms Float64PtrSupplier into Supplier
func (s Float64PtrSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentFloat64PtrSupplier transforms Float64PtrSupplier into SilentFloat64PtrSupplier
func (s Float64PtrSupplier) ToSilentFloat64PtrSupplier() SilentFloat64PtrSupplier {
	return func() *float64 {
		v, _ := s()
		return v
	}
}

// MustFloat64PtrSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustFloat64PtrSupplier func() *float64

// ToMustFloat64PtrSupplier transforms Float64PtrSupplier into MustFloat64PtrSupplier
func (s Float64PtrSupplier) ToMustFloat64PtrSupplier() MustFloat64PtrSupplier {
	return func() *float64 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToFloat64PtrSupplier transforms MustFloat64PtrSupplier into Float64PtrSupplier
func (ms MustFloat64PtrSupplier) ToFloat64PtrSupplier() Float64PtrSupplier {
	return func() (v *float64, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentFloat64PtrSupplier transforms MustFloat64PtrSupplier into SilentFloat64PtrSupplier
func (ms MustFloat64PtrSupplier) ToSilentFloat64PtrSupplier() SilentFloat64PtrSupplier {
	s := ms.ToFloat64PtrSupplier()
	return s.ToSilentFloat64PtrSupplier()
}
