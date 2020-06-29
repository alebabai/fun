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
	testInt64SliceConsumerValue []int64
	testInt64SliceConsumerError = errors.New("error")
)

type testInt64SliceConsumerFactory func(t *testing.T) Int64SliceConsumer

func TestInt64SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64SliceConsumer {
				return func(v []int64) error {
					require.Equal(t, testInt64SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64SliceConsumer {
				return func(v []int64) error {
					require.Equal(t, testInt64SliceConsumerValue, v)
					return testInt64SliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testInt64SliceConsumerValue)
			if err != nil {
				r.EqualError(err, testInt64SliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt64SliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64SliceConsumer {
				return func(v []int64) error {
					require.Equal(t, testInt64SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64SliceConsumer {
				return func(v []int64) error {
					require.Equal(t, testInt64SliceConsumerValue, v)
					return testInt64SliceConsumerError
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

			err := c(testInt64SliceConsumerValue)
			if err != nil {
				r.EqualError(err, testInt64SliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt64SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt64SliceConsumerFactory
		cf2   testInt64SliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int64SliceConsumer {
				return func(v []int64) error {
					calls++
					require.Equal(t, testInt64SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64SliceConsumer {
				return func(v []int64) error {
					calls++
					require.Equal(t, testInt64SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int64SliceConsumer {
				return func(v []int64) error {
					calls++
					require.Equal(t, testInt64SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt64SliceConsumerError
				}
			},
			cf2: func(t *testing.T) Int64SliceConsumer {
				return func(v []int64) error {
					calls++
					require.Equal(t, testInt64SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int64SliceConsumer {
				return func(v []int64) error {
					calls++
					require.Equal(t, testInt64SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64SliceConsumer {
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
			err := cc(testInt64SliceConsumerValue)
			if err != nil {
				r.EqualError(err, testInt64SliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestInt64SliceConsumer_ToSilentInt64SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64SliceConsumer {
				return func(v []int64) error {
					require.Equal(t, testInt64SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64SliceConsumer {
				return func(v []int64) error {
					require.Equal(t, testInt64SliceConsumerValue, v)
					return testInt64SliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentInt64SliceConsumer()
			r.NotNil(sc)

			sc(testInt64SliceConsumerValue)
		})
	}
}

func TestInt64SliceConsumer_ToMustInt64SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64SliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64SliceConsumer {
				return func(v []int64) error {
					require.Equal(t, testInt64SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64SliceConsumer {
				return func(v []int64) error {
					require.Equal(t, testInt64SliceConsumerValue, v)
					return testInt64SliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustInt64SliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testInt64SliceConsumerError.Error(), func() {
					mc(testInt64SliceConsumerValue)
				})
			} else {
				mc(testInt64SliceConsumerValue)
			}
		})
	}
}

func TestSilentInt64SliceConsumer(t *testing.T) {
	var sc SilentInt64SliceConsumer = func(v []int64) {
		require.Equal(t, testInt64SliceConsumerValue, v)
		return
	}
	sc(testInt64SliceConsumerValue)
}

func TestSilentInt64SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt64SliceConsumerFactory
		cf2   testInt64SliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int64SliceConsumer {
				return func(v []int64) error {
					calls++
					require.Equal(t, testInt64SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64SliceConsumer {
				return func(v []int64) error {
					calls++
					require.Equal(t, testInt64SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int64SliceConsumer {
				return func(v []int64) error {
					calls++
					require.Equal(t, testInt64SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt64SliceConsumerError
				}
			},
			cf2: func(t *testing.T) Int64SliceConsumer {
				return func(v []int64) error {
					calls++
					require.Equal(t, testInt64SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int64SliceConsumer {
				return func(v []int64) error {
					calls++
					require.Equal(t, testInt64SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64SliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentInt64SliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentInt64SliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentInt64SliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testInt64SliceConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt64SliceConsumer(t *testing.T) {
	var sc SilentInt64SliceConsumer = func(v []int64) {
		require.Equal(t, testInt64SliceConsumerValue, v)
		return
	}
	sc(testInt64SliceConsumerValue)
}

func TestMustInt64SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt64SliceConsumerFactory
		cf2   testInt64SliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int64SliceConsumer {
				return func(v []int64) error {
					calls++
					require.Equal(t, testInt64SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64SliceConsumer {
				return func(v []int64) error {
					calls++
					require.Equal(t, testInt64SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int64SliceConsumer {
				return func(v []int64) error {
					calls++
					require.Equal(t, testInt64SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt64SliceConsumerError
				}
			},
			cf2: func(t *testing.T) Int64SliceConsumer {
				return func(v []int64) error {
					calls++
					require.Equal(t, testInt64SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int64SliceConsumer {
				return func(v []int64) error {
					calls++
					require.Equal(t, testInt64SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64SliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustInt64SliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustInt64SliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustInt64SliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testInt64SliceConsumerError.Error(), func() {
					cmc(testInt64SliceConsumerValue)
				})
			} else {
				cmc(testInt64SliceConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt64SliceConsumer_ToSilentInt64SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64SliceConsumer {
				return func(v []int64) error {
					require.Equal(t, testInt64SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64SliceConsumer {
				return func(v []int64) error {
					require.Equal(t, testInt64SliceConsumerValue, v)
					return testInt64SliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt64SliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentInt64SliceConsumer()
			r.NotNil(sc)

			sc(testInt64SliceConsumerValue)
		})
	}
}

func TestMustInt64SliceConsumer_ToInt64SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64SliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64SliceConsumer {
				return func(v []int64) error {
					require.Equal(t, testInt64SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64SliceConsumer {
				return func(v []int64) error {
					require.Equal(t, testInt64SliceConsumerValue, v)
					return testInt64SliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt64SliceConsumer()
			r.NotNil(mc)

			c = mc.ToInt64SliceConsumer()
			r.NotNil(c)

			err := c(testInt64SliceConsumerValue)
			if tt.err {
				r.EqualError(err, testInt64SliceConsumerError.Error())
			}
		})
	}
}
