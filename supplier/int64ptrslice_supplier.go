// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Int64PtrSliceSupplier represents a supplier of results or an error.
type Int64PtrSliceSupplier func() ([]*int64, error)

// ToSupplier transforms Int64PtrSliceSupplier into Supplier
func (s Int64PtrSliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// SilentInt64PtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentInt64PtrSliceSupplier func() []*int64

// ToSilentInt64PtrSliceSupplier transforms Int64PtrSliceSupplier into SilentInt64PtrSliceSupplier
func (s Int64PtrSliceSupplier) ToSilentInt64PtrSliceSupplier() SilentInt64PtrSliceSupplier {
	return func() []*int64 {
		v, _ := s()
		return v
	}
}

// ToSilentSupplier transforms SilentInt64PtrSliceSupplier into SilentSupplier
func (ss SilentInt64PtrSliceSupplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		return ss()
	}
}

// MustInt64PtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustInt64PtrSliceSupplier func() []*int64

// ToMustInt64PtrSliceSupplier transforms Int64PtrSliceSupplier into MustInt64PtrSliceSupplier
func (s Int64PtrSliceSupplier) ToMustInt64PtrSliceSupplier() MustInt64PtrSliceSupplier {
	return func() []*int64 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToMustSupplier transforms MustInt64PtrSliceSupplier into MustSupplier
func (ms MustInt64PtrSliceSupplier) ToMustSupplier() MustSupplier {
	return func() interface{} {
		return ms()
	}
}

// ToInt64PtrSliceSupplier transforms MustInt64PtrSliceSupplier into Int64PtrSliceSupplier
func (ms MustInt64PtrSliceSupplier) ToInt64PtrSliceSupplier() Int64PtrSliceSupplier {
	return func() (v []*int64, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentInt64PtrSliceSupplier transforms MustInt64PtrSliceSupplier into SilentInt64PtrSliceSupplier
func (ms MustInt64PtrSliceSupplier) ToSilentInt64PtrSliceSupplier() SilentInt64PtrSliceSupplier {
	s := ms.ToInt64PtrSliceSupplier()
	return s.ToSilentInt64PtrSliceSupplier()
}
