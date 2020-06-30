// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// IntConsumer represents an operation that accepts a single input argument or returns an error.
type IntConsumer func(v int) error

// ToConsumer transforms IntConsumer into Consumer
func (c IntConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(int))
	}
}

// AndThen returns a composed IntConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c IntConsumer) AndThen(after IntConsumer) IntConsumer {
	if after != nil {
		return func(v int) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentIntConsumer represents an operation that accepts a single input argument without returning an error.
type SilentIntConsumer func(v int)

// ToSilentIntConsumer transforms IntConsumer into SilentIntConsumer
func (c IntConsumer) ToSilentIntConsumer() SilentIntConsumer {
	return func(v int) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentIntConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentIntConsumer) AndThen(after SilentIntConsumer) SilentIntConsumer {
	if after != nil {
		return func(v int) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustIntConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustIntConsumer func(v int)

// ToMustIntConsumer transforms IntConsumer into MustIntConsumer
func (c IntConsumer) ToMustIntConsumer() MustIntConsumer {
	return func(v int) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustIntConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustIntConsumer) AndThen(after MustIntConsumer) MustIntConsumer {
	if after != nil {
		return func(v int) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToIntConsumer transforms MustIntConsumer into IntConsumer
func (mc MustIntConsumer) ToIntConsumer() IntConsumer {
	return func(v int) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentIntConsumer transforms MustIntConsumer into SilentIntConsumer
func (mc MustIntConsumer) ToSilentIntConsumer() SilentIntConsumer {
	c := mc.ToIntConsumer()
	return c.ToSilentIntConsumer()
}
