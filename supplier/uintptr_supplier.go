// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// UintPtrSupplier represents a supplier of results or an error.
type UintPtrSupplier func() (*uint, error)

// ToSupplier transforms UintPtrSupplier into Supplier
func (s UintPtrSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// SilentUintPtrSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentUintPtrSupplier func() *uint

// ToSilentUintPtrSupplier transforms UintPtrSupplier into SilentUintPtrSupplier
func (s UintPtrSupplier) ToSilentUintPtrSupplier() SilentUintPtrSupplier {
	return func() *uint {
		v, _ := s()
		return v
	}
}

// ToSilentSupplier transforms SilentUintPtrSupplier into SilentSupplier
func (ss SilentUintPtrSupplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		return ss()
	}
}

// MustUintPtrSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustUintPtrSupplier func() *uint

// ToMustUintPtrSupplier transforms UintPtrSupplier into MustUintPtrSupplier
func (s UintPtrSupplier) ToMustUintPtrSupplier() MustUintPtrSupplier {
	return func() *uint {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToMustSupplier transforms MustUintPtrSupplier into MustSupplier
func (ms MustUintPtrSupplier) ToMustSupplier() MustSupplier {
	return func() interface{} {
		return ms()
	}
}

// ToUintPtrSupplier transforms MustUintPtrSupplier into UintPtrSupplier
func (ms MustUintPtrSupplier) ToUintPtrSupplier() UintPtrSupplier {
	return func() (v *uint, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentUintPtrSupplier transforms MustUintPtrSupplier into SilentUintPtrSupplier
func (ms MustUintPtrSupplier) ToSilentUintPtrSupplier() SilentUintPtrSupplier {
	s := ms.ToUintPtrSupplier()
	return s.ToSilentUintPtrSupplier()
}
