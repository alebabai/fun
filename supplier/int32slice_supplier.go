// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Int32SliceSupplier represents a supplier of results or an error.
type Int32SliceSupplier func() ([]int32, error)

// SilentInt32SliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentInt32SliceSupplier func() []int32

// ToSupplier transforms Int32SliceSupplier into Supplier
func (s Int32SliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentInt32SliceSupplier transforms Int32SliceSupplier into SilentInt32SliceSupplier
func (s Int32SliceSupplier) ToSilentInt32SliceSupplier() SilentInt32SliceSupplier {
	return func() []int32 {
		v, _ := s()
		return v
	}
}

// MustInt32SliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustInt32SliceSupplier func() []int32

// ToMustInt32SliceSupplier transforms Int32SliceSupplier into MustInt32SliceSupplier
func (s Int32SliceSupplier) ToMustInt32SliceSupplier() MustInt32SliceSupplier {
	return func() []int32 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToInt32SliceSupplier transforms MustInt32SliceSupplier into Int32SliceSupplier
func (ms MustInt32SliceSupplier) ToInt32SliceSupplier() Int32SliceSupplier {
	return func() (v []int32, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentInt32SliceSupplier transforms MustInt32SliceSupplier into SilentInt32SliceSupplier
func (ms MustInt32SliceSupplier) ToSilentInt32SliceSupplier() SilentInt32SliceSupplier {
	s := ms.ToInt32SliceSupplier()
	return s.ToSilentInt32SliceSupplier()
}
