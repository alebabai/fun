// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Uint64PtrConsumer represents an operation that accepts a single input argument or returns an error.
type Uint64PtrConsumer func(v *uint64) error

// ToConsumer transforms Uint64PtrConsumer into Consumer
func (c Uint64PtrConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(*uint64))
	}
}

// AndThen returns a composed Uint64PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Uint64PtrConsumer) AndThen(after Uint64PtrConsumer) Uint64PtrConsumer {
	if after != nil {
		return func(v *uint64) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentUint64PtrConsumer represents an operation that accepts a single input argument without returning an error.
type SilentUint64PtrConsumer func(v *uint64)

// ToSilentUint64PtrConsumer transforms Uint64PtrConsumer into SilentUint64PtrConsumer
func (c Uint64PtrConsumer) ToSilentUint64PtrConsumer() SilentUint64PtrConsumer {
	return func(v *uint64) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentUint64PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentUint64PtrConsumer) AndThen(after SilentUint64PtrConsumer) SilentUint64PtrConsumer {
	if after != nil {
		return func(v *uint64) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustUint64PtrConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustUint64PtrConsumer func(v *uint64)

// ToMustUint64PtrConsumer transforms Uint64PtrConsumer into MustUint64PtrConsumer
func (c Uint64PtrConsumer) ToMustUint64PtrConsumer() MustUint64PtrConsumer {
	return func(v *uint64) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustUint64PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustUint64PtrConsumer) AndThen(after MustUint64PtrConsumer) MustUint64PtrConsumer {
	if after != nil {
		return func(v *uint64) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToUint64PtrConsumer transforms MustUint64PtrConsumer into Uint64PtrConsumer
func (mc MustUint64PtrConsumer) ToUint64PtrConsumer() Uint64PtrConsumer {
	return func(v *uint64) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// SilentUint64PtrConsumer transforms MustUint64PtrConsumer into SilentUint64PtrConsumer
func (mc MustUint64PtrConsumer) ToSilentUint64PtrConsumer() SilentUint64PtrConsumer {
	c := mc.ToUint64PtrConsumer()
	return c.ToSilentUint64PtrConsumer()
}
