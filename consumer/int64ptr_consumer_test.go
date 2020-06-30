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
	valTestInt64PtrConsumer *int64
	errTestInt64PtrConsumer = errors.New("error")
)

type testInt64PtrConsumerFactory func(t *testing.T) Int64PtrConsumer

func TestInt64PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64PtrConsumer {
				return func(v *int64) error {
					require.Equal(t, valTestInt64PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64PtrConsumer {
				return func(v *int64) error {
					require.Equal(t, valTestInt64PtrConsumer, v)
					return errTestInt64PtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestInt64PtrConsumer)
			if err != nil {
				r.EqualError(err, errTestInt64PtrConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt64PtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64PtrConsumer {
				return func(v *int64) error {
					require.Equal(t, valTestInt64PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64PtrConsumer {
				return func(v *int64) error {
					require.Equal(t, valTestInt64PtrConsumer, v)
					return errTestInt64PtrConsumer
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

			err := c(valTestInt64PtrConsumer)
			if err != nil {
				r.EqualError(err, errTestInt64PtrConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt64PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt64PtrConsumerFactory
		cf2   testInt64PtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int64PtrConsumer {
				return func(v *int64) error {
					calls++
					require.Equal(t, valTestInt64PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64PtrConsumer {
				return func(v *int64) error {
					calls++
					require.Equal(t, valTestInt64PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int64PtrConsumer {
				return func(v *int64) error {
					calls++
					require.Equal(t, valTestInt64PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestInt64PtrConsumer
				}
			},
			cf2: func(t *testing.T) Int64PtrConsumer {
				return func(v *int64) error {
					calls++
					require.Equal(t, valTestInt64PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int64PtrConsumer {
				return func(v *int64) error {
					calls++
					require.Equal(t, valTestInt64PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64PtrConsumer {
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
			err := cc(valTestInt64PtrConsumer)
			if err != nil {
				r.EqualError(err, errTestInt64PtrConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestInt64PtrConsumer_ToSilentInt64PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64PtrConsumer {
				return func(v *int64) error {
					require.Equal(t, valTestInt64PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64PtrConsumer {
				return func(v *int64) error {
					require.Equal(t, valTestInt64PtrConsumer, v)
					return errTestInt64PtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentInt64PtrConsumer()
			r.NotNil(sc)

			sc(valTestInt64PtrConsumer)
		})
	}
}

func TestInt64PtrConsumer_ToMustInt64PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64PtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64PtrConsumer {
				return func(v *int64) error {
					require.Equal(t, valTestInt64PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64PtrConsumer {
				return func(v *int64) error {
					require.Equal(t, valTestInt64PtrConsumer, v)
					return errTestInt64PtrConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustInt64PtrConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestInt64PtrConsumer.Error(), func() {
					mc(valTestInt64PtrConsumer)
				})
			} else {
				mc(valTestInt64PtrConsumer)
			}
		})
	}
}

func TestSilentInt64PtrConsumer(t *testing.T) {
	var sc SilentInt64PtrConsumer = func(v *int64) {
		require.Equal(t, valTestInt64PtrConsumer, v)
		return
	}
	sc(valTestInt64PtrConsumer)
}

func TestSilentInt64PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt64PtrConsumerFactory
		cf2   testInt64PtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int64PtrConsumer {
				return func(v *int64) error {
					calls++
					require.Equal(t, valTestInt64PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64PtrConsumer {
				return func(v *int64) error {
					calls++
					require.Equal(t, valTestInt64PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int64PtrConsumer {
				return func(v *int64) error {
					calls++
					require.Equal(t, valTestInt64PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestInt64PtrConsumer
				}
			},
			cf2: func(t *testing.T) Int64PtrConsumer {
				return func(v *int64) error {
					calls++
					require.Equal(t, valTestInt64PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int64PtrConsumer {
				return func(v *int64) error {
					calls++
					require.Equal(t, valTestInt64PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64PtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentInt64PtrConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentInt64PtrConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentInt64PtrConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestInt64PtrConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt64PtrConsumer(t *testing.T) {
	var sc SilentInt64PtrConsumer = func(v *int64) {
		require.Equal(t, valTestInt64PtrConsumer, v)
		return
	}
	sc(valTestInt64PtrConsumer)
}

func TestMustInt64PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt64PtrConsumerFactory
		cf2   testInt64PtrConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int64PtrConsumer {
				return func(v *int64) error {
					calls++
					require.Equal(t, valTestInt64PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64PtrConsumer {
				return func(v *int64) error {
					calls++
					require.Equal(t, valTestInt64PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int64PtrConsumer {
				return func(v *int64) error {
					calls++
					require.Equal(t, valTestInt64PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestInt64PtrConsumer
				}
			},
			cf2: func(t *testing.T) Int64PtrConsumer {
				return func(v *int64) error {
					calls++
					require.Equal(t, valTestInt64PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int64PtrConsumer {
				return func(v *int64) error {
					calls++
					require.Equal(t, valTestInt64PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64PtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustInt64PtrConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustInt64PtrConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustInt64PtrConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(errTestInt64PtrConsumer.Error(), func() {
					cmc(valTestInt64PtrConsumer)
				})
			} else {
				cmc(valTestInt64PtrConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt64PtrConsumer_ToSilentInt64PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64PtrConsumer {
				return func(v *int64) error {
					require.Equal(t, valTestInt64PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64PtrConsumer {
				return func(v *int64) error {
					require.Equal(t, valTestInt64PtrConsumer, v)
					return errTestInt64PtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt64PtrConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentInt64PtrConsumer()
			r.NotNil(sc)

			sc(valTestInt64PtrConsumer)
		})
	}
}

func TestMustInt64PtrConsumer_ToInt64PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64PtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64PtrConsumer {
				return func(v *int64) error {
					require.Equal(t, valTestInt64PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64PtrConsumer {
				return func(v *int64) error {
					require.Equal(t, valTestInt64PtrConsumer, v)
					return errTestInt64PtrConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt64PtrConsumer()
			r.NotNil(mc)

			c = mc.ToInt64PtrConsumer()
			r.NotNil(c)

			err := c(valTestInt64PtrConsumer)
			if tt.err {
				r.EqualError(err, errTestInt64PtrConsumer.Error())
			}
		})
	}
}
