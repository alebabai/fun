// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Int32PtrConsumer represents an operation that accepts a single input argument or returns an error.
type Int32PtrConsumer func(v *int32) error

// ToConsumer transforms Int32PtrConsumer into Consumer
func (c Int32PtrConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(*int32))
	}
}

// AndThen returns a composed Int32PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Int32PtrConsumer) AndThen(after Int32PtrConsumer) Int32PtrConsumer {
	if after != nil {
		return func(v *int32) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentInt32PtrConsumer represents an operation that accepts a single input argument without returning an error.
type SilentInt32PtrConsumer func(v *int32)

// ToSilentInt32PtrConsumer transforms Int32PtrConsumer into SilentInt32PtrConsumer
func (c Int32PtrConsumer) ToSilentInt32PtrConsumer() SilentInt32PtrConsumer {
	return func(v *int32) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentInt32PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentInt32PtrConsumer) AndThen(after SilentInt32PtrConsumer) SilentInt32PtrConsumer {
	if after != nil {
		return func(v *int32) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustInt32PtrConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustInt32PtrConsumer func(v *int32)

// ToMustInt32PtrConsumer transforms Int32PtrConsumer into MustInt32PtrConsumer
func (c Int32PtrConsumer) ToMustInt32PtrConsumer() MustInt32PtrConsumer {
	return func(v *int32) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustInt32PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustInt32PtrConsumer) AndThen(after MustInt32PtrConsumer) MustInt32PtrConsumer {
	if after != nil {
		return func(v *int32) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToInt32PtrConsumer transforms MustInt32PtrConsumer into Int32PtrConsumer
func (mc MustInt32PtrConsumer) ToInt32PtrConsumer() Int32PtrConsumer {
	return func(v *int32) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentInt32PtrConsumer transforms MustInt32PtrConsumer into SilentInt32PtrConsumer
func (mc MustInt32PtrConsumer) ToSilentInt32PtrConsumer() SilentInt32PtrConsumer {
	c := mc.ToInt32PtrConsumer()
	return c.ToSilentInt32PtrConsumer()
}
