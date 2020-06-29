// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Int32Consumer represents an operation that accepts a single input argument or returns an error.
type Int32Consumer func(v int32) error

// ToConsumer transforms Int32Consumer into Consumer
func (c Int32Consumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(int32))
	}
}

// AndThen returns a composed Int32Consumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Int32Consumer) AndThen(after Int32Consumer) Int32Consumer {
	if after != nil {
		return func(v int32) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentInt32Consumer represents an operation that accepts a single input argument without returning an error.
type SilentInt32Consumer func(v int32)

// ToSilentInt32Consumer transforms Int32Consumer into SilentInt32Consumer
func (c Int32Consumer) ToSilentInt32Consumer() SilentInt32Consumer {
	return func(v int32) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentInt32Consumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentInt32Consumer) AndThen(after SilentInt32Consumer) SilentInt32Consumer {
	if after != nil {
		return func(v int32) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustInt32Consumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustInt32Consumer func(v int32)

// ToMustInt32Consumer transforms Int32Consumer into MustInt32Consumer
func (c Int32Consumer) ToMustInt32Consumer() MustInt32Consumer {
	return func(v int32) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustInt32Consumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustInt32Consumer) AndThen(after MustInt32Consumer) MustInt32Consumer {
	if after != nil {
		return func(v int32) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToInt32Consumer transforms MustInt32Consumer into Int32Consumer
func (mc MustInt32Consumer) ToInt32Consumer() Int32Consumer {
	return func(v int32) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// SilentInt32Consumer transforms MustInt32Consumer into SilentInt32Consumer
func (mc MustInt32Consumer) ToSilentInt32Consumer() SilentInt32Consumer {
	c := mc.ToInt32Consumer()
	return c.ToSilentInt32Consumer()
}
