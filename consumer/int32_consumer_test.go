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
	valTestInt32Consumer int32
	errTestInt32Consumer = errors.New("error")
)

type testInt32ConsumerFactory func(t *testing.T) Int32Consumer

func TestInt32Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					require.Equal(t, valTestInt32Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					require.Equal(t, valTestInt32Consumer, v)
					return errTestInt32Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestInt32Consumer)
			if err != nil {
				r.EqualError(err, errTestInt32Consumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt32Consumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					require.Equal(t, valTestInt32Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					require.Equal(t, valTestInt32Consumer, v)
					return errTestInt32Consumer
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

			err := c(valTestInt32Consumer)
			if err != nil {
				r.EqualError(err, errTestInt32Consumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt32Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt32ConsumerFactory
		cf2   testInt32ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, valTestInt32Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, valTestInt32Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, valTestInt32Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestInt32Consumer
				}
			},
			cf2: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, valTestInt32Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, valTestInt32Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32Consumer {
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
			err := cc(valTestInt32Consumer)
			if err != nil {
				r.EqualError(err, errTestInt32Consumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestInt32Consumer_ToSilentInt32Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					require.Equal(t, valTestInt32Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					require.Equal(t, valTestInt32Consumer, v)
					return errTestInt32Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentInt32Consumer()
			r.NotNil(sc)

			sc(valTestInt32Consumer)
		})
	}
}

func TestInt32Consumer_ToMustInt32Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					require.Equal(t, valTestInt32Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					require.Equal(t, valTestInt32Consumer, v)
					return errTestInt32Consumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustInt32Consumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestInt32Consumer.Error(), func() {
					mc(valTestInt32Consumer)
				})
			} else {
				mc(valTestInt32Consumer)
			}
		})
	}
}

func TestSilentInt32Consumer(t *testing.T) {
	var sc SilentInt32Consumer = func(v int32) {
		require.Equal(t, valTestInt32Consumer, v)
		return
	}
	sc(valTestInt32Consumer)
}

func TestSilentInt32Consumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					require.Equal(t, valTestInt32Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					require.Equal(t, valTestInt32Consumer, v)
					return errTestInt32Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentInt32Consumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestInt32Consumer)
		})
	}
}

func TestSilentInt32Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt32ConsumerFactory
		cf2   testInt32ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, valTestInt32Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, valTestInt32Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, valTestInt32Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestInt32Consumer
				}
			},
			cf2: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, valTestInt32Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, valTestInt32Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentInt32Consumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentInt32Consumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentInt32Consumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestInt32Consumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt32Consumer(t *testing.T) {
	var mc MustInt32Consumer = func(v int32) {
		require.Equal(t, valTestInt32Consumer, v)
		return
	}
	mc(valTestInt32Consumer)
}

func TestMustInt32Consumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					require.Equal(t, valTestInt32Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					require.Equal(t, valTestInt32Consumer, v)
					return errTestInt32Consumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustInt32Consumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestInt32Consumer.Error(), func() {
					mc(valTestInt32Consumer)
				})
			} else {
				mc(valTestInt32Consumer)
			}
		})
	}
}

func TestMustInt32Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt32ConsumerFactory
		cf2   testInt32ConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, valTestInt32Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, valTestInt32Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, valTestInt32Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestInt32Consumer
				}
			},
			cf2: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, valTestInt32Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, valTestInt32Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustInt32Consumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustInt32Consumer = nil
			if c2 != nil {
				mc2 = c2.ToMustInt32Consumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(errTestInt32Consumer.Error(), func() {
					cmc(valTestInt32Consumer)
				})
			} else {
				cmc(valTestInt32Consumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt32Consumer_ToSilentInt32Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					require.Equal(t, valTestInt32Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					require.Equal(t, valTestInt32Consumer, v)
					return errTestInt32Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt32Consumer()
			r.NotNil(mc)

			sc := mc.ToSilentInt32Consumer()
			r.NotNil(sc)

			sc(valTestInt32Consumer)
		})
	}
}

func TestMustInt32Consumer_ToInt32Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					require.Equal(t, valTestInt32Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					require.Equal(t, valTestInt32Consumer, v)
					return errTestInt32Consumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt32Consumer()
			r.NotNil(mc)

			c = mc.ToInt32Consumer()
			r.NotNil(c)

			err := c(valTestInt32Consumer)
			if tt.err {
				r.EqualError(err, errTestInt32Consumer.Error())
			}
		})
	}
}
