// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// StringPtrSliceConsumer represents an operation that accepts a single input argument or returns an error.
type StringPtrSliceConsumer func(v []*string) error

// ToConsumer transforms StringPtrSliceConsumer into Consumer
func (c StringPtrSliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]*string))
	}
}

// AndThen returns a composed StringPtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c StringPtrSliceConsumer) AndThen(after StringPtrSliceConsumer) StringPtrSliceConsumer {
	if after != nil {
		return func(v []*string) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentStringPtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentStringPtrSliceConsumer func(v []*string)

// ToSilentStringPtrSliceConsumer transforms StringPtrSliceConsumer into SilentStringPtrSliceConsumer
func (c StringPtrSliceConsumer) ToSilentStringPtrSliceConsumer() SilentStringPtrSliceConsumer {
	return func(v []*string) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentStringPtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentStringPtrSliceConsumer) AndThen(after SilentStringPtrSliceConsumer) SilentStringPtrSliceConsumer {
	if after != nil {
		return func(v []*string) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustStringPtrSliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustStringPtrSliceConsumer func(v []*string)

// ToMustStringPtrSliceConsumer transforms StringPtrSliceConsumer into MustStringPtrSliceConsumer
func (c StringPtrSliceConsumer) ToMustStringPtrSliceConsumer() MustStringPtrSliceConsumer {
	return func(v []*string) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustStringPtrSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustStringPtrSliceConsumer) AndThen(after MustStringPtrSliceConsumer) MustStringPtrSliceConsumer {
	if after != nil {
		return func(v []*string) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToStringPtrSliceConsumer transforms MustStringPtrSliceConsumer into StringPtrSliceConsumer
func (mc MustStringPtrSliceConsumer) ToStringPtrSliceConsumer() StringPtrSliceConsumer {
	return func(v []*string) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// SilentStringPtrSliceConsumer transforms MustStringPtrSliceConsumer into SilentStringPtrSliceConsumer
func (mc MustStringPtrSliceConsumer) ToSilentStringPtrSliceConsumer() SilentStringPtrSliceConsumer {
	c := mc.ToStringPtrSliceConsumer()
	return c.ToSilentStringPtrSliceConsumer()
}
