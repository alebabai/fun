// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Int8PtrSupplier represents a supplier of results or an error.
type Int8PtrSupplier func() (*int8, error)

// ToSupplier transforms Int8PtrSupplier into Supplier
func (s Int8PtrSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// SilentInt8PtrSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentInt8PtrSupplier func() *int8

// ToSilentInt8PtrSupplier transforms Int8PtrSupplier into SilentInt8PtrSupplier
func (s Int8PtrSupplier) ToSilentInt8PtrSupplier() SilentInt8PtrSupplier {
	return func() *int8 {
		v, _ := s()
		return v
	}
}

// ToSilentSupplier transforms SilentInt8PtrSupplier into SilentSupplier
func (ss SilentInt8PtrSupplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		return ss()
	}
}

// MustInt8PtrSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustInt8PtrSupplier func() *int8

// ToMustInt8PtrSupplier transforms Int8PtrSupplier into MustInt8PtrSupplier
func (s Int8PtrSupplier) ToMustInt8PtrSupplier() MustInt8PtrSupplier {
	return func() *int8 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToMustSupplier transforms MustInt8PtrSupplier into MustSupplier
func (ms MustInt8PtrSupplier) ToMustSupplier() MustSupplier {
	return func() interface{} {
		return ms()
	}
}

// ToInt8PtrSupplier transforms MustInt8PtrSupplier into Int8PtrSupplier
func (ms MustInt8PtrSupplier) ToInt8PtrSupplier() Int8PtrSupplier {
	return func() (v *int8, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentInt8PtrSupplier transforms MustInt8PtrSupplier into SilentInt8PtrSupplier
func (ms MustInt8PtrSupplier) ToSilentInt8PtrSupplier() SilentInt8PtrSupplier {
	s := ms.ToInt8PtrSupplier()
	return s.ToSilentInt8PtrSupplier()
}
