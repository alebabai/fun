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
	valTestStringConsumer string
	errTestStringConsumer = errors.New("error")
)

type testStringConsumerFactory func(t *testing.T) StringConsumer

func TestStringConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringConsumer {
				return func(v string) error {
					require.Equal(t, valTestStringConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) StringConsumer {
				return func(v string) error {
					require.Equal(t, valTestStringConsumer, v)
					return errTestStringConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestStringConsumer)
			if err != nil {
				r.EqualError(err, errTestStringConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestStringConsumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringConsumer {
				return func(v string) error {
					require.Equal(t, valTestStringConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) StringConsumer {
				return func(v string) error {
					require.Equal(t, valTestStringConsumer, v)
					return errTestStringConsumer
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

			err := c(valTestStringConsumer)
			if err != nil {
				r.EqualError(err, errTestStringConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestStringConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testStringConsumerFactory
		cf2   testStringConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) StringConsumer {
				return func(v string) error {
					calls++
					require.Equal(t, valTestStringConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringConsumer {
				return func(v string) error {
					calls++
					require.Equal(t, valTestStringConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) StringConsumer {
				return func(v string) error {
					calls++
					require.Equal(t, valTestStringConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestStringConsumer
				}
			},
			cf2: func(t *testing.T) StringConsumer {
				return func(v string) error {
					calls++
					require.Equal(t, valTestStringConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) StringConsumer {
				return func(v string) error {
					calls++
					require.Equal(t, valTestStringConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringConsumer {
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
			err := cc(valTestStringConsumer)
			if err != nil {
				r.EqualError(err, errTestStringConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestStringConsumer_ToSilentStringConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringConsumer {
				return func(v string) error {
					require.Equal(t, valTestStringConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) StringConsumer {
				return func(v string) error {
					require.Equal(t, valTestStringConsumer, v)
					return errTestStringConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentStringConsumer()
			r.NotNil(sc)

			sc(valTestStringConsumer)
		})
	}
}

func TestStringConsumer_ToMustStringConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringConsumer {
				return func(v string) error {
					require.Equal(t, valTestStringConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) StringConsumer {
				return func(v string) error {
					require.Equal(t, valTestStringConsumer, v)
					return errTestStringConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustStringConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestStringConsumer.Error(), func() {
					mc(valTestStringConsumer)
				})
			} else {
				mc(valTestStringConsumer)
			}
		})
	}
}

func TestSilentStringConsumer(t *testing.T) {
	var sc SilentStringConsumer = func(v string) {
		require.Equal(t, valTestStringConsumer, v)
		return
	}
	sc(valTestStringConsumer)
}

func TestSilentStringConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringConsumer {
				return func(v string) error {
					require.Equal(t, valTestStringConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) StringConsumer {
				return func(v string) error {
					require.Equal(t, valTestStringConsumer, v)
					return errTestStringConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentStringConsumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestStringConsumer)
		})
	}
}

func TestSilentStringConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testStringConsumerFactory
		cf2   testStringConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) StringConsumer {
				return func(v string) error {
					calls++
					require.Equal(t, valTestStringConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringConsumer {
				return func(v string) error {
					calls++
					require.Equal(t, valTestStringConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) StringConsumer {
				return func(v string) error {
					calls++
					require.Equal(t, valTestStringConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestStringConsumer
				}
			},
			cf2: func(t *testing.T) StringConsumer {
				return func(v string) error {
					calls++
					require.Equal(t, valTestStringConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) StringConsumer {
				return func(v string) error {
					calls++
					require.Equal(t, valTestStringConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentStringConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentStringConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentStringConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestStringConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustStringConsumer(t *testing.T) {
	var mc MustStringConsumer = func(v string) {
		require.Equal(t, valTestStringConsumer, v)
		return
	}
	mc(valTestStringConsumer)
}

func TestMustStringConsumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringConsumer {
				return func(v string) error {
					require.Equal(t, valTestStringConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) StringConsumer {
				return func(v string) error {
					require.Equal(t, valTestStringConsumer, v)
					return errTestStringConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustStringConsumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestStringConsumer.Error(), func() {
					mc(valTestStringConsumer)
				})
			} else {
				mc(valTestStringConsumer)
			}
		})
	}
}

func TestMustStringConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testStringConsumerFactory
		cf2   testStringConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) StringConsumer {
				return func(v string) error {
					calls++
					require.Equal(t, valTestStringConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringConsumer {
				return func(v string) error {
					calls++
					require.Equal(t, valTestStringConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) StringConsumer {
				return func(v string) error {
					calls++
					require.Equal(t, valTestStringConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestStringConsumer
				}
			},
			cf2: func(t *testing.T) StringConsumer {
				return func(v string) error {
					calls++
					require.Equal(t, valTestStringConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) StringConsumer {
				return func(v string) error {
					calls++
					require.Equal(t, valTestStringConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustStringConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustStringConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustStringConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(errTestStringConsumer.Error(), func() {
					cmc(valTestStringConsumer)
				})
			} else {
				cmc(valTestStringConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustStringConsumer_ToSilentStringConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringConsumer {
				return func(v string) error {
					require.Equal(t, valTestStringConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) StringConsumer {
				return func(v string) error {
					require.Equal(t, valTestStringConsumer, v)
					return errTestStringConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustStringConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentStringConsumer()
			r.NotNil(sc)

			sc(valTestStringConsumer)
		})
	}
}

func TestMustStringConsumer_ToStringConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringConsumer {
				return func(v string) error {
					require.Equal(t, valTestStringConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) StringConsumer {
				return func(v string) error {
					require.Equal(t, valTestStringConsumer, v)
					return errTestStringConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustStringConsumer()
			r.NotNil(mc)

			c = mc.ToStringConsumer()
			r.NotNil(c)

			err := c(valTestStringConsumer)
			if tt.err {
				r.EqualError(err, errTestStringConsumer.Error())
			}
		})
	}
}
