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
	valTestUint32Consumer uint32
	errTestUint32Consumer = errors.New("error")
)

type testUint32ConsumerFactory func(t *testing.T) Uint32Consumer

func TestUint32Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint32ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32Consumer {
				return func(v uint32) error {
					require.Equal(t, valTestUint32Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint32Consumer {
				return func(v uint32) error {
					require.Equal(t, valTestUint32Consumer, v)
					return errTestUint32Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestUint32Consumer)
			if err != nil {
				r.EqualError(err, errTestUint32Consumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint32Supplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint32ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32Consumer {
				return func(v uint32) error {
					require.Equal(t, valTestUint32Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint32Consumer {
				return func(v uint32) error {
					require.Equal(t, valTestUint32Consumer, v)
					return errTestUint32Consumer
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

			err := c(valTestUint32Consumer)
			if err != nil {
				r.EqualError(err, errTestUint32Consumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint32Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint32ConsumerFactory
		cf2   testUint32ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint32Consumer {
				return func(v uint32) error {
					calls++
					require.Equal(t, valTestUint32Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint32Consumer {
				return func(v uint32) error {
					calls++
					require.Equal(t, valTestUint32Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint32Consumer {
				return func(v uint32) error {
					calls++
					require.Equal(t, valTestUint32Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint32Consumer
				}
			},
			cf2: func(t *testing.T) Uint32Consumer {
				return func(v uint32) error {
					calls++
					require.Equal(t, valTestUint32Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint32Consumer {
				return func(v uint32) error {
					calls++
					require.Equal(t, valTestUint32Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint32Consumer {
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
			err := cc(valTestUint32Consumer)
			if err != nil {
				r.EqualError(err, errTestUint32Consumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestUint32Consumer_ToSilentUint32Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint32ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32Consumer {
				return func(v uint32) error {
					require.Equal(t, valTestUint32Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint32Consumer {
				return func(v uint32) error {
					require.Equal(t, valTestUint32Consumer, v)
					return errTestUint32Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentUint32Consumer()
			r.NotNil(sc)

			sc(valTestUint32Consumer)
		})
	}
}

func TestUint32Consumer_ToMustUint32Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint32ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32Consumer {
				return func(v uint32) error {
					require.Equal(t, valTestUint32Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint32Consumer {
				return func(v uint32) error {
					require.Equal(t, valTestUint32Consumer, v)
					return errTestUint32Consumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustUint32Consumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestUint32Consumer.Error(), func() {
					mc(valTestUint32Consumer)
				})
			} else {
				mc(valTestUint32Consumer)
			}
		})
	}
}

func TestSilentUint32Consumer(t *testing.T) {
	var sc SilentUint32Consumer = func(v uint32) {
		require.Equal(t, valTestUint32Consumer, v)
		return
	}
	sc(valTestUint32Consumer)
}

func TestSilentUint32Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint32ConsumerFactory
		cf2   testUint32ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint32Consumer {
				return func(v uint32) error {
					calls++
					require.Equal(t, valTestUint32Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint32Consumer {
				return func(v uint32) error {
					calls++
					require.Equal(t, valTestUint32Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint32Consumer {
				return func(v uint32) error {
					calls++
					require.Equal(t, valTestUint32Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint32Consumer
				}
			},
			cf2: func(t *testing.T) Uint32Consumer {
				return func(v uint32) error {
					calls++
					require.Equal(t, valTestUint32Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint32Consumer {
				return func(v uint32) error {
					calls++
					require.Equal(t, valTestUint32Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint32Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentUint32Consumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentUint32Consumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentUint32Consumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestUint32Consumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint32Consumer(t *testing.T) {
	var sc SilentUint32Consumer = func(v uint32) {
		require.Equal(t, valTestUint32Consumer, v)
		return
	}
	sc(valTestUint32Consumer)
}

func TestMustUint32Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint32ConsumerFactory
		cf2   testUint32ConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint32Consumer {
				return func(v uint32) error {
					calls++
					require.Equal(t, valTestUint32Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint32Consumer {
				return func(v uint32) error {
					calls++
					require.Equal(t, valTestUint32Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint32Consumer {
				return func(v uint32) error {
					calls++
					require.Equal(t, valTestUint32Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint32Consumer
				}
			},
			cf2: func(t *testing.T) Uint32Consumer {
				return func(v uint32) error {
					calls++
					require.Equal(t, valTestUint32Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint32Consumer {
				return func(v uint32) error {
					calls++
					require.Equal(t, valTestUint32Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint32Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustUint32Consumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustUint32Consumer = nil
			if c2 != nil {
				mc2 = c2.ToMustUint32Consumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(errTestUint32Consumer.Error(), func() {
					cmc(valTestUint32Consumer)
				})
			} else {
				cmc(valTestUint32Consumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint32Consumer_ToSilentUint32Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint32ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32Consumer {
				return func(v uint32) error {
					require.Equal(t, valTestUint32Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint32Consumer {
				return func(v uint32) error {
					require.Equal(t, valTestUint32Consumer, v)
					return errTestUint32Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint32Consumer()
			r.NotNil(mc)

			sc := mc.ToSilentUint32Consumer()
			r.NotNil(sc)

			sc(valTestUint32Consumer)
		})
	}
}

func TestMustUint32Consumer_ToUint32Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint32ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32Consumer {
				return func(v uint32) error {
					require.Equal(t, valTestUint32Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint32Consumer {
				return func(v uint32) error {
					require.Equal(t, valTestUint32Consumer, v)
					return errTestUint32Consumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint32Consumer()
			r.NotNil(mc)

			c = mc.ToUint32Consumer()
			r.NotNil(c)

			err := c(valTestUint32Consumer)
			if tt.err {
				r.EqualError(err, errTestUint32Consumer.Error())
			}
		})
	}
}
