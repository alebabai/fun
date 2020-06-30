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
	valTestInt8PtrSliceConsumer []*int8
	errTestInt8PtrSliceConsumer = errors.New("error")
)

type testInt8PtrSliceConsumerFactory func(t *testing.T) Int8PtrSliceConsumer

func TestInt8PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					require.Equal(t, valTestInt8PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					require.Equal(t, valTestInt8PtrSliceConsumer, v)
					return errTestInt8PtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestInt8PtrSliceConsumer)
			if err != nil {
				r.EqualError(err, errTestInt8PtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt8PtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					require.Equal(t, valTestInt8PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					require.Equal(t, valTestInt8PtrSliceConsumer, v)
					return errTestInt8PtrSliceConsumer
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

			err := c(valTestInt8PtrSliceConsumer)
			if err != nil {
				r.EqualError(err, errTestInt8PtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt8PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt8PtrSliceConsumerFactory
		cf2   testInt8PtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, valTestInt8PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, valTestInt8PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, valTestInt8PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestInt8PtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, valTestInt8PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, valTestInt8PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8PtrSliceConsumer {
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
			err := cc(valTestInt8PtrSliceConsumer)
			if err != nil {
				r.EqualError(err, errTestInt8PtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestInt8PtrSliceConsumer_ToSilentInt8PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					require.Equal(t, valTestInt8PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					require.Equal(t, valTestInt8PtrSliceConsumer, v)
					return errTestInt8PtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentInt8PtrSliceConsumer()
			r.NotNil(sc)

			sc(valTestInt8PtrSliceConsumer)
		})
	}
}

func TestInt8PtrSliceConsumer_ToMustInt8PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8PtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					require.Equal(t, valTestInt8PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					require.Equal(t, valTestInt8PtrSliceConsumer, v)
					return errTestInt8PtrSliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustInt8PtrSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestInt8PtrSliceConsumer.Error(), func() {
					mc(valTestInt8PtrSliceConsumer)
				})
			} else {
				mc(valTestInt8PtrSliceConsumer)
			}
		})
	}
}

func TestSilentInt8PtrSliceConsumer(t *testing.T) {
	var sc SilentInt8PtrSliceConsumer = func(v []*int8) {
		require.Equal(t, valTestInt8PtrSliceConsumer, v)
		return
	}
	sc(valTestInt8PtrSliceConsumer)
}

func TestSilentInt8PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt8PtrSliceConsumerFactory
		cf2   testInt8PtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, valTestInt8PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, valTestInt8PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, valTestInt8PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestInt8PtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, valTestInt8PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, valTestInt8PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8PtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentInt8PtrSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentInt8PtrSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentInt8PtrSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestInt8PtrSliceConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt8PtrSliceConsumer(t *testing.T) {
	var sc SilentInt8PtrSliceConsumer = func(v []*int8) {
		require.Equal(t, valTestInt8PtrSliceConsumer, v)
		return
	}
	sc(valTestInt8PtrSliceConsumer)
}

func TestMustInt8PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt8PtrSliceConsumerFactory
		cf2   testInt8PtrSliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, valTestInt8PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, valTestInt8PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, valTestInt8PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestInt8PtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, valTestInt8PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, valTestInt8PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8PtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustInt8PtrSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustInt8PtrSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustInt8PtrSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(errTestInt8PtrSliceConsumer.Error(), func() {
					cmc(valTestInt8PtrSliceConsumer)
				})
			} else {
				cmc(valTestInt8PtrSliceConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt8PtrSliceConsumer_ToSilentInt8PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					require.Equal(t, valTestInt8PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					require.Equal(t, valTestInt8PtrSliceConsumer, v)
					return errTestInt8PtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt8PtrSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentInt8PtrSliceConsumer()
			r.NotNil(sc)

			sc(valTestInt8PtrSliceConsumer)
		})
	}
}

func TestMustInt8PtrSliceConsumer_ToInt8PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8PtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					require.Equal(t, valTestInt8PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					require.Equal(t, valTestInt8PtrSliceConsumer, v)
					return errTestInt8PtrSliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt8PtrSliceConsumer()
			r.NotNil(mc)

			c = mc.ToInt8PtrSliceConsumer()
			r.NotNil(c)

			err := c(valTestInt8PtrSliceConsumer)
			if tt.err {
				r.EqualError(err, errTestInt8PtrSliceConsumer.Error())
			}
		})
	}
}
