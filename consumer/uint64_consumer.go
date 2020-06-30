// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Uint64Consumer represents an operation that accepts a single input argument or returns an error.
type Uint64Consumer func(v uint64) error

// ToConsumer transforms Uint64Consumer into Consumer
func (c Uint64Consumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(uint64))
	}
}

// AndThen returns a composed Uint64Consumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Uint64Consumer) AndThen(after Uint64Consumer) Uint64Consumer {
	if after != nil {
		return func(v uint64) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentUint64Consumer represents an operation that accepts a single input argument without returning an error.
type SilentUint64Consumer func(v uint64)

// ToSilentUint64Consumer transforms Uint64Consumer into SilentUint64Consumer
func (c Uint64Consumer) ToSilentUint64Consumer() SilentUint64Consumer {
	return func(v uint64) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentUint64Consumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentUint64Consumer) AndThen(after SilentUint64Consumer) SilentUint64Consumer {
	if after != nil {
		return func(v uint64) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustUint64Consumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustUint64Consumer func(v uint64)

// ToMustUint64Consumer transforms Uint64Consumer into MustUint64Consumer
func (c Uint64Consumer) ToMustUint64Consumer() MustUint64Consumer {
	return func(v uint64) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustUint64Consumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustUint64Consumer) AndThen(after MustUint64Consumer) MustUint64Consumer {
	if after != nil {
		return func(v uint64) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToUint64Consumer transforms MustUint64Consumer into Uint64Consumer
func (mc MustUint64Consumer) ToUint64Consumer() Uint64Consumer {
	return func(v uint64) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentUint64Consumer transforms MustUint64Consumer into SilentUint64Consumer
func (mc MustUint64Consumer) ToSilentUint64Consumer() SilentUint64Consumer {
	c := mc.ToUint64Consumer()
	return c.ToSilentUint64Consumer()
}
