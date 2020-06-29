// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Uint32SliceConsumer represents an operation that accepts a single input argument or returns an error.
type Uint32SliceConsumer func(v []uint32) error

// ToConsumer transforms Uint32SliceConsumer into Consumer
func (c Uint32SliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]uint32))
	}
}

// AndThen returns a composed Uint32SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Uint32SliceConsumer) AndThen(after Uint32SliceConsumer) Uint32SliceConsumer {
	if after != nil {
		return func(v []uint32) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentUint32SliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentUint32SliceConsumer func(v []uint32)

// ToSilentUint32SliceConsumer transforms Uint32SliceConsumer into SilentUint32SliceConsumer
func (c Uint32SliceConsumer) ToSilentUint32SliceConsumer() SilentUint32SliceConsumer {
	return func(v []uint32) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentUint32SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentUint32SliceConsumer) AndThen(after SilentUint32SliceConsumer) SilentUint32SliceConsumer {
	if after != nil {
		return func(v []uint32) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustUint32SliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustUint32SliceConsumer func(v []uint32)

// ToMustUint32SliceConsumer transforms Uint32SliceConsumer into MustUint32SliceConsumer
func (c Uint32SliceConsumer) ToMustUint32SliceConsumer() MustUint32SliceConsumer {
	return func(v []uint32) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustUint32SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustUint32SliceConsumer) AndThen(after MustUint32SliceConsumer) MustUint32SliceConsumer {
	if after != nil {
		return func(v []uint32) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToUint32SliceConsumer transforms MustUint32SliceConsumer into Uint32SliceConsumer
func (mc MustUint32SliceConsumer) ToUint32SliceConsumer() Uint32SliceConsumer {
	return func(v []uint32) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// SilentUint32SliceConsumer transforms MustUint32SliceConsumer into SilentUint32SliceConsumer
func (mc MustUint32SliceConsumer) ToSilentUint32SliceConsumer() SilentUint32SliceConsumer {
	c := mc.ToUint32SliceConsumer()
	return c.ToSilentUint32SliceConsumer()
}
