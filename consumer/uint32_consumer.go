// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Uint32Consumer represents an operation that accepts a single input argument or returns an error.
type Uint32Consumer func(v uint32) error

// ToConsumer transforms Uint32Consumer into Consumer
func (c Uint32Consumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(uint32))
	}
}

// AndThen returns a composed Uint32Consumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Uint32Consumer) AndThen(after Uint32Consumer) Uint32Consumer {
	if after != nil {
		return func(v uint32) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentUint32Consumer represents an operation that accepts a single input argument without returning an error.
type SilentUint32Consumer func(v uint32)

// ToSilentUint32Consumer transforms Uint32Consumer into SilentUint32Consumer
func (c Uint32Consumer) ToSilentUint32Consumer() SilentUint32Consumer {
	return func(v uint32) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentUint32Consumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentUint32Consumer) AndThen(after SilentUint32Consumer) SilentUint32Consumer {
	if after != nil {
		return func(v uint32) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustUint32Consumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustUint32Consumer func(v uint32)

// ToMustUint32Consumer transforms Uint32Consumer into MustUint32Consumer
func (c Uint32Consumer) ToMustUint32Consumer() MustUint32Consumer {
	return func(v uint32) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustUint32Consumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustUint32Consumer) AndThen(after MustUint32Consumer) MustUint32Consumer {
	if after != nil {
		return func(v uint32) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToUint32Consumer transforms MustUint32Consumer into Uint32Consumer
func (mc MustUint32Consumer) ToUint32Consumer() Uint32Consumer {
	return func(v uint32) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// ToSilentUint32Consumer transforms MustUint32Consumer into SilentUint32Consumer
func (mc MustUint32Consumer) ToSilentUint32Consumer() SilentUint32Consumer {
	c := mc.ToUint32Consumer()
	return c.ToSilentUint32Consumer()
}
