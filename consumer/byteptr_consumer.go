// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// BytePtrConsumer represents an operation that accepts a single input argument or returns an error.
type BytePtrConsumer func(v *byte) error

// ToConsumer transforms BytePtrConsumer into Consumer
func (c BytePtrConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(*byte))
	}
}

// AndThen returns a composed BytePtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c BytePtrConsumer) AndThen(after BytePtrConsumer) BytePtrConsumer {
	if after != nil {
		return func(v *byte) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentBytePtrConsumer represents an operation that accepts a single input argument without returning an error.
type SilentBytePtrConsumer func(v *byte)

// ToSilentBytePtrConsumer transforms BytePtrConsumer into SilentBytePtrConsumer
func (c BytePtrConsumer) ToSilentBytePtrConsumer() SilentBytePtrConsumer {
	return func(v *byte) {
		_ = c(v)
		return
	}
}

// ToSilentConsumer transforms SilentBytePtrConsumer into SilentConsumer
func (sc SilentBytePtrConsumer) ToSilentConsumer() SilentConsumer {
	return func(v interface{}) {
		sc(v.(*byte))
	}
}

// AndThen returns a composed SilentBytePtrConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentBytePtrConsumer) AndThen(after SilentBytePtrConsumer) SilentBytePtrConsumer {
	if after != nil {
		return func(v *byte) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustBytePtrConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustBytePtrConsumer func(v *byte)

// ToMustBytePtrConsumer transforms BytePtrConsumer into MustBytePtrConsumer
func (c BytePtrConsumer) ToMustBytePtrConsumer() MustBytePtrConsumer {
	return func(v *byte) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// ToMustConsumer transforms MustBytePtrConsumer into MustConsumer
func (mc MustBytePtrConsumer) ToMustConsumer() MustConsumer {
	return func(v interface{}) {
		mc(v.(*byte))
	}
}

// AndThen returns a composed MustBytePtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustBytePtrConsumer) AndThen(after MustBytePtrConsumer) MustBytePtrConsumer {
	if after != nil {
		return func(v *byte) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToBytePtrConsumer transforms MustBytePtrConsumer into BytePtrConsumer
func (mc MustBytePtrConsumer) ToBytePtrConsumer() BytePtrConsumer {
	return func(v *byte) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentBytePtrConsumer transforms MustBytePtrConsumer into SilentBytePtrConsumer
func (mc MustBytePtrConsumer) ToSilentBytePtrConsumer() SilentBytePtrConsumer {
	c := mc.ToBytePtrConsumer()
	return c.ToSilentBytePtrConsumer()
}
