// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer.go.tmpl
// DO NOT EDIT

package consumer

// Float32Consumer represents an operation that accepts a single input argument or returns an error.
type Float32Consumer func(v float32) error

// ToConsumer transforms Float32Consumer into Consumer
func (c Float32Consumer) ToConsumer() Consumer {
	return func(v interface{}) error {
		return c(v.(float32))
	}
}

// AndThen returns a composed Float32Consumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Float32Consumer) AndThen(after Float32Consumer) Float32Consumer {
	if after != nil {
		return func(v float32) error {
			err := c(v)
			if err != nil {
				return err
			}
			return after(v)
		}
	}
	return c
}

// SilentFloat32Consumer represents an operation that accepts a single input argument without returning an error.
type SilentFloat32Consumer func(v float32)

// ToSilentFloat32Consumer transforms Float32Consumer into SilentFloat32Consumer
func (c Float32Consumer) ToSilentFloat32Consumer() SilentFloat32Consumer {
	return func(v float32) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentFloat32Consumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentFloat32Consumer) AndThen(after SilentFloat32Consumer) SilentFloat32Consumer {
	if after != nil {
		return func(v float32) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustFloat32Consumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustFloat32Consumer func(v float32)

// ToMustFloat32Consumer transforms Float32Consumer into MustFloat32Consumer
func (c Float32Consumer) ToMustFloat32Consumer() MustFloat32Consumer {
	return func(v float32) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustFloat32Consumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustFloat32Consumer) AndThen(after MustFloat32Consumer) MustFloat32Consumer {
	if after != nil {
		return func(v float32) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToFloat32Consumer transforms MustFloat32Consumer into Float32Consumer
func (mc MustFloat32Consumer) ToFloat32Consumer() Float32Consumer {
	return func(v float32) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// SilentFloat32Consumer transforms MustFloat32Consumer into SilentFloat32Consumer
func (mc MustFloat32Consumer) ToSilentFloat32Consumer() SilentFloat32Consumer {
	c := mc.ToFloat32Consumer()
	return c.ToSilentFloat32Consumer()
}
