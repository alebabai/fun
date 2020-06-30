// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Uint16Consumer represents an operation that accepts a single input argument or returns an error.
type Uint16Consumer func(v uint16) error

// ToConsumer transforms Uint16Consumer into Consumer
func (c Uint16Consumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(uint16))
	}
}

// AndThen returns a composed Uint16Consumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Uint16Consumer) AndThen(after Uint16Consumer) Uint16Consumer {
	if after != nil {
		return func(v uint16) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentUint16Consumer represents an operation that accepts a single input argument without returning an error.
type SilentUint16Consumer func(v uint16)

// ToSilentUint16Consumer transforms Uint16Consumer into SilentUint16Consumer
func (c Uint16Consumer) ToSilentUint16Consumer() SilentUint16Consumer {
	return func(v uint16) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentUint16Consumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentUint16Consumer) AndThen(after SilentUint16Consumer) SilentUint16Consumer {
	if after != nil {
		return func(v uint16) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustUint16Consumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustUint16Consumer func(v uint16)

// ToMustUint16Consumer transforms Uint16Consumer into MustUint16Consumer
func (c Uint16Consumer) ToMustUint16Consumer() MustUint16Consumer {
	return func(v uint16) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustUint16Consumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustUint16Consumer) AndThen(after MustUint16Consumer) MustUint16Consumer {
	if after != nil {
		return func(v uint16) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToUint16Consumer transforms MustUint16Consumer into Uint16Consumer
func (mc MustUint16Consumer) ToUint16Consumer() Uint16Consumer {
	return func(v uint16) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentUint16Consumer transforms MustUint16Consumer into SilentUint16Consumer
func (mc MustUint16Consumer) ToSilentUint16Consumer() SilentUint16Consumer {
	c := mc.ToUint16Consumer()
	return c.ToSilentUint16Consumer()
}
