// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Int16SliceConsumer represents an operation that accepts a single input argument or returns an error.
type Int16SliceConsumer func(v []int16) error

// ToConsumer transforms Int16SliceConsumer into Consumer
func (c Int16SliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]int16))
	}
}

// AndThen returns a composed Int16SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Int16SliceConsumer) AndThen(after Int16SliceConsumer) Int16SliceConsumer {
	if after != nil {
		return func(v []int16) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentInt16SliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentInt16SliceConsumer func(v []int16)

// ToSilentInt16SliceConsumer transforms Int16SliceConsumer into SilentInt16SliceConsumer
func (c Int16SliceConsumer) ToSilentInt16SliceConsumer() SilentInt16SliceConsumer {
	return func(v []int16) {
		_ = c(v)
		return
	}
}

// ToSilentConsumer transforms SilentInt16SliceConsumer into SilentConsumer
func (sc SilentInt16SliceConsumer) ToSilentConsumer() SilentConsumer {
	return func(v interface{}) {
		sc(v.([]int16))
	}
}

// AndThen returns a composed SilentInt16SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentInt16SliceConsumer) AndThen(after SilentInt16SliceConsumer) SilentInt16SliceConsumer {
	if after != nil {
		return func(v []int16) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustInt16SliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustInt16SliceConsumer func(v []int16)

// ToMustInt16SliceConsumer transforms Int16SliceConsumer into MustInt16SliceConsumer
func (c Int16SliceConsumer) ToMustInt16SliceConsumer() MustInt16SliceConsumer {
	return func(v []int16) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// ToMustConsumer transforms MustInt16SliceConsumer into MustConsumer
func (mc MustInt16SliceConsumer) ToMustConsumer() MustConsumer {
	return func(v interface{}) {
		mc(v.([]int16))
	}
}

// AndThen returns a composed MustInt16SliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustInt16SliceConsumer) AndThen(after MustInt16SliceConsumer) MustInt16SliceConsumer {
	if after != nil {
		return func(v []int16) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToInt16SliceConsumer transforms MustInt16SliceConsumer into Int16SliceConsumer
func (mc MustInt16SliceConsumer) ToInt16SliceConsumer() Int16SliceConsumer {
	return func(v []int16) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentInt16SliceConsumer transforms MustInt16SliceConsumer into SilentInt16SliceConsumer
func (mc MustInt16SliceConsumer) ToSilentInt16SliceConsumer() SilentInt16SliceConsumer {
	c := mc.ToInt16SliceConsumer()
	return c.ToSilentInt16SliceConsumer()
}
