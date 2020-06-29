// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Int16SliceSupplier represents a supplier of results or an error.
type Int16SliceSupplier func() ([]int16, error)

// SilentInt16SliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentInt16SliceSupplier func() []int16

// ToSupplier transforms Int16SliceSupplier into Supplier
func (s Int16SliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentInt16SliceSupplier transforms Int16SliceSupplier into SilentInt16SliceSupplier
func (s Int16SliceSupplier) ToSilentInt16SliceSupplier() SilentInt16SliceSupplier {
	return func() []int16 {
		v, _ := s()
		return v
	}
}

// MustInt16SliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustInt16SliceSupplier func() []int16

// ToMustInt16SliceSupplier transforms Int16SliceSupplier into MustInt16SliceSupplier
func (s Int16SliceSupplier) ToMustInt16SliceSupplier() MustInt16SliceSupplier {
	return func() []int16 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToInt16SliceSupplier transforms MustInt16SliceSupplier into Int16SliceSupplier
func (ms MustInt16SliceSupplier) ToInt16SliceSupplier() Int16SliceSupplier {
	return func() (v []int16, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentInt16SliceSupplier transforms MustInt16SliceSupplier into SilentInt16SliceSupplier
func (ms MustInt16SliceSupplier) ToSilentInt16SliceSupplier() SilentInt16SliceSupplier {
	s := ms.ToInt16SliceSupplier()
	return s.ToSilentInt16SliceSupplier()
}
