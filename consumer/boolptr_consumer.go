// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// BoolPtrConsumer represents an operation that accepts a single input argument or returns an error.
type BoolPtrConsumer func(v *bool) error

// ToConsumer transforms BoolPtrConsumer into Consumer
func (c BoolPtrConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(*bool))
	}
}

// AndThen returns a composed BoolPtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c BoolPtrConsumer) AndThen(after BoolPtrConsumer) BoolPtrConsumer {
	if after != nil {
		return func(v *bool) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentBoolPtrConsumer represents an operation that accepts a single input argument without returning an error.
type SilentBoolPtrConsumer func(v *bool)

// ToSilentBoolPtrConsumer transforms BoolPtrConsumer into SilentBoolPtrConsumer
func (c BoolPtrConsumer) ToSilentBoolPtrConsumer() SilentBoolPtrConsumer {
	return func(v *bool) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentBoolPtrConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentBoolPtrConsumer) AndThen(after SilentBoolPtrConsumer) SilentBoolPtrConsumer {
	if after != nil {
		return func(v *bool) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustBoolPtrConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustBoolPtrConsumer func(v *bool)

// ToMustBoolPtrConsumer transforms BoolPtrConsumer into MustBoolPtrConsumer
func (c BoolPtrConsumer) ToMustBoolPtrConsumer() MustBoolPtrConsumer {
	return func(v *bool) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustBoolPtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustBoolPtrConsumer) AndThen(after MustBoolPtrConsumer) MustBoolPtrConsumer {
	if after != nil {
		return func(v *bool) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToBoolPtrConsumer transforms MustBoolPtrConsumer into BoolPtrConsumer
func (mc MustBoolPtrConsumer) ToBoolPtrConsumer() BoolPtrConsumer {
	return func(v *bool) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentBoolPtrConsumer transforms MustBoolPtrConsumer into SilentBoolPtrConsumer
func (mc MustBoolPtrConsumer) ToSilentBoolPtrConsumer() SilentBoolPtrConsumer {
	c := mc.ToBoolPtrConsumer()
	return c.ToSilentBoolPtrConsumer()
}
