// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Int8SliceConsumer represents an operation that accepts a single input argument or returns an error.
type Int8SliceConsumer func(v []int8) error

// ToConsumer transforms Int8SliceConsumer into Consumer
func (c Int8SliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]int8))
	}
}

// AndThen returns a composed Int8SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Int8SliceConsumer) AndThen(after Int8SliceConsumer) Int8SliceConsumer {
	if after != nil {
		return func(v []int8) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentInt8SliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentInt8SliceConsumer func(v []int8)

// ToSilentInt8SliceConsumer transforms Int8SliceConsumer into SilentInt8SliceConsumer
func (c Int8SliceConsumer) ToSilentInt8SliceConsumer() SilentInt8SliceConsumer {
	return func(v []int8) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentInt8SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentInt8SliceConsumer) AndThen(after SilentInt8SliceConsumer) SilentInt8SliceConsumer {
	if after != nil {
		return func(v []int8) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustInt8SliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustInt8SliceConsumer func(v []int8)

// ToMustInt8SliceConsumer transforms Int8SliceConsumer into MustInt8SliceConsumer
func (c Int8SliceConsumer) ToMustInt8SliceConsumer() MustInt8SliceConsumer {
	return func(v []int8) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustInt8SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustInt8SliceConsumer) AndThen(after MustInt8SliceConsumer) MustInt8SliceConsumer {
	if after != nil {
		return func(v []int8) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToInt8SliceConsumer transforms MustInt8SliceConsumer into Int8SliceConsumer
func (mc MustInt8SliceConsumer) ToInt8SliceConsumer() Int8SliceConsumer {
	return func(v []int8) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentInt8SliceConsumer transforms MustInt8SliceConsumer into SilentInt8SliceConsumer
func (mc MustInt8SliceConsumer) ToSilentInt8SliceConsumer() SilentInt8SliceConsumer {
	c := mc.ToInt8SliceConsumer()
	return c.ToSilentInt8SliceConsumer()
}
