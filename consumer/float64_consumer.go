// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Float64Consumer represents an operation that accepts a single input argument or returns an error.
type Float64Consumer func(v float64) error

// ToConsumer transforms Float64Consumer into Consumer
func (c Float64Consumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(float64))
	}
}

// AndThen returns a composed Float64Consumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Float64Consumer) AndThen(after Float64Consumer) Float64Consumer {
	if after != nil {
		return func(v float64) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentFloat64Consumer represents an operation that accepts a single input argument without returning an error.
type SilentFloat64Consumer func(v float64)

// ToSilentFloat64Consumer transforms Float64Consumer into SilentFloat64Consumer
func (c Float64Consumer) ToSilentFloat64Consumer() SilentFloat64Consumer {
	return func(v float64) {
		_ = c(v)
		return
	}
}

// ToSilentConsumer transforms SilentFloat64Consumer into SilentConsumer
func (sc SilentFloat64Consumer) ToSilentConsumer() SilentConsumer {
	return func(v interface{}) {
		sc(v.(float64))
	}
}

// AndThen returns a composed SilentFloat64Consumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentFloat64Consumer) AndThen(after SilentFloat64Consumer) SilentFloat64Consumer {
	if after != nil {
		return func(v float64) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustFloat64Consumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustFloat64Consumer func(v float64)

// ToMustFloat64Consumer transforms Float64Consumer into MustFloat64Consumer
func (c Float64Consumer) ToMustFloat64Consumer() MustFloat64Consumer {
	return func(v float64) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// ToMustConsumer transforms MustFloat64Consumer into MustConsumer
func (mc MustFloat64Consumer) ToMustConsumer() MustConsumer {
	return func(v interface{}) {
		mc(v.(float64))
	}
}

// AndThen returns a composed MustFloat64Consumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustFloat64Consumer) AndThen(after MustFloat64Consumer) MustFloat64Consumer {
	if after != nil {
		return func(v float64) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToFloat64Consumer transforms MustFloat64Consumer into Float64Consumer
func (mc MustFloat64Consumer) ToFloat64Consumer() Float64Consumer {
	return func(v float64) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentFloat64Consumer transforms MustFloat64Consumer into SilentFloat64Consumer
func (mc MustFloat64Consumer) ToSilentFloat64Consumer() SilentFloat64Consumer {
	c := mc.ToFloat64Consumer()
	return c.ToSilentFloat64Consumer()
}
