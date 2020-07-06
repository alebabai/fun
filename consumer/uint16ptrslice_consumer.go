// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Uint16PtrSliceConsumer represents an operation that accepts a single input argument or returns an error.
type Uint16PtrSliceConsumer func(v []*uint16) error

// ToConsumer transforms Uint16PtrSliceConsumer into Consumer
func (c Uint16PtrSliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]*uint16))
	}
}

// AndThen returns a composed Uint16PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Uint16PtrSliceConsumer) AndThen(after Uint16PtrSliceConsumer) Uint16PtrSliceConsumer {
	if after != nil {
		return func(v []*uint16) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentUint16PtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentUint16PtrSliceConsumer func(v []*uint16)

// ToSilentUint16PtrSliceConsumer transforms Uint16PtrSliceConsumer into SilentUint16PtrSliceConsumer
func (c Uint16PtrSliceConsumer) ToSilentUint16PtrSliceConsumer() SilentUint16PtrSliceConsumer {
	return func(v []*uint16) {
		_ = c(v)
		return
	}
}

// ToSilentConsumer transforms SilentUint16PtrSliceConsumer into SilentConsumer
func (sc SilentUint16PtrSliceConsumer) ToSilentConsumer() SilentConsumer {
	return func(v interface{}) {
		sc(v.([]*uint16))
	}
}

// AndThen returns a composed SilentUint16PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentUint16PtrSliceConsumer) AndThen(after SilentUint16PtrSliceConsumer) SilentUint16PtrSliceConsumer {
	if after != nil {
		return func(v []*uint16) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustUint16PtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustUint16PtrSliceConsumer func(v []*uint16)

// ToMustUint16PtrSliceConsumer transforms Uint16PtrSliceConsumer into MustUint16PtrSliceConsumer
func (c Uint16PtrSliceConsumer) ToMustUint16PtrSliceConsumer() MustUint16PtrSliceConsumer {
	return func(v []*uint16) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// ToMustConsumer transforms MustUint16PtrSliceConsumer into MustConsumer
func (mc MustUint16PtrSliceConsumer) ToMustConsumer() MustConsumer {
	return func(v interface{}) {
		mc(v.([]*uint16))
	}
}

// AndThen returns a composed MustUint16PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustUint16PtrSliceConsumer) AndThen(after MustUint16PtrSliceConsumer) MustUint16PtrSliceConsumer {
	if after != nil {
		return func(v []*uint16) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToUint16PtrSliceConsumer transforms MustUint16PtrSliceConsumer into Uint16PtrSliceConsumer
func (mc MustUint16PtrSliceConsumer) ToUint16PtrSliceConsumer() Uint16PtrSliceConsumer {
	return func(v []*uint16) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentUint16PtrSliceConsumer transforms MustUint16PtrSliceConsumer into SilentUint16PtrSliceConsumer
func (mc MustUint16PtrSliceConsumer) ToSilentUint16PtrSliceConsumer() SilentUint16PtrSliceConsumer {
	c := mc.ToUint16PtrSliceConsumer()
	return c.ToSilentUint16PtrSliceConsumer()
}
