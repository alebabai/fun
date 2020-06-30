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
	valTestRuneSliceConsumer []rune
	errTestRuneSliceConsumer = errors.New("error")
)

type testRuneSliceConsumerFactory func(t *testing.T) RuneSliceConsumer

func TestRuneSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testRuneSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RuneSliceConsumer {
				return func(v []rune) error {
					require.Equal(t, valTestRuneSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) RuneSliceConsumer {
				return func(v []rune) error {
					require.Equal(t, valTestRuneSliceConsumer, v)
					return errTestRuneSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestRuneSliceConsumer)
			if err != nil {
				r.EqualError(err, errTestRuneSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestRuneSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testRuneSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RuneSliceConsumer {
				return func(v []rune) error {
					require.Equal(t, valTestRuneSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) RuneSliceConsumer {
				return func(v []rune) error {
					require.Equal(t, valTestRuneSliceConsumer, v)
					return errTestRuneSliceConsumer
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

			err := c(valTestRuneSliceConsumer)
			if err != nil {
				r.EqualError(err, errTestRuneSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestRuneSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testRuneSliceConsumerFactory
		cf2   testRuneSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) RuneSliceConsumer {
				return func(v []rune) error {
					calls++
					require.Equal(t, valTestRuneSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RuneSliceConsumer {
				return func(v []rune) error {
					calls++
					require.Equal(t, valTestRuneSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) RuneSliceConsumer {
				return func(v []rune) error {
					calls++
					require.Equal(t, valTestRuneSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestRuneSliceConsumer
				}
			},
			cf2: func(t *testing.T) RuneSliceConsumer {
				return func(v []rune) error {
					calls++
					require.Equal(t, valTestRuneSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) RuneSliceConsumer {
				return func(v []rune) error {
					calls++
					require.Equal(t, valTestRuneSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RuneSliceConsumer {
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
			err := cc(valTestRuneSliceConsumer)
			if err != nil {
				r.EqualError(err, errTestRuneSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestRuneSliceConsumer_ToSilentRuneSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testRuneSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RuneSliceConsumer {
				return func(v []rune) error {
					require.Equal(t, valTestRuneSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) RuneSliceConsumer {
				return func(v []rune) error {
					require.Equal(t, valTestRuneSliceConsumer, v)
					return errTestRuneSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentRuneSliceConsumer()
			r.NotNil(sc)

			sc(valTestRuneSliceConsumer)
		})
	}
}

func TestRuneSliceConsumer_ToMustRuneSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testRuneSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RuneSliceConsumer {
				return func(v []rune) error {
					require.Equal(t, valTestRuneSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) RuneSliceConsumer {
				return func(v []rune) error {
					require.Equal(t, valTestRuneSliceConsumer, v)
					return errTestRuneSliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustRuneSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestRuneSliceConsumer.Error(), func() {
					mc(valTestRuneSliceConsumer)
				})
			} else {
				mc(valTestRuneSliceConsumer)
			}
		})
	}
}

func TestSilentRuneSliceConsumer(t *testing.T) {
	var sc SilentRuneSliceConsumer = func(v []rune) {
		require.Equal(t, valTestRuneSliceConsumer, v)
		return
	}
	sc(valTestRuneSliceConsumer)
}

func TestSilentRuneSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testRuneSliceConsumerFactory
		cf2   testRuneSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) RuneSliceConsumer {
				return func(v []rune) error {
					calls++
					require.Equal(t, valTestRuneSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RuneSliceConsumer {
				return func(v []rune) error {
					calls++
					require.Equal(t, valTestRuneSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) RuneSliceConsumer {
				return func(v []rune) error {
					calls++
					require.Equal(t, valTestRuneSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestRuneSliceConsumer
				}
			},
			cf2: func(t *testing.T) RuneSliceConsumer {
				return func(v []rune) error {
					calls++
					require.Equal(t, valTestRuneSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) RuneSliceConsumer {
				return func(v []rune) error {
					calls++
					require.Equal(t, valTestRuneSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RuneSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentRuneSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentRuneSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentRuneSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestRuneSliceConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustRuneSliceConsumer(t *testing.T) {
	var sc SilentRuneSliceConsumer = func(v []rune) {
		require.Equal(t, valTestRuneSliceConsumer, v)
		return
	}
	sc(valTestRuneSliceConsumer)
}

func TestMustRuneSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testRuneSliceConsumerFactory
		cf2   testRuneSliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) RuneSliceConsumer {
				return func(v []rune) error {
					calls++
					require.Equal(t, valTestRuneSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RuneSliceConsumer {
				return func(v []rune) error {
					calls++
					require.Equal(t, valTestRuneSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) RuneSliceConsumer {
				return func(v []rune) error {
					calls++
					require.Equal(t, valTestRuneSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestRuneSliceConsumer
				}
			},
			cf2: func(t *testing.T) RuneSliceConsumer {
				return func(v []rune) error {
					calls++
					require.Equal(t, valTestRuneSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) RuneSliceConsumer {
				return func(v []rune) error {
					calls++
					require.Equal(t, valTestRuneSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RuneSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustRuneSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustRuneSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustRuneSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(errTestRuneSliceConsumer.Error(), func() {
					cmc(valTestRuneSliceConsumer)
				})
			} else {
				cmc(valTestRuneSliceConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustRuneSliceConsumer_ToSilentRuneSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testRuneSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RuneSliceConsumer {
				return func(v []rune) error {
					require.Equal(t, valTestRuneSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) RuneSliceConsumer {
				return func(v []rune) error {
					require.Equal(t, valTestRuneSliceConsumer, v)
					return errTestRuneSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustRuneSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentRuneSliceConsumer()
			r.NotNil(sc)

			sc(valTestRuneSliceConsumer)
		})
	}
}

func TestMustRuneSliceConsumer_ToRuneSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testRuneSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RuneSliceConsumer {
				return func(v []rune) error {
					require.Equal(t, valTestRuneSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) RuneSliceConsumer {
				return func(v []rune) error {
					require.Equal(t, valTestRuneSliceConsumer, v)
					return errTestRuneSliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustRuneSliceConsumer()
			r.NotNil(mc)

			c = mc.ToRuneSliceConsumer()
			r.NotNil(c)

			err := c(valTestRuneSliceConsumer)
			if tt.err {
				r.EqualError(err, errTestRuneSliceConsumer.Error())
			}
		})
	}
}
