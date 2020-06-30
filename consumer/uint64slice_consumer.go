// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Uint64SliceConsumer represents an operation that accepts a single input argument or returns an error.
type Uint64SliceConsumer func(v []uint64) error

// ToConsumer transforms Uint64SliceConsumer into Consumer
func (c Uint64SliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]uint64))
	}
}

// AndThen returns a composed Uint64SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Uint64SliceConsumer) AndThen(after Uint64SliceConsumer) Uint64SliceConsumer {
	if after != nil {
		return func(v []uint64) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentUint64SliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentUint64SliceConsumer func(v []uint64)

// ToSilentUint64SliceConsumer transforms Uint64SliceConsumer into SilentUint64SliceConsumer
func (c Uint64SliceConsumer) ToSilentUint64SliceConsumer() SilentUint64SliceConsumer {
	return func(v []uint64) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentUint64SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentUint64SliceConsumer) AndThen(after SilentUint64SliceConsumer) SilentUint64SliceConsumer {
	if after != nil {
		return func(v []uint64) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustUint64SliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustUint64SliceConsumer func(v []uint64)

// ToMustUint64SliceConsumer transforms Uint64SliceConsumer into MustUint64SliceConsumer
func (c Uint64SliceConsumer) ToMustUint64SliceConsumer() MustUint64SliceConsumer {
	return func(v []uint64) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustUint64SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustUint64SliceConsumer) AndThen(after MustUint64SliceConsumer) MustUint64SliceConsumer {
	if after != nil {
		return func(v []uint64) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToUint64SliceConsumer transforms MustUint64SliceConsumer into Uint64SliceConsumer
func (mc MustUint64SliceConsumer) ToUint64SliceConsumer() Uint64SliceConsumer {
	return func(v []uint64) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentUint64SliceConsumer transforms MustUint64SliceConsumer into SilentUint64SliceConsumer
func (mc MustUint64SliceConsumer) ToSilentUint64SliceConsumer() SilentUint64SliceConsumer {
	c := mc.ToUint64SliceConsumer()
	return c.ToSilentUint64SliceConsumer()
}
