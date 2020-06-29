// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// Int8SliceSupplier represents a supplier of results or an error.
type Int8SliceSupplier func() ([]int8, error)

// SilentInt8SliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentInt8SliceSupplier func() []int8

// ToSupplier transforms Int8SliceSupplier into Supplier
func (s Int8SliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentInt8SliceSupplier transforms Int8SliceSupplier into SilentInt8SliceSupplier
func (s Int8SliceSupplier) ToSilentInt8SliceSupplier() SilentInt8SliceSupplier {
	return func() []int8 {
		v, _ := s()
		return v
	}
}

// MustInt8SliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustInt8SliceSupplier func() []int8

// ToMustInt8SliceSupplier transforms Int8SliceSupplier into MustInt8SliceSupplier
func (s Int8SliceSupplier) ToMustInt8SliceSupplier() MustInt8SliceSupplier {
	return func() []int8 {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToInt8SliceSupplier transforms MustInt8SliceSupplier into Int8SliceSupplier
func (ms MustInt8SliceSupplier) ToInt8SliceSupplier() Int8SliceSupplier {
	return func() (v []int8, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentInt8SliceSupplier transforms MustInt8SliceSupplier into SilentInt8SliceSupplier
func (ms MustInt8SliceSupplier) ToSilentInt8SliceSupplier() SilentInt8SliceSupplier {
	s := ms.ToInt8SliceSupplier()
	return s.ToSilentInt8SliceSupplier()
}
