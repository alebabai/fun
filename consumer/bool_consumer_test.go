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
	testBoolConsumerValue bool
	testBoolConsumerError = errors.New("error")
)

type testBoolConsumerFactory func(t *testing.T) BoolConsumer

func TestBoolConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolConsumer {
				return func(v bool) error {
					require.Equal(t, testBoolConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BoolConsumer {
				return func(v bool) error {
					require.Equal(t, testBoolConsumerValue, v)
					return testBoolConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testBoolConsumerValue)
			if err != nil {
				r.EqualError(err, testBoolConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestBoolSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolConsumer {
				return func(v bool) error {
					require.Equal(t, testBoolConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BoolConsumer {
				return func(v bool) error {
					require.Equal(t, testBoolConsumerValue, v)
					return testBoolConsumerError
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

			err := c(testBoolConsumerValue)
			if err != nil {
				r.EqualError(err, testBoolConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestBoolConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testBoolConsumerFactory
		cf2   testBoolConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) BoolConsumer {
				return func(v bool) error {
					calls++
					require.Equal(t, testBoolConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolConsumer {
				return func(v bool) error {
					calls++
					require.Equal(t, testBoolConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) BoolConsumer {
				return func(v bool) error {
					calls++
					require.Equal(t, testBoolConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testBoolConsumerError
				}
			},
			cf2: func(t *testing.T) BoolConsumer {
				return func(v bool) error {
					calls++
					require.Equal(t, testBoolConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) BoolConsumer {
				return func(v bool) error {
					calls++
					require.Equal(t, testBoolConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolConsumer {
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
			err := cc(testBoolConsumerValue)
			if err != nil {
				r.EqualError(err, testBoolConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestBoolConsumer_ToSilentBoolConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolConsumer {
				return func(v bool) error {
					require.Equal(t, testBoolConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BoolConsumer {
				return func(v bool) error {
					require.Equal(t, testBoolConsumerValue, v)
					return testBoolConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentBoolConsumer()
			r.NotNil(sc)

			sc(testBoolConsumerValue)
		})
	}
}

func TestBoolConsumer_ToMustBoolConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolConsumer {
				return func(v bool) error {
					require.Equal(t, testBoolConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BoolConsumer {
				return func(v bool) error {
					require.Equal(t, testBoolConsumerValue, v)
					return testBoolConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustBoolConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testBoolConsumerError.Error(), func() {
					mc(testBoolConsumerValue)
				})
			} else {
				mc(testBoolConsumerValue)
			}
		})
	}
}

func TestSilentBoolConsumer(t *testing.T) {
	var sc SilentBoolConsumer = func(v bool) {
		require.Equal(t, testBoolConsumerValue, v)
		return
	}
	sc(testBoolConsumerValue)
}

func TestSilentBoolConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testBoolConsumerFactory
		cf2   testBoolConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) BoolConsumer {
				return func(v bool) error {
					calls++
					require.Equal(t, testBoolConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolConsumer {
				return func(v bool) error {
					calls++
					require.Equal(t, testBoolConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) BoolConsumer {
				return func(v bool) error {
					calls++
					require.Equal(t, testBoolConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testBoolConsumerError
				}
			},
			cf2: func(t *testing.T) BoolConsumer {
				return func(v bool) error {
					calls++
					require.Equal(t, testBoolConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) BoolConsumer {
				return func(v bool) error {
					calls++
					require.Equal(t, testBoolConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentBoolConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentBoolConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentBoolConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testBoolConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustBoolConsumer(t *testing.T) {
	var sc SilentBoolConsumer = func(v bool) {
		require.Equal(t, testBoolConsumerValue, v)
		return
	}
	sc(testBoolConsumerValue)
}

func TestMustBoolConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testBoolConsumerFactory
		cf2   testBoolConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) BoolConsumer {
				return func(v bool) error {
					calls++
					require.Equal(t, testBoolConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolConsumer {
				return func(v bool) error {
					calls++
					require.Equal(t, testBoolConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) BoolConsumer {
				return func(v bool) error {
					calls++
					require.Equal(t, testBoolConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testBoolConsumerError
				}
			},
			cf2: func(t *testing.T) BoolConsumer {
				return func(v bool) error {
					calls++
					require.Equal(t, testBoolConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) BoolConsumer {
				return func(v bool) error {
					calls++
					require.Equal(t, testBoolConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustBoolConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustBoolConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustBoolConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testBoolConsumerError.Error(), func() {
					cmc(testBoolConsumerValue)
				})
			} else {
				cmc(testBoolConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustBoolConsumer_ToSilentBoolConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolConsumer {
				return func(v bool) error {
					require.Equal(t, testBoolConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BoolConsumer {
				return func(v bool) error {
					require.Equal(t, testBoolConsumerValue, v)
					return testBoolConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustBoolConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentBoolConsumer()
			r.NotNil(sc)

			sc(testBoolConsumerValue)
		})
	}
}

func TestMustBoolConsumer_ToBoolConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolConsumer {
				return func(v bool) error {
					require.Equal(t, testBoolConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BoolConsumer {
				return func(v bool) error {
					require.Equal(t, testBoolConsumerValue, v)
					return testBoolConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustBoolConsumer()
			r.NotNil(mc)

			c = mc.ToBoolConsumer()
			r.NotNil(c)

			err := c(testBoolConsumerValue)
			if tt.err {
				r.EqualError(err, testBoolConsumerError.Error())
			}
		})
	}
}
