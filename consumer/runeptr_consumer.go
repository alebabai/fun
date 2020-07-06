// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// RunePtrConsumer represents an operation that accepts a single input argument or returns an error.
type RunePtrConsumer func(v *rune) error

// ToConsumer transforms RunePtrConsumer into Consumer
func (c RunePtrConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(*rune))
	}
}

// AndThen returns a composed RunePtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c RunePtrConsumer) AndThen(after RunePtrConsumer) RunePtrConsumer {
	if after != nil {
		return func(v *rune) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentRunePtrConsumer represents an operation that accepts a single input argument without returning an error.
type SilentRunePtrConsumer func(v *rune)

// ToSilentRunePtrConsumer transforms RunePtrConsumer into SilentRunePtrConsumer
func (c RunePtrConsumer) ToSilentRunePtrConsumer() SilentRunePtrConsumer {
	return func(v *rune) {
		_ = c(v)
		return
	}
}

// ToSilentConsumer transforms SilentRunePtrConsumer into SilentConsumer
func (sc SilentRunePtrConsumer) ToSilentConsumer() SilentConsumer {
	return func(v interface{}) {
		sc(v.(*rune))
	}
}

// AndThen returns a composed SilentRunePtrConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentRunePtrConsumer) AndThen(after SilentRunePtrConsumer) SilentRunePtrConsumer {
	if after != nil {
		return func(v *rune) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustRunePtrConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustRunePtrConsumer func(v *rune)

// ToMustRunePtrConsumer transforms RunePtrConsumer into MustRunePtrConsumer
func (c RunePtrConsumer) ToMustRunePtrConsumer() MustRunePtrConsumer {
	return func(v *rune) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// ToMustConsumer transforms MustRunePtrConsumer into MustConsumer
func (mc MustRunePtrConsumer) ToMustConsumer() MustConsumer {
	return func(v interface{}) {
		mc(v.(*rune))
	}
}

// AndThen returns a composed MustRunePtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustRunePtrConsumer) AndThen(after MustRunePtrConsumer) MustRunePtrConsumer {
	if after != nil {
		return func(v *rune) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToRunePtrConsumer transforms MustRunePtrConsumer into RunePtrConsumer
func (mc MustRunePtrConsumer) ToRunePtrConsumer() RunePtrConsumer {
	return func(v *rune) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentRunePtrConsumer transforms MustRunePtrConsumer into SilentRunePtrConsumer
func (mc MustRunePtrConsumer) ToSilentRunePtrConsumer() SilentRunePtrConsumer {
	c := mc.ToRunePtrConsumer()
	return c.ToSilentRunePtrConsumer()
}
