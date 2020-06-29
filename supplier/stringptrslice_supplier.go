// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// StringPtrSliceSupplier represents a supplier of results or an error.
type StringPtrSliceSupplier func() ([]*string, error)

// SilentStringPtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentStringPtrSliceSupplier func() []*string

// ToSupplier transforms StringPtrSliceSupplier into Supplier
func (s StringPtrSliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentStringPtrSliceSupplier transforms StringPtrSliceSupplier into SilentStringPtrSliceSupplier
func (s StringPtrSliceSupplier) ToSilentStringPtrSliceSupplier() SilentStringPtrSliceSupplier {
	return func() []*string {
		v, _ := s()
		return v
	}
}

// MustStringPtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustStringPtrSliceSupplier func() []*string

// ToMustStringPtrSliceSupplier transforms StringPtrSliceSupplier into MustStringPtrSliceSupplier
func (s StringPtrSliceSupplier) ToMustStringPtrSliceSupplier() MustStringPtrSliceSupplier {
	return func() []*string {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToStringPtrSliceSupplier transforms MustStringPtrSliceSupplier into StringPtrSliceSupplier
func (ms MustStringPtrSliceSupplier) ToStringPtrSliceSupplier() StringPtrSliceSupplier {
	return func() (v []*string, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentStringPtrSliceSupplier transforms MustStringPtrSliceSupplier into SilentStringPtrSliceSupplier
func (ms MustStringPtrSliceSupplier) ToSilentStringPtrSliceSupplier() SilentStringPtrSliceSupplier {
	s := ms.ToStringPtrSliceSupplier()
	return s.ToSilentStringPtrSliceSupplier()
}
