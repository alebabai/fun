// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Int16PtrSliceSupplier represents a supplier of results or an error.
type Int16PtrSliceSupplier func() ([]*int16, error)

// ToSupplier transforms Int16PtrSliceSupplier into Supplier
func (s Int16PtrSliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// SilentInt16PtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentInt16PtrSliceSupplier func() []*int16

// ToSilentInt16PtrSliceSupplier transforms Int16PtrSliceSupplier into SilentInt16PtrSliceSupplier
func (s Int16PtrSliceSupplier) ToSilentInt16PtrSliceSupplier() SilentInt16PtrSliceSupplier {
	return func() []*int16 {
		v, _ := s()
		return v
	}
}

// ToSilentSupplier transforms SilentInt16PtrSliceSupplier into SilentSupplier
func (ss SilentInt16PtrSliceSupplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		return ss()
	}
}

// MustInt16PtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustInt16PtrSliceSupplier func() []*int16

// ToMustInt16PtrSliceSupplier transforms Int16PtrSliceSupplier into MustInt16PtrSliceSupplier
func (s Int16PtrSliceSupplier) ToMustInt16PtrSliceSupplier() MustInt16PtrSliceSupplier {
	return func() []*int16 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToMustSupplier transforms MustInt16PtrSliceSupplier into MustSupplier
func (ms MustInt16PtrSliceSupplier) ToMustSupplier() MustSupplier {
	return func() interface{} {
		return ms()
	}
}

// ToInt16PtrSliceSupplier transforms MustInt16PtrSliceSupplier into Int16PtrSliceSupplier
func (ms MustInt16PtrSliceSupplier) ToInt16PtrSliceSupplier() Int16PtrSliceSupplier {
	return func() (v []*int16, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentInt16PtrSliceSupplier transforms MustInt16PtrSliceSupplier into SilentInt16PtrSliceSupplier
func (ms MustInt16PtrSliceSupplier) ToSilentInt16PtrSliceSupplier() SilentInt16PtrSliceSupplier {
	s := ms.ToInt16PtrSliceSupplier()
	return s.ToSilentInt16PtrSliceSupplier()
}
