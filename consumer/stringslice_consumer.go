// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// StringSliceConsumer represents an operation that accepts a single input argument or returns an error.
type StringSliceConsumer func(v []string) error

// ToConsumer transforms StringSliceConsumer into Consumer
func (c StringSliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]string))
	}
}

// AndThen returns a composed StringSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c StringSliceConsumer) AndThen(after StringSliceConsumer) StringSliceConsumer {
	if after != nil {
		return func(v []string) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentStringSliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentStringSliceConsumer func(v []string)

// ToSilentStringSliceConsumer transforms StringSliceConsumer into SilentStringSliceConsumer
func (c StringSliceConsumer) ToSilentStringSliceConsumer() SilentStringSliceConsumer {
	return func(v []string) {
		_ = c(v)
		return
	}
}

// ToSilentConsumer transforms SilentStringSliceConsumer into SilentConsumer
func (sc SilentStringSliceConsumer) ToSilentConsumer() SilentConsumer {
	return func(v interface{}) {
		sc(v.([]string))
	}
}

// AndThen returns a composed SilentStringSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentStringSliceConsumer) AndThen(after SilentStringSliceConsumer) SilentStringSliceConsumer {
	if after != nil {
		return func(v []string) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustStringSliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustStringSliceConsumer func(v []string)

// ToMustStringSliceConsumer transforms StringSliceConsumer into MustStringSliceConsumer
func (c StringSliceConsumer) ToMustStringSliceConsumer() MustStringSliceConsumer {
	return func(v []string) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// ToMustConsumer transforms MustStringSliceConsumer into MustConsumer
func (mc MustStringSliceConsumer) ToMustConsumer() MustConsumer {
	return func(v interface{}) {
		mc(v.([]string))
	}
}

// AndThen returns a composed MustStringSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustStringSliceConsumer) AndThen(after MustStringSliceConsumer) MustStringSliceConsumer {
	if after != nil {
		return func(v []string) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToStringSliceConsumer transforms MustStringSliceConsumer into StringSliceConsumer
func (mc MustStringSliceConsumer) ToStringSliceConsumer() StringSliceConsumer {
	return func(v []string) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentStringSliceConsumer transforms MustStringSliceConsumer into SilentStringSliceConsumer
func (mc MustStringSliceConsumer) ToSilentStringSliceConsumer() SilentStringSliceConsumer {
	c := mc.ToStringSliceConsumer()
	return c.ToSilentStringSliceConsumer()
}
