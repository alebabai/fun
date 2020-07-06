// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Int16PtrConsumer represents an operation that accepts a single input argument or returns an error.
type Int16PtrConsumer func(v *int16) error

// ToConsumer transforms Int16PtrConsumer into Consumer
func (c Int16PtrConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(*int16))
	}
}

// AndThen returns a composed Int16PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Int16PtrConsumer) AndThen(after Int16PtrConsumer) Int16PtrConsumer {
	if after != nil {
		return func(v *int16) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentInt16PtrConsumer represents an operation that accepts a single input argument without returning an error.
type SilentInt16PtrConsumer func(v *int16)

// ToSilentInt16PtrConsumer transforms Int16PtrConsumer into SilentInt16PtrConsumer
func (c Int16PtrConsumer) ToSilentInt16PtrConsumer() SilentInt16PtrConsumer {
	return func(v *int16) {
		_ = c(v)
		return
	}
}

// ToSilentConsumer transforms SilentInt16PtrConsumer into SilentConsumer
func (sc SilentInt16PtrConsumer) ToSilentConsumer() SilentConsumer {
	return func(v interface{}) {
		sc(v.(*int16))
	}
}

// AndThen returns a composed SilentInt16PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentInt16PtrConsumer) AndThen(after SilentInt16PtrConsumer) SilentInt16PtrConsumer {
	if after != nil {
		return func(v *int16) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustInt16PtrConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustInt16PtrConsumer func(v *int16)

// ToMustInt16PtrConsumer transforms Int16PtrConsumer into MustInt16PtrConsumer
func (c Int16PtrConsumer) ToMustInt16PtrConsumer() MustInt16PtrConsumer {
	return func(v *int16) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// ToMustConsumer transforms MustInt16PtrConsumer into MustConsumer
func (mc MustInt16PtrConsumer) ToMustConsumer() MustConsumer {
	return func(v interface{}) {
		mc(v.(*int16))
	}
}

// AndThen returns a composed MustInt16PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustInt16PtrConsumer) AndThen(after MustInt16PtrConsumer) MustInt16PtrConsumer {
	if after != nil {
		return func(v *int16) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToInt16PtrConsumer transforms MustInt16PtrConsumer into Int16PtrConsumer
func (mc MustInt16PtrConsumer) ToInt16PtrConsumer() Int16PtrConsumer {
	return func(v *int16) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentInt16PtrConsumer transforms MustInt16PtrConsumer into SilentInt16PtrConsumer
func (mc MustInt16PtrConsumer) ToSilentInt16PtrConsumer() SilentInt16PtrConsumer {
	c := mc.ToInt16PtrConsumer()
	return c.ToSilentInt16PtrConsumer()
}
