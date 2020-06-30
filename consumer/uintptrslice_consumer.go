// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// UintPtrSliceConsumer represents an operation that accepts a single input argument or returns an error.
type UintPtrSliceConsumer func(v []*uint) error

// ToConsumer transforms UintPtrSliceConsumer into Consumer
func (c UintPtrSliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]*uint))
	}
}

// AndThen returns a composed UintPtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c UintPtrSliceConsumer) AndThen(after UintPtrSliceConsumer) UintPtrSliceConsumer {
	if after != nil {
		return func(v []*uint) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentUintPtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentUintPtrSliceConsumer func(v []*uint)

// ToSilentUintPtrSliceConsumer transforms UintPtrSliceConsumer into SilentUintPtrSliceConsumer
func (c UintPtrSliceConsumer) ToSilentUintPtrSliceConsumer() SilentUintPtrSliceConsumer {
	return func(v []*uint) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentUintPtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentUintPtrSliceConsumer) AndThen(after SilentUintPtrSliceConsumer) SilentUintPtrSliceConsumer {
	if after != nil {
		return func(v []*uint) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustUintPtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustUintPtrSliceConsumer func(v []*uint)

// ToMustUintPtrSliceConsumer transforms UintPtrSliceConsumer into MustUintPtrSliceConsumer
func (c UintPtrSliceConsumer) ToMustUintPtrSliceConsumer() MustUintPtrSliceConsumer {
	return func(v []*uint) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustUintPtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustUintPtrSliceConsumer) AndThen(after MustUintPtrSliceConsumer) MustUintPtrSliceConsumer {
	if after != nil {
		return func(v []*uint) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToUintPtrSliceConsumer transforms MustUintPtrSliceConsumer into UintPtrSliceConsumer
func (mc MustUintPtrSliceConsumer) ToUintPtrSliceConsumer() UintPtrSliceConsumer {
	return func(v []*uint) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentUintPtrSliceConsumer transforms MustUintPtrSliceConsumer into SilentUintPtrSliceConsumer
func (mc MustUintPtrSliceConsumer) ToSilentUintPtrSliceConsumer() SilentUintPtrSliceConsumer {
	c := mc.ToUintPtrSliceConsumer()
	return c.ToSilentUintPtrSliceConsumer()
}
