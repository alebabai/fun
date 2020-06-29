// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Int16Consumer represents an operation that accepts a single input argument or returns an error.
type Int16Consumer func(v int16) error

// ToConsumer transforms Int16Consumer into Consumer
func (c Int16Consumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(int16))
	}
}

// AndThen returns a composed Int16Consumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Int16Consumer) AndThen(after Int16Consumer) Int16Consumer {
	if after != nil {
		return func(v int16) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentInt16Consumer represents an operation that accepts a single input argument without returning an error.
type SilentInt16Consumer func(v int16)

// ToSilentInt16Consumer transforms Int16Consumer into SilentInt16Consumer
func (c Int16Consumer) ToSilentInt16Consumer() SilentInt16Consumer {
	return func(v int16) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentInt16Consumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentInt16Consumer) AndThen(after SilentInt16Consumer) SilentInt16Consumer {
	if after != nil {
		return func(v int16) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustInt16Consumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustInt16Consumer func(v int16)

// ToMustInt16Consumer transforms Int16Consumer into MustInt16Consumer
func (c Int16Consumer) ToMustInt16Consumer() MustInt16Consumer {
	return func(v int16) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustInt16Consumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustInt16Consumer) AndThen(after MustInt16Consumer) MustInt16Consumer {
	if after != nil {
		return func(v int16) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToInt16Consumer transforms MustInt16Consumer into Int16Consumer
func (mc MustInt16Consumer) ToInt16Consumer() Int16Consumer {
	return func(v int16) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// SilentInt16Consumer transforms MustInt16Consumer into SilentInt16Consumer
func (mc MustInt16Consumer) ToSilentInt16Consumer() SilentInt16Consumer {
	c := mc.ToInt16Consumer()
	return c.ToSilentInt16Consumer()
}
