// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// IntPtrSliceConsumer represents an operation that accepts a single input argument or returns an error.
type IntPtrSliceConsumer func(v []*int) error

// ToConsumer transforms IntPtrSliceConsumer into Consumer
func (c IntPtrSliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]*int))
	}
}

// AndThen returns a composed IntPtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c IntPtrSliceConsumer) AndThen(after IntPtrSliceConsumer) IntPtrSliceConsumer {
	if after != nil {
		return func(v []*int) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentIntPtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentIntPtrSliceConsumer func(v []*int)

// ToSilentIntPtrSliceConsumer transforms IntPtrSliceConsumer into SilentIntPtrSliceConsumer
func (c IntPtrSliceConsumer) ToSilentIntPtrSliceConsumer() SilentIntPtrSliceConsumer {
	return func(v []*int) {
		_ = c(v)
		return
	}
}

// ToSilentConsumer transforms SilentIntPtrSliceConsumer into SilentConsumer
func (sc SilentIntPtrSliceConsumer) ToSilentConsumer() SilentConsumer {
	return func(v interface{}) {
		sc(v.([]*int))
	}
}

// AndThen returns a composed SilentIntPtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentIntPtrSliceConsumer) AndThen(after SilentIntPtrSliceConsumer) SilentIntPtrSliceConsumer {
	if after != nil {
		return func(v []*int) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustIntPtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustIntPtrSliceConsumer func(v []*int)

// ToMustIntPtrSliceConsumer transforms IntPtrSliceConsumer into MustIntPtrSliceConsumer
func (c IntPtrSliceConsumer) ToMustIntPtrSliceConsumer() MustIntPtrSliceConsumer {
	return func(v []*int) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// ToMustConsumer transforms MustIntPtrSliceConsumer into MustConsumer
func (mc MustIntPtrSliceConsumer) ToMustConsumer() MustConsumer {
	return func(v interface{}) {
		mc(v.([]*int))
	}
}

// AndThen returns a composed MustIntPtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustIntPtrSliceConsumer) AndThen(after MustIntPtrSliceConsumer) MustIntPtrSliceConsumer {
	if after != nil {
		return func(v []*int) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToIntPtrSliceConsumer transforms MustIntPtrSliceConsumer into IntPtrSliceConsumer
func (mc MustIntPtrSliceConsumer) ToIntPtrSliceConsumer() IntPtrSliceConsumer {
	return func(v []*int) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentIntPtrSliceConsumer transforms MustIntPtrSliceConsumer into SilentIntPtrSliceConsumer
func (mc MustIntPtrSliceConsumer) ToSilentIntPtrSliceConsumer() SilentIntPtrSliceConsumer {
	c := mc.ToIntPtrSliceConsumer()
	return c.ToSilentIntPtrSliceConsumer()
}
