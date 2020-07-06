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
	valTestInt8SliceConsumer []int8
	errTestInt8SliceConsumer = errors.New("error")
)

type testInt8SliceConsumerFactory func(t *testing.T) Int8SliceConsumer

func TestInt8SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					require.Equal(t, valTestInt8SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					require.Equal(t, valTestInt8SliceConsumer, v)
					return errTestInt8SliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestInt8SliceConsumer)
			if err != nil {
				r.EqualError(err, errTestInt8SliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt8SliceConsumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					require.Equal(t, valTestInt8SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					require.Equal(t, valTestInt8SliceConsumer, v)
					return errTestInt8SliceConsumer
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

			err := c(valTestInt8SliceConsumer)
			if err != nil {
				r.EqualError(err, errTestInt8SliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt8SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt8SliceConsumerFactory
		cf2   testInt8SliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, valTestInt8SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, valTestInt8SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, valTestInt8SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestInt8SliceConsumer
				}
			},
			cf2: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, valTestInt8SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, valTestInt8SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8SliceConsumer {
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
			err := cc(valTestInt8SliceConsumer)
			if err != nil {
				r.EqualError(err, errTestInt8SliceConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestInt8SliceConsumer_ToSilentInt8SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					require.Equal(t, valTestInt8SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					require.Equal(t, valTestInt8SliceConsumer, v)
					return errTestInt8SliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentInt8SliceConsumer()
			r.NotNil(sc)

			sc(valTestInt8SliceConsumer)
		})
	}
}

func TestInt8SliceConsumer_ToMustInt8SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8SliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					require.Equal(t, valTestInt8SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					require.Equal(t, valTestInt8SliceConsumer, v)
					return errTestInt8SliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustInt8SliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestInt8SliceConsumer.Error(), func() {
					mc(valTestInt8SliceConsumer)
				})
			} else {
				mc(valTestInt8SliceConsumer)
			}
		})
	}
}

func TestSilentInt8SliceConsumer(t *testing.T) {
	var sc SilentInt8SliceConsumer = func(v []int8) {
		require.Equal(t, valTestInt8SliceConsumer, v)
		return
	}
	sc(valTestInt8SliceConsumer)
}

func TestSilentInt8SliceConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					require.Equal(t, valTestInt8SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					require.Equal(t, valTestInt8SliceConsumer, v)
					return errTestInt8SliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentInt8SliceConsumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestInt8SliceConsumer)
		})
	}
}

func TestSilentInt8SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt8SliceConsumerFactory
		cf2   testInt8SliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, valTestInt8SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, valTestInt8SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, valTestInt8SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestInt8SliceConsumer
				}
			},
			cf2: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, valTestInt8SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, valTestInt8SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8SliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentInt8SliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentInt8SliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentInt8SliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestInt8SliceConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt8SliceConsumer(t *testing.T) {
	var mc MustInt8SliceConsumer = func(v []int8) {
		require.Equal(t, valTestInt8SliceConsumer, v)
		return
	}
	mc(valTestInt8SliceConsumer)
}

func TestMustInt8SliceConsumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8SliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					require.Equal(t, valTestInt8SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					require.Equal(t, valTestInt8SliceConsumer, v)
					return errTestInt8SliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustInt8SliceConsumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestInt8SliceConsumer.Error(), func() {
					mc(valTestInt8SliceConsumer)
				})
			} else {
				mc(valTestInt8SliceConsumer)
			}
		})
	}
}

func TestMustInt8SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt8SliceConsumerFactory
		cf2   testInt8SliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, valTestInt8SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, valTestInt8SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, valTestInt8SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestInt8SliceConsumer
				}
			},
			cf2: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, valTestInt8SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, valTestInt8SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8SliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustInt8SliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustInt8SliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustInt8SliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(errTestInt8SliceConsumer.Error(), func() {
					cmc(valTestInt8SliceConsumer)
				})
			} else {
				cmc(valTestInt8SliceConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt8SliceConsumer_ToSilentInt8SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					require.Equal(t, valTestInt8SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					require.Equal(t, valTestInt8SliceConsumer, v)
					return errTestInt8SliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt8SliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentInt8SliceConsumer()
			r.NotNil(sc)

			sc(valTestInt8SliceConsumer)
		})
	}
}

func TestMustInt8SliceConsumer_ToInt8SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8SliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					require.Equal(t, valTestInt8SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					require.Equal(t, valTestInt8SliceConsumer, v)
					return errTestInt8SliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt8SliceConsumer()
			r.NotNil(mc)

			c = mc.ToInt8SliceConsumer()
			r.NotNil(c)

			err := c(valTestInt8SliceConsumer)
			if tt.err {
				r.EqualError(err, errTestInt8SliceConsumer.Error())
			}
		})
	}
}
