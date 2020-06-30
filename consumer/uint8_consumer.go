// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Uint8Consumer represents an operation that accepts a single input argument or returns an error.
type Uint8Consumer func(v uint8) error

// ToConsumer transforms Uint8Consumer into Consumer
func (c Uint8Consumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(uint8))
	}
}

// AndThen returns a composed Uint8Consumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Uint8Consumer) AndThen(after Uint8Consumer) Uint8Consumer {
	if after != nil {
		return func(v uint8) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentUint8Consumer represents an operation that accepts a single input argument without returning an error.
type SilentUint8Consumer func(v uint8)

// ToSilentUint8Consumer transforms Uint8Consumer into SilentUint8Consumer
func (c Uint8Consumer) ToSilentUint8Consumer() SilentUint8Consumer {
	return func(v uint8) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentUint8Consumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentUint8Consumer) AndThen(after SilentUint8Consumer) SilentUint8Consumer {
	if after != nil {
		return func(v uint8) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustUint8Consumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustUint8Consumer func(v uint8)

// ToMustUint8Consumer transforms Uint8Consumer into MustUint8Consumer
func (c Uint8Consumer) ToMustUint8Consumer() MustUint8Consumer {
	return func(v uint8) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustUint8Consumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustUint8Consumer) AndThen(after MustUint8Consumer) MustUint8Consumer {
	if after != nil {
		return func(v uint8) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToUint8Consumer transforms MustUint8Consumer into Uint8Consumer
func (mc MustUint8Consumer) ToUint8Consumer() Uint8Consumer {
	return func(v uint8) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentUint8Consumer transforms MustUint8Consumer into SilentUint8Consumer
func (mc MustUint8Consumer) ToSilentUint8Consumer() SilentUint8Consumer {
	c := mc.ToUint8Consumer()
	return c.ToSilentUint8Consumer()
}
