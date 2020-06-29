// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Uint8PtrConsumer represents an operation that accepts a single input argument or returns an error.
type Uint8PtrConsumer func(v *uint8) error

// ToConsumer transforms Uint8PtrConsumer into Consumer
func (c Uint8PtrConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(*uint8))
	}
}

// AndThen returns a composed Uint8PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Uint8PtrConsumer) AndThen(after Uint8PtrConsumer) Uint8PtrConsumer {
	if after != nil {
		return func(v *uint8) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentUint8PtrConsumer represents an operation that accepts a single input argument without returning an error.
type SilentUint8PtrConsumer func(v *uint8)

// ToSilentUint8PtrConsumer transforms Uint8PtrConsumer into SilentUint8PtrConsumer
func (c Uint8PtrConsumer) ToSilentUint8PtrConsumer() SilentUint8PtrConsumer {
	return func(v *uint8) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentUint8PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentUint8PtrConsumer) AndThen(after SilentUint8PtrConsumer) SilentUint8PtrConsumer {
	if after != nil {
		return func(v *uint8) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustUint8PtrConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustUint8PtrConsumer func(v *uint8)

// ToMustUint8PtrConsumer transforms Uint8PtrConsumer into MustUint8PtrConsumer
func (c Uint8PtrConsumer) ToMustUint8PtrConsumer() MustUint8PtrConsumer {
	return func(v *uint8) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustUint8PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustUint8PtrConsumer) AndThen(after MustUint8PtrConsumer) MustUint8PtrConsumer {
	if after != nil {
		return func(v *uint8) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToUint8PtrConsumer transforms MustUint8PtrConsumer into Uint8PtrConsumer
func (mc MustUint8PtrConsumer) ToUint8PtrConsumer() Uint8PtrConsumer {
	return func(v *uint8) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// SilentUint8PtrConsumer transforms MustUint8PtrConsumer into SilentUint8PtrConsumer
func (mc MustUint8PtrConsumer) ToSilentUint8PtrConsumer() SilentUint8PtrConsumer {
	c := mc.ToUint8PtrConsumer()
	return c.ToSilentUint8PtrConsumer()
}
