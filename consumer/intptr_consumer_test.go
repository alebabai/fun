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
	valTestIntPtrConsumer *int
	errTestIntPtrConsumer = errors.New("error")
)

type testIntPtrConsumerFactory func(t *testing.T) IntPtrConsumer

func TestIntPtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntPtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntPtrConsumer {
				return func(v *int) error {
					require.Equal(t, valTestIntPtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntPtrConsumer {
				return func(v *int) error {
					require.Equal(t, valTestIntPtrConsumer, v)
					return errTestIntPtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestIntPtrConsumer)
			if err != nil {
				r.EqualError(err, errTestIntPtrConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestIntPtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntPtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntPtrConsumer {
				return func(v *int) error {
					require.Equal(t, valTestIntPtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntPtrConsumer {
				return func(v *int) error {
					require.Equal(t, valTestIntPtrConsumer, v)
					return errTestIntPtrConsumer
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

			err := c(valTestIntPtrConsumer)
			if err != nil {
				r.EqualError(err, errTestIntPtrConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestIntPtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testIntPtrConsumerFactory
		cf2   testIntPtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) IntPtrConsumer {
				return func(v *int) error {
					calls++
					require.Equal(t, valTestIntPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntPtrConsumer {
				return func(v *int) error {
					calls++
					require.Equal(t, valTestIntPtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) IntPtrConsumer {
				return func(v *int) error {
					calls++
					require.Equal(t, valTestIntPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestIntPtrConsumer
				}
			},
			cf2: func(t *testing.T) IntPtrConsumer {
				return func(v *int) error {
					calls++
					require.Equal(t, valTestIntPtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) IntPtrConsumer {
				return func(v *int) error {
					calls++
					require.Equal(t, valTestIntPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntPtrConsumer {
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
			err := cc(valTestIntPtrConsumer)
			if err != nil {
				r.EqualError(err, errTestIntPtrConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestIntPtrConsumer_ToSilentIntPtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntPtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntPtrConsumer {
				return func(v *int) error {
					require.Equal(t, valTestIntPtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntPtrConsumer {
				return func(v *int) error {
					require.Equal(t, valTestIntPtrConsumer, v)
					return errTestIntPtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentIntPtrConsumer()
			r.NotNil(sc)

			sc(valTestIntPtrConsumer)
		})
	}
}

func TestIntPtrConsumer_ToMustIntPtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntPtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntPtrConsumer {
				return func(v *int) error {
					require.Equal(t, valTestIntPtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntPtrConsumer {
				return func(v *int) error {
					require.Equal(t, valTestIntPtrConsumer, v)
					return errTestIntPtrConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustIntPtrConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestIntPtrConsumer.Error(), func() {
					mc(valTestIntPtrConsumer)
				})
			} else {
				mc(valTestIntPtrConsumer)
			}
		})
	}
}

func TestSilentIntPtrConsumer(t *testing.T) {
	var sc SilentIntPtrConsumer = func(v *int) {
		require.Equal(t, valTestIntPtrConsumer, v)
		return
	}
	sc(valTestIntPtrConsumer)
}

func TestSilentIntPtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testIntPtrConsumerFactory
		cf2   testIntPtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) IntPtrConsumer {
				return func(v *int) error {
					calls++
					require.Equal(t, valTestIntPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntPtrConsumer {
				return func(v *int) error {
					calls++
					require.Equal(t, valTestIntPtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) IntPtrConsumer {
				return func(v *int) error {
					calls++
					require.Equal(t, valTestIntPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestIntPtrConsumer
				}
			},
			cf2: func(t *testing.T) IntPtrConsumer {
				return func(v *int) error {
					calls++
					require.Equal(t, valTestIntPtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) IntPtrConsumer {
				return func(v *int) error {
					calls++
					require.Equal(t, valTestIntPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntPtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentIntPtrConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentIntPtrConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentIntPtrConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestIntPtrConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustIntPtrConsumer(t *testing.T) {
	var sc SilentIntPtrConsumer = func(v *int) {
		require.Equal(t, valTestIntPtrConsumer, v)
		return
	}
	sc(valTestIntPtrConsumer)
}

func TestMustIntPtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testIntPtrConsumerFactory
		cf2   testIntPtrConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) IntPtrConsumer {
				return func(v *int) error {
					calls++
					require.Equal(t, valTestIntPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntPtrConsumer {
				return func(v *int) error {
					calls++
					require.Equal(t, valTestIntPtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) IntPtrConsumer {
				return func(v *int) error {
					calls++
					require.Equal(t, valTestIntPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestIntPtrConsumer
				}
			},
			cf2: func(t *testing.T) IntPtrConsumer {
				return func(v *int) error {
					calls++
					require.Equal(t, valTestIntPtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) IntPtrConsumer {
				return func(v *int) error {
					calls++
					require.Equal(t, valTestIntPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntPtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustIntPtrConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustIntPtrConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustIntPtrConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(errTestIntPtrConsumer.Error(), func() {
					cmc(valTestIntPtrConsumer)
				})
			} else {
				cmc(valTestIntPtrConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustIntPtrConsumer_ToSilentIntPtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntPtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntPtrConsumer {
				return func(v *int) error {
					require.Equal(t, valTestIntPtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntPtrConsumer {
				return func(v *int) error {
					require.Equal(t, valTestIntPtrConsumer, v)
					return errTestIntPtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustIntPtrConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentIntPtrConsumer()
			r.NotNil(sc)

			sc(valTestIntPtrConsumer)
		})
	}
}

func TestMustIntPtrConsumer_ToIntPtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntPtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntPtrConsumer {
				return func(v *int) error {
					require.Equal(t, valTestIntPtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntPtrConsumer {
				return func(v *int) error {
					require.Equal(t, valTestIntPtrConsumer, v)
					return errTestIntPtrConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustIntPtrConsumer()
			r.NotNil(mc)

			c = mc.ToIntPtrConsumer()
			r.NotNil(c)

			err := c(valTestIntPtrConsumer)
			if tt.err {
				r.EqualError(err, errTestIntPtrConsumer.Error())
			}
		})
	}
}
