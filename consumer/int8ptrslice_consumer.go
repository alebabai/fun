// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Int8PtrSliceConsumer represents an operation that accepts a single input argument or returns an error.
type Int8PtrSliceConsumer func(v []*int8) error

// ToConsumer transforms Int8PtrSliceConsumer into Consumer
func (c Int8PtrSliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]*int8))
	}
}

// AndThen returns a composed Int8PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Int8PtrSliceConsumer) AndThen(after Int8PtrSliceConsumer) Int8PtrSliceConsumer {
	if after != nil {
		return func(v []*int8) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentInt8PtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentInt8PtrSliceConsumer func(v []*int8)

// ToSilentInt8PtrSliceConsumer transforms Int8PtrSliceConsumer into SilentInt8PtrSliceConsumer
func (c Int8PtrSliceConsumer) ToSilentInt8PtrSliceConsumer() SilentInt8PtrSliceConsumer {
	return func(v []*int8) {
		_ = c(v)
		return
	}
}

// ToSilentConsumer transforms SilentInt8PtrSliceConsumer into SilentConsumer
func (sc SilentInt8PtrSliceConsumer) ToSilentConsumer() SilentConsumer {
	return func(v interface{}) {
		sc(v.([]*int8))
	}
}

// AndThen returns a composed SilentInt8PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentInt8PtrSliceConsumer) AndThen(after SilentInt8PtrSliceConsumer) SilentInt8PtrSliceConsumer {
	if after != nil {
		return func(v []*int8) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustInt8PtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustInt8PtrSliceConsumer func(v []*int8)

// ToMustInt8PtrSliceConsumer transforms Int8PtrSliceConsumer into MustInt8PtrSliceConsumer
func (c Int8PtrSliceConsumer) ToMustInt8PtrSliceConsumer() MustInt8PtrSliceConsumer {
	return func(v []*int8) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// ToMustConsumer transforms MustInt8PtrSliceConsumer into MustConsumer
func (mc MustInt8PtrSliceConsumer) ToMustConsumer() MustConsumer {
	return func(v interface{}) {
		mc(v.([]*int8))
	}
}

// AndThen returns a composed MustInt8PtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustInt8PtrSliceConsumer) AndThen(after MustInt8PtrSliceConsumer) MustInt8PtrSliceConsumer {
	if after != nil {
		return func(v []*int8) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToInt8PtrSliceConsumer transforms MustInt8PtrSliceConsumer into Int8PtrSliceConsumer
func (mc MustInt8PtrSliceConsumer) ToInt8PtrSliceConsumer() Int8PtrSliceConsumer {
	return func(v []*int8) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentInt8PtrSliceConsumer transforms MustInt8PtrSliceConsumer into SilentInt8PtrSliceConsumer
func (mc MustInt8PtrSliceConsumer) ToSilentInt8PtrSliceConsumer() SilentInt8PtrSliceConsumer {
	c := mc.ToInt8PtrSliceConsumer()
	return c.ToSilentInt8PtrSliceConsumer()
}
