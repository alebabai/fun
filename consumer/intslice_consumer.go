// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// IntSliceConsumer represents an operation that accepts a single input argument or returns an error.
type IntSliceConsumer func(v []int) error

// ToConsumer transforms IntSliceConsumer into Consumer
func (c IntSliceConsumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.([]int))
	}
}

// AndThen returns a composed IntSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c IntSliceConsumer) AndThen(after IntSliceConsumer) IntSliceConsumer {
	if after != nil {
		return func(v []int) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentIntSliceConsumer represents an operation that accepts a single input argument without returning an error.
type SilentIntSliceConsumer func(v []int)

// ToSilentIntSliceConsumer transforms IntSliceConsumer into SilentIntSliceConsumer
func (c IntSliceConsumer) ToSilentIntSliceConsumer() SilentIntSliceConsumer {
	return func(v []int) {
		_ = c(v)
		return
	}
}

// ToSilentConsumer transforms SilentIntSliceConsumer into SilentConsumer
func (sc SilentIntSliceConsumer) ToSilentConsumer() SilentConsumer {
	return func(v interface{}) {
		sc(v.([]int))
	}
}

// AndThen returns a composed SilentIntSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentIntSliceConsumer) AndThen(after SilentIntSliceConsumer) SilentIntSliceConsumer {
	if after != nil {
		return func(v []int) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustIntSliceConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustIntSliceConsumer func(v []int)

// ToMustIntSliceConsumer transforms IntSliceConsumer into MustIntSliceConsumer
func (c IntSliceConsumer) ToMustIntSliceConsumer() MustIntSliceConsumer {
	return func(v []int) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// ToMustConsumer transforms MustIntSliceConsumer into MustConsumer
func (mc MustIntSliceConsumer) ToMustConsumer() MustConsumer {
	return func(v interface{}) {
		mc(v.([]int))
	}
}

// AndThen returns a composed MustIntSliceConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustIntSliceConsumer) AndThen(after MustIntSliceConsumer) MustIntSliceConsumer {
	if after != nil {
		return func(v []int) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToIntSliceConsumer transforms MustIntSliceConsumer into IntSliceConsumer
func (mc MustIntSliceConsumer) ToIntSliceConsumer() IntSliceConsumer {
	return func(v []int) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentIntSliceConsumer transforms MustIntSliceConsumer into SilentIntSliceConsumer
func (mc MustIntSliceConsumer) ToSilentIntSliceConsumer() SilentIntSliceConsumer {
	c := mc.ToIntSliceConsumer()
	return c.ToSilentIntSliceConsumer()
}
