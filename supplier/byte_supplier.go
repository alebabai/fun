// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier.go.tmpl
// DO NOT EDIT

package supplier

// ByteSupplier represents a supplier of results or an error.
type ByteSupplier func() (byte, error)

// ToSupplier transforms ByteSupplier into Supplier
func (s ByteSupplier) ToSupplier() Supplier {
	return func() (interface{}, error) {
		return s()
	}
}

// SilentByteSupplier represents a supplier of results without returning an error.
// In case of an error it should just return the default value of the type.
type SilentByteSupplier func() byte

// ToSilentByteSupplier transforms ByteSupplier into SilentByteSupplier
func (s ByteSupplier) ToSilentByteSupplier() SilentByteSupplier {
	return func() byte {
		v, _ := s()
		return v
	}
}

// ToSilentSupplier transforms SilentByteSupplier into SilentSupplier
func (ss SilentByteSupplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		return ss()
	}
}

// MustByteSupplier represents a supplier of results without returning an error.
// In case of an error it should panic with error value.
type MustByteSupplier func() byte

// ToMustByteSupplier transforms ByteSupplier into MustByteSupplier
func (s ByteSupplier) ToMustByteSupplier() MustByteSupplier {
	return func() byte {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// ToMustSupplier transforms MustByteSupplier into MustSupplier
func (ms MustByteSupplier) ToMustSupplier() MustSupplier {
	return func() interface{} {
		return ms()
	}
}

// ToByteSupplier transforms MustByteSupplier into ByteSupplier
func (ms MustByteSupplier) ToByteSupplier() ByteSupplier {
	return func() (v byte, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

// ToSilentByteSupplier transforms MustByteSupplier into SilentByteSupplier
func (ms MustByteSupplier) ToSilentByteSupplier() SilentByteSupplier {
	s := ms.ToByteSupplier()
	return s.ToSilentByteSupplier()
}
