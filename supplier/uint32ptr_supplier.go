// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Uint32PtrSupplier represents a supplier of results or an error.
type Uint32PtrSupplier func() (*uint32, error)

// SilentUint32PtrSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentUint32PtrSupplier func() *uint32

// ToSupplier transforms Uint32PtrSupplier into Supplier
func (s Uint32PtrSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentUint32PtrSupplier transforms Uint32PtrSupplier into SilentUint32PtrSupplier
func (s Uint32PtrSupplier) ToSilentUint32PtrSupplier() SilentUint32PtrSupplier {
	return func() *uint32 {
		v, _ := s()
		return v
	}
}

// MustUint32PtrSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustUint32PtrSupplier func() *uint32

// ToMustUint32PtrSupplier transforms Uint32PtrSupplier into MustUint32PtrSupplier
func (s Uint32PtrSupplier) ToMustUint32PtrSupplier() MustUint32PtrSupplier {
	return func() *uint32 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToUint32PtrSupplier transforms MustUint32PtrSupplier into Uint32PtrSupplier
func (ms MustUint32PtrSupplier) ToUint32PtrSupplier() Uint32PtrSupplier {
	return func() (v *uint32, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentUint32PtrSupplier transforms MustUint32PtrSupplier into SilentUint32PtrSupplier
func (ms MustUint32PtrSupplier) ToSilentUint32PtrSupplier() SilentUint32PtrSupplier {
	s := ms.ToUint32PtrSupplier()
	return s.ToSilentUint32PtrSupplier()
}
