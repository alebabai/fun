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
	testInt8SliceConsumerValue []int8
	testInt8SliceConsumerError = errors.New("error")
)

type testInt8SliceConsumerFactory func(t *testing.T) Int8SliceConsumer

func TestInt8SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					require.Equal(t, testInt8SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					require.Equal(t, testInt8SliceConsumerValue, v)
					return testInt8SliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testInt8SliceConsumerValue)
			if err != nil {
				r.EqualError(err, testInt8SliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt8SliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					require.Equal(t, testInt8SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					require.Equal(t, testInt8SliceConsumerValue, v)
					return testInt8SliceConsumerError
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

			err := c(testInt8SliceConsumerValue)
			if err != nil {
				r.EqualError(err, testInt8SliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt8SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt8SliceConsumerFactory
		cf2   testInt8SliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, testInt8SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, testInt8SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, testInt8SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt8SliceConsumerError
				}
			},
			cf2: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, testInt8SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, testInt8SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8SliceConsumer {
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
			err := cc(testInt8SliceConsumerValue)
			if err != nil {
				r.EqualError(err, testInt8SliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestInt8SliceConsumer_ToSilentInt8SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					require.Equal(t, testInt8SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					require.Equal(t, testInt8SliceConsumerValue, v)
					return testInt8SliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentInt8SliceConsumer()
			r.NotNil(sc)

			sc(testInt8SliceConsumerValue)
		})
	}
}

func TestInt8SliceConsumer_ToMustInt8SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8SliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					require.Equal(t, testInt8SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					require.Equal(t, testInt8SliceConsumerValue, v)
					return testInt8SliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustInt8SliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testInt8SliceConsumerError.Error(), func() {
					mc(testInt8SliceConsumerValue)
				})
			} else {
				mc(testInt8SliceConsumerValue)
			}
		})
	}
}

func TestSilentInt8SliceConsumer(t *testing.T) {
	var sc SilentInt8SliceConsumer = func(v []int8) {
		require.Equal(t, testInt8SliceConsumerValue, v)
		return
	}
	sc(testInt8SliceConsumerValue)
}

func TestSilentInt8SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt8SliceConsumerFactory
		cf2   testInt8SliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, testInt8SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, testInt8SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, testInt8SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt8SliceConsumerError
				}
			},
			cf2: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, testInt8SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, testInt8SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8SliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentInt8SliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentInt8SliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentInt8SliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testInt8SliceConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt8SliceConsumer(t *testing.T) {
	var sc SilentInt8SliceConsumer = func(v []int8) {
		require.Equal(t, testInt8SliceConsumerValue, v)
		return
	}
	sc(testInt8SliceConsumerValue)
}

func TestMustInt8SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt8SliceConsumerFactory
		cf2   testInt8SliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, testInt8SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, testInt8SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, testInt8SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt8SliceConsumerError
				}
			},
			cf2: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, testInt8SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					calls++
					require.Equal(t, testInt8SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8SliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustInt8SliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustInt8SliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustInt8SliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testInt8SliceConsumerError.Error(), func() {
					cmc(testInt8SliceConsumerValue)
				})
			} else {
				cmc(testInt8SliceConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt8SliceConsumer_ToSilentInt8SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					require.Equal(t, testInt8SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					require.Equal(t, testInt8SliceConsumerValue, v)
					return testInt8SliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt8SliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentInt8SliceConsumer()
			r.NotNil(sc)

			sc(testInt8SliceConsumerValue)
		})
	}
}

func TestMustInt8SliceConsumer_ToInt8SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8SliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					require.Equal(t, testInt8SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8SliceConsumer {
				return func(v []int8) error {
					require.Equal(t, testInt8SliceConsumerValue, v)
					return testInt8SliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt8SliceConsumer()
			r.NotNil(mc)

			c = mc.ToInt8SliceConsumer()
			r.NotNil(c)

			err := c(testInt8SliceConsumerValue)
			if tt.err {
				r.EqualError(err, testInt8SliceConsumerError.Error())
			}
		})
	}
}
