// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Int32SliceConsumer represents an operation that accepts a single input argument or returns an error.
type Int32SliceConsumer func(v []int32) error

// ToConsumer transforms Int32SliceConsumer into Consumer
func (c Int32SliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]int32))
	}
}

// AndThen returns a composed Int32SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Int32SliceConsumer) AndThen(after Int32SliceConsumer) Int32SliceConsumer {
	if after != nil {
		return func(v []int32) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentInt32SliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentInt32SliceConsumer func(v []int32)

// ToSilentInt32SliceConsumer transforms Int32SliceConsumer into SilentInt32SliceConsumer
func (c Int32SliceConsumer) ToSilentInt32SliceConsumer() SilentInt32SliceConsumer {
	return func(v []int32) {
		_ = c(v)
		return
	}
}

// ToSilentConsumer transforms SilentInt32SliceConsumer into SilentConsumer
func (sc SilentInt32SliceConsumer) ToSilentConsumer() SilentConsumer {
	return func(v interface{}) {
		sc(v.([]int32))
	}
}

// AndThen returns a composed SilentInt32SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentInt32SliceConsumer) AndThen(after SilentInt32SliceConsumer) SilentInt32SliceConsumer {
	if after != nil {
		return func(v []int32) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustInt32SliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustInt32SliceConsumer func(v []int32)

// ToMustInt32SliceConsumer transforms Int32SliceConsumer into MustInt32SliceConsumer
func (c Int32SliceConsumer) ToMustInt32SliceConsumer() MustInt32SliceConsumer {
	return func(v []int32) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// ToMustConsumer transforms MustInt32SliceConsumer into MustConsumer
func (mc MustInt32SliceConsumer) ToMustConsumer() MustConsumer {
	return func(v interface{}) {
		mc(v.([]int32))
	}
}

// AndThen returns a composed MustInt32SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustInt32SliceConsumer) AndThen(after MustInt32SliceConsumer) MustInt32SliceConsumer {
	if after != nil {
		return func(v []int32) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToInt32SliceConsumer transforms MustInt32SliceConsumer into Int32SliceConsumer
func (mc MustInt32SliceConsumer) ToInt32SliceConsumer() Int32SliceConsumer {
	return func(v []int32) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentInt32SliceConsumer transforms MustInt32SliceConsumer into SilentInt32SliceConsumer
func (mc MustInt32SliceConsumer) ToSilentInt32SliceConsumer() SilentInt32SliceConsumer {
	c := mc.ToInt32SliceConsumer()
	return c.ToSilentInt32SliceConsumer()
}
