// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Float32Supplier represents a supplier of results or an error.
type Float32Supplier func() (float32, error)

// ToSupplier transforms Float32Supplier into Supplier
func (s Float32Supplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// SilentFloat32Supplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentFloat32Supplier func() float32

// ToSilentFloat32Supplier transforms Float32Supplier into SilentFloat32Supplier
func (s Float32Supplier) ToSilentFloat32Supplier() SilentFloat32Supplier {
	return func() float32 {
		v, _ := s()
		return v
	}
}

// ToSilentSupplier transforms SilentFloat32Supplier into SilentSupplier
func (ss SilentFloat32Supplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		return ss()
	}
}

// MustFloat32Supplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustFloat32Supplier func() float32

// ToMustFloat32Supplier transforms Float32Supplier into MustFloat32Supplier
func (s Float32Supplier) ToMustFloat32Supplier() MustFloat32Supplier {
	return func() float32 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToMustSupplier transforms MustFloat32Supplier into MustSupplier
func (ms MustFloat32Supplier) ToMustSupplier() MustSupplier {
	return func() interface{} {
		return ms()
	}
}

// ToFloat32Supplier transforms MustFloat32Supplier into Float32Supplier
func (ms MustFloat32Supplier) ToFloat32Supplier() Float32Supplier {
	return func() (v float32, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentFloat32Supplier transforms MustFloat32Supplier into SilentFloat32Supplier
func (ms MustFloat32Supplier) ToSilentFloat32Supplier() SilentFloat32Supplier {
	s := ms.ToFloat32Supplier()
	return s.ToSilentFloat32Supplier()
}
