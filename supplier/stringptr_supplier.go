// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// StringPtrSupplier represents a supplier of results or an error.
type StringPtrSupplier func() (*string, error)

// ToSupplier transforms StringPtrSupplier into Supplier
func (s StringPtrSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// SilentStringPtrSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentStringPtrSupplier func() *string

// ToSilentStringPtrSupplier transforms StringPtrSupplier into SilentStringPtrSupplier
func (s StringPtrSupplier) ToSilentStringPtrSupplier() SilentStringPtrSupplier {
	return func() *string {
		v, _ := s()
		return v
	}
}

// ToSilentSupplier transforms SilentStringPtrSupplier into SilentSupplier
func (ss SilentStringPtrSupplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		return ss()
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

// ToMustSupplier transforms MustStringPtrSupplier into MustSupplier
func (ms MustStringPtrSupplier) ToMustSupplier() MustSupplier {
	return func() interface{} {
		return ms()
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
