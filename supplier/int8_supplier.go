// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Int8Supplier represents a supplier of results or an error.
type Int8Supplier func() (int8, error)

// SilentInt8Supplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentInt8Supplier func() int8

// ToSupplier transforms Int8Supplier into Supplier
func (s Int8Supplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentInt8Supplier transforms Int8Supplier into SilentInt8Supplier
func (s Int8Supplier) ToSilentInt8Supplier() SilentInt8Supplier {
	return func() int8 {
		v, _ := s()
		return v
	}
}

// MustInt8Supplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustInt8Supplier func() int8

// ToMustInt8Supplier transforms Int8Supplier into MustInt8Supplier
func (s Int8Supplier) ToMustInt8Supplier() MustInt8Supplier {
	return func() int8 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToInt8Supplier transforms MustInt8Supplier into Int8Supplier
func (ms MustInt8Supplier) ToInt8Supplier() Int8Supplier {
	return func() (v int8, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentInt8Supplier transforms MustInt8Supplier into SilentInt8Supplier
func (ms MustInt8Supplier) ToSilentInt8Supplier() SilentInt8Supplier {
	s := ms.ToInt8Supplier()
	return s.ToSilentInt8Supplier()
}
