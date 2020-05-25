package consumer

// Consumer represents an operation that accepts a single input argument or returns an error.
type Consumer func(v interface{}) error

// AndThen returns a composed Consumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (c Consumer) AndThen(after Consumer) Consumer {
	if after != nil {
		return func(v interface{}) error {
			err := c(v)
			if err != nil {
				return err
			}
			err = after(v)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return c
}

// SilentConsumer represents an operation that accepts a single input argument without returning an error.
type SilentConsumer func(v interface{})

// ToSilentConsumer transforms Consumer into SilentConsumer
func (c Consumer) ToSilentConsumer() SilentConsumer {
	return func(v interface{}) {
		_ = c(v)
		return
	}
}

// AndThen returns a composed SilentConsumer that performs, in sequence, this operation followed by the after operation.
// If after is nil, it returns original consumer.
func (sc SilentConsumer) AndThen(after SilentConsumer) SilentConsumer {
	if after != nil {
		return func(v interface{}) {
			sc(v)
			after(v)
		}
	}
	return sc
}

// MustConsumer represents an operation that accepts a single input argument without returning an error.
// In case of an error it should panic with error value.
type MustConsumer func(v interface{})

// ToMustConsumer transforms Consumer into MustConsumer
func (c Consumer) ToMustConsumer() MustConsumer {
	return func(v interface{}) {
		err := c(v)
		if err != nil {
			panic(err)
		}
		return
	}
}

// AndThen returns a composed MustConsumer that performs, in sequence, this operation followed by the after operation.
// If performing this operation returns an error, the after operation will not be performed.
// If after is nil, it returns original consumer.
func (mc MustConsumer) AndThen(after MustConsumer) MustConsumer {
	if after != nil {
		return func(v interface{}) {
			mc(v)
			after(v)
		}
	}
	return mc
}

// ToConsumer transforms MustConsumer into Consumer
func (mc MustConsumer) ToConsumer() Consumer {
	return func(v interface{}) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()
		mc(v)
		return
	}
}

// SilentConsumer transforms MustConsumer into SilentConsumer
func (mc MustConsumer) ToSilentConsumer() SilentConsumer {
	c := mc.ToConsumer()
	return c.ToSilentConsumer()
}
