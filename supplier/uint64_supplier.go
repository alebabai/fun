// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Uint64Supplier represents a supplier of results or an error.
type Uint64Supplier func() (uint64, error)

// ToSupplier transforms Uint64Supplier into Supplier
func (s Uint64Supplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// SilentUint64Supplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentUint64Supplier func() uint64

// ToSilentUint64Supplier transforms Uint64Supplier into SilentUint64Supplier
func (s Uint64Supplier) ToSilentUint64Supplier() SilentUint64Supplier {
	return func() uint64 {
		v, _ := s()
		return v
	}
}

// ToSilentSupplier transforms SilentUint64Supplier into SilentSupplier
func (ss SilentUint64Supplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		return ss()
	}
}

// MustUint64Supplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustUint64Supplier func() uint64

// ToMustUint64Supplier transforms Uint64Supplier into MustUint64Supplier
func (s Uint64Supplier) ToMustUint64Supplier() MustUint64Supplier {
	return func() uint64 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToMustSupplier transforms MustUint64Supplier into MustSupplier
func (ms MustUint64Supplier) ToMustSupplier() MustSupplier {
	return func() interface{} {
		return ms()
	}
}

// ToUint64Supplier transforms MustUint64Supplier into Uint64Supplier
func (ms MustUint64Supplier) ToUint64Supplier() Uint64Supplier {
	return func() (v uint64, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentUint64Supplier transforms MustUint64Supplier into SilentUint64Supplier
func (ms MustUint64Supplier) ToSilentUint64Supplier() SilentUint64Supplier {
	s := ms.ToUint64Supplier()
	return s.ToSilentUint64Supplier()
}
