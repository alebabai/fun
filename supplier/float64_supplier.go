// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Float64Supplier represents a supplier of results or an error.
type Float64Supplier func() (float64, error)

// SilentFloat64Supplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentFloat64Supplier func() float64

// ToSupplier transforms Float64Supplier into Supplier
func (s Float64Supplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentFloat64Supplier transforms Float64Supplier into SilentFloat64Supplier
func (s Float64Supplier) ToSilentFloat64Supplier() SilentFloat64Supplier {
	return func() float64 {
		v, _ := s()
		return v
	}
}

// MustFloat64Supplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustFloat64Supplier func() float64

// ToMustFloat64Supplier transforms Float64Supplier into MustFloat64Supplier
func (s Float64Supplier) ToMustFloat64Supplier() MustFloat64Supplier {
	return func() float64 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToFloat64Supplier transforms MustFloat64Supplier into Float64Supplier
func (ms MustFloat64Supplier) ToFloat64Supplier() Float64Supplier {
	return func() (v float64, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentFloat64Supplier transforms MustFloat64Supplier into SilentFloat64Supplier
func (ms MustFloat64Supplier) ToSilentFloat64Supplier() SilentFloat64Supplier {
	s := ms.ToFloat64Supplier()
	return s.ToSilentFloat64Supplier()
}
