// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// BoolPtrSliceConsumer represents an operation that accepts a single input argument or returns an error.
type BoolPtrSliceConsumer func(v []*bool) error

// ToConsumer transforms BoolPtrSliceConsumer into Consumer
func (c BoolPtrSliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]*bool))
	}
}

// AndThen returns a composed BoolPtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c BoolPtrSliceConsumer) AndThen(after BoolPtrSliceConsumer) BoolPtrSliceConsumer {
	if after != nil {
		return func(v []*bool) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentBoolPtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentBoolPtrSliceConsumer func(v []*bool)

// ToSilentBoolPtrSliceConsumer transforms BoolPtrSliceConsumer into SilentBoolPtrSliceConsumer
func (c BoolPtrSliceConsumer) ToSilentBoolPtrSliceConsumer() SilentBoolPtrSliceConsumer {
	return func(v []*bool) {
		_ = c(v)
		return
	}
}

// ToSilentConsumer transforms SilentBoolPtrSliceConsumer into SilentConsumer
func (sc SilentBoolPtrSliceConsumer) ToSilentConsumer() SilentConsumer {
	return func(v interface{}) {
		sc(v.([]*bool))
	}
}

// AndThen returns a composed SilentBoolPtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentBoolPtrSliceConsumer) AndThen(after SilentBoolPtrSliceConsumer) SilentBoolPtrSliceConsumer {
	if after != nil {
		return func(v []*bool) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustBoolPtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustBoolPtrSliceConsumer func(v []*bool)

// ToMustBoolPtrSliceConsumer transforms BoolPtrSliceConsumer into MustBoolPtrSliceConsumer
func (c BoolPtrSliceConsumer) ToMustBoolPtrSliceConsumer() MustBoolPtrSliceConsumer {
	return func(v []*bool) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// ToMustConsumer transforms MustBoolPtrSliceConsumer into MustConsumer
func (mc MustBoolPtrSliceConsumer) ToMustConsumer() MustConsumer {
	return func(v interface{}) {
		mc(v.([]*bool))
	}
}

// AndThen returns a composed MustBoolPtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustBoolPtrSliceConsumer) AndThen(after MustBoolPtrSliceConsumer) MustBoolPtrSliceConsumer {
	if after != nil {
		return func(v []*bool) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToBoolPtrSliceConsumer transforms MustBoolPtrSliceConsumer into BoolPtrSliceConsumer
func (mc MustBoolPtrSliceConsumer) ToBoolPtrSliceConsumer() BoolPtrSliceConsumer {
	return func(v []*bool) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentBoolPtrSliceConsumer transforms MustBoolPtrSliceConsumer into SilentBoolPtrSliceConsumer
func (mc MustBoolPtrSliceConsumer) ToSilentBoolPtrSliceConsumer() SilentBoolPtrSliceConsumer {
	c := mc.ToBoolPtrSliceConsumer()
	return c.ToSilentBoolPtrSliceConsumer()
}
