// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Int16Supplier represents a supplier of results or an error.
type Int16Supplier func() (int16, error)

// ToSupplier transforms Int16Supplier into Supplier
func (s Int16Supplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// SilentInt16Supplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentInt16Supplier func() int16

// ToSilentInt16Supplier transforms Int16Supplier into SilentInt16Supplier
func (s Int16Supplier) ToSilentInt16Supplier() SilentInt16Supplier {
	return func() int16 {
		v, _ := s()
		return v
	}
}

// ToSilentSupplier transforms SilentInt16Supplier into SilentSupplier
func (ss SilentInt16Supplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		return ss()
	}
}

// MustInt16Supplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustInt16Supplier func() int16

// ToMustInt16Supplier transforms Int16Supplier into MustInt16Supplier
func (s Int16Supplier) ToMustInt16Supplier() MustInt16Supplier {
	return func() int16 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToMustSupplier transforms MustInt16Supplier into MustSupplier
func (ms MustInt16Supplier) ToMustSupplier() MustSupplier {
	return func() interface{} {
		return ms()
	}
}

// ToInt16Supplier transforms MustInt16Supplier into Int16Supplier
func (ms MustInt16Supplier) ToInt16Supplier() Int16Supplier {
	return func() (v int16, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentInt16Supplier transforms MustInt16Supplier into SilentInt16Supplier
func (ms MustInt16Supplier) ToSilentInt16Supplier() SilentInt16Supplier {
	s := ms.ToInt16Supplier()
	return s.ToSilentInt16Supplier()
}
