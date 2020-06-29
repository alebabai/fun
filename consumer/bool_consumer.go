// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// BoolConsumer represents an operation that accepts a single input argument or returns an error.
type BoolConsumer func(v bool) error

// ToConsumer transforms BoolConsumer into Consumer
func (c BoolConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(bool))
	}
}

// AndThen returns a composed BoolConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c BoolConsumer) AndThen(after BoolConsumer) BoolConsumer {
	if after != nil {
		return func(v bool) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentBoolConsumer represents an operation that accepts a single input argument without returning an error.
type SilentBoolConsumer func(v bool)

// ToSilentBoolConsumer transforms BoolConsumer into SilentBoolConsumer
func (c BoolConsumer) ToSilentBoolConsumer() SilentBoolConsumer {
	return func(v bool) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentBoolConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentBoolConsumer) AndThen(after SilentBoolConsumer) SilentBoolConsumer {
	if after != nil {
		return func(v bool) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustBoolConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustBoolConsumer func(v bool)

// ToMustBoolConsumer transforms BoolConsumer into MustBoolConsumer
func (c BoolConsumer) ToMustBoolConsumer() MustBoolConsumer {
	return func(v bool) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustBoolConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustBoolConsumer) AndThen(after MustBoolConsumer) MustBoolConsumer {
	if after != nil {
		return func(v bool) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToBoolConsumer transforms MustBoolConsumer into BoolConsumer
func (mc MustBoolConsumer) ToBoolConsumer() BoolConsumer {
	return func(v bool) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// SilentBoolConsumer transforms MustBoolConsumer into SilentBoolConsumer
func (mc MustBoolConsumer) ToSilentBoolConsumer() SilentBoolConsumer {
	c := mc.ToBoolConsumer()
	return c.ToSilentBoolConsumer()
}
