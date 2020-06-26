// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Int32Supplier represents a supplier of results or an error.
type Int32Supplier func() (int32, error)

// SilentInt32Supplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentInt32Supplier func() int32

// ToSupplier transforms Int32Supplier into Supplier
func (s Int32Supplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentInt32Supplier transforms Int32Supplier into SilentInt32Supplier
func (s Int32Supplier) ToSilentInt32Supplier() SilentInt32Supplier {
	return func() int32 {
		v, _ := s()
		return v
	}
}

// MustInt32Supplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustInt32Supplier func() int32

// ToMustInt32Supplier transforms Int32Supplier into MustInt32Supplier
func (s Int32Supplier) ToMustInt32Supplier() MustInt32Supplier {
	return func() int32 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToInt32Supplier transforms MustInt32Supplier into Int32Supplier
func (ms MustInt32Supplier) ToInt32Supplier() Int32Supplier {
	return func() (v int32, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentInt32Supplier transforms MustInt32Supplier into SilentInt32Supplier
func (ms MustInt32Supplier) ToSilentInt32Supplier() SilentInt32Supplier {
	s := ms.ToInt32Supplier()
	return s.ToSilentInt32Supplier()
}
