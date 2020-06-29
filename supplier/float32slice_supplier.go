// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Float32SliceSupplier represents a supplier of results or an error.
type Float32SliceSupplier func() ([]float32, error)

// SilentFloat32SliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentFloat32SliceSupplier func() []float32

// ToSupplier transforms Float32SliceSupplier into Supplier
func (s Float32SliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentFloat32SliceSupplier transforms Float32SliceSupplier into SilentFloat32SliceSupplier
func (s Float32SliceSupplier) ToSilentFloat32SliceSupplier() SilentFloat32SliceSupplier {
	return func() []float32 {
		v, _ := s()
		return v
	}
}

// MustFloat32SliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustFloat32SliceSupplier func() []float32

// ToMustFloat32SliceSupplier transforms Float32SliceSupplier into MustFloat32SliceSupplier
func (s Float32SliceSupplier) ToMustFloat32SliceSupplier() MustFloat32SliceSupplier {
	return func() []float32 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToFloat32SliceSupplier transforms MustFloat32SliceSupplier into Float32SliceSupplier
func (ms MustFloat32SliceSupplier) ToFloat32SliceSupplier() Float32SliceSupplier {
	return func() (v []float32, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentFloat32SliceSupplier transforms MustFloat32SliceSupplier into SilentFloat32SliceSupplier
func (ms MustFloat32SliceSupplier) ToSilentFloat32SliceSupplier() SilentFloat32SliceSupplier {
	s := ms.ToFloat32SliceSupplier()
	return s.ToSilentFloat32SliceSupplier()
}
