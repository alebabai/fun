// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// BytePtrSupplier represents a supplier of results or an error.
type BytePtrSupplier func() (*byte, error)

// ToSupplier transforms BytePtrSupplier into Supplier
func (s BytePtrSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// SilentBytePtrSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentBytePtrSupplier func() *byte

// ToSilentBytePtrSupplier transforms BytePtrSupplier into SilentBytePtrSupplier
func (s BytePtrSupplier) ToSilentBytePtrSupplier() SilentBytePtrSupplier {
	return func() *byte {
		v, _ := s()
		return v
	}
}

// ToSilentSupplier transforms SilentBytePtrSupplier into SilentSupplier
func (ss SilentBytePtrSupplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		return ss()
	}
}

// MustBytePtrSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustBytePtrSupplier func() *byte

// ToMustBytePtrSupplier transforms BytePtrSupplier into MustBytePtrSupplier
func (s BytePtrSupplier) ToMustBytePtrSupplier() MustBytePtrSupplier {
	return func() *byte {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToMustSupplier transforms MustBytePtrSupplier into MustSupplier
func (ms MustBytePtrSupplier) ToMustSupplier() MustSupplier {
	return func() interface{} {
		return ms()
	}
}

// ToBytePtrSupplier transforms MustBytePtrSupplier into BytePtrSupplier
func (ms MustBytePtrSupplier) ToBytePtrSupplier() BytePtrSupplier {
	return func() (v *byte, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentBytePtrSupplier transforms MustBytePtrSupplier into SilentBytePtrSupplier
func (ms MustBytePtrSupplier) ToSilentBytePtrSupplier() SilentBytePtrSupplier {
	s := ms.ToBytePtrSupplier()
	return s.ToSilentBytePtrSupplier()
}
