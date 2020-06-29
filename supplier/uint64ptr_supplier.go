// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Uint64PtrSupplier represents a supplier of results or an error.
type Uint64PtrSupplier func() (*uint64, error)

// SilentUint64PtrSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentUint64PtrSupplier func() *uint64

// ToSupplier transforms Uint64PtrSupplier into Supplier
func (s Uint64PtrSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentUint64PtrSupplier transforms Uint64PtrSupplier into SilentUint64PtrSupplier
func (s Uint64PtrSupplier) ToSilentUint64PtrSupplier() SilentUint64PtrSupplier {
	return func() *uint64 {
		v, _ := s()
		return v
	}
}

// MustUint64PtrSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustUint64PtrSupplier func() *uint64

// ToMustUint64PtrSupplier transforms Uint64PtrSupplier into MustUint64PtrSupplier
func (s Uint64PtrSupplier) ToMustUint64PtrSupplier() MustUint64PtrSupplier {
	return func() *uint64 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToUint64PtrSupplier transforms MustUint64PtrSupplier into Uint64PtrSupplier
func (ms MustUint64PtrSupplier) ToUint64PtrSupplier() Uint64PtrSupplier {
	return func() (v *uint64, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentUint64PtrSupplier transforms MustUint64PtrSupplier into SilentUint64PtrSupplier
func (ms MustUint64PtrSupplier) ToSilentUint64PtrSupplier() SilentUint64PtrSupplier {
	s := ms.ToUint64PtrSupplier()
	return s.ToSilentUint64PtrSupplier()
}
