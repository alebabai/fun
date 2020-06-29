// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Uint8Supplier represents a supplier of results or an error.
type Uint8Supplier func() (uint8, error)

// SilentUint8Supplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentUint8Supplier func() uint8

// ToSupplier transforms Uint8Supplier into Supplier
func (s Uint8Supplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentUint8Supplier transforms Uint8Supplier into SilentUint8Supplier
func (s Uint8Supplier) ToSilentUint8Supplier() SilentUint8Supplier {
	return func() uint8 {
		v, _ := s()
		return v
	}
}

// MustUint8Supplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustUint8Supplier func() uint8

// ToMustUint8Supplier transforms Uint8Supplier into MustUint8Supplier
func (s Uint8Supplier) ToMustUint8Supplier() MustUint8Supplier {
	return func() uint8 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToUint8Supplier transforms MustUint8Supplier into Uint8Supplier
func (ms MustUint8Supplier) ToUint8Supplier() Uint8Supplier {
	return func() (v uint8, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentUint8Supplier transforms MustUint8Supplier into SilentUint8Supplier
func (ms MustUint8Supplier) ToSilentUint8Supplier() SilentUint8Supplier {
	s := ms.ToUint8Supplier()
	return s.ToSilentUint8Supplier()
}
