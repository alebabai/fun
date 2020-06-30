// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Int64SliceConsumer represents an operation that accepts a single input argument or returns an error.
type Int64SliceConsumer func(v []int64) error

// ToConsumer transforms Int64SliceConsumer into Consumer
func (c Int64SliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]int64))
	}
}

// AndThen returns a composed Int64SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Int64SliceConsumer) AndThen(after Int64SliceConsumer) Int64SliceConsumer {
	if after != nil {
		return func(v []int64) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentInt64SliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentInt64SliceConsumer func(v []int64)

// ToSilentInt64SliceConsumer transforms Int64SliceConsumer into SilentInt64SliceConsumer
func (c Int64SliceConsumer) ToSilentInt64SliceConsumer() SilentInt64SliceConsumer {
	return func(v []int64) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentInt64SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentInt64SliceConsumer) AndThen(after SilentInt64SliceConsumer) SilentInt64SliceConsumer {
	if after != nil {
		return func(v []int64) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustInt64SliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustInt64SliceConsumer func(v []int64)

// ToMustInt64SliceConsumer transforms Int64SliceConsumer into MustInt64SliceConsumer
func (c Int64SliceConsumer) ToMustInt64SliceConsumer() MustInt64SliceConsumer {
	return func(v []int64) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustInt64SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustInt64SliceConsumer) AndThen(after MustInt64SliceConsumer) MustInt64SliceConsumer {
	if after != nil {
		return func(v []int64) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToInt64SliceConsumer transforms MustInt64SliceConsumer into Int64SliceConsumer
func (mc MustInt64SliceConsumer) ToInt64SliceConsumer() Int64SliceConsumer {
	return func(v []int64) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentInt64SliceConsumer transforms MustInt64SliceConsumer into SilentInt64SliceConsumer
func (mc MustInt64SliceConsumer) ToSilentInt64SliceConsumer() SilentInt64SliceConsumer {
	c := mc.ToInt64SliceConsumer()
	return c.ToSilentInt64SliceConsumer()
}
