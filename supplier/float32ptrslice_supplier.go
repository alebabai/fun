// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Float32PtrSliceSupplier represents a supplier of results or an error.
type Float32PtrSliceSupplier func() ([]*float32, error)

// SilentFloat32PtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentFloat32PtrSliceSupplier func() []*float32

// ToSupplier transforms Float32PtrSliceSupplier into Supplier
func (s Float32PtrSliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentFloat32PtrSliceSupplier transforms Float32PtrSliceSupplier into SilentFloat32PtrSliceSupplier
func (s Float32PtrSliceSupplier) ToSilentFloat32PtrSliceSupplier() SilentFloat32PtrSliceSupplier {
	return func() []*float32 {
		v, _ := s()
		return v
	}
}

// MustFloat32PtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustFloat32PtrSliceSupplier func() []*float32

// ToMustFloat32PtrSliceSupplier transforms Float32PtrSliceSupplier into MustFloat32PtrSliceSupplier
func (s Float32PtrSliceSupplier) ToMustFloat32PtrSliceSupplier() MustFloat32PtrSliceSupplier {
	return func() []*float32 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToFloat32PtrSliceSupplier transforms MustFloat32PtrSliceSupplier into Float32PtrSliceSupplier
func (ms MustFloat32PtrSliceSupplier) ToFloat32PtrSliceSupplier() Float32PtrSliceSupplier {
	return func() (v []*float32, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentFloat32PtrSliceSupplier transforms MustFloat32PtrSliceSupplier into SilentFloat32PtrSliceSupplier
func (ms MustFloat32PtrSliceSupplier) ToSilentFloat32PtrSliceSupplier() SilentFloat32PtrSliceSupplier {
	s := ms.ToFloat32PtrSliceSupplier()
	return s.ToSilentFloat32PtrSliceSupplier()
}
