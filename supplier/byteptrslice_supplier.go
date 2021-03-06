// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// BytePtrSliceSupplier represents a supplier of results or an error.
type BytePtrSliceSupplier func() ([]*byte, error)

// ToSupplier transforms BytePtrSliceSupplier into Supplier
func (s BytePtrSliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// SilentBytePtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentBytePtrSliceSupplier func() []*byte

// ToSilentBytePtrSliceSupplier transforms BytePtrSliceSupplier into SilentBytePtrSliceSupplier
func (s BytePtrSliceSupplier) ToSilentBytePtrSliceSupplier() SilentBytePtrSliceSupplier {
	return func() []*byte {
		v, _ := s()
		return v
	}
}

// ToSilentSupplier transforms SilentBytePtrSliceSupplier into SilentSupplier
func (ss SilentBytePtrSliceSupplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		return ss()
	}
}

// MustBytePtrSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustBytePtrSliceSupplier func() []*byte

// ToMustBytePtrSliceSupplier transforms BytePtrSliceSupplier into MustBytePtrSliceSupplier
func (s BytePtrSliceSupplier) ToMustBytePtrSliceSupplier() MustBytePtrSliceSupplier {
	return func() []*byte {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToMustSupplier transforms MustBytePtrSliceSupplier into MustSupplier
func (ms MustBytePtrSliceSupplier) ToMustSupplier() MustSupplier {
	return func() interface{} {
		return ms()
	}
}

// ToBytePtrSliceSupplier transforms MustBytePtrSliceSupplier into BytePtrSliceSupplier
func (ms MustBytePtrSliceSupplier) ToBytePtrSliceSupplier() BytePtrSliceSupplier {
	return func() (v []*byte, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentBytePtrSliceSupplier transforms MustBytePtrSliceSupplier into SilentBytePtrSliceSupplier
func (ms MustBytePtrSliceSupplier) ToSilentBytePtrSliceSupplier() SilentBytePtrSliceSupplier {
	s := ms.ToBytePtrSliceSupplier()
	return s.ToSilentBytePtrSliceSupplier()
}
