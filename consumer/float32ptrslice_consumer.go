// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Float32PtrSliceConsumer represents an operation that accepts a single input argument or returns an error.
type Float32PtrSliceConsumer func(v []*float32) error

// ToConsumer transforms Float32PtrSliceConsumer into Consumer
func (c Float32PtrSliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]*float32))
	}
}

// AndThen returns a composed Float32PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Float32PtrSliceConsumer) AndThen(after Float32PtrSliceConsumer) Float32PtrSliceConsumer {
	if after != nil {
		return func(v []*float32) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentFloat32PtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentFloat32PtrSliceConsumer func(v []*float32)

// ToSilentFloat32PtrSliceConsumer transforms Float32PtrSliceConsumer into SilentFloat32PtrSliceConsumer
func (c Float32PtrSliceConsumer) ToSilentFloat32PtrSliceConsumer() SilentFloat32PtrSliceConsumer {
	return func(v []*float32) {
		_ = c(v)
		return
	}
}

// ToSilentConsumer transforms SilentFloat32PtrSliceConsumer into SilentConsumer
func (sc SilentFloat32PtrSliceConsumer) ToSilentConsumer() SilentConsumer {
	return func(v interface{}) {
		sc(v.([]*float32))
	}
}

// AndThen returns a composed SilentFloat32PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentFloat32PtrSliceConsumer) AndThen(after SilentFloat32PtrSliceConsumer) SilentFloat32PtrSliceConsumer {
	if after != nil {
		return func(v []*float32) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustFloat32PtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustFloat32PtrSliceConsumer func(v []*float32)

// ToMustFloat32PtrSliceConsumer transforms Float32PtrSliceConsumer into MustFloat32PtrSliceConsumer
func (c Float32PtrSliceConsumer) ToMustFloat32PtrSliceConsumer() MustFloat32PtrSliceConsumer {
	return func(v []*float32) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// ToMustConsumer transforms MustFloat32PtrSliceConsumer into MustConsumer
func (mc MustFloat32PtrSliceConsumer) ToMustConsumer() MustConsumer {
	return func(v interface{}) {
		mc(v.([]*float32))
	}
}

// AndThen returns a composed MustFloat32PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustFloat32PtrSliceConsumer) AndThen(after MustFloat32PtrSliceConsumer) MustFloat32PtrSliceConsumer {
	if after != nil {
		return func(v []*float32) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToFloat32PtrSliceConsumer transforms MustFloat32PtrSliceConsumer into Float32PtrSliceConsumer
func (mc MustFloat32PtrSliceConsumer) ToFloat32PtrSliceConsumer() Float32PtrSliceConsumer {
	return func(v []*float32) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentFloat32PtrSliceConsumer transforms MustFloat32PtrSliceConsumer into SilentFloat32PtrSliceConsumer
func (mc MustFloat32PtrSliceConsumer) ToSilentFloat32PtrSliceConsumer() SilentFloat32PtrSliceConsumer {
	c := mc.ToFloat32PtrSliceConsumer()
	return c.ToSilentFloat32PtrSliceConsumer()
}
