// CODE GENERATED AUTOMATICALLY
// SOURCE: consumer_test.go.tmpl
// DO NOT EDIT

package consumer

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	valTestConsumer interface{}
	errTestConsumer = errors.New("error")
)

type testConsumerFactory func(t *testing.T) Consumer

func TestConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					require.Equal(t, valTestConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					require.Equal(t, valTestConsumer, v)
					return errTestConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestConsumer)
			if err != nil {
				r.EqualError(err, errTestConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testConsumerFactory
		cf2   testConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					calls++
					require.Equal(t, valTestConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					calls++
					require.Equal(t, valTestConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					calls++
					require.Equal(t, valTestConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestConsumer
				}
			},
			cf2: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					calls++
					require.Equal(t, valTestConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					calls++
					require.Equal(t, valTestConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			c2 := tt.cf2(t)

			cc := c1.AndThen(c2)
			r.NotNil(cc)

			calls = 0
			err := cc(valTestConsumer)
			if err != nil {
				r.EqualError(err, errTestConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					require.Equal(t, valTestConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					require.Equal(t, valTestConsumer, v)
					return errTestConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestConsumer)
		})
	}
}

func TestConsumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					require.Equal(t, valTestConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					require.Equal(t, valTestConsumer, v)
					return errTestConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestConsumer.Error(), func() {
					mc(valTestConsumer)
				})
			} else {
				mc(valTestConsumer)
			}
		})
	}
}

func TestSilentConsumer(t *testing.T) {
	var sc SilentConsumer = func(v interface{}) {
		require.Equal(t, valTestConsumer, v)
		return
	}
	sc(valTestConsumer)
}

func TestSilentConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testConsumerFactory
		cf2   testConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					calls++
					require.Equal(t, valTestConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					calls++
					require.Equal(t, valTestConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					calls++
					require.Equal(t, valTestConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestConsumer
				}
			},
			cf2: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					calls++
					require.Equal(t, valTestConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					calls++
					require.Equal(t, valTestConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustConsumer(t *testing.T) {
	var mc MustConsumer = func(v interface{}) {
		require.Equal(t, valTestConsumer, v)
		return
	}
	mc(valTestConsumer)
}

func TestMustConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testConsumerFactory
		cf2   testConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					calls++
					require.Equal(t, valTestConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					calls++
					require.Equal(t, valTestConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					calls++
					require.Equal(t, valTestConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestConsumer
				}
			},
			cf2: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					calls++
					require.Equal(t, valTestConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					calls++
					require.Equal(t, valTestConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(errTestConsumer.Error(), func() {
					cmc(valTestConsumer)
				})
			} else {
				cmc(valTestConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					require.Equal(t, valTestConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					require.Equal(t, valTestConsumer, v)
					return errTestConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestConsumer)
		})
	}
}

func TestMustConsumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					require.Equal(t, valTestConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					require.Equal(t, valTestConsumer, v)
					return errTestConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustConsumer()
			r.NotNil(mc)

			c = mc.ToConsumer()
			r.NotNil(c)

			err := c(valTestConsumer)
			if tt.err {
				r.EqualError(err, errTestConsumer.Error())
			}
		})
	}
}
