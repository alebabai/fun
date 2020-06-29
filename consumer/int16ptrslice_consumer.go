// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Int16PtrSliceConsumer represents an operation that accepts a single input argument or returns an error.
type Int16PtrSliceConsumer func(v []*int16) error

// ToConsumer transforms Int16PtrSliceConsumer into Consumer
func (c Int16PtrSliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]*int16))
	}
}

// AndThen returns a composed Int16PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Int16PtrSliceConsumer) AndThen(after Int16PtrSliceConsumer) Int16PtrSliceConsumer {
	if after != nil {
		return func(v []*int16) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentInt16PtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentInt16PtrSliceConsumer func(v []*int16)

// ToSilentInt16PtrSliceConsumer transforms Int16PtrSliceConsumer into SilentInt16PtrSliceConsumer
func (c Int16PtrSliceConsumer) ToSilentInt16PtrSliceConsumer() SilentInt16PtrSliceConsumer {
	return func(v []*int16) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentInt16PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentInt16PtrSliceConsumer) AndThen(after SilentInt16PtrSliceConsumer) SilentInt16PtrSliceConsumer {
	if after != nil {
		return func(v []*int16) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustInt16PtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustInt16PtrSliceConsumer func(v []*int16)

// ToMustInt16PtrSliceConsumer transforms Int16PtrSliceConsumer into MustInt16PtrSliceConsumer
func (c Int16PtrSliceConsumer) ToMustInt16PtrSliceConsumer() MustInt16PtrSliceConsumer {
	return func(v []*int16) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustInt16PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustInt16PtrSliceConsumer) AndThen(after MustInt16PtrSliceConsumer) MustInt16PtrSliceConsumer {
	if after != nil {
		return func(v []*int16) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToInt16PtrSliceConsumer transforms MustInt16PtrSliceConsumer into Int16PtrSliceConsumer
func (mc MustInt16PtrSliceConsumer) ToInt16PtrSliceConsumer() Int16PtrSliceConsumer {
	return func(v []*int16) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// SilentInt16PtrSliceConsumer transforms MustInt16PtrSliceConsumer into SilentInt16PtrSliceConsumer
func (mc MustInt16PtrSliceConsumer) ToSilentInt16PtrSliceConsumer() SilentInt16PtrSliceConsumer {
	c := mc.ToInt16PtrSliceConsumer()
	return c.ToSilentInt16PtrSliceConsumer()
}
