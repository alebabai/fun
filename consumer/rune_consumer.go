// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// RuneConsumer represents an operation that accepts a single input argument or returns an error.
type RuneConsumer func(v rune) error

// ToConsumer transforms RuneConsumer into Consumer
func (c RuneConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(rune))
	}
}

// AndThen returns a composed RuneConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c RuneConsumer) AndThen(after RuneConsumer) RuneConsumer {
	if after != nil {
		return func(v rune) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentRuneConsumer represents an operation that accepts a single input argument without returning an error.
type SilentRuneConsumer func(v rune)

// ToSilentRuneConsumer transforms RuneConsumer into SilentRuneConsumer
func (c RuneConsumer) ToSilentRuneConsumer() SilentRuneConsumer {
	return func(v rune) {
		_ = c(v)
		return
	}
}

// ToSilentConsumer transforms SilentRuneConsumer into SilentConsumer
func (sc SilentRuneConsumer) ToSilentConsumer() SilentConsumer {
	return func(v interface{}) {
		sc(v.(rune))
	}
}

// AndThen returns a composed SilentRuneConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentRuneConsumer) AndThen(after SilentRuneConsumer) SilentRuneConsumer {
	if after != nil {
		return func(v rune) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustRuneConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustRuneConsumer func(v rune)

// ToMustRuneConsumer transforms RuneConsumer into MustRuneConsumer
func (c RuneConsumer) ToMustRuneConsumer() MustRuneConsumer {
	return func(v rune) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// ToMustConsumer transforms MustRuneConsumer into MustConsumer
func (mc MustRuneConsumer) ToMustConsumer() MustConsumer {
	return func(v interface{}) {
		mc(v.(rune))
	}
}

// AndThen returns a composed MustRuneConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustRuneConsumer) AndThen(after MustRuneConsumer) MustRuneConsumer {
	if after != nil {
		return func(v rune) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToRuneConsumer transforms MustRuneConsumer into RuneConsumer
func (mc MustRuneConsumer) ToRuneConsumer() RuneConsumer {
	return func(v rune) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentRuneConsumer transforms MustRuneConsumer into SilentRuneConsumer
func (mc MustRuneConsumer) ToSilentRuneConsumer() SilentRuneConsumer {
	c := mc.ToRuneConsumer()
	return c.ToSilentRuneConsumer()
}
