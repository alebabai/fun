// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Uint16Supplier represents a supplier of results or an error.
type Uint16Supplier func() (uint16, error)

// SilentUint16Supplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentUint16Supplier func() uint16

// ToSupplier transforms Uint16Supplier into Supplier
func (s Uint16Supplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentUint16Supplier transforms Uint16Supplier into SilentUint16Supplier
func (s Uint16Supplier) ToSilentUint16Supplier() SilentUint16Supplier {
	return func() uint16 {
		v, _ := s()
		return v
	}
}

// MustUint16Supplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustUint16Supplier func() uint16

// ToMustUint16Supplier transforms Uint16Supplier into MustUint16Supplier
func (s Uint16Supplier) ToMustUint16Supplier() MustUint16Supplier {
	return func() uint16 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToUint16Supplier transforms MustUint16Supplier into Uint16Supplier
func (ms MustUint16Supplier) ToUint16Supplier() Uint16Supplier {
	return func() (v uint16, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentUint16Supplier transforms MustUint16Supplier into SilentUint16Supplier
func (ms MustUint16Supplier) ToSilentUint16Supplier() SilentUint16Supplier {
	s := ms.ToUint16Supplier()
	return s.ToSilentUint16Supplier()
}
