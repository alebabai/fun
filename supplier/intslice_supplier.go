// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// IntSliceSupplier represents a supplier of results or an error.
type IntSliceSupplier func() ([]int, error)

// ToSupplier transforms IntSliceSupplier into Supplier
func (s IntSliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// SilentIntSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentIntSliceSupplier func() []int

// ToSilentIntSliceSupplier transforms IntSliceSupplier into SilentIntSliceSupplier
func (s IntSliceSupplier) ToSilentIntSliceSupplier() SilentIntSliceSupplier {
	return func() []int {
		v, _ := s()
		return v
	}
}

// ToSilentSupplier transforms SilentIntSliceSupplier into SilentSupplier
func (ss SilentIntSliceSupplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		return ss()
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

// ToMustSupplier transforms MustIntSliceSupplier into MustSupplier
func (ms MustIntSliceSupplier) ToMustSupplier() MustSupplier {
	return func() interface{} {
		return ms()
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
