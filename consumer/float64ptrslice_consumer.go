// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Float64PtrSliceConsumer represents an operation that accepts a single input argument or returns an error.
type Float64PtrSliceConsumer func(v []*float64) error

// ToConsumer transforms Float64PtrSliceConsumer into Consumer
func (c Float64PtrSliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]*float64))
	}
}

// AndThen returns a composed Float64PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Float64PtrSliceConsumer) AndThen(after Float64PtrSliceConsumer) Float64PtrSliceConsumer {
	if after != nil {
		return func(v []*float64) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentFloat64PtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentFloat64PtrSliceConsumer func(v []*float64)

// ToSilentFloat64PtrSliceConsumer transforms Float64PtrSliceConsumer into SilentFloat64PtrSliceConsumer
func (c Float64PtrSliceConsumer) ToSilentFloat64PtrSliceConsumer() SilentFloat64PtrSliceConsumer {
	return func(v []*float64) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentFloat64PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentFloat64PtrSliceConsumer) AndThen(after SilentFloat64PtrSliceConsumer) SilentFloat64PtrSliceConsumer {
	if after != nil {
		return func(v []*float64) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustFloat64PtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustFloat64PtrSliceConsumer func(v []*float64)

// ToMustFloat64PtrSliceConsumer transforms Float64PtrSliceConsumer into MustFloat64PtrSliceConsumer
func (c Float64PtrSliceConsumer) ToMustFloat64PtrSliceConsumer() MustFloat64PtrSliceConsumer {
	return func(v []*float64) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustFloat64PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustFloat64PtrSliceConsumer) AndThen(after MustFloat64PtrSliceConsumer) MustFloat64PtrSliceConsumer {
	if after != nil {
		return func(v []*float64) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToFloat64PtrSliceConsumer transforms MustFloat64PtrSliceConsumer into Float64PtrSliceConsumer
func (mc MustFloat64PtrSliceConsumer) ToFloat64PtrSliceConsumer() Float64PtrSliceConsumer {
	return func(v []*float64) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// SilentFloat64PtrSliceConsumer transforms MustFloat64PtrSliceConsumer into SilentFloat64PtrSliceConsumer
func (mc MustFloat64PtrSliceConsumer) ToSilentFloat64PtrSliceConsumer() SilentFloat64PtrSliceConsumer {
	c := mc.ToFloat64PtrSliceConsumer()
	return c.ToSilentFloat64PtrSliceConsumer()
}
