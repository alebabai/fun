// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Int32PtrSliceConsumer represents an operation that accepts a single input argument or returns an error.
type Int32PtrSliceConsumer func(v []*int32) error

// ToConsumer transforms Int32PtrSliceConsumer into Consumer
func (c Int32PtrSliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]*int32))
	}
}

// AndThen returns a composed Int32PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Int32PtrSliceConsumer) AndThen(after Int32PtrSliceConsumer) Int32PtrSliceConsumer {
	if after != nil {
		return func(v []*int32) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentInt32PtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentInt32PtrSliceConsumer func(v []*int32)

// ToSilentInt32PtrSliceConsumer transforms Int32PtrSliceConsumer into SilentInt32PtrSliceConsumer
func (c Int32PtrSliceConsumer) ToSilentInt32PtrSliceConsumer() SilentInt32PtrSliceConsumer {
	return func(v []*int32) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentInt32PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentInt32PtrSliceConsumer) AndThen(after SilentInt32PtrSliceConsumer) SilentInt32PtrSliceConsumer {
	if after != nil {
		return func(v []*int32) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustInt32PtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustInt32PtrSliceConsumer func(v []*int32)

// ToMustInt32PtrSliceConsumer transforms Int32PtrSliceConsumer into MustInt32PtrSliceConsumer
func (c Int32PtrSliceConsumer) ToMustInt32PtrSliceConsumer() MustInt32PtrSliceConsumer {
	return func(v []*int32) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustInt32PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustInt32PtrSliceConsumer) AndThen(after MustInt32PtrSliceConsumer) MustInt32PtrSliceConsumer {
	if after != nil {
		return func(v []*int32) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToInt32PtrSliceConsumer transforms MustInt32PtrSliceConsumer into Int32PtrSliceConsumer
func (mc MustInt32PtrSliceConsumer) ToInt32PtrSliceConsumer() Int32PtrSliceConsumer {
	return func(v []*int32) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentInt32PtrSliceConsumer transforms MustInt32PtrSliceConsumer into SilentInt32PtrSliceConsumer
func (mc MustInt32PtrSliceConsumer) ToSilentInt32PtrSliceConsumer() SilentInt32PtrSliceConsumer {
	c := mc.ToInt32PtrSliceConsumer()
	return c.ToSilentInt32PtrSliceConsumer()
}
