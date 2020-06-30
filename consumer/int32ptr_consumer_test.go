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
	valTestInt32PtrConsumer *int32
	errTestInt32PtrConsumer = errors.New("error")
)

type testInt32PtrConsumerFactory func(t *testing.T) Int32PtrConsumer

func TestInt32PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					require.Equal(t, valTestInt32PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					require.Equal(t, valTestInt32PtrConsumer, v)
					return errTestInt32PtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestInt32PtrConsumer)
			if err != nil {
				r.EqualError(err, errTestInt32PtrConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt32PtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					require.Equal(t, valTestInt32PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					require.Equal(t, valTestInt32PtrConsumer, v)
					return errTestInt32PtrConsumer
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

			err := c(valTestInt32PtrConsumer)
			if err != nil {
				r.EqualError(err, errTestInt32PtrConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt32PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt32PtrConsumerFactory
		cf2   testInt32PtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, valTestInt32PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, valTestInt32PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, valTestInt32PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestInt32PtrConsumer
				}
			},
			cf2: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, valTestInt32PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, valTestInt32PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32PtrConsumer {
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
			err := cc(valTestInt32PtrConsumer)
			if err != nil {
				r.EqualError(err, errTestInt32PtrConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestInt32PtrConsumer_ToSilentInt32PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					require.Equal(t, valTestInt32PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					require.Equal(t, valTestInt32PtrConsumer, v)
					return errTestInt32PtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentInt32PtrConsumer()
			r.NotNil(sc)

			sc(valTestInt32PtrConsumer)
		})
	}
}

func TestInt32PtrConsumer_ToMustInt32PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32PtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					require.Equal(t, valTestInt32PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					require.Equal(t, valTestInt32PtrConsumer, v)
					return errTestInt32PtrConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustInt32PtrConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestInt32PtrConsumer.Error(), func() {
					mc(valTestInt32PtrConsumer)
				})
			} else {
				mc(valTestInt32PtrConsumer)
			}
		})
	}
}

func TestSilentInt32PtrConsumer(t *testing.T) {
	var sc SilentInt32PtrConsumer = func(v *int32) {
		require.Equal(t, valTestInt32PtrConsumer, v)
		return
	}
	sc(valTestInt32PtrConsumer)
}

func TestSilentInt32PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt32PtrConsumerFactory
		cf2   testInt32PtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, valTestInt32PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, valTestInt32PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, valTestInt32PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestInt32PtrConsumer
				}
			},
			cf2: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, valTestInt32PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, valTestInt32PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32PtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentInt32PtrConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentInt32PtrConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentInt32PtrConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestInt32PtrConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt32PtrConsumer(t *testing.T) {
	var sc SilentInt32PtrConsumer = func(v *int32) {
		require.Equal(t, valTestInt32PtrConsumer, v)
		return
	}
	sc(valTestInt32PtrConsumer)
}

func TestMustInt32PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt32PtrConsumerFactory
		cf2   testInt32PtrConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, valTestInt32PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, valTestInt32PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, valTestInt32PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestInt32PtrConsumer
				}
			},
			cf2: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, valTestInt32PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, valTestInt32PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32PtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustInt32PtrConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustInt32PtrConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustInt32PtrConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(errTestInt32PtrConsumer.Error(), func() {
					cmc(valTestInt32PtrConsumer)
				})
			} else {
				cmc(valTestInt32PtrConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt32PtrConsumer_ToSilentInt32PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					require.Equal(t, valTestInt32PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					require.Equal(t, valTestInt32PtrConsumer, v)
					return errTestInt32PtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt32PtrConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentInt32PtrConsumer()
			r.NotNil(sc)

			sc(valTestInt32PtrConsumer)
		})
	}
}

func TestMustInt32PtrConsumer_ToInt32PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32PtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					require.Equal(t, valTestInt32PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					require.Equal(t, valTestInt32PtrConsumer, v)
					return errTestInt32PtrConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt32PtrConsumer()
			r.NotNil(mc)

			c = mc.ToInt32PtrConsumer()
			r.NotNil(c)

			err := c(valTestInt32PtrConsumer)
			if tt.err {
				r.EqualError(err, errTestInt32PtrConsumer.Error())
			}
		})
	}
}
