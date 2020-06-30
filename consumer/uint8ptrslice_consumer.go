// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Uint8PtrSliceConsumer represents an operation that accepts a single input argument or returns an error.
type Uint8PtrSliceConsumer func(v []*uint8) error

// ToConsumer transforms Uint8PtrSliceConsumer into Consumer
func (c Uint8PtrSliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]*uint8))
	}
}

// AndThen returns a composed Uint8PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Uint8PtrSliceConsumer) AndThen(after Uint8PtrSliceConsumer) Uint8PtrSliceConsumer {
	if after != nil {
		return func(v []*uint8) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentUint8PtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentUint8PtrSliceConsumer func(v []*uint8)

// ToSilentUint8PtrSliceConsumer transforms Uint8PtrSliceConsumer into SilentUint8PtrSliceConsumer
func (c Uint8PtrSliceConsumer) ToSilentUint8PtrSliceConsumer() SilentUint8PtrSliceConsumer {
	return func(v []*uint8) {
		_ = c(v)
		return
	}
}

// ToSilentConsumer transforms SilentUint8PtrSliceConsumer into SilentConsumer
func (sc SilentUint8PtrSliceConsumer) ToSilentConsumer() SilentConsumer {
	return func(v interface{}) {
		sc(v.([]*uint8))
	}
}

// AndThen returns a composed SilentUint8PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentUint8PtrSliceConsumer) AndThen(after SilentUint8PtrSliceConsumer) SilentUint8PtrSliceConsumer {
	if after != nil {
		return func(v []*uint8) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustUint8PtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustUint8PtrSliceConsumer func(v []*uint8)

// ToMustUint8PtrSliceConsumer transforms Uint8PtrSliceConsumer into MustUint8PtrSliceConsumer
func (c Uint8PtrSliceConsumer) ToMustUint8PtrSliceConsumer() MustUint8PtrSliceConsumer {
	return func(v []*uint8) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// ToMustConsumer transforms MustUint8PtrSliceConsumer into MustConsumer
func (mc MustUint8PtrSliceConsumer) ToMustConsumer() MustConsumer {
	return func(v interface{}) {
		mc(v.([]*uint8))
	}
}

// AndThen returns a composed MustUint8PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustUint8PtrSliceConsumer) AndThen(after MustUint8PtrSliceConsumer) MustUint8PtrSliceConsumer {
	if after != nil {
		return func(v []*uint8) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToUint8PtrSliceConsumer transforms MustUint8PtrSliceConsumer into Uint8PtrSliceConsumer
func (mc MustUint8PtrSliceConsumer) ToUint8PtrSliceConsumer() Uint8PtrSliceConsumer {
	return func(v []*uint8) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentUint8PtrSliceConsumer transforms MustUint8PtrSliceConsumer into SilentUint8PtrSliceConsumer
func (mc MustUint8PtrSliceConsumer) ToSilentUint8PtrSliceConsumer() SilentUint8PtrSliceConsumer {
	c := mc.ToUint8PtrSliceConsumer()
	return c.ToSilentUint8PtrSliceConsumer()
}
