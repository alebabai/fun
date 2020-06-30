// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Int8Consumer represents an operation that accepts a single input argument or returns an error.
type Int8Consumer func(v int8) error

// ToConsumer transforms Int8Consumer into Consumer
func (c Int8Consumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(int8))
	}
}

// AndThen returns a composed Int8Consumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Int8Consumer) AndThen(after Int8Consumer) Int8Consumer {
	if after != nil {
		return func(v int8) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentInt8Consumer represents an operation that accepts a single input argument without returning an error.
type SilentInt8Consumer func(v int8)

// ToSilentInt8Consumer transforms Int8Consumer into SilentInt8Consumer
func (c Int8Consumer) ToSilentInt8Consumer() SilentInt8Consumer {
	return func(v int8) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentInt8Consumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentInt8Consumer) AndThen(after SilentInt8Consumer) SilentInt8Consumer {
	if after != nil {
		return func(v int8) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustInt8Consumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustInt8Consumer func(v int8)

// ToMustInt8Consumer transforms Int8Consumer into MustInt8Consumer
func (c Int8Consumer) ToMustInt8Consumer() MustInt8Consumer {
	return func(v int8) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustInt8Consumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustInt8Consumer) AndThen(after MustInt8Consumer) MustInt8Consumer {
	if after != nil {
		return func(v int8) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToInt8Consumer transforms MustInt8Consumer into Int8Consumer
func (mc MustInt8Consumer) ToInt8Consumer() Int8Consumer {
	return func(v int8) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentInt8Consumer transforms MustInt8Consumer into SilentInt8Consumer
func (mc MustInt8Consumer) ToSilentInt8Consumer() SilentInt8Consumer {
	c := mc.ToInt8Consumer()
	return c.ToSilentInt8Consumer()
}
