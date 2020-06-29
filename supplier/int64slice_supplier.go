// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Int64SliceSupplier represents a supplier of results or an error.
type Int64SliceSupplier func() ([]int64, error)

// SilentInt64SliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentInt64SliceSupplier func() []int64

// ToSupplier transforms Int64SliceSupplier into Supplier
func (s Int64SliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentInt64SliceSupplier transforms Int64SliceSupplier into SilentInt64SliceSupplier
func (s Int64SliceSupplier) ToSilentInt64SliceSupplier() SilentInt64SliceSupplier {
	return func() []int64 {
		v, _ := s()
		return v
	}
}

// MustInt64SliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustInt64SliceSupplier func() []int64

// ToMustInt64SliceSupplier transforms Int64SliceSupplier into MustInt64SliceSupplier
func (s Int64SliceSupplier) ToMustInt64SliceSupplier() MustInt64SliceSupplier {
	return func() []int64 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToInt64SliceSupplier transforms MustInt64SliceSupplier into Int64SliceSupplier
func (ms MustInt64SliceSupplier) ToInt64SliceSupplier() Int64SliceSupplier {
	return func() (v []int64, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentInt64SliceSupplier transforms MustInt64SliceSupplier into SilentInt64SliceSupplier
func (ms MustInt64SliceSupplier) ToSilentInt64SliceSupplier() SilentInt64SliceSupplier {
	s := ms.ToInt64SliceSupplier()
	return s.ToSilentInt64SliceSupplier()
}
