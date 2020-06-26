// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Int64Supplier represents a supplier of results or an error.
type Int64Supplier func() (int64, error)

// SilentInt64Supplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentInt64Supplier func() int64

// ToSupplier transforms Int64Supplier into Supplier
func (s Int64Supplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentInt64Supplier transforms Int64Supplier into SilentInt64Supplier
func (s Int64Supplier) ToSilentInt64Supplier() SilentInt64Supplier {
	return func() int64 {
		v, _ := s()
		return v
	}
}

// MustInt64Supplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustInt64Supplier func() int64

// ToMustInt64Supplier transforms Int64Supplier into MustInt64Supplier
func (s Int64Supplier) ToMustInt64Supplier() MustInt64Supplier {
	return func() int64 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToInt64Supplier transforms MustInt64Supplier into Int64Supplier
func (ms MustInt64Supplier) ToInt64Supplier() Int64Supplier {
	return func() (v int64, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentInt64Supplier transforms MustInt64Supplier into SilentInt64Supplier
func (ms MustInt64Supplier) ToSilentInt64Supplier() SilentInt64Supplier {
	s := ms.ToInt64Supplier()
	return s.ToSilentInt64Supplier()
}
