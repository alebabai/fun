// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// ByteConsumer represents an operation that accepts a single input argument or returns an error.
type ByteConsumer func(v byte) error

// ToConsumer transforms ByteConsumer into Consumer
func (c ByteConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(byte))
	}
}

// AndThen returns a composed ByteConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c ByteConsumer) AndThen(after ByteConsumer) ByteConsumer {
	if after != nil {
		return func(v byte) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentByteConsumer represents an operation that accepts a single input argument without returning an error.
type SilentByteConsumer func(v byte)

// ToSilentByteConsumer transforms ByteConsumer into SilentByteConsumer
func (c ByteConsumer) ToSilentByteConsumer() SilentByteConsumer {
	return func(v byte) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentByteConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentByteConsumer) AndThen(after SilentByteConsumer) SilentByteConsumer {
	if after != nil {
		return func(v byte) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustByteConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustByteConsumer func(v byte)

// ToMustByteConsumer transforms ByteConsumer into MustByteConsumer
func (c ByteConsumer) ToMustByteConsumer() MustByteConsumer {
	return func(v byte) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustByteConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustByteConsumer) AndThen(after MustByteConsumer) MustByteConsumer {
	if after != nil {
		return func(v byte) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToByteConsumer transforms MustByteConsumer into ByteConsumer
func (mc MustByteConsumer) ToByteConsumer() ByteConsumer {
	return func(v byte) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentByteConsumer transforms MustByteConsumer into SilentByteConsumer
func (mc MustByteConsumer) ToSilentByteConsumer() SilentByteConsumer {
	c := mc.ToByteConsumer()
	return c.ToSilentByteConsumer()
}
