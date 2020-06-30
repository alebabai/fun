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
	valTestIntSliceConsumer []int
	errTestIntSliceConsumer = errors.New("error")
)

type testIntSliceConsumerFactory func(t *testing.T) IntSliceConsumer

func TestIntSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					require.Equal(t, valTestIntSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					require.Equal(t, valTestIntSliceConsumer, v)
					return errTestIntSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestIntSliceConsumer)
			if err != nil {
				r.EqualError(err, errTestIntSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestIntSliceConsumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					require.Equal(t, valTestIntSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					require.Equal(t, valTestIntSliceConsumer, v)
					return errTestIntSliceConsumer
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

			err := c(valTestIntSliceConsumer)
			if err != nil {
				r.EqualError(err, errTestIntSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestIntSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testIntSliceConsumerFactory
		cf2   testIntSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, valTestIntSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, valTestIntSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, valTestIntSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestIntSliceConsumer
				}
			},
			cf2: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, valTestIntSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, valTestIntSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntSliceConsumer {
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
			err := cc(valTestIntSliceConsumer)
			if err != nil {
				r.EqualError(err, errTestIntSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestIntSliceConsumer_ToSilentIntSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					require.Equal(t, valTestIntSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					require.Equal(t, valTestIntSliceConsumer, v)
					return errTestIntSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentIntSliceConsumer()
			r.NotNil(sc)

			sc(valTestIntSliceConsumer)
		})
	}
}

func TestIntSliceConsumer_ToMustIntSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					require.Equal(t, valTestIntSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					require.Equal(t, valTestIntSliceConsumer, v)
					return errTestIntSliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustIntSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestIntSliceConsumer.Error(), func() {
					mc(valTestIntSliceConsumer)
				})
			} else {
				mc(valTestIntSliceConsumer)
			}
		})
	}
}

func TestSilentIntSliceConsumer(t *testing.T) {
	var sc SilentIntSliceConsumer = func(v []int) {
		require.Equal(t, valTestIntSliceConsumer, v)
		return
	}
	sc(valTestIntSliceConsumer)
}

func TestSilentIntSliceConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					require.Equal(t, valTestIntSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					require.Equal(t, valTestIntSliceConsumer, v)
					return errTestIntSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentIntSliceConsumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestIntSliceConsumer)
		})
	}
}

func TestSilentIntSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testIntSliceConsumerFactory
		cf2   testIntSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, valTestIntSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, valTestIntSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, valTestIntSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestIntSliceConsumer
				}
			},
			cf2: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, valTestIntSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, valTestIntSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentIntSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentIntSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentIntSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestIntSliceConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustIntSliceConsumer(t *testing.T) {
	var mc MustIntSliceConsumer = func(v []int) {
		require.Equal(t, valTestIntSliceConsumer, v)
		return
	}
	mc(valTestIntSliceConsumer)
}

func TestMustIntSliceConsumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					require.Equal(t, valTestIntSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					require.Equal(t, valTestIntSliceConsumer, v)
					return errTestIntSliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustIntSliceConsumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestIntSliceConsumer.Error(), func() {
					mc(valTestIntSliceConsumer)
				})
			} else {
				mc(valTestIntSliceConsumer)
			}
		})
	}
}

func TestMustIntSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testIntSliceConsumerFactory
		cf2   testIntSliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, valTestIntSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, valTestIntSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, valTestIntSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestIntSliceConsumer
				}
			},
			cf2: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, valTestIntSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, valTestIntSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustIntSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustIntSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustIntSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(errTestIntSliceConsumer.Error(), func() {
					cmc(valTestIntSliceConsumer)
				})
			} else {
				cmc(valTestIntSliceConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustIntSliceConsumer_ToSilentIntSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					require.Equal(t, valTestIntSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					require.Equal(t, valTestIntSliceConsumer, v)
					return errTestIntSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustIntSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentIntSliceConsumer()
			r.NotNil(sc)

			sc(valTestIntSliceConsumer)
		})
	}
}

func TestMustIntSliceConsumer_ToIntSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					require.Equal(t, valTestIntSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					require.Equal(t, valTestIntSliceConsumer, v)
					return errTestIntSliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustIntSliceConsumer()
			r.NotNil(mc)

			c = mc.ToIntSliceConsumer()
			r.NotNil(c)

			err := c(valTestIntSliceConsumer)
			if tt.err {
				r.EqualError(err, errTestIntSliceConsumer.Error())
			}
		})
	}
}
