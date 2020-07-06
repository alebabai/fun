// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// ByteSliceConsumer represents an operation that accepts a single input argument or returns an error.
type ByteSliceConsumer func(v []byte) error

// ToConsumer transforms ByteSliceConsumer into Consumer
func (c ByteSliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]byte))
	}
}

// AndThen returns a composed ByteSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c ByteSliceConsumer) AndThen(after ByteSliceConsumer) ByteSliceConsumer {
	if after != nil {
		return func(v []byte) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentByteSliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentByteSliceConsumer func(v []byte)

// ToSilentByteSliceConsumer transforms ByteSliceConsumer into SilentByteSliceConsumer
func (c ByteSliceConsumer) ToSilentByteSliceConsumer() SilentByteSliceConsumer {
	return func(v []byte) {
		_ = c(v)
		return
	}
}

// ToSilentConsumer transforms SilentByteSliceConsumer into SilentConsumer
func (sc SilentByteSliceConsumer) ToSilentConsumer() SilentConsumer {
	return func(v interface{}) {
		sc(v.([]byte))
	}
}

// AndThen returns a composed SilentByteSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentByteSliceConsumer) AndThen(after SilentByteSliceConsumer) SilentByteSliceConsumer {
	if after != nil {
		return func(v []byte) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustByteSliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustByteSliceConsumer func(v []byte)

// ToMustByteSliceConsumer transforms ByteSliceConsumer into MustByteSliceConsumer
func (c ByteSliceConsumer) ToMustByteSliceConsumer() MustByteSliceConsumer {
	return func(v []byte) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// ToMustConsumer transforms MustByteSliceConsumer into MustConsumer
func (mc MustByteSliceConsumer) ToMustConsumer() MustConsumer {
	return func(v interface{}) {
		mc(v.([]byte))
	}
}

// AndThen returns a composed MustByteSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustByteSliceConsumer) AndThen(after MustByteSliceConsumer) MustByteSliceConsumer {
	if after != nil {
		return func(v []byte) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToByteSliceConsumer transforms MustByteSliceConsumer into ByteSliceConsumer
func (mc MustByteSliceConsumer) ToByteSliceConsumer() ByteSliceConsumer {
	return func(v []byte) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentByteSliceConsumer transforms MustByteSliceConsumer into SilentByteSliceConsumer
func (mc MustByteSliceConsumer) ToSilentByteSliceConsumer() SilentByteSliceConsumer {
	c := mc.ToByteSliceConsumer()
	return c.ToSilentByteSliceConsumer()
}
