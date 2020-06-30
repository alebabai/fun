// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// RuneSliceConsumer represents an operation that accepts a single input argument or returns an error.
type RuneSliceConsumer func(v []rune) error

// ToConsumer transforms RuneSliceConsumer into Consumer
func (c RuneSliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]rune))
	}
}

// AndThen returns a composed RuneSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c RuneSliceConsumer) AndThen(after RuneSliceConsumer) RuneSliceConsumer {
	if after != nil {
		return func(v []rune) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentRuneSliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentRuneSliceConsumer func(v []rune)

// ToSilentRuneSliceConsumer transforms RuneSliceConsumer into SilentRuneSliceConsumer
func (c RuneSliceConsumer) ToSilentRuneSliceConsumer() SilentRuneSliceConsumer {
	return func(v []rune) {
		_ = c(v)
		return
	}
}

// ToSilentConsumer transforms SilentRuneSliceConsumer into SilentConsumer
func (sc SilentRuneSliceConsumer) ToSilentConsumer() SilentConsumer {
	return func(v interface{}) {
		sc(v.([]rune))
	}
}

// AndThen returns a composed SilentRuneSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentRuneSliceConsumer) AndThen(after SilentRuneSliceConsumer) SilentRuneSliceConsumer {
	if after != nil {
		return func(v []rune) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustRuneSliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustRuneSliceConsumer func(v []rune)

// ToMustRuneSliceConsumer transforms RuneSliceConsumer into MustRuneSliceConsumer
func (c RuneSliceConsumer) ToMustRuneSliceConsumer() MustRuneSliceConsumer {
	return func(v []rune) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// ToMustConsumer transforms MustRuneSliceConsumer into MustConsumer
func (mc MustRuneSliceConsumer) ToMustConsumer() MustConsumer {
	return func(v interface{}) {
		mc(v.([]rune))
	}
}

// AndThen returns a composed MustRuneSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustRuneSliceConsumer) AndThen(after MustRuneSliceConsumer) MustRuneSliceConsumer {
	if after != nil {
		return func(v []rune) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToRuneSliceConsumer transforms MustRuneSliceConsumer into RuneSliceConsumer
func (mc MustRuneSliceConsumer) ToRuneSliceConsumer() RuneSliceConsumer {
	return func(v []rune) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentRuneSliceConsumer transforms MustRuneSliceConsumer into SilentRuneSliceConsumer
func (mc MustRuneSliceConsumer) ToSilentRuneSliceConsumer() SilentRuneSliceConsumer {
	c := mc.ToRuneSliceConsumer()
	return c.ToSilentRuneSliceConsumer()
}
