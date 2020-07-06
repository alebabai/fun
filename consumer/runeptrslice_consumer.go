// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// RunePtrSliceConsumer represents an operation that accepts a single input argument or returns an error.
type RunePtrSliceConsumer func(v []*rune) error

// ToConsumer transforms RunePtrSliceConsumer into Consumer
func (c RunePtrSliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]*rune))
	}
}

// AndThen returns a composed RunePtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c RunePtrSliceConsumer) AndThen(after RunePtrSliceConsumer) RunePtrSliceConsumer {
	if after != nil {
		return func(v []*rune) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentRunePtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentRunePtrSliceConsumer func(v []*rune)

// ToSilentRunePtrSliceConsumer transforms RunePtrSliceConsumer into SilentRunePtrSliceConsumer
func (c RunePtrSliceConsumer) ToSilentRunePtrSliceConsumer() SilentRunePtrSliceConsumer {
	return func(v []*rune) {
		_ = c(v)
		return
	}
}

// ToSilentConsumer transforms SilentRunePtrSliceConsumer into SilentConsumer
func (sc SilentRunePtrSliceConsumer) ToSilentConsumer() SilentConsumer {
	return func(v interface{}) {
		sc(v.([]*rune))
	}
}

// AndThen returns a composed SilentRunePtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentRunePtrSliceConsumer) AndThen(after SilentRunePtrSliceConsumer) SilentRunePtrSliceConsumer {
	if after != nil {
		return func(v []*rune) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustRunePtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustRunePtrSliceConsumer func(v []*rune)

// ToMustRunePtrSliceConsumer transforms RunePtrSliceConsumer into MustRunePtrSliceConsumer
func (c RunePtrSliceConsumer) ToMustRunePtrSliceConsumer() MustRunePtrSliceConsumer {
	return func(v []*rune) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// ToMustConsumer transforms MustRunePtrSliceConsumer into MustConsumer
func (mc MustRunePtrSliceConsumer) ToMustConsumer() MustConsumer {
	return func(v interface{}) {
		mc(v.([]*rune))
	}
}

// AndThen returns a composed MustRunePtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustRunePtrSliceConsumer) AndThen(after MustRunePtrSliceConsumer) MustRunePtrSliceConsumer {
	if after != nil {
		return func(v []*rune) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToRunePtrSliceConsumer transforms MustRunePtrSliceConsumer into RunePtrSliceConsumer
func (mc MustRunePtrSliceConsumer) ToRunePtrSliceConsumer() RunePtrSliceConsumer {
	return func(v []*rune) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentRunePtrSliceConsumer transforms MustRunePtrSliceConsumer into SilentRunePtrSliceConsumer
func (mc MustRunePtrSliceConsumer) ToSilentRunePtrSliceConsumer() SilentRunePtrSliceConsumer {
	c := mc.ToRunePtrSliceConsumer()
	return c.ToSilentRunePtrSliceConsumer()
}
