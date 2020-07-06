// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Uint32Supplier represents a supplier of results or an error.
type Uint32Supplier func() (uint32, error)

// ToSupplier transforms Uint32Supplier into Supplier
func (s Uint32Supplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// SilentUint32Supplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentUint32Supplier func() uint32

// ToSilentUint32Supplier transforms Uint32Supplier into SilentUint32Supplier
func (s Uint32Supplier) ToSilentUint32Supplier() SilentUint32Supplier {
	return func() uint32 {
		v, _ := s()
		return v
	}
}

// ToSilentSupplier transforms SilentUint32Supplier into SilentSupplier
func (ss SilentUint32Supplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		return ss()
	}
}

// MustUint32Supplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustUint32Supplier func() uint32

// ToMustUint32Supplier transforms Uint32Supplier into MustUint32Supplier
func (s Uint32Supplier) ToMustUint32Supplier() MustUint32Supplier {
	return func() uint32 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToMustSupplier transforms MustUint32Supplier into MustSupplier
func (ms MustUint32Supplier) ToMustSupplier() MustSupplier {
	return func() interface{} {
		return ms()
	}
}

// ToUint32Supplier transforms MustUint32Supplier into Uint32Supplier
func (ms MustUint32Supplier) ToUint32Supplier() Uint32Supplier {
	return func() (v uint32, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentUint32Supplier transforms MustUint32Supplier into SilentUint32Supplier
func (ms MustUint32Supplier) ToSilentUint32Supplier() SilentUint32Supplier {
	s := ms.ToUint32Supplier()
	return s.ToSilentUint32Supplier()
}
