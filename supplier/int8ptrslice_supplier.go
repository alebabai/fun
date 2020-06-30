// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Int8PtrSliceSupplier represents a supplier of results or an error.
type Int8PtrSliceSupplier func() ([]*int8, error)

// ToSupplier transforms Int8PtrSliceSupplier into Supplier
func (s Int8PtrSliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// SilentInt8PtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentInt8PtrSliceSupplier func() []*int8

// ToSilentInt8PtrSliceSupplier transforms Int8PtrSliceSupplier into SilentInt8PtrSliceSupplier
func (s Int8PtrSliceSupplier) ToSilentInt8PtrSliceSupplier() SilentInt8PtrSliceSupplier {
	return func() []*int8 {
		v, _ := s()
		return v
	}
}

// ToSilentSupplier transforms SilentInt8PtrSliceSupplier into SilentSupplier
func (ss SilentInt8PtrSliceSupplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		return ss()
	}
}

// MustInt8PtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustInt8PtrSliceSupplier func() []*int8

// ToMustInt8PtrSliceSupplier transforms Int8PtrSliceSupplier into MustInt8PtrSliceSupplier
func (s Int8PtrSliceSupplier) ToMustInt8PtrSliceSupplier() MustInt8PtrSliceSupplier {
	return func() []*int8 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToMustSupplier transforms MustInt8PtrSliceSupplier into MustSupplier
func (ms MustInt8PtrSliceSupplier) ToMustSupplier() MustSupplier {
	return func() interface{} {
		return ms()
	}
}

// ToInt8PtrSliceSupplier transforms MustInt8PtrSliceSupplier into Int8PtrSliceSupplier
func (ms MustInt8PtrSliceSupplier) ToInt8PtrSliceSupplier() Int8PtrSliceSupplier {
	return func() (v []*int8, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentInt8PtrSliceSupplier transforms MustInt8PtrSliceSupplier into SilentInt8PtrSliceSupplier
func (ms MustInt8PtrSliceSupplier) ToSilentInt8PtrSliceSupplier() SilentInt8PtrSliceSupplier {
	s := ms.ToInt8PtrSliceSupplier()
	return s.ToSilentInt8PtrSliceSupplier()
}
