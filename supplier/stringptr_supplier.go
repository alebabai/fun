// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// StringPtrSupplier represents a supplier of results or an error.
type StringPtrSupplier func() (*string, error)

// SilentStringPtrSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentStringPtrSupplier func() *string

// ToSupplier transforms StringPtrSupplier into Supplier
func (s StringPtrSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentStringPtrSupplier transforms StringPtrSupplier into SilentStringPtrSupplier
func (s StringPtrSupplier) ToSilentStringPtrSupplier() SilentStringPtrSupplier {
	return func() *string {
		v, _ := s()
		return v
	}
}

// MustStringPtrSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustStringPtrSupplier func() *string

// ToMustStringPtrSupplier transforms StringPtrSupplier into MustStringPtrSupplier
func (s StringPtrSupplier) ToMustStringPtrSupplier() MustStringPtrSupplier {
	return func() *string {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToStringPtrSupplier transforms MustStringPtrSupplier into StringPtrSupplier
func (ms MustStringPtrSupplier) ToStringPtrSupplier() StringPtrSupplier {
	return func() (v *string, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentStringPtrSupplier transforms MustStringPtrSupplier into SilentStringPtrSupplier
func (ms MustStringPtrSupplier) ToSilentStringPtrSupplier() SilentStringPtrSupplier {
	s := ms.ToStringPtrSupplier()
	return s.ToSilentStringPtrSupplier()
}
