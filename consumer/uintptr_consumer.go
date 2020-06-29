// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// UintPtrConsumer represents an operation that accepts a single input argument or returns an error.
type UintPtrConsumer func(v *uint) error

// ToConsumer transforms UintPtrConsumer into Consumer
func (c UintPtrConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(*uint))
	}
}

// AndThen returns a composed UintPtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c UintPtrConsumer) AndThen(after UintPtrConsumer) UintPtrConsumer {
	if after != nil {
		return func(v *uint) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentUintPtrConsumer represents an operation that accepts a single input argument without returning an error.
type SilentUintPtrConsumer func(v *uint)

// ToSilentUintPtrConsumer transforms UintPtrConsumer into SilentUintPtrConsumer
func (c UintPtrConsumer) ToSilentUintPtrConsumer() SilentUintPtrConsumer {
	return func(v *uint) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentUintPtrConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentUintPtrConsumer) AndThen(after SilentUintPtrConsumer) SilentUintPtrConsumer {
	if after != nil {
		return func(v *uint) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustUintPtrConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustUintPtrConsumer func(v *uint)

// ToMustUintPtrConsumer transforms UintPtrConsumer into MustUintPtrConsumer
func (c UintPtrConsumer) ToMustUintPtrConsumer() MustUintPtrConsumer {
	return func(v *uint) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustUintPtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustUintPtrConsumer) AndThen(after MustUintPtrConsumer) MustUintPtrConsumer {
	if after != nil {
		return func(v *uint) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToUintPtrConsumer transforms MustUintPtrConsumer into UintPtrConsumer
func (mc MustUintPtrConsumer) ToUintPtrConsumer() UintPtrConsumer {
	return func(v *uint) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// SilentUintPtrConsumer transforms MustUintPtrConsumer into SilentUintPtrConsumer
func (mc MustUintPtrConsumer) ToSilentUintPtrConsumer() SilentUintPtrConsumer {
	c := mc.ToUintPtrConsumer()
	return c.ToSilentUintPtrConsumer()
}
