// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// StringConsumer represents an operation that accepts a single input argument or returns an error.
type StringConsumer func(v string) error

// ToConsumer transforms StringConsumer into Consumer
func (c StringConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(string))
	}
}

// AndThen returns a composed StringConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c StringConsumer) AndThen(after StringConsumer) StringConsumer {
	if after != nil {
		return func(v string) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentStringConsumer represents an operation that accepts a single input argument without returning an error.
type SilentStringConsumer func(v string)

// ToSilentStringConsumer transforms StringConsumer into SilentStringConsumer
func (c StringConsumer) ToSilentStringConsumer() SilentStringConsumer {
	return func(v string) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentStringConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentStringConsumer) AndThen(after SilentStringConsumer) SilentStringConsumer {
	if after != nil {
		return func(v string) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustStringConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustStringConsumer func(v string)

// ToMustStringConsumer transforms StringConsumer into MustStringConsumer
func (c StringConsumer) ToMustStringConsumer() MustStringConsumer {
	return func(v string) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustStringConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustStringConsumer) AndThen(after MustStringConsumer) MustStringConsumer {
	if after != nil {
		return func(v string) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToStringConsumer transforms MustStringConsumer into StringConsumer
func (mc MustStringConsumer) ToStringConsumer() StringConsumer {
	return func(v string) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// SilentStringConsumer transforms MustStringConsumer into SilentStringConsumer
func (mc MustStringConsumer) ToSilentStringConsumer() SilentStringConsumer {
	c := mc.ToStringConsumer()
	return c.ToSilentStringConsumer()
}
