// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Int32PtrSliceSupplier represents a supplier of results or an error.
type Int32PtrSliceSupplier func() ([]*int32, error)

// SilentInt32PtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentInt32PtrSliceSupplier func() []*int32

// ToSupplier transforms Int32PtrSliceSupplier into Supplier
func (s Int32PtrSliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentInt32PtrSliceSupplier transforms Int32PtrSliceSupplier into SilentInt32PtrSliceSupplier
func (s Int32PtrSliceSupplier) ToSilentInt32PtrSliceSupplier() SilentInt32PtrSliceSupplier {
	return func() []*int32 {
		v, _ := s()
		return v
	}
}

// MustInt32PtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustInt32PtrSliceSupplier func() []*int32

// ToMustInt32PtrSliceSupplier transforms Int32PtrSliceSupplier into MustInt32PtrSliceSupplier
func (s Int32PtrSliceSupplier) ToMustInt32PtrSliceSupplier() MustInt32PtrSliceSupplier {
	return func() []*int32 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToInt32PtrSliceSupplier transforms MustInt32PtrSliceSupplier into Int32PtrSliceSupplier
func (ms MustInt32PtrSliceSupplier) ToInt32PtrSliceSupplier() Int32PtrSliceSupplier {
	return func() (v []*int32, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentInt32PtrSliceSupplier transforms MustInt32PtrSliceSupplier into SilentInt32PtrSliceSupplier
func (ms MustInt32PtrSliceSupplier) ToSilentInt32PtrSliceSupplier() SilentInt32PtrSliceSupplier {
	s := ms.ToInt32PtrSliceSupplier()
	return s.ToSilentInt32PtrSliceSupplier()
}
