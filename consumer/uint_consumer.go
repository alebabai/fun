// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// UintConsumer represents an operation that accepts a single input argument or returns an error.
type UintConsumer func(v uint) error

// ToConsumer transforms UintConsumer into Consumer
func (c UintConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(uint))
	}
}

// AndThen returns a composed UintConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c UintConsumer) AndThen(after UintConsumer) UintConsumer {
	if after != nil {
		return func(v uint) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentUintConsumer represents an operation that accepts a single input argument without returning an error.
type SilentUintConsumer func(v uint)

// ToSilentUintConsumer transforms UintConsumer into SilentUintConsumer
func (c UintConsumer) ToSilentUintConsumer() SilentUintConsumer {
	return func(v uint) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentUintConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentUintConsumer) AndThen(after SilentUintConsumer) SilentUintConsumer {
	if after != nil {
		return func(v uint) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustUintConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustUintConsumer func(v uint)

// ToMustUintConsumer transforms UintConsumer into MustUintConsumer
func (c UintConsumer) ToMustUintConsumer() MustUintConsumer {
	return func(v uint) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustUintConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustUintConsumer) AndThen(after MustUintConsumer) MustUintConsumer {
	if after != nil {
		return func(v uint) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToUintConsumer transforms MustUintConsumer into UintConsumer
func (mc MustUintConsumer) ToUintConsumer() UintConsumer {
	return func(v uint) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// SilentUintConsumer transforms MustUintConsumer into SilentUintConsumer
func (mc MustUintConsumer) ToSilentUintConsumer() SilentUintConsumer {
	c := mc.ToUintConsumer()
	return c.ToSilentUintConsumer()
}
