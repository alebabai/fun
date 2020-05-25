package supplier

type Supplier func() (interface{}, error)

type SilentSupplier func() interface{}

func (s Supplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		v, _ := s()
		return v
	}
}

type MustSupplier func() interface{}

func (s Supplier) ToMustSupplier() MustSupplier {
	return func() interface{} {
		v, err := s()
		if err != nil {
			panic(err)
		}
		return v
	}
}

func (ms MustSupplier) ToSupplier() Supplier {
	return func() (v interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		v = ms()
		return
	}
}

func (ms MustSupplier) ToSilentSupplier() SilentSupplier {
	s := ms.ToSupplier()
	return s.ToSilentSupplier()
}
