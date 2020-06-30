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
	valTestFloat64Consumer float64
	errTestFloat64Consumer = errors.New("error")
)

type testFloat64ConsumerFactory func(t *testing.T) Float64Consumer

func TestFloat64Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat64ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64Consumer {
				return func(v float64) error {
					require.Equal(t, valTestFloat64Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float64Consumer {
				return func(v float64) error {
					require.Equal(t, valTestFloat64Consumer, v)
					return errTestFloat64Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestFloat64Consumer)
			if err != nil {
				r.EqualError(err, errTestFloat64Consumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestFloat64Supplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat64ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64Consumer {
				return func(v float64) error {
					require.Equal(t, valTestFloat64Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float64Consumer {
				return func(v float64) error {
					require.Equal(t, valTestFloat64Consumer, v)
					return errTestFloat64Consumer
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

			err := c(valTestFloat64Consumer)
			if err != nil {
				r.EqualError(err, errTestFloat64Consumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestFloat64Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testFloat64ConsumerFactory
		cf2   testFloat64ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float64Consumer {
				return func(v float64) error {
					calls++
					require.Equal(t, valTestFloat64Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64Consumer {
				return func(v float64) error {
					calls++
					require.Equal(t, valTestFloat64Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Float64Consumer {
				return func(v float64) error {
					calls++
					require.Equal(t, valTestFloat64Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestFloat64Consumer
				}
			},
			cf2: func(t *testing.T) Float64Consumer {
				return func(v float64) error {
					calls++
					require.Equal(t, valTestFloat64Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float64Consumer {
				return func(v float64) error {
					calls++
					require.Equal(t, valTestFloat64Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64Consumer {
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
			err := cc(valTestFloat64Consumer)
			if err != nil {
				r.EqualError(err, errTestFloat64Consumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestFloat64Consumer_ToSilentFloat64Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat64ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64Consumer {
				return func(v float64) error {
					require.Equal(t, valTestFloat64Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float64Consumer {
				return func(v float64) error {
					require.Equal(t, valTestFloat64Consumer, v)
					return errTestFloat64Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentFloat64Consumer()
			r.NotNil(sc)

			sc(valTestFloat64Consumer)
		})
	}
}

func TestFloat64Consumer_ToMustFloat64Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat64ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64Consumer {
				return func(v float64) error {
					require.Equal(t, valTestFloat64Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float64Consumer {
				return func(v float64) error {
					require.Equal(t, valTestFloat64Consumer, v)
					return errTestFloat64Consumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustFloat64Consumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestFloat64Consumer.Error(), func() {
					mc(valTestFloat64Consumer)
				})
			} else {
				mc(valTestFloat64Consumer)
			}
		})
	}
}

func TestSilentFloat64Consumer(t *testing.T) {
	var sc SilentFloat64Consumer = func(v float64) {
		require.Equal(t, valTestFloat64Consumer, v)
		return
	}
	sc(valTestFloat64Consumer)
}

func TestSilentFloat64Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testFloat64ConsumerFactory
		cf2   testFloat64ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float64Consumer {
				return func(v float64) error {
					calls++
					require.Equal(t, valTestFloat64Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64Consumer {
				return func(v float64) error {
					calls++
					require.Equal(t, valTestFloat64Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Float64Consumer {
				return func(v float64) error {
					calls++
					require.Equal(t, valTestFloat64Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestFloat64Consumer
				}
			},
			cf2: func(t *testing.T) Float64Consumer {
				return func(v float64) error {
					calls++
					require.Equal(t, valTestFloat64Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float64Consumer {
				return func(v float64) error {
					calls++
					require.Equal(t, valTestFloat64Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentFloat64Consumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentFloat64Consumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentFloat64Consumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestFloat64Consumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustFloat64Consumer(t *testing.T) {
	var sc SilentFloat64Consumer = func(v float64) {
		require.Equal(t, valTestFloat64Consumer, v)
		return
	}
	sc(valTestFloat64Consumer)
}

func TestMustFloat64Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testFloat64ConsumerFactory
		cf2   testFloat64ConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float64Consumer {
				return func(v float64) error {
					calls++
					require.Equal(t, valTestFloat64Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64Consumer {
				return func(v float64) error {
					calls++
					require.Equal(t, valTestFloat64Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Float64Consumer {
				return func(v float64) error {
					calls++
					require.Equal(t, valTestFloat64Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestFloat64Consumer
				}
			},
			cf2: func(t *testing.T) Float64Consumer {
				return func(v float64) error {
					calls++
					require.Equal(t, valTestFloat64Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float64Consumer {
				return func(v float64) error {
					calls++
					require.Equal(t, valTestFloat64Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustFloat64Consumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustFloat64Consumer = nil
			if c2 != nil {
				mc2 = c2.ToMustFloat64Consumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(errTestFloat64Consumer.Error(), func() {
					cmc(valTestFloat64Consumer)
				})
			} else {
				cmc(valTestFloat64Consumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustFloat64Consumer_ToSilentFloat64Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat64ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64Consumer {
				return func(v float64) error {
					require.Equal(t, valTestFloat64Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float64Consumer {
				return func(v float64) error {
					require.Equal(t, valTestFloat64Consumer, v)
					return errTestFloat64Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustFloat64Consumer()
			r.NotNil(mc)

			sc := mc.ToSilentFloat64Consumer()
			r.NotNil(sc)

			sc(valTestFloat64Consumer)
		})
	}
}

func TestMustFloat64Consumer_ToFloat64Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat64ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64Consumer {
				return func(v float64) error {
					require.Equal(t, valTestFloat64Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float64Consumer {
				return func(v float64) error {
					require.Equal(t, valTestFloat64Consumer, v)
					return errTestFloat64Consumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustFloat64Consumer()
			r.NotNil(mc)

			c = mc.ToFloat64Consumer()
			r.NotNil(c)

			err := c(valTestFloat64Consumer)
			if tt.err {
				r.EqualError(err, errTestFloat64Consumer.Error())
			}
		})
	}
}
