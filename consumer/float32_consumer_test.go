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
	valTestFloat32Consumer float32
	errTestFloat32Consumer = errors.New("error")
)

type testFloat32ConsumerFactory func(t *testing.T) Float32Consumer

func TestFloat32Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					require.Equal(t, valTestFloat32Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					require.Equal(t, valTestFloat32Consumer, v)
					return errTestFloat32Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestFloat32Consumer)
			if err != nil {
				r.EqualError(err, errTestFloat32Consumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestFloat32Supplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					require.Equal(t, valTestFloat32Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					require.Equal(t, valTestFloat32Consumer, v)
					return errTestFloat32Consumer
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

			err := c(valTestFloat32Consumer)
			if err != nil {
				r.EqualError(err, errTestFloat32Consumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestFloat32Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testFloat32ConsumerFactory
		cf2   testFloat32ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, valTestFloat32Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, valTestFloat32Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, valTestFloat32Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestFloat32Consumer
				}
			},
			cf2: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, valTestFloat32Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, valTestFloat32Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32Consumer {
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
			err := cc(valTestFloat32Consumer)
			if err != nil {
				r.EqualError(err, errTestFloat32Consumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestFloat32Consumer_ToSilentFloat32Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					require.Equal(t, valTestFloat32Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					require.Equal(t, valTestFloat32Consumer, v)
					return errTestFloat32Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentFloat32Consumer()
			r.NotNil(sc)

			sc(valTestFloat32Consumer)
		})
	}
}

func TestFloat32Consumer_ToMustFloat32Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					require.Equal(t, valTestFloat32Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					require.Equal(t, valTestFloat32Consumer, v)
					return errTestFloat32Consumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustFloat32Consumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestFloat32Consumer.Error(), func() {
					mc(valTestFloat32Consumer)
				})
			} else {
				mc(valTestFloat32Consumer)
			}
		})
	}
}

func TestSilentFloat32Consumer(t *testing.T) {
	var sc SilentFloat32Consumer = func(v float32) {
		require.Equal(t, valTestFloat32Consumer, v)
		return
	}
	sc(valTestFloat32Consumer)
}

func TestSilentFloat32Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testFloat32ConsumerFactory
		cf2   testFloat32ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, valTestFloat32Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, valTestFloat32Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, valTestFloat32Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestFloat32Consumer
				}
			},
			cf2: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, valTestFloat32Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, valTestFloat32Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentFloat32Consumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentFloat32Consumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentFloat32Consumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestFloat32Consumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustFloat32Consumer(t *testing.T) {
	var sc SilentFloat32Consumer = func(v float32) {
		require.Equal(t, valTestFloat32Consumer, v)
		return
	}
	sc(valTestFloat32Consumer)
}

func TestMustFloat32Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testFloat32ConsumerFactory
		cf2   testFloat32ConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, valTestFloat32Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, valTestFloat32Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, valTestFloat32Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestFloat32Consumer
				}
			},
			cf2: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, valTestFloat32Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, valTestFloat32Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustFloat32Consumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustFloat32Consumer = nil
			if c2 != nil {
				mc2 = c2.ToMustFloat32Consumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(errTestFloat32Consumer.Error(), func() {
					cmc(valTestFloat32Consumer)
				})
			} else {
				cmc(valTestFloat32Consumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustFloat32Consumer_ToSilentFloat32Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					require.Equal(t, valTestFloat32Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					require.Equal(t, valTestFloat32Consumer, v)
					return errTestFloat32Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustFloat32Consumer()
			r.NotNil(mc)

			sc := mc.ToSilentFloat32Consumer()
			r.NotNil(sc)

			sc(valTestFloat32Consumer)
		})
	}
}

func TestMustFloat32Consumer_ToFloat32Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					require.Equal(t, valTestFloat32Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					require.Equal(t, valTestFloat32Consumer, v)
					return errTestFloat32Consumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustFloat32Consumer()
			r.NotNil(mc)

			c = mc.ToFloat32Consumer()
			r.NotNil(c)

			err := c(valTestFloat32Consumer)
			if tt.err {
				r.EqualError(err, errTestFloat32Consumer.Error())
			}
		})
	}
}
