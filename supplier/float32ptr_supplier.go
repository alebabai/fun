// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Float32PtrSupplier represents a supplier of results or an error.
type Float32PtrSupplier func() (*float32, error)

// SilentFloat32PtrSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentFloat32PtrSupplier func() *float32

// ToSupplier transforms Float32PtrSupplier into Supplier
func (s Float32PtrSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentFloat32PtrSupplier transforms Float32PtrSupplier into SilentFloat32PtrSupplier
func (s Float32PtrSupplier) ToSilentFloat32PtrSupplier() SilentFloat32PtrSupplier {
	return func() *float32 {
		v, _ := s()
		return v
	}
}

// MustFloat32PtrSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustFloat32PtrSupplier func() *float32

// ToMustFloat32PtrSupplier transforms Float32PtrSupplier into MustFloat32PtrSupplier
func (s Float32PtrSupplier) ToMustFloat32PtrSupplier() MustFloat32PtrSupplier {
	return func() *float32 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToFloat32PtrSupplier transforms MustFloat32PtrSupplier into Float32PtrSupplier
func (ms MustFloat32PtrSupplier) ToFloat32PtrSupplier() Float32PtrSupplier {
	return func() (v *float32, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentFloat32PtrSupplier transforms MustFloat32PtrSupplier into SilentFloat32PtrSupplier
func (ms MustFloat32PtrSupplier) ToSilentFloat32PtrSupplier() SilentFloat32PtrSupplier {
	s := ms.ToFloat32PtrSupplier()
	return s.ToSilentFloat32PtrSupplier()
}
