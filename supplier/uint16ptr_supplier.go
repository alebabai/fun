// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Uint16PtrSupplier represents a supplier of results or an error.
type Uint16PtrSupplier func() (*uint16, error)

// SilentUint16PtrSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentUint16PtrSupplier func() *uint16

// ToSupplier transforms Uint16PtrSupplier into Supplier
func (s Uint16PtrSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentUint16PtrSupplier transforms Uint16PtrSupplier into SilentUint16PtrSupplier
func (s Uint16PtrSupplier) ToSilentUint16PtrSupplier() SilentUint16PtrSupplier {
	return func() *uint16 {
		v, _ := s()
		return v
	}
}

// MustUint16PtrSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustUint16PtrSupplier func() *uint16

// ToMustUint16PtrSupplier transforms Uint16PtrSupplier into MustUint16PtrSupplier
func (s Uint16PtrSupplier) ToMustUint16PtrSupplier() MustUint16PtrSupplier {
	return func() *uint16 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToUint16PtrSupplier transforms MustUint16PtrSupplier into Uint16PtrSupplier
func (ms MustUint16PtrSupplier) ToUint16PtrSupplier() Uint16PtrSupplier {
	return func() (v *uint16, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentUint16PtrSupplier transforms MustUint16PtrSupplier into SilentUint16PtrSupplier
func (ms MustUint16PtrSupplier) ToSilentUint16PtrSupplier() SilentUint16PtrSupplier {
	s := ms.ToUint16PtrSupplier()
	return s.ToSilentUint16PtrSupplier()
}
