// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// BoolSliceConsumer represents an operation that accepts a single input argument or returns an error.
type BoolSliceConsumer func(v []bool) error

// ToConsumer transforms BoolSliceConsumer into Consumer
func (c BoolSliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]bool))
	}
}

// AndThen returns a composed BoolSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c BoolSliceConsumer) AndThen(after BoolSliceConsumer) BoolSliceConsumer {
	if after != nil {
		return func(v []bool) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentBoolSliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentBoolSliceConsumer func(v []bool)

// ToSilentBoolSliceConsumer transforms BoolSliceConsumer into SilentBoolSliceConsumer
func (c BoolSliceConsumer) ToSilentBoolSliceConsumer() SilentBoolSliceConsumer {
	return func(v []bool) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentBoolSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentBoolSliceConsumer) AndThen(after SilentBoolSliceConsumer) SilentBoolSliceConsumer {
	if after != nil {
		return func(v []bool) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustBoolSliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustBoolSliceConsumer func(v []bool)

// ToMustBoolSliceConsumer transforms BoolSliceConsumer into MustBoolSliceConsumer
func (c BoolSliceConsumer) ToMustBoolSliceConsumer() MustBoolSliceConsumer {
	return func(v []bool) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustBoolSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustBoolSliceConsumer) AndThen(after MustBoolSliceConsumer) MustBoolSliceConsumer {
	if after != nil {
		return func(v []bool) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToBoolSliceConsumer transforms MustBoolSliceConsumer into BoolSliceConsumer
func (mc MustBoolSliceConsumer) ToBoolSliceConsumer() BoolSliceConsumer {
	return func(v []bool) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// SilentBoolSliceConsumer transforms MustBoolSliceConsumer into SilentBoolSliceConsumer
func (mc MustBoolSliceConsumer) ToSilentBoolSliceConsumer() SilentBoolSliceConsumer {
	c := mc.ToBoolSliceConsumer()
	return c.ToSilentBoolSliceConsumer()
}
