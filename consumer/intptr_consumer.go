// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// IntPtrConsumer represents an operation that accepts a single input argument or returns an error.
type IntPtrConsumer func(v *int) error

// ToConsumer transforms IntPtrConsumer into Consumer
func (c IntPtrConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(*int))
	}
}

// AndThen returns a composed IntPtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c IntPtrConsumer) AndThen(after IntPtrConsumer) IntPtrConsumer {
	if after != nil {
		return func(v *int) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentIntPtrConsumer represents an operation that accepts a single input argument without returning an error.
type SilentIntPtrConsumer func(v *int)

// ToSilentIntPtrConsumer transforms IntPtrConsumer into SilentIntPtrConsumer
func (c IntPtrConsumer) ToSilentIntPtrConsumer() SilentIntPtrConsumer {
	return func(v *int) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentIntPtrConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentIntPtrConsumer) AndThen(after SilentIntPtrConsumer) SilentIntPtrConsumer {
	if after != nil {
		return func(v *int) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustIntPtrConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustIntPtrConsumer func(v *int)

// ToMustIntPtrConsumer transforms IntPtrConsumer into MustIntPtrConsumer
func (c IntPtrConsumer) ToMustIntPtrConsumer() MustIntPtrConsumer {
	return func(v *int) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustIntPtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustIntPtrConsumer) AndThen(after MustIntPtrConsumer) MustIntPtrConsumer {
	if after != nil {
		return func(v *int) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToIntPtrConsumer transforms MustIntPtrConsumer into IntPtrConsumer
func (mc MustIntPtrConsumer) ToIntPtrConsumer() IntPtrConsumer {
	return func(v *int) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentIntPtrConsumer transforms MustIntPtrConsumer into SilentIntPtrConsumer
func (mc MustIntPtrConsumer) ToSilentIntPtrConsumer() SilentIntPtrConsumer {
	c := mc.ToIntPtrConsumer()
	return c.ToSilentIntPtrConsumer()
}
