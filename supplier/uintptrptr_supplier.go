// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// UintptrPtrSupplier represents a supplier of results or an error.
type UintptrPtrSupplier func() (*uintptr, error)

// SilentUintptrPtrSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentUintptrPtrSupplier func() *uintptr

// ToSupplier transforms UintptrPtrSupplier into Supplier
func (s UintptrPtrSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentUintptrPtrSupplier transforms UintptrPtrSupplier into SilentUintptrPtrSupplier
func (s UintptrPtrSupplier) ToSilentUintptrPtrSupplier() SilentUintptrPtrSupplier {
	return func() *uintptr {
		v, _ := s()
		return v
	}
}

// MustUintptrPtrSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustUintptrPtrSupplier func() *uintptr

// ToMustUintptrPtrSupplier transforms UintptrPtrSupplier into MustUintptrPtrSupplier
func (s UintptrPtrSupplier) ToMustUintptrPtrSupplier() MustUintptrPtrSupplier {
	return func() *uintptr {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToUintptrPtrSupplier transforms MustUintptrPtrSupplier into UintptrPtrSupplier
func (ms MustUintptrPtrSupplier) ToUintptrPtrSupplier() UintptrPtrSupplier {
	return func() (v *uintptr, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentUintptrPtrSupplier transforms MustUintptrPtrSupplier into SilentUintptrPtrSupplier
func (ms MustUintptrPtrSupplier) ToSilentUintptrPtrSupplier() SilentUintptrPtrSupplier {
	s := ms.ToUintptrPtrSupplier()
	return s.ToSilentUintptrPtrSupplier()
}
