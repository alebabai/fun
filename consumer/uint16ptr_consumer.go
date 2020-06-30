// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Uint16PtrConsumer represents an operation that accepts a single input argument or returns an error.
type Uint16PtrConsumer func(v *uint16) error

// ToConsumer transforms Uint16PtrConsumer into Consumer
func (c Uint16PtrConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(*uint16))
	}
}

// AndThen returns a composed Uint16PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Uint16PtrConsumer) AndThen(after Uint16PtrConsumer) Uint16PtrConsumer {
	if after != nil {
		return func(v *uint16) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentUint16PtrConsumer represents an operation that accepts a single input argument without returning an error.
type SilentUint16PtrConsumer func(v *uint16)

// ToSilentUint16PtrConsumer transforms Uint16PtrConsumer into SilentUint16PtrConsumer
func (c Uint16PtrConsumer) ToSilentUint16PtrConsumer() SilentUint16PtrConsumer {
	return func(v *uint16) {
		_ = c(v)
		return
	}
}

// ToSilentConsumer transforms SilentUint16PtrConsumer into SilentConsumer
func (sc SilentUint16PtrConsumer) ToSilentConsumer() SilentConsumer {
	return func(v interface{}) {
		sc(v.(*uint16))
	}
}

// AndThen returns a composed SilentUint16PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentUint16PtrConsumer) AndThen(after SilentUint16PtrConsumer) SilentUint16PtrConsumer {
	if after != nil {
		return func(v *uint16) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustUint16PtrConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustUint16PtrConsumer func(v *uint16)

// ToMustUint16PtrConsumer transforms Uint16PtrConsumer into MustUint16PtrConsumer
func (c Uint16PtrConsumer) ToMustUint16PtrConsumer() MustUint16PtrConsumer {
	return func(v *uint16) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// ToMustConsumer transforms MustUint16PtrConsumer into MustConsumer
func (mc MustUint16PtrConsumer) ToMustConsumer() MustConsumer {
	return func(v interface{}) {
		mc(v.(*uint16))
	}
}

// AndThen returns a composed MustUint16PtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustUint16PtrConsumer) AndThen(after MustUint16PtrConsumer) MustUint16PtrConsumer {
	if after != nil {
		return func(v *uint16) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToUint16PtrConsumer transforms MustUint16PtrConsumer into Uint16PtrConsumer
func (mc MustUint16PtrConsumer) ToUint16PtrConsumer() Uint16PtrConsumer {
	return func(v *uint16) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentUint16PtrConsumer transforms MustUint16PtrConsumer into SilentUint16PtrConsumer
func (mc MustUint16PtrConsumer) ToSilentUint16PtrConsumer() SilentUint16PtrConsumer {
	c := mc.ToUint16PtrConsumer()
	return c.ToSilentUint16PtrConsumer()
}
