package supplier

// Supplier represents a supplier of results or an error.
type Supplier func() (interface{}, error)

// SilentSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentSupplier func() interface{}

// ToSilentSupplier transforms Supplier into SilentSupplier
func (s Supplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		v, _ := s()
		return v
	}
}

// MustSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustSupplier func() interface{}

// ToMustSupplier transforms Supplier into MustSupplier
func (s Supplier) ToMustSupplier() MustSupplier {
	return func() interface{} {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToSupplier transforms MustSupplier into Supplier
func (ms MustSupplier) ToSupplier() Supplier {
	return func() (v interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentSupplier transforms MustSupplier into SilentSupplier
func (ms MustSupplier) ToSilentSupplier() SilentSupplier {
	s := ms.ToSupplier()
	return s.ToSilentSupplier()
}
