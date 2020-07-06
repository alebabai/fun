// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Int8PtrConsumer represents an operation that accepts a single input argument or returns an error.
type Int8PtrConsumer func(v *int8) error

// ToConsumer transforms Int8PtrConsumer into Consumer
func (c Int8PtrConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(*int8))
	}
}

// AndThen returns a composed Int8PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Int8PtrConsumer) AndThen(after Int8PtrConsumer) Int8PtrConsumer {
	if after != nil {
		return func(v *int8) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentInt8PtrConsumer represents an operation that accepts a single input argument without returning an error.
type SilentInt8PtrConsumer func(v *int8)

// ToSilentInt8PtrConsumer transforms Int8PtrConsumer into SilentInt8PtrConsumer
func (c Int8PtrConsumer) ToSilentInt8PtrConsumer() SilentInt8PtrConsumer {
	return func(v *int8) {
		_ = c(v)
		return
	}
}

// ToSilentConsumer transforms SilentInt8PtrConsumer into SilentConsumer
func (sc SilentInt8PtrConsumer) ToSilentConsumer() SilentConsumer {
	return func(v interface{}) {
		sc(v.(*int8))
	}
}

// AndThen returns a composed SilentInt8PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentInt8PtrConsumer) AndThen(after SilentInt8PtrConsumer) SilentInt8PtrConsumer {
	if after != nil {
		return func(v *int8) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustInt8PtrConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustInt8PtrConsumer func(v *int8)

// ToMustInt8PtrConsumer transforms Int8PtrConsumer into MustInt8PtrConsumer
func (c Int8PtrConsumer) ToMustInt8PtrConsumer() MustInt8PtrConsumer {
	return func(v *int8) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// ToMustConsumer transforms MustInt8PtrConsumer into MustConsumer
func (mc MustInt8PtrConsumer) ToMustConsumer() MustConsumer {
	return func(v interface{}) {
		mc(v.(*int8))
	}
}

// AndThen returns a composed MustInt8PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustInt8PtrConsumer) AndThen(after MustInt8PtrConsumer) MustInt8PtrConsumer {
	if after != nil {
		return func(v *int8) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToInt8PtrConsumer transforms MustInt8PtrConsumer into Int8PtrConsumer
func (mc MustInt8PtrConsumer) ToInt8PtrConsumer() Int8PtrConsumer {
	return func(v *int8) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentInt8PtrConsumer transforms MustInt8PtrConsumer into SilentInt8PtrConsumer
func (mc MustInt8PtrConsumer) ToSilentInt8PtrConsumer() SilentInt8PtrConsumer {
	c := mc.ToInt8PtrConsumer()
	return c.ToSilentInt8PtrConsumer()
}
