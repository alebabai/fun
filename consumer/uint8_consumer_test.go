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
	valTestUint8Consumer uint8
	errTestUint8Consumer = errors.New("error")
)

type testUint8ConsumerFactory func(t *testing.T) Uint8Consumer

func TestUint8Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint8ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					require.Equal(t, valTestUint8Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					require.Equal(t, valTestUint8Consumer, v)
					return errTestUint8Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestUint8Consumer)
			if err != nil {
				r.EqualError(err, errTestUint8Consumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint8Consumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint8ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					require.Equal(t, valTestUint8Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					require.Equal(t, valTestUint8Consumer, v)
					return errTestUint8Consumer
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

			err := c(valTestUint8Consumer)
			if err != nil {
				r.EqualError(err, errTestUint8Consumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint8Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint8ConsumerFactory
		cf2   testUint8ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					calls++
					require.Equal(t, valTestUint8Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					calls++
					require.Equal(t, valTestUint8Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					calls++
					require.Equal(t, valTestUint8Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint8Consumer
				}
			},
			cf2: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					calls++
					require.Equal(t, valTestUint8Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					calls++
					require.Equal(t, valTestUint8Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint8Consumer {
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
			err := cc(valTestUint8Consumer)
			if err != nil {
				r.EqualError(err, errTestUint8Consumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestUint8Consumer_ToSilentUint8Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint8ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					require.Equal(t, valTestUint8Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					require.Equal(t, valTestUint8Consumer, v)
					return errTestUint8Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentUint8Consumer()
			r.NotNil(sc)

			sc(valTestUint8Consumer)
		})
	}
}

func TestUint8Consumer_ToMustUint8Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint8ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					require.Equal(t, valTestUint8Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					require.Equal(t, valTestUint8Consumer, v)
					return errTestUint8Consumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustUint8Consumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestUint8Consumer.Error(), func() {
					mc(valTestUint8Consumer)
				})
			} else {
				mc(valTestUint8Consumer)
			}
		})
	}
}

func TestSilentUint8Consumer(t *testing.T) {
	var sc SilentUint8Consumer = func(v uint8) {
		require.Equal(t, valTestUint8Consumer, v)
		return
	}
	sc(valTestUint8Consumer)
}

func TestSilentUint8Consumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint8ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					require.Equal(t, valTestUint8Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					require.Equal(t, valTestUint8Consumer, v)
					return errTestUint8Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentUint8Consumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestUint8Consumer)
		})
	}
}

func TestSilentUint8Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint8ConsumerFactory
		cf2   testUint8ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					calls++
					require.Equal(t, valTestUint8Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					calls++
					require.Equal(t, valTestUint8Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					calls++
					require.Equal(t, valTestUint8Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint8Consumer
				}
			},
			cf2: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					calls++
					require.Equal(t, valTestUint8Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					calls++
					require.Equal(t, valTestUint8Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint8Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentUint8Consumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentUint8Consumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentUint8Consumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestUint8Consumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint8Consumer(t *testing.T) {
	var mc MustUint8Consumer = func(v uint8) {
		require.Equal(t, valTestUint8Consumer, v)
		return
	}
	mc(valTestUint8Consumer)
}

func TestMustUint8Consumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint8ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					require.Equal(t, valTestUint8Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					require.Equal(t, valTestUint8Consumer, v)
					return errTestUint8Consumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustUint8Consumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestUint8Consumer.Error(), func() {
					mc(valTestUint8Consumer)
				})
			} else {
				mc(valTestUint8Consumer)
			}
		})
	}
}

func TestMustUint8Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint8ConsumerFactory
		cf2   testUint8ConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					calls++
					require.Equal(t, valTestUint8Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					calls++
					require.Equal(t, valTestUint8Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					calls++
					require.Equal(t, valTestUint8Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint8Consumer
				}
			},
			cf2: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					calls++
					require.Equal(t, valTestUint8Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					calls++
					require.Equal(t, valTestUint8Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint8Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustUint8Consumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustUint8Consumer = nil
			if c2 != nil {
				mc2 = c2.ToMustUint8Consumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(errTestUint8Consumer.Error(), func() {
					cmc(valTestUint8Consumer)
				})
			} else {
				cmc(valTestUint8Consumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint8Consumer_ToSilentUint8Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint8ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					require.Equal(t, valTestUint8Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					require.Equal(t, valTestUint8Consumer, v)
					return errTestUint8Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint8Consumer()
			r.NotNil(mc)

			sc := mc.ToSilentUint8Consumer()
			r.NotNil(sc)

			sc(valTestUint8Consumer)
		})
	}
}

func TestMustUint8Consumer_ToUint8Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint8ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					require.Equal(t, valTestUint8Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint8Consumer {
				return func(v uint8) error {
					require.Equal(t, valTestUint8Consumer, v)
					return errTestUint8Consumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint8Consumer()
			r.NotNil(mc)

			c = mc.ToUint8Consumer()
			r.NotNil(c)

			err := c(valTestUint8Consumer)
			if tt.err {
				r.EqualError(err, errTestUint8Consumer.Error())
			}
		})
	}
}
