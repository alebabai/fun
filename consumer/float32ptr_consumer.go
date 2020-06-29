// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Float32PtrConsumer represents an operation that accepts a single input argument or returns an error.
type Float32PtrConsumer func(v *float32) error

// ToConsumer transforms Float32PtrConsumer into Consumer
func (c Float32PtrConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(*float32))
	}
}

// AndThen returns a composed Float32PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Float32PtrConsumer) AndThen(after Float32PtrConsumer) Float32PtrConsumer {
	if after != nil {
		return func(v *float32) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentFloat32PtrConsumer represents an operation that accepts a single input argument without returning an error.
type SilentFloat32PtrConsumer func(v *float32)

// ToSilentFloat32PtrConsumer transforms Float32PtrConsumer into SilentFloat32PtrConsumer
func (c Float32PtrConsumer) ToSilentFloat32PtrConsumer() SilentFloat32PtrConsumer {
	return func(v *float32) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentFloat32PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentFloat32PtrConsumer) AndThen(after SilentFloat32PtrConsumer) SilentFloat32PtrConsumer {
	if after != nil {
		return func(v *float32) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustFloat32PtrConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustFloat32PtrConsumer func(v *float32)

// ToMustFloat32PtrConsumer transforms Float32PtrConsumer into MustFloat32PtrConsumer
func (c Float32PtrConsumer) ToMustFloat32PtrConsumer() MustFloat32PtrConsumer {
	return func(v *float32) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustFloat32PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustFloat32PtrConsumer) AndThen(after MustFloat32PtrConsumer) MustFloat32PtrConsumer {
	if after != nil {
		return func(v *float32) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToFloat32PtrConsumer transforms MustFloat32PtrConsumer into Float32PtrConsumer
func (mc MustFloat32PtrConsumer) ToFloat32PtrConsumer() Float32PtrConsumer {
	return func(v *float32) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// SilentFloat32PtrConsumer transforms MustFloat32PtrConsumer into SilentFloat32PtrConsumer
func (mc MustFloat32PtrConsumer) ToSilentFloat32PtrConsumer() SilentFloat32PtrConsumer {
	c := mc.ToFloat32PtrConsumer()
	return c.ToSilentFloat32PtrConsumer()
}
