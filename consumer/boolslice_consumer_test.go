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
	testBoolSliceConsumerValue []bool
	testBoolSliceConsumerError = errors.New("error")
)

type testBoolSliceConsumerFactory func(t *testing.T) BoolSliceConsumer

func TestBoolSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					require.Equal(t, testBoolSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					require.Equal(t, testBoolSliceConsumerValue, v)
					return testBoolSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testBoolSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testBoolSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestBoolSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					require.Equal(t, testBoolSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					require.Equal(t, testBoolSliceConsumerValue, v)
					return testBoolSliceConsumerError
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

			err := c(testBoolSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testBoolSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestBoolSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testBoolSliceConsumerFactory
		cf2   testBoolSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, testBoolSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, testBoolSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, testBoolSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testBoolSliceConsumerError
				}
			},
			cf2: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, testBoolSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, testBoolSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolSliceConsumer {
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
			err := cc(testBoolSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testBoolSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestBoolSliceConsumer_ToSilentBoolSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					require.Equal(t, testBoolSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					require.Equal(t, testBoolSliceConsumerValue, v)
					return testBoolSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentBoolSliceConsumer()
			r.NotNil(sc)

			sc(testBoolSliceConsumerValue)
		})
	}
}

func TestBoolSliceConsumer_ToMustBoolSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					require.Equal(t, testBoolSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					require.Equal(t, testBoolSliceConsumerValue, v)
					return testBoolSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustBoolSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testBoolSliceConsumerError.Error(), func() {
					mc(testBoolSliceConsumerValue)
				})
			} else {
				mc(testBoolSliceConsumerValue)
			}
		})
	}
}

func TestSilentBoolSliceConsumer(t *testing.T) {
	var sc SilentBoolSliceConsumer = func(v []bool) {
		require.Equal(t, testBoolSliceConsumerValue, v)
		return
	}
	sc(testBoolSliceConsumerValue)
}

func TestSilentBoolSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testBoolSliceConsumerFactory
		cf2   testBoolSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, testBoolSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, testBoolSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, testBoolSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testBoolSliceConsumerError
				}
			},
			cf2: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, testBoolSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, testBoolSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentBoolSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentBoolSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentBoolSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testBoolSliceConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustBoolSliceConsumer(t *testing.T) {
	var sc SilentBoolSliceConsumer = func(v []bool) {
		require.Equal(t, testBoolSliceConsumerValue, v)
		return
	}
	sc(testBoolSliceConsumerValue)
}

func TestMustBoolSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testBoolSliceConsumerFactory
		cf2   testBoolSliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, testBoolSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, testBoolSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, testBoolSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testBoolSliceConsumerError
				}
			},
			cf2: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, testBoolSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, testBoolSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustBoolSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustBoolSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustBoolSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testBoolSliceConsumerError.Error(), func() {
					cmc(testBoolSliceConsumerValue)
				})
			} else {
				cmc(testBoolSliceConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustBoolSliceConsumer_ToSilentBoolSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					require.Equal(t, testBoolSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					require.Equal(t, testBoolSliceConsumerValue, v)
					return testBoolSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustBoolSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentBoolSliceConsumer()
			r.NotNil(sc)

			sc(testBoolSliceConsumerValue)
		})
	}
}

func TestMustBoolSliceConsumer_ToBoolSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					require.Equal(t, testBoolSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					require.Equal(t, testBoolSliceConsumerValue, v)
					return testBoolSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustBoolSliceConsumer()
			r.NotNil(mc)

			c = mc.ToBoolSliceConsumer()
			r.NotNil(c)

			err := c(testBoolSliceConsumerValue)
			if tt.err {
				r.EqualError(err, testBoolSliceConsumerError.Error())
			}
		})
	}
}
