// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Uint64PtrSliceConsumer represents an operation that accepts a single input argument or returns an error.
type Uint64PtrSliceConsumer func(v []*uint64) error

// ToConsumer transforms Uint64PtrSliceConsumer into Consumer
func (c Uint64PtrSliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]*uint64))
	}
}

// AndThen returns a composed Uint64PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Uint64PtrSliceConsumer) AndThen(after Uint64PtrSliceConsumer) Uint64PtrSliceConsumer {
	if after != nil {
		return func(v []*uint64) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentUint64PtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentUint64PtrSliceConsumer func(v []*uint64)

// ToSilentUint64PtrSliceConsumer transforms Uint64PtrSliceConsumer into SilentUint64PtrSliceConsumer
func (c Uint64PtrSliceConsumer) ToSilentUint64PtrSliceConsumer() SilentUint64PtrSliceConsumer {
	return func(v []*uint64) {
		_ = c(v)
		return
	}
}

// ToSilentConsumer transforms SilentUint64PtrSliceConsumer into SilentConsumer
func (sc SilentUint64PtrSliceConsumer) ToSilentConsumer() SilentConsumer {
	return func(v interface{}) {
		sc(v.([]*uint64))
	}
}

// AndThen returns a composed SilentUint64PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentUint64PtrSliceConsumer) AndThen(after SilentUint64PtrSliceConsumer) SilentUint64PtrSliceConsumer {
	if after != nil {
		return func(v []*uint64) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustUint64PtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustUint64PtrSliceConsumer func(v []*uint64)

// ToMustUint64PtrSliceConsumer transforms Uint64PtrSliceConsumer into MustUint64PtrSliceConsumer
func (c Uint64PtrSliceConsumer) ToMustUint64PtrSliceConsumer() MustUint64PtrSliceConsumer {
	return func(v []*uint64) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// ToMustConsumer transforms MustUint64PtrSliceConsumer into MustConsumer
func (mc MustUint64PtrSliceConsumer) ToMustConsumer() MustConsumer {
	return func(v interface{}) {
		mc(v.([]*uint64))
	}
}

// AndThen returns a composed MustUint64PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustUint64PtrSliceConsumer) AndThen(after MustUint64PtrSliceConsumer) MustUint64PtrSliceConsumer {
	if after != nil {
		return func(v []*uint64) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToUint64PtrSliceConsumer transforms MustUint64PtrSliceConsumer into Uint64PtrSliceConsumer
func (mc MustUint64PtrSliceConsumer) ToUint64PtrSliceConsumer() Uint64PtrSliceConsumer {
	return func(v []*uint64) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentUint64PtrSliceConsumer transforms MustUint64PtrSliceConsumer into SilentUint64PtrSliceConsumer
func (mc MustUint64PtrSliceConsumer) ToSilentUint64PtrSliceConsumer() SilentUint64PtrSliceConsumer {
	c := mc.ToUint64PtrSliceConsumer()
	return c.ToSilentUint64PtrSliceConsumer()
}
