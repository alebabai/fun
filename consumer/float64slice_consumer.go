// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Float64SliceConsumer represents an operation that accepts a single input argument or returns an error.
type Float64SliceConsumer func(v []float64) error

// ToConsumer transforms Float64SliceConsumer into Consumer
func (c Float64SliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]float64))
	}
}

// AndThen returns a composed Float64SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Float64SliceConsumer) AndThen(after Float64SliceConsumer) Float64SliceConsumer {
	if after != nil {
		return func(v []float64) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentFloat64SliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentFloat64SliceConsumer func(v []float64)

// ToSilentFloat64SliceConsumer transforms Float64SliceConsumer into SilentFloat64SliceConsumer
func (c Float64SliceConsumer) ToSilentFloat64SliceConsumer() SilentFloat64SliceConsumer {
	return func(v []float64) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentFloat64SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentFloat64SliceConsumer) AndThen(after SilentFloat64SliceConsumer) SilentFloat64SliceConsumer {
	if after != nil {
		return func(v []float64) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustFloat64SliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustFloat64SliceConsumer func(v []float64)

// ToMustFloat64SliceConsumer transforms Float64SliceConsumer into MustFloat64SliceConsumer
func (c Float64SliceConsumer) ToMustFloat64SliceConsumer() MustFloat64SliceConsumer {
	return func(v []float64) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustFloat64SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustFloat64SliceConsumer) AndThen(after MustFloat64SliceConsumer) MustFloat64SliceConsumer {
	if after != nil {
		return func(v []float64) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToFloat64SliceConsumer transforms MustFloat64SliceConsumer into Float64SliceConsumer
func (mc MustFloat64SliceConsumer) ToFloat64SliceConsumer() Float64SliceConsumer {
	return func(v []float64) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// SilentFloat64SliceConsumer transforms MustFloat64SliceConsumer into SilentFloat64SliceConsumer
func (mc MustFloat64SliceConsumer) ToSilentFloat64SliceConsumer() SilentFloat64SliceConsumer {
	c := mc.ToFloat64SliceConsumer()
	return c.ToSilentFloat64SliceConsumer()
}
