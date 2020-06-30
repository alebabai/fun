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
	valTestInt64Consumer int64
	errTestInt64Consumer = errors.New("error")
)

type testInt64ConsumerFactory func(t *testing.T) Int64Consumer

func TestInt64Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					require.Equal(t, valTestInt64Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					require.Equal(t, valTestInt64Consumer, v)
					return errTestInt64Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestInt64Consumer)
			if err != nil {
				r.EqualError(err, errTestInt64Consumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt64Consumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					require.Equal(t, valTestInt64Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					require.Equal(t, valTestInt64Consumer, v)
					return errTestInt64Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			c := tc.ToConsumer()
			r.NotNil(c)

			err := c(valTestInt64Consumer)
			if err != nil {
				r.EqualError(err, errTestInt64Consumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt64Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt64ConsumerFactory
		cf2   testInt64ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, valTestInt64Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, valTestInt64Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, valTestInt64Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestInt64Consumer
				}
			},
			cf2: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, valTestInt64Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, valTestInt64Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64Consumer {
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
			err := cc(valTestInt64Consumer)
			if err != nil {
				r.EqualError(err, errTestInt64Consumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestInt64Consumer_ToSilentInt64Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					require.Equal(t, valTestInt64Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					require.Equal(t, valTestInt64Consumer, v)
					return errTestInt64Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentInt64Consumer()
			r.NotNil(sc)

			sc(valTestInt64Consumer)
		})
	}
}

func TestInt64Consumer_ToMustInt64Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					require.Equal(t, valTestInt64Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					require.Equal(t, valTestInt64Consumer, v)
					return errTestInt64Consumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustInt64Consumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestInt64Consumer.Error(), func() {
					mc(valTestInt64Consumer)
				})
			} else {
				mc(valTestInt64Consumer)
			}
		})
	}
}

func TestSilentInt64Consumer(t *testing.T) {
	var sc SilentInt64Consumer = func(v int64) {
		require.Equal(t, valTestInt64Consumer, v)
		return
	}
	sc(valTestInt64Consumer)
}

func TestSilentInt64Consumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					require.Equal(t, valTestInt64Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					require.Equal(t, valTestInt64Consumer, v)
					return errTestInt64Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentInt64Consumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestInt64Consumer)
		})
	}
}

func TestSilentInt64Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt64ConsumerFactory
		cf2   testInt64ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, valTestInt64Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, valTestInt64Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, valTestInt64Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestInt64Consumer
				}
			},
			cf2: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, valTestInt64Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, valTestInt64Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentInt64Consumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentInt64Consumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentInt64Consumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestInt64Consumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt64Consumer(t *testing.T) {
	var mc MustInt64Consumer = func(v int64) {
		require.Equal(t, valTestInt64Consumer, v)
		return
	}
	mc(valTestInt64Consumer)
}

func TestMustInt64Consumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					require.Equal(t, valTestInt64Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					require.Equal(t, valTestInt64Consumer, v)
					return errTestInt64Consumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustInt64Consumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestInt64Consumer.Error(), func() {
					mc(valTestInt64Consumer)
				})
			} else {
				mc(valTestInt64Consumer)
			}
		})
	}
}

func TestMustInt64Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt64ConsumerFactory
		cf2   testInt64ConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, valTestInt64Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, valTestInt64Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, valTestInt64Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestInt64Consumer
				}
			},
			cf2: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, valTestInt64Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, valTestInt64Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustInt64Consumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustInt64Consumer = nil
			if c2 != nil {
				mc2 = c2.ToMustInt64Consumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(errTestInt64Consumer.Error(), func() {
					cmc(valTestInt64Consumer)
				})
			} else {
				cmc(valTestInt64Consumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt64Consumer_ToSilentInt64Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					require.Equal(t, valTestInt64Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					require.Equal(t, valTestInt64Consumer, v)
					return errTestInt64Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt64Consumer()
			r.NotNil(mc)

			sc := mc.ToSilentInt64Consumer()
			r.NotNil(sc)

			sc(valTestInt64Consumer)
		})
	}
}

func TestMustInt64Consumer_ToInt64Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					require.Equal(t, valTestInt64Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					require.Equal(t, valTestInt64Consumer, v)
					return errTestInt64Consumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt64Consumer()
			r.NotNil(mc)

			c = mc.ToInt64Consumer()
			r.NotNil(c)

			err := c(valTestInt64Consumer)
			if tt.err {
				r.EqualError(err, errTestInt64Consumer.Error())
			}
		})
	}
}
