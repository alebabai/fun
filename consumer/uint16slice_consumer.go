// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Uint16SliceConsumer represents an operation that accepts a single input argument or returns an error.
type Uint16SliceConsumer func(v []uint16) error

// ToConsumer transforms Uint16SliceConsumer into Consumer
func (c Uint16SliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]uint16))
	}
}

// AndThen returns a composed Uint16SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Uint16SliceConsumer) AndThen(after Uint16SliceConsumer) Uint16SliceConsumer {
	if after != nil {
		return func(v []uint16) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentUint16SliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentUint16SliceConsumer func(v []uint16)

// ToSilentUint16SliceConsumer transforms Uint16SliceConsumer into SilentUint16SliceConsumer
func (c Uint16SliceConsumer) ToSilentUint16SliceConsumer() SilentUint16SliceConsumer {
	return func(v []uint16) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentUint16SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentUint16SliceConsumer) AndThen(after SilentUint16SliceConsumer) SilentUint16SliceConsumer {
	if after != nil {
		return func(v []uint16) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustUint16SliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustUint16SliceConsumer func(v []uint16)

// ToMustUint16SliceConsumer transforms Uint16SliceConsumer into MustUint16SliceConsumer
func (c Uint16SliceConsumer) ToMustUint16SliceConsumer() MustUint16SliceConsumer {
	return func(v []uint16) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustUint16SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustUint16SliceConsumer) AndThen(after MustUint16SliceConsumer) MustUint16SliceConsumer {
	if after != nil {
		return func(v []uint16) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToUint16SliceConsumer transforms MustUint16SliceConsumer into Uint16SliceConsumer
func (mc MustUint16SliceConsumer) ToUint16SliceConsumer() Uint16SliceConsumer {
	return func(v []uint16) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentUint16SliceConsumer transforms MustUint16SliceConsumer into SilentUint16SliceConsumer
func (mc MustUint16SliceConsumer) ToSilentUint16SliceConsumer() SilentUint16SliceConsumer {
	c := mc.ToUint16SliceConsumer()
	return c.ToSilentUint16SliceConsumer()
}
