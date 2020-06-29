// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Int64PtrConsumer represents an operation that accepts a single input argument or returns an error.
type Int64PtrConsumer func(v *int64) error

// ToConsumer transforms Int64PtrConsumer into Consumer
func (c Int64PtrConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(*int64))
	}
}

// AndThen returns a composed Int64PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Int64PtrConsumer) AndThen(after Int64PtrConsumer) Int64PtrConsumer {
	if after != nil {
		return func(v *int64) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentInt64PtrConsumer represents an operation that accepts a single input argument without returning an error.
type SilentInt64PtrConsumer func(v *int64)

// ToSilentInt64PtrConsumer transforms Int64PtrConsumer into SilentInt64PtrConsumer
func (c Int64PtrConsumer) ToSilentInt64PtrConsumer() SilentInt64PtrConsumer {
	return func(v *int64) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentInt64PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentInt64PtrConsumer) AndThen(after SilentInt64PtrConsumer) SilentInt64PtrConsumer {
	if after != nil {
		return func(v *int64) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustInt64PtrConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustInt64PtrConsumer func(v *int64)

// ToMustInt64PtrConsumer transforms Int64PtrConsumer into MustInt64PtrConsumer
func (c Int64PtrConsumer) ToMustInt64PtrConsumer() MustInt64PtrConsumer {
	return func(v *int64) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustInt64PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustInt64PtrConsumer) AndThen(after MustInt64PtrConsumer) MustInt64PtrConsumer {
	if after != nil {
		return func(v *int64) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToInt64PtrConsumer transforms MustInt64PtrConsumer into Int64PtrConsumer
func (mc MustInt64PtrConsumer) ToInt64PtrConsumer() Int64PtrConsumer {
	return func(v *int64) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// SilentInt64PtrConsumer transforms MustInt64PtrConsumer into SilentInt64PtrConsumer
func (mc MustInt64PtrConsumer) ToSilentInt64PtrConsumer() SilentInt64PtrConsumer {
	c := mc.ToInt64PtrConsumer()
	return c.ToSilentInt64PtrConsumer()
}
