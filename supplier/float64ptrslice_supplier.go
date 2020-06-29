// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Float64PtrSliceSupplier represents a supplier of results or an error.
type Float64PtrSliceSupplier func() ([]*float64, error)

// SilentFloat64PtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentFloat64PtrSliceSupplier func() []*float64

// ToSupplier transforms Float64PtrSliceSupplier into Supplier
func (s Float64PtrSliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentFloat64PtrSliceSupplier transforms Float64PtrSliceSupplier into SilentFloat64PtrSliceSupplier
func (s Float64PtrSliceSupplier) ToSilentFloat64PtrSliceSupplier() SilentFloat64PtrSliceSupplier {
	return func() []*float64 {
		v, _ := s()
		return v
	}
}

// MustFloat64PtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustFloat64PtrSliceSupplier func() []*float64

// ToMustFloat64PtrSliceSupplier transforms Float64PtrSliceSupplier into MustFloat64PtrSliceSupplier
func (s Float64PtrSliceSupplier) ToMustFloat64PtrSliceSupplier() MustFloat64PtrSliceSupplier {
	return func() []*float64 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToFloat64PtrSliceSupplier transforms MustFloat64PtrSliceSupplier into Float64PtrSliceSupplier
func (ms MustFloat64PtrSliceSupplier) ToFloat64PtrSliceSupplier() Float64PtrSliceSupplier {
	return func() (v []*float64, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentFloat64PtrSliceSupplier transforms MustFloat64PtrSliceSupplier into SilentFloat64PtrSliceSupplier
func (ms MustFloat64PtrSliceSupplier) ToSilentFloat64PtrSliceSupplier() SilentFloat64PtrSliceSupplier {
	s := ms.ToFloat64PtrSliceSupplier()
	return s.ToSilentFloat64PtrSliceSupplier()
}
