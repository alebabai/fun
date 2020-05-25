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

func (s MustSupplier) ToSupplier() Supplier {
	return func() (v interface{}, err error) {
		defer func() {
			if re := recover(); re != nil {
				err = re.(error)
			}
		}()
		v = s()
		return
	}
}

func (s MustSupplier) ToSilentSupplier() SilentSupplier {
	return func() interface{} {
		defer func() {
			recover()
		}()
		return s()
	}
}
