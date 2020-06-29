// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Uint8SliceConsumer represents an operation that accepts a single input argument or returns an error.
type Uint8SliceConsumer func(v []uint8) error

// ToConsumer transforms Uint8SliceConsumer into Consumer
func (c Uint8SliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]uint8))
	}
}

// AndThen returns a composed Uint8SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Uint8SliceConsumer) AndThen(after Uint8SliceConsumer) Uint8SliceConsumer {
	if after != nil {
		return func(v []uint8) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentUint8SliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentUint8SliceConsumer func(v []uint8)

// ToSilentUint8SliceConsumer transforms Uint8SliceConsumer into SilentUint8SliceConsumer
func (c Uint8SliceConsumer) ToSilentUint8SliceConsumer() SilentUint8SliceConsumer {
	return func(v []uint8) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentUint8SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentUint8SliceConsumer) AndThen(after SilentUint8SliceConsumer) SilentUint8SliceConsumer {
	if after != nil {
		return func(v []uint8) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustUint8SliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustUint8SliceConsumer func(v []uint8)

// ToMustUint8SliceConsumer transforms Uint8SliceConsumer into MustUint8SliceConsumer
func (c Uint8SliceConsumer) ToMustUint8SliceConsumer() MustUint8SliceConsumer {
	return func(v []uint8) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustUint8SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustUint8SliceConsumer) AndThen(after MustUint8SliceConsumer) MustUint8SliceConsumer {
	if after != nil {
		return func(v []uint8) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToUint8SliceConsumer transforms MustUint8SliceConsumer into Uint8SliceConsumer
func (mc MustUint8SliceConsumer) ToUint8SliceConsumer() Uint8SliceConsumer {
	return func(v []uint8) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// SilentUint8SliceConsumer transforms MustUint8SliceConsumer into SilentUint8SliceConsumer
func (mc MustUint8SliceConsumer) ToSilentUint8SliceConsumer() SilentUint8SliceConsumer {
	c := mc.ToUint8SliceConsumer()
	return c.ToSilentUint8SliceConsumer()
}
