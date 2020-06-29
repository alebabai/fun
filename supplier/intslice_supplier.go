// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// IntSliceSupplier represents a supplier of results or an error.
type IntSliceSupplier func() ([]int, error)

// SilentIntSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentIntSliceSupplier func() []int

// ToSupplier transforms IntSliceSupplier into Supplier
func (s IntSliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentIntSliceSupplier transforms IntSliceSupplier into SilentIntSliceSupplier
func (s IntSliceSupplier) ToSilentIntSliceSupplier() SilentIntSliceSupplier {
	return func() []int {
		v, _ := s()
		return v
	}
}

// MustIntSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustIntSliceSupplier func() []int

// ToMustIntSliceSupplier transforms IntSliceSupplier into MustIntSliceSupplier
func (s IntSliceSupplier) ToMustIntSliceSupplier() MustIntSliceSupplier {
	return func() []int {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToIntSliceSupplier transforms MustIntSliceSupplier into IntSliceSupplier
func (ms MustIntSliceSupplier) ToIntSliceSupplier() IntSliceSupplier {
	return func() (v []int, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentIntSliceSupplier transforms MustIntSliceSupplier into SilentIntSliceSupplier
func (ms MustIntSliceSupplier) ToSilentIntSliceSupplier() SilentIntSliceSupplier {
	s := ms.ToIntSliceSupplier()
	return s.ToSilentIntSliceSupplier()
}
