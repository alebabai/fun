// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Int64PtrSupplier represents a supplier of results or an error.
type Int64PtrSupplier func() (*int64, error)

// ToSupplier transforms Int64PtrSupplier into Supplier
func (s Int64PtrSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// SilentInt64PtrSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentInt64PtrSupplier func() *int64

// ToSilentInt64PtrSupplier transforms Int64PtrSupplier into SilentInt64PtrSupplier
func (s Int64PtrSupplier) ToSilentInt64PtrSupplier() SilentInt64PtrSupplier {
	return func() *int64 {
		v, _ := s()
		return v
	}
}

// ToSilentSupplier transforms SilentInt64PtrSupplier into SilentSupplier
func (ss SilentInt64PtrSupplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		return ss()
	}
}

// MustInt64PtrSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustInt64PtrSupplier func() *int64

// ToMustInt64PtrSupplier transforms Int64PtrSupplier into MustInt64PtrSupplier
func (s Int64PtrSupplier) ToMustInt64PtrSupplier() MustInt64PtrSupplier {
	return func() *int64 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToMustSupplier transforms MustInt64PtrSupplier into MustSupplier
func (ms MustInt64PtrSupplier) ToMustSupplier() MustSupplier {
	return func() interface{} {
		return ms()
	}
}

// ToInt64PtrSupplier transforms MustInt64PtrSupplier into Int64PtrSupplier
func (ms MustInt64PtrSupplier) ToInt64PtrSupplier() Int64PtrSupplier {
	return func() (v *int64, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentInt64PtrSupplier transforms MustInt64PtrSupplier into SilentInt64PtrSupplier
func (ms MustInt64PtrSupplier) ToSilentInt64PtrSupplier() SilentInt64PtrSupplier {
	s := ms.ToInt64PtrSupplier()
	return s.ToSilentInt64PtrSupplier()
}
