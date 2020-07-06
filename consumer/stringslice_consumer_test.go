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
	valTestStringSliceConsumer []string
	errTestStringSliceConsumer = errors.New("error")
)

type testStringSliceConsumerFactory func(t *testing.T) StringSliceConsumer

func TestStringSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					require.Equal(t, valTestStringSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					require.Equal(t, valTestStringSliceConsumer, v)
					return errTestStringSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestStringSliceConsumer)
			if err != nil {
				r.EqualError(err, errTestStringSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestStringSliceConsumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					require.Equal(t, valTestStringSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					require.Equal(t, valTestStringSliceConsumer, v)
					return errTestStringSliceConsumer
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

			err := c(valTestStringSliceConsumer)
			if err != nil {
				r.EqualError(err, errTestStringSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestStringSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testStringSliceConsumerFactory
		cf2   testStringSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					calls++
					require.Equal(t, valTestStringSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					calls++
					require.Equal(t, valTestStringSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					calls++
					require.Equal(t, valTestStringSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestStringSliceConsumer
				}
			},
			cf2: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					calls++
					require.Equal(t, valTestStringSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					calls++
					require.Equal(t, valTestStringSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringSliceConsumer {
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
			err := cc(valTestStringSliceConsumer)
			if err != nil {
				r.EqualError(err, errTestStringSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestStringSliceConsumer_ToSilentStringSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					require.Equal(t, valTestStringSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					require.Equal(t, valTestStringSliceConsumer, v)
					return errTestStringSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentStringSliceConsumer()
			r.NotNil(sc)

			sc(valTestStringSliceConsumer)
		})
	}
}

func TestStringSliceConsumer_ToMustStringSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					require.Equal(t, valTestStringSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					require.Equal(t, valTestStringSliceConsumer, v)
					return errTestStringSliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustStringSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestStringSliceConsumer.Error(), func() {
					mc(valTestStringSliceConsumer)
				})
			} else {
				mc(valTestStringSliceConsumer)
			}
		})
	}
}

func TestSilentStringSliceConsumer(t *testing.T) {
	var sc SilentStringSliceConsumer = func(v []string) {
		require.Equal(t, valTestStringSliceConsumer, v)
		return
	}
	sc(valTestStringSliceConsumer)
}

func TestSilentStringSliceConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					require.Equal(t, valTestStringSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					require.Equal(t, valTestStringSliceConsumer, v)
					return errTestStringSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentStringSliceConsumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestStringSliceConsumer)
		})
	}
}

func TestSilentStringSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testStringSliceConsumerFactory
		cf2   testStringSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					calls++
					require.Equal(t, valTestStringSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					calls++
					require.Equal(t, valTestStringSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					calls++
					require.Equal(t, valTestStringSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestStringSliceConsumer
				}
			},
			cf2: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					calls++
					require.Equal(t, valTestStringSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					calls++
					require.Equal(t, valTestStringSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentStringSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentStringSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentStringSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestStringSliceConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustStringSliceConsumer(t *testing.T) {
	var mc MustStringSliceConsumer = func(v []string) {
		require.Equal(t, valTestStringSliceConsumer, v)
		return
	}
	mc(valTestStringSliceConsumer)
}

func TestMustStringSliceConsumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					require.Equal(t, valTestStringSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					require.Equal(t, valTestStringSliceConsumer, v)
					return errTestStringSliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustStringSliceConsumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestStringSliceConsumer.Error(), func() {
					mc(valTestStringSliceConsumer)
				})
			} else {
				mc(valTestStringSliceConsumer)
			}
		})
	}
}

func TestMustStringSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testStringSliceConsumerFactory
		cf2   testStringSliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					calls++
					require.Equal(t, valTestStringSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					calls++
					require.Equal(t, valTestStringSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					calls++
					require.Equal(t, valTestStringSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestStringSliceConsumer
				}
			},
			cf2: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					calls++
					require.Equal(t, valTestStringSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					calls++
					require.Equal(t, valTestStringSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustStringSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustStringSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustStringSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(errTestStringSliceConsumer.Error(), func() {
					cmc(valTestStringSliceConsumer)
				})
			} else {
				cmc(valTestStringSliceConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustStringSliceConsumer_ToSilentStringSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					require.Equal(t, valTestStringSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					require.Equal(t, valTestStringSliceConsumer, v)
					return errTestStringSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustStringSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentStringSliceConsumer()
			r.NotNil(sc)

			sc(valTestStringSliceConsumer)
		})
	}
}

func TestMustStringSliceConsumer_ToStringSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					require.Equal(t, valTestStringSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) StringSliceConsumer {
				return func(v []string) error {
					require.Equal(t, valTestStringSliceConsumer, v)
					return errTestStringSliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustStringSliceConsumer()
			r.NotNil(mc)

			c = mc.ToStringSliceConsumer()
			r.NotNil(c)

			err := c(valTestStringSliceConsumer)
			if tt.err {
				r.EqualError(err, errTestStringSliceConsumer.Error())
			}
		})
	}
}
