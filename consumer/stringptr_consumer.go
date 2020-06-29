// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// StringPtrConsumer represents an operation that accepts a single input argument or returns an error.
type StringPtrConsumer func(v *string) error

// ToConsumer transforms StringPtrConsumer into Consumer
func (c StringPtrConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(*string))
	}
}

// AndThen returns a composed StringPtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c StringPtrConsumer) AndThen(after StringPtrConsumer) StringPtrConsumer {
	if after != nil {
		return func(v *string) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentStringPtrConsumer represents an operation that accepts a single input argument without returning an error.
type SilentStringPtrConsumer func(v *string)

// ToSilentStringPtrConsumer transforms StringPtrConsumer into SilentStringPtrConsumer
func (c StringPtrConsumer) ToSilentStringPtrConsumer() SilentStringPtrConsumer {
	return func(v *string) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentStringPtrConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentStringPtrConsumer) AndThen(after SilentStringPtrConsumer) SilentStringPtrConsumer {
	if after != nil {
		return func(v *string) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustStringPtrConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustStringPtrConsumer func(v *string)

// ToMustStringPtrConsumer transforms StringPtrConsumer into MustStringPtrConsumer
func (c StringPtrConsumer) ToMustStringPtrConsumer() MustStringPtrConsumer {
	return func(v *string) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustStringPtrConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustStringPtrConsumer) AndThen(after MustStringPtrConsumer) MustStringPtrConsumer {
	if after != nil {
		return func(v *string) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToStringPtrConsumer transforms MustStringPtrConsumer into StringPtrConsumer
func (mc MustStringPtrConsumer) ToStringPtrConsumer() StringPtrConsumer {
	return func(v *string) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// SilentStringPtrConsumer transforms MustStringPtrConsumer into SilentStringPtrConsumer
func (mc MustStringPtrConsumer) ToSilentStringPtrConsumer() SilentStringPtrConsumer {
	c := mc.ToStringPtrConsumer()
	return c.ToSilentStringPtrConsumer()
}
