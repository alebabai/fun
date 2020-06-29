// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// ByteSliceSupplier represents a supplier of results or an error.
type ByteSliceSupplier func() ([]byte, error)

// SilentByteSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentByteSliceSupplier func() []byte

// ToSupplier transforms ByteSliceSupplier into Supplier
func (s ByteSliceSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// ToSilentByteSliceSupplier transforms ByteSliceSupplier into SilentByteSliceSupplier
func (s ByteSliceSupplier) ToSilentByteSliceSupplier() SilentByteSliceSupplier {
	return func() []byte {
		v, _ := s()
		return v
	}
}

// MustByteSliceSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustByteSliceSupplier func() []byte

// ToMustByteSliceSupplier transforms ByteSliceSupplier into MustByteSliceSupplier
func (s ByteSliceSupplier) ToMustByteSliceSupplier() MustByteSliceSupplier {
	return func() []byte {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToByteSliceSupplier transforms MustByteSliceSupplier into ByteSliceSupplier
func (ms MustByteSliceSupplier) ToByteSliceSupplier() ByteSliceSupplier {
	return func() (v []byte, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentByteSliceSupplier transforms MustByteSliceSupplier into SilentByteSliceSupplier
func (ms MustByteSliceSupplier) ToSilentByteSliceSupplier() SilentByteSliceSupplier {
	s := ms.ToByteSliceSupplier()
	return s.ToSilentByteSliceSupplier()
}
