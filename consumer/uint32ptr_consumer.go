// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Uint32PtrConsumer represents an operation that accepts a single input argument or returns an error.
type Uint32PtrConsumer func(v *uint32) error

// ToConsumer transforms Uint32PtrConsumer into Consumer
func (c Uint32PtrConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(*uint32))
	}
}

// AndThen returns a composed Uint32PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Uint32PtrConsumer) AndThen(after Uint32PtrConsumer) Uint32PtrConsumer {
	if after != nil {
		return func(v *uint32) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentUint32PtrConsumer represents an operation that accepts a single input argument without returning an error.
type SilentUint32PtrConsumer func(v *uint32)

// ToSilentUint32PtrConsumer transforms Uint32PtrConsumer into SilentUint32PtrConsumer
func (c Uint32PtrConsumer) ToSilentUint32PtrConsumer() SilentUint32PtrConsumer {
	return func(v *uint32) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentUint32PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentUint32PtrConsumer) AndThen(after SilentUint32PtrConsumer) SilentUint32PtrConsumer {
	if after != nil {
		return func(v *uint32) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustUint32PtrConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustUint32PtrConsumer func(v *uint32)

// ToMustUint32PtrConsumer transforms Uint32PtrConsumer into MustUint32PtrConsumer
func (c Uint32PtrConsumer) ToMustUint32PtrConsumer() MustUint32PtrConsumer {
	return func(v *uint32) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustUint32PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustUint32PtrConsumer) AndThen(after MustUint32PtrConsumer) MustUint32PtrConsumer {
	if after != nil {
		return func(v *uint32) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToUint32PtrConsumer transforms MustUint32PtrConsumer into Uint32PtrConsumer
func (mc MustUint32PtrConsumer) ToUint32PtrConsumer() Uint32PtrConsumer {
	return func(v *uint32) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentUint32PtrConsumer transforms MustUint32PtrConsumer into SilentUint32PtrConsumer
func (mc MustUint32PtrConsumer) ToSilentUint32PtrConsumer() SilentUint32PtrConsumer {
	c := mc.ToUint32PtrConsumer()
	return c.ToSilentUint32PtrConsumer()
}
