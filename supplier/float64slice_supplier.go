// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Float64SliceSupplier represents a supplier of results or an error.
type Float64SliceSupplier func() ([]float64, error)

// SilentFloat64SliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentFloat64SliceSupplier func() []float64

// ToSupplier transforms Float64SliceSupplier into Supplier
func (s Float64SliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentFloat64SliceSupplier transforms Float64SliceSupplier into SilentFloat64SliceSupplier
func (s Float64SliceSupplier) ToSilentFloat64SliceSupplier() SilentFloat64SliceSupplier {
	return func() []float64 {
		v, _ := s()
		return v
	}
}

// MustFloat64SliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustFloat64SliceSupplier func() []float64

// ToMustFloat64SliceSupplier transforms Float64SliceSupplier into MustFloat64SliceSupplier
func (s Float64SliceSupplier) ToMustFloat64SliceSupplier() MustFloat64SliceSupplier {
	return func() []float64 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToFloat64SliceSupplier transforms MustFloat64SliceSupplier into Float64SliceSupplier
func (ms MustFloat64SliceSupplier) ToFloat64SliceSupplier() Float64SliceSupplier {
	return func() (v []float64, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentFloat64SliceSupplier transforms MustFloat64SliceSupplier into SilentFloat64SliceSupplier
func (ms MustFloat64SliceSupplier) ToSilentFloat64SliceSupplier() SilentFloat64SliceSupplier {
	s := ms.ToFloat64SliceSupplier()
	return s.ToSilentFloat64SliceSupplier()
}
