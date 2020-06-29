// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Int64PtrSliceConsumer represents an operation that accepts a single input argument or returns an error.
type Int64PtrSliceConsumer func(v []*int64) error

// ToConsumer transforms Int64PtrSliceConsumer into Consumer
func (c Int64PtrSliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]*int64))
	}
}

// AndThen returns a composed Int64PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Int64PtrSliceConsumer) AndThen(after Int64PtrSliceConsumer) Int64PtrSliceConsumer {
	if after != nil {
		return func(v []*int64) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentInt64PtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentInt64PtrSliceConsumer func(v []*int64)

// ToSilentInt64PtrSliceConsumer transforms Int64PtrSliceConsumer into SilentInt64PtrSliceConsumer
func (c Int64PtrSliceConsumer) ToSilentInt64PtrSliceConsumer() SilentInt64PtrSliceConsumer {
	return func(v []*int64) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentInt64PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentInt64PtrSliceConsumer) AndThen(after SilentInt64PtrSliceConsumer) SilentInt64PtrSliceConsumer {
	if after != nil {
		return func(v []*int64) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustInt64PtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustInt64PtrSliceConsumer func(v []*int64)

// ToMustInt64PtrSliceConsumer transforms Int64PtrSliceConsumer into MustInt64PtrSliceConsumer
func (c Int64PtrSliceConsumer) ToMustInt64PtrSliceConsumer() MustInt64PtrSliceConsumer {
	return func(v []*int64) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustInt64PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustInt64PtrSliceConsumer) AndThen(after MustInt64PtrSliceConsumer) MustInt64PtrSliceConsumer {
	if after != nil {
		return func(v []*int64) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToInt64PtrSliceConsumer transforms MustInt64PtrSliceConsumer into Int64PtrSliceConsumer
func (mc MustInt64PtrSliceConsumer) ToInt64PtrSliceConsumer() Int64PtrSliceConsumer {
	return func(v []*int64) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// SilentInt64PtrSliceConsumer transforms MustInt64PtrSliceConsumer into SilentInt64PtrSliceConsumer
func (mc MustInt64PtrSliceConsumer) ToSilentInt64PtrSliceConsumer() SilentInt64PtrSliceConsumer {
	c := mc.ToInt64PtrSliceConsumer()
	return c.ToSilentInt64PtrSliceConsumer()
}
