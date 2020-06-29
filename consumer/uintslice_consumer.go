// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// UintSliceConsumer represents an operation that accepts a single input argument or returns an error.
type UintSliceConsumer func(v []uint) error

// ToConsumer transforms UintSliceConsumer into Consumer
func (c UintSliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]uint))
	}
}

// AndThen returns a composed UintSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c UintSliceConsumer) AndThen(after UintSliceConsumer) UintSliceConsumer {
	if after != nil {
		return func(v []uint) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentUintSliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentUintSliceConsumer func(v []uint)

// ToSilentUintSliceConsumer transforms UintSliceConsumer into SilentUintSliceConsumer
func (c UintSliceConsumer) ToSilentUintSliceConsumer() SilentUintSliceConsumer {
	return func(v []uint) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentUintSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentUintSliceConsumer) AndThen(after SilentUintSliceConsumer) SilentUintSliceConsumer {
	if after != nil {
		return func(v []uint) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustUintSliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustUintSliceConsumer func(v []uint)

// ToMustUintSliceConsumer transforms UintSliceConsumer into MustUintSliceConsumer
func (c UintSliceConsumer) ToMustUintSliceConsumer() MustUintSliceConsumer {
	return func(v []uint) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustUintSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustUintSliceConsumer) AndThen(after MustUintSliceConsumer) MustUintSliceConsumer {
	if after != nil {
		return func(v []uint) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToUintSliceConsumer transforms MustUintSliceConsumer into UintSliceConsumer
func (mc MustUintSliceConsumer) ToUintSliceConsumer() UintSliceConsumer {
	return func(v []uint) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// SilentUintSliceConsumer transforms MustUintSliceConsumer into SilentUintSliceConsumer
func (mc MustUintSliceConsumer) ToSilentUintSliceConsumer() SilentUintSliceConsumer {
	c := mc.ToUintSliceConsumer()
	return c.ToSilentUintSliceConsumer()
}
