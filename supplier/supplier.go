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
