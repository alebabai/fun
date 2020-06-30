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
	valTestRuneConsumer rune
	errTestRuneConsumer = errors.New("error")
)

type testRuneConsumerFactory func(t *testing.T) RuneConsumer

func TestRuneConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testRuneConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					require.Equal(t, valTestRuneConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					require.Equal(t, valTestRuneConsumer, v)
					return errTestRuneConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestRuneConsumer)
			if err != nil {
				r.EqualError(err, errTestRuneConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestRuneConsumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testRuneConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					require.Equal(t, valTestRuneConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					require.Equal(t, valTestRuneConsumer, v)
					return errTestRuneConsumer
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

			err := c(valTestRuneConsumer)
			if err != nil {
				r.EqualError(err, errTestRuneConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestRuneConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testRuneConsumerFactory
		cf2   testRuneConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					calls++
					require.Equal(t, valTestRuneConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					calls++
					require.Equal(t, valTestRuneConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					calls++
					require.Equal(t, valTestRuneConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestRuneConsumer
				}
			},
			cf2: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					calls++
					require.Equal(t, valTestRuneConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					calls++
					require.Equal(t, valTestRuneConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RuneConsumer {
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
			err := cc(valTestRuneConsumer)
			if err != nil {
				r.EqualError(err, errTestRuneConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestRuneConsumer_ToSilentRuneConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testRuneConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					require.Equal(t, valTestRuneConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					require.Equal(t, valTestRuneConsumer, v)
					return errTestRuneConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentRuneConsumer()
			r.NotNil(sc)

			sc(valTestRuneConsumer)
		})
	}
}

func TestRuneConsumer_ToMustRuneConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testRuneConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					require.Equal(t, valTestRuneConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					require.Equal(t, valTestRuneConsumer, v)
					return errTestRuneConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustRuneConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestRuneConsumer.Error(), func() {
					mc(valTestRuneConsumer)
				})
			} else {
				mc(valTestRuneConsumer)
			}
		})
	}
}

func TestSilentRuneConsumer(t *testing.T) {
	var sc SilentRuneConsumer = func(v rune) {
		require.Equal(t, valTestRuneConsumer, v)
		return
	}
	sc(valTestRuneConsumer)
}

func TestSilentRuneConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testRuneConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					require.Equal(t, valTestRuneConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					require.Equal(t, valTestRuneConsumer, v)
					return errTestRuneConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentRuneConsumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestRuneConsumer)
		})
	}
}

func TestSilentRuneConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testRuneConsumerFactory
		cf2   testRuneConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					calls++
					require.Equal(t, valTestRuneConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					calls++
					require.Equal(t, valTestRuneConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					calls++
					require.Equal(t, valTestRuneConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestRuneConsumer
				}
			},
			cf2: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					calls++
					require.Equal(t, valTestRuneConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					calls++
					require.Equal(t, valTestRuneConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RuneConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentRuneConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentRuneConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentRuneConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestRuneConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustRuneConsumer(t *testing.T) {
	var mc MustRuneConsumer = func(v rune) {
		require.Equal(t, valTestRuneConsumer, v)
		return
	}
	mc(valTestRuneConsumer)
}

func TestMustRuneConsumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testRuneConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					require.Equal(t, valTestRuneConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					require.Equal(t, valTestRuneConsumer, v)
					return errTestRuneConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustRuneConsumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestRuneConsumer.Error(), func() {
					mc(valTestRuneConsumer)
				})
			} else {
				mc(valTestRuneConsumer)
			}
		})
	}
}

func TestMustRuneConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testRuneConsumerFactory
		cf2   testRuneConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					calls++
					require.Equal(t, valTestRuneConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					calls++
					require.Equal(t, valTestRuneConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					calls++
					require.Equal(t, valTestRuneConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestRuneConsumer
				}
			},
			cf2: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					calls++
					require.Equal(t, valTestRuneConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					calls++
					require.Equal(t, valTestRuneConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RuneConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustRuneConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustRuneConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustRuneConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(errTestRuneConsumer.Error(), func() {
					cmc(valTestRuneConsumer)
				})
			} else {
				cmc(valTestRuneConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustRuneConsumer_ToSilentRuneConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testRuneConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					require.Equal(t, valTestRuneConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					require.Equal(t, valTestRuneConsumer, v)
					return errTestRuneConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustRuneConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentRuneConsumer()
			r.NotNil(sc)

			sc(valTestRuneConsumer)
		})
	}
}

func TestMustRuneConsumer_ToRuneConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testRuneConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					require.Equal(t, valTestRuneConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) RuneConsumer {
				return func(v rune) error {
					require.Equal(t, valTestRuneConsumer, v)
					return errTestRuneConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustRuneConsumer()
			r.NotNil(mc)

			c = mc.ToRuneConsumer()
			r.NotNil(c)

			err := c(valTestRuneConsumer)
			if tt.err {
				r.EqualError(err, errTestRuneConsumer.Error())
			}
		})
	}
}
