// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// UintSupplier represents a supplier of results or an error.
type UintSupplier func() (uint, error)

// ToSupplier transforms UintSupplier into Supplier
func (s UintSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// SilentUintSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentUintSupplier func() uint

// ToSilentUintSupplier transforms UintSupplier into SilentUintSupplier
func (s UintSupplier) ToSilentUintSupplier() SilentUintSupplier {
	return func() uint {
		v, _ := s()
		return v
	}
}

// ToSilentSupplier transforms SilentUintSupplier into SilentSupplier
func (ss SilentUintSupplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		return ss()
	}
}

// MustUintSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustUintSupplier func() uint

// ToMustUintSupplier transforms UintSupplier into MustUintSupplier
func (s UintSupplier) ToMustUintSupplier() MustUintSupplier {
	return func() uint {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToMustSupplier transforms MustUintSupplier into MustSupplier
func (ms MustUintSupplier) ToMustSupplier() MustSupplier {
	return func() interface{} {
		return ms()
	}
}

// ToUintSupplier transforms MustUintSupplier into UintSupplier
func (ms MustUintSupplier) ToUintSupplier() UintSupplier {
	return func() (v uint, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentUintSupplier transforms MustUintSupplier into SilentUintSupplier
func (ms MustUintSupplier) ToSilentUintSupplier() SilentUintSupplier {
	s := ms.ToUintSupplier()
	return s.ToSilentUintSupplier()
}
