// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Uint8PtrSupplier represents a supplier of results or an error.
type Uint8PtrSupplier func() (*uint8, error)

// SilentUint8PtrSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentUint8PtrSupplier func() *uint8

// ToSupplier transforms Uint8PtrSupplier into Supplier
func (s Uint8PtrSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentUint8PtrSupplier transforms Uint8PtrSupplier into SilentUint8PtrSupplier
func (s Uint8PtrSupplier) ToSilentUint8PtrSupplier() SilentUint8PtrSupplier {
	return func() *uint8 {
		v, _ := s()
		return v
	}
}

// MustUint8PtrSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustUint8PtrSupplier func() *uint8

// ToMustUint8PtrSupplier transforms Uint8PtrSupplier into MustUint8PtrSupplier
func (s Uint8PtrSupplier) ToMustUint8PtrSupplier() MustUint8PtrSupplier {
	return func() *uint8 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToUint8PtrSupplier transforms MustUint8PtrSupplier into Uint8PtrSupplier
func (ms MustUint8PtrSupplier) ToUint8PtrSupplier() Uint8PtrSupplier {
	return func() (v *uint8, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentUint8PtrSupplier transforms MustUint8PtrSupplier into SilentUint8PtrSupplier
func (ms MustUint8PtrSupplier) ToSilentUint8PtrSupplier() SilentUint8PtrSupplier {
	s := ms.ToUint8PtrSupplier()
	return s.ToSilentUint8PtrSupplier()
}
