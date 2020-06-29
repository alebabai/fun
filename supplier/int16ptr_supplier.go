// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Int16PtrSupplier represents a supplier of results or an error.
type Int16PtrSupplier func() (*int16, error)

// SilentInt16PtrSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentInt16PtrSupplier func() *int16

// ToSupplier transforms Int16PtrSupplier into Supplier
func (s Int16PtrSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentInt16PtrSupplier transforms Int16PtrSupplier into SilentInt16PtrSupplier
func (s Int16PtrSupplier) ToSilentInt16PtrSupplier() SilentInt16PtrSupplier {
	return func() *int16 {
		v, _ := s()
		return v
	}
}

// MustInt16PtrSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustInt16PtrSupplier func() *int16

// ToMustInt16PtrSupplier transforms Int16PtrSupplier into MustInt16PtrSupplier
func (s Int16PtrSupplier) ToMustInt16PtrSupplier() MustInt16PtrSupplier {
	return func() *int16 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToInt16PtrSupplier transforms MustInt16PtrSupplier into Int16PtrSupplier
func (ms MustInt16PtrSupplier) ToInt16PtrSupplier() Int16PtrSupplier {
	return func() (v *int16, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentInt16PtrSupplier transforms MustInt16PtrSupplier into SilentInt16PtrSupplier
func (ms MustInt16PtrSupplier) ToSilentInt16PtrSupplier() SilentInt16PtrSupplier {
	s := ms.ToInt16PtrSupplier()
	return s.ToSilentInt16PtrSupplier()
}
