// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Float64PtrConsumer represents an operation that accepts a single input argument or returns an error.
type Float64PtrConsumer func(v *float64) error

// ToConsumer transforms Float64PtrConsumer into Consumer
func (c Float64PtrConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(*float64))
	}
}

// AndThen returns a composed Float64PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Float64PtrConsumer) AndThen(after Float64PtrConsumer) Float64PtrConsumer {
	if after != nil {
		return func(v *float64) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentFloat64PtrConsumer represents an operation that accepts a single input argument without returning an error.
type SilentFloat64PtrConsumer func(v *float64)

// ToSilentFloat64PtrConsumer transforms Float64PtrConsumer into SilentFloat64PtrConsumer
func (c Float64PtrConsumer) ToSilentFloat64PtrConsumer() SilentFloat64PtrConsumer {
	return func(v *float64) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentFloat64PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentFloat64PtrConsumer) AndThen(after SilentFloat64PtrConsumer) SilentFloat64PtrConsumer {
	if after != nil {
		return func(v *float64) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustFloat64PtrConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustFloat64PtrConsumer func(v *float64)

// ToMustFloat64PtrConsumer transforms Float64PtrConsumer into MustFloat64PtrConsumer
func (c Float64PtrConsumer) ToMustFloat64PtrConsumer() MustFloat64PtrConsumer {
	return func(v *float64) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustFloat64PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustFloat64PtrConsumer) AndThen(after MustFloat64PtrConsumer) MustFloat64PtrConsumer {
	if after != nil {
		return func(v *float64) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToFloat64PtrConsumer transforms MustFloat64PtrConsumer into Float64PtrConsumer
func (mc MustFloat64PtrConsumer) ToFloat64PtrConsumer() Float64PtrConsumer {
	return func(v *float64) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentFloat64PtrConsumer transforms MustFloat64PtrConsumer into SilentFloat64PtrConsumer
func (mc MustFloat64PtrConsumer) ToSilentFloat64PtrConsumer() SilentFloat64PtrConsumer {
	c := mc.ToFloat64PtrConsumer()
	return c.ToSilentFloat64PtrConsumer()
}
