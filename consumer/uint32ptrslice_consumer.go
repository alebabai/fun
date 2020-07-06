// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Uint32PtrSliceConsumer represents an operation that accepts a single input argument or returns an error.
type Uint32PtrSliceConsumer func(v []*uint32) error

// ToConsumer transforms Uint32PtrSliceConsumer into Consumer
func (c Uint32PtrSliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]*uint32))
	}
}

// AndThen returns a composed Uint32PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Uint32PtrSliceConsumer) AndThen(after Uint32PtrSliceConsumer) Uint32PtrSliceConsumer {
	if after != nil {
		return func(v []*uint32) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentUint32PtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentUint32PtrSliceConsumer func(v []*uint32)

// ToSilentUint32PtrSliceConsumer transforms Uint32PtrSliceConsumer into SilentUint32PtrSliceConsumer
func (c Uint32PtrSliceConsumer) ToSilentUint32PtrSliceConsumer() SilentUint32PtrSliceConsumer {
	return func(v []*uint32) {
		_ = c(v)
		return
	}
}

// ToSilentConsumer transforms SilentUint32PtrSliceConsumer into SilentConsumer
func (sc SilentUint32PtrSliceConsumer) ToSilentConsumer() SilentConsumer {
	return func(v interface{}) {
		sc(v.([]*uint32))
	}
}

// AndThen returns a composed SilentUint32PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentUint32PtrSliceConsumer) AndThen(after SilentUint32PtrSliceConsumer) SilentUint32PtrSliceConsumer {
	if after != nil {
		return func(v []*uint32) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustUint32PtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustUint32PtrSliceConsumer func(v []*uint32)

// ToMustUint32PtrSliceConsumer transforms Uint32PtrSliceConsumer into MustUint32PtrSliceConsumer
func (c Uint32PtrSliceConsumer) ToMustUint32PtrSliceConsumer() MustUint32PtrSliceConsumer {
	return func(v []*uint32) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// ToMustConsumer transforms MustUint32PtrSliceConsumer into MustConsumer
func (mc MustUint32PtrSliceConsumer) ToMustConsumer() MustConsumer {
	return func(v interface{}) {
		mc(v.([]*uint32))
	}
}

// AndThen returns a composed MustUint32PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustUint32PtrSliceConsumer) AndThen(after MustUint32PtrSliceConsumer) MustUint32PtrSliceConsumer {
	if after != nil {
		return func(v []*uint32) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToUint32PtrSliceConsumer transforms MustUint32PtrSliceConsumer into Uint32PtrSliceConsumer
func (mc MustUint32PtrSliceConsumer) ToUint32PtrSliceConsumer() Uint32PtrSliceConsumer {
	return func(v []*uint32) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentUint32PtrSliceConsumer transforms MustUint32PtrSliceConsumer into SilentUint32PtrSliceConsumer
func (mc MustUint32PtrSliceConsumer) ToSilentUint32PtrSliceConsumer() SilentUint32PtrSliceConsumer {
	c := mc.ToUint32PtrSliceConsumer()
	return c.ToSilentUint32PtrSliceConsumer()
}
