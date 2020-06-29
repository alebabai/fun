// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Float32SliceConsumer represents an operation that accepts a single input argument or returns an error.
type Float32SliceConsumer func(v []float32) error

// ToConsumer transforms Float32SliceConsumer into Consumer
func (c Float32SliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]float32))
	}
}

// AndThen returns a composed Float32SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Float32SliceConsumer) AndThen(after Float32SliceConsumer) Float32SliceConsumer {
	if after != nil {
		return func(v []float32) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentFloat32SliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentFloat32SliceConsumer func(v []float32)

// ToSilentFloat32SliceConsumer transforms Float32SliceConsumer into SilentFloat32SliceConsumer
func (c Float32SliceConsumer) ToSilentFloat32SliceConsumer() SilentFloat32SliceConsumer {
	return func(v []float32) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentFloat32SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentFloat32SliceConsumer) AndThen(after SilentFloat32SliceConsumer) SilentFloat32SliceConsumer {
	if after != nil {
		return func(v []float32) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustFloat32SliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustFloat32SliceConsumer func(v []float32)

// ToMustFloat32SliceConsumer transforms Float32SliceConsumer into MustFloat32SliceConsumer
func (c Float32SliceConsumer) ToMustFloat32SliceConsumer() MustFloat32SliceConsumer {
	return func(v []float32) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustFloat32SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustFloat32SliceConsumer) AndThen(after MustFloat32SliceConsumer) MustFloat32SliceConsumer {
	if after != nil {
		return func(v []float32) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToFloat32SliceConsumer transforms MustFloat32SliceConsumer into Float32SliceConsumer
func (mc MustFloat32SliceConsumer) ToFloat32SliceConsumer() Float32SliceConsumer {
	return func(v []float32) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// SilentFloat32SliceConsumer transforms MustFloat32SliceConsumer into SilentFloat32SliceConsumer
func (mc MustFloat32SliceConsumer) ToSilentFloat32SliceConsumer() SilentFloat32SliceConsumer {
	c := mc.ToFloat32SliceConsumer()
	return c.ToSilentFloat32SliceConsumer()
}
