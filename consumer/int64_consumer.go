// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Int64Consumer represents an operation that accepts a single input argument or returns an error.
type Int64Consumer func(v int64) error

// ToConsumer transforms Int64Consumer into Consumer
func (c Int64Consumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(int64))
	}
}

// AndThen returns a composed Int64Consumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Int64Consumer) AndThen(after Int64Consumer) Int64Consumer {
	if after != nil {
		return func(v int64) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentInt64Consumer represents an operation that accepts a single input argument without returning an error.
type SilentInt64Consumer func(v int64)

// ToSilentInt64Consumer transforms Int64Consumer into SilentInt64Consumer
func (c Int64Consumer) ToSilentInt64Consumer() SilentInt64Consumer {
	return func(v int64) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentInt64Consumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentInt64Consumer) AndThen(after SilentInt64Consumer) SilentInt64Consumer {
	if after != nil {
		return func(v int64) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustInt64Consumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustInt64Consumer func(v int64)

// ToMustInt64Consumer transforms Int64Consumer into MustInt64Consumer
func (c Int64Consumer) ToMustInt64Consumer() MustInt64Consumer {
	return func(v int64) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustInt64Consumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustInt64Consumer) AndThen(after MustInt64Consumer) MustInt64Consumer {
	if after != nil {
		return func(v int64) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToInt64Consumer transforms MustInt64Consumer into Int64Consumer
func (mc MustInt64Consumer) ToInt64Consumer() Int64Consumer {
	return func(v int64) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentInt64Consumer transforms MustInt64Consumer into SilentInt64Consumer
func (mc MustInt64Consumer) ToSilentInt64Consumer() SilentInt64Consumer {
	c := mc.ToInt64Consumer()
	return c.ToSilentInt64Consumer()
}
