// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Int32PtrSupplier represents a supplier of results or an error.
type Int32PtrSupplier func() (*int32, error)

// ToSupplier transforms Int32PtrSupplier into Supplier
func (s Int32PtrSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// SilentInt32PtrSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentInt32PtrSupplier func() *int32

// ToSilentInt32PtrSupplier transforms Int32PtrSupplier into SilentInt32PtrSupplier
func (s Int32PtrSupplier) ToSilentInt32PtrSupplier() SilentInt32PtrSupplier {
	return func() *int32 {
		v, _ := s()
		return v
	}
}

// ToSilentSupplier transforms SilentInt32PtrSupplier into SilentSupplier
func (ss SilentInt32PtrSupplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		return ss()
	}
}

// MustInt32PtrSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustInt32PtrSupplier func() *int32

// ToMustInt32PtrSupplier transforms Int32PtrSupplier into MustInt32PtrSupplier
func (s Int32PtrSupplier) ToMustInt32PtrSupplier() MustInt32PtrSupplier {
	return func() *int32 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToMustSupplier transforms MustInt32PtrSupplier into MustSupplier
func (ms MustInt32PtrSupplier) ToMustSupplier() MustSupplier {
	return func() interface{} {
		return ms()
	}
}

// ToInt32PtrSupplier transforms MustInt32PtrSupplier into Int32PtrSupplier
func (ms MustInt32PtrSupplier) ToInt32PtrSupplier() Int32PtrSupplier {
	return func() (v *int32, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentInt32PtrSupplier transforms MustInt32PtrSupplier into SilentInt32PtrSupplier
func (ms MustInt32PtrSupplier) ToSilentInt32PtrSupplier() SilentInt32PtrSupplier {
	s := ms.ToInt32PtrSupplier()
	return s.ToSilentInt32PtrSupplier()
}
