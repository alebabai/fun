// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// BytePtrSliceConsumer represents an operation that accepts a single input argument or returns an error.
type BytePtrSliceConsumer func(v []*byte) error

// ToConsumer transforms BytePtrSliceConsumer into Consumer
func (c BytePtrSliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]*byte))
	}
}

// AndThen returns a composed BytePtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c BytePtrSliceConsumer) AndThen(after BytePtrSliceConsumer) BytePtrSliceConsumer {
	if after != nil {
		return func(v []*byte) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentBytePtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentBytePtrSliceConsumer func(v []*byte)

// ToSilentBytePtrSliceConsumer transforms BytePtrSliceConsumer into SilentBytePtrSliceConsumer
func (c BytePtrSliceConsumer) ToSilentBytePtrSliceConsumer() SilentBytePtrSliceConsumer {
	return func(v []*byte) {
		_ = c(v)
		return
	}
}

// ToSilentConsumer transforms SilentBytePtrSliceConsumer into SilentConsumer
func (sc SilentBytePtrSliceConsumer) ToSilentConsumer() SilentConsumer {
	return func(v interface{}) {
		sc(v.([]*byte))
	}
}

// AndThen returns a composed SilentBytePtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentBytePtrSliceConsumer) AndThen(after SilentBytePtrSliceConsumer) SilentBytePtrSliceConsumer {
	if after != nil {
		return func(v []*byte) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustBytePtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustBytePtrSliceConsumer func(v []*byte)

// ToMustBytePtrSliceConsumer transforms BytePtrSliceConsumer into MustBytePtrSliceConsumer
func (c BytePtrSliceConsumer) ToMustBytePtrSliceConsumer() MustBytePtrSliceConsumer {
	return func(v []*byte) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// ToMustConsumer transforms MustBytePtrSliceConsumer into MustConsumer
func (mc MustBytePtrSliceConsumer) ToMustConsumer() MustConsumer {
	return func(v interface{}) {
		mc(v.([]*byte))
	}
}

// AndThen returns a composed MustBytePtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustBytePtrSliceConsumer) AndThen(after MustBytePtrSliceConsumer) MustBytePtrSliceConsumer {
	if after != nil {
		return func(v []*byte) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToBytePtrSliceConsumer transforms MustBytePtrSliceConsumer into BytePtrSliceConsumer
func (mc MustBytePtrSliceConsumer) ToBytePtrSliceConsumer() BytePtrSliceConsumer {
	return func(v []*byte) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentBytePtrSliceConsumer transforms MustBytePtrSliceConsumer into SilentBytePtrSliceConsumer
func (mc MustBytePtrSliceConsumer) ToSilentBytePtrSliceConsumer() SilentBytePtrSliceConsumer {
	c := mc.ToBytePtrSliceConsumer()
	return c.ToSilentBytePtrSliceConsumer()
}
