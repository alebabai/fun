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
	testInt16SliceConsumerValue []int16
	testInt16SliceConsumerError = errors.New("error")
)

type testInt16SliceConsumerFactory func(t *testing.T) Int16SliceConsumer

func TestInt16SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt16SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16SliceConsumer {
				return func(v []int16) error {
					require.Equal(t, testInt16SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int16SliceConsumer {
				return func(v []int16) error {
					require.Equal(t, testInt16SliceConsumerValue, v)
					return testInt16SliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testInt16SliceConsumerValue)
			if err != nil {
				r.EqualError(err, testInt16SliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt16SliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt16SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16SliceConsumer {
				return func(v []int16) error {
					require.Equal(t, testInt16SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int16SliceConsumer {
				return func(v []int16) error {
					require.Equal(t, testInt16SliceConsumerValue, v)
					return testInt16SliceConsumerError
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

			err := c(testInt16SliceConsumerValue)
			if err != nil {
				r.EqualError(err, testInt16SliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt16SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt16SliceConsumerFactory
		cf2   testInt16SliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int16SliceConsumer {
				return func(v []int16) error {
					calls++
					require.Equal(t, testInt16SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16SliceConsumer {
				return func(v []int16) error {
					calls++
					require.Equal(t, testInt16SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int16SliceConsumer {
				return func(v []int16) error {
					calls++
					require.Equal(t, testInt16SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt16SliceConsumerError
				}
			},
			cf2: func(t *testing.T) Int16SliceConsumer {
				return func(v []int16) error {
					calls++
					require.Equal(t, testInt16SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int16SliceConsumer {
				return func(v []int16) error {
					calls++
					require.Equal(t, testInt16SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16SliceConsumer {
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
			err := cc(testInt16SliceConsumerValue)
			if err != nil {
				r.EqualError(err, testInt16SliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestInt16SliceConsumer_ToSilentInt16SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt16SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16SliceConsumer {
				return func(v []int16) error {
					require.Equal(t, testInt16SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int16SliceConsumer {
				return func(v []int16) error {
					require.Equal(t, testInt16SliceConsumerValue, v)
					return testInt16SliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentInt16SliceConsumer()
			r.NotNil(sc)

			sc(testInt16SliceConsumerValue)
		})
	}
}

func TestInt16SliceConsumer_ToMustInt16SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt16SliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16SliceConsumer {
				return func(v []int16) error {
					require.Equal(t, testInt16SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int16SliceConsumer {
				return func(v []int16) error {
					require.Equal(t, testInt16SliceConsumerValue, v)
					return testInt16SliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustInt16SliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testInt16SliceConsumerError.Error(), func() {
					mc(testInt16SliceConsumerValue)
				})
			} else {
				mc(testInt16SliceConsumerValue)
			}
		})
	}
}

func TestSilentInt16SliceConsumer(t *testing.T) {
	var sc SilentInt16SliceConsumer = func(v []int16) {
		require.Equal(t, testInt16SliceConsumerValue, v)
		return
	}
	sc(testInt16SliceConsumerValue)
}

func TestSilentInt16SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt16SliceConsumerFactory
		cf2   testInt16SliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int16SliceConsumer {
				return func(v []int16) error {
					calls++
					require.Equal(t, testInt16SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16SliceConsumer {
				return func(v []int16) error {
					calls++
					require.Equal(t, testInt16SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int16SliceConsumer {
				return func(v []int16) error {
					calls++
					require.Equal(t, testInt16SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt16SliceConsumerError
				}
			},
			cf2: func(t *testing.T) Int16SliceConsumer {
				return func(v []int16) error {
					calls++
					require.Equal(t, testInt16SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int16SliceConsumer {
				return func(v []int16) error {
					calls++
					require.Equal(t, testInt16SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16SliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentInt16SliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentInt16SliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentInt16SliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testInt16SliceConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt16SliceConsumer(t *testing.T) {
	var sc SilentInt16SliceConsumer = func(v []int16) {
		require.Equal(t, testInt16SliceConsumerValue, v)
		return
	}
	sc(testInt16SliceConsumerValue)
}

func TestMustInt16SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt16SliceConsumerFactory
		cf2   testInt16SliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int16SliceConsumer {
				return func(v []int16) error {
					calls++
					require.Equal(t, testInt16SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16SliceConsumer {
				return func(v []int16) error {
					calls++
					require.Equal(t, testInt16SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int16SliceConsumer {
				return func(v []int16) error {
					calls++
					require.Equal(t, testInt16SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt16SliceConsumerError
				}
			},
			cf2: func(t *testing.T) Int16SliceConsumer {
				return func(v []int16) error {
					calls++
					require.Equal(t, testInt16SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int16SliceConsumer {
				return func(v []int16) error {
					calls++
					require.Equal(t, testInt16SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16SliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustInt16SliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustInt16SliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustInt16SliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testInt16SliceConsumerError.Error(), func() {
					cmc(testInt16SliceConsumerValue)
				})
			} else {
				cmc(testInt16SliceConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt16SliceConsumer_ToSilentInt16SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt16SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16SliceConsumer {
				return func(v []int16) error {
					require.Equal(t, testInt16SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int16SliceConsumer {
				return func(v []int16) error {
					require.Equal(t, testInt16SliceConsumerValue, v)
					return testInt16SliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt16SliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentInt16SliceConsumer()
			r.NotNil(sc)

			sc(testInt16SliceConsumerValue)
		})
	}
}

func TestMustInt16SliceConsumer_ToInt16SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt16SliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16SliceConsumer {
				return func(v []int16) error {
					require.Equal(t, testInt16SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int16SliceConsumer {
				return func(v []int16) error {
					require.Equal(t, testInt16SliceConsumerValue, v)
					return testInt16SliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt16SliceConsumer()
			r.NotNil(mc)

			c = mc.ToInt16SliceConsumer()
			r.NotNil(c)

			err := c(testInt16SliceConsumerValue)
			if tt.err {
				r.EqualError(err, testInt16SliceConsumerError.Error())
			}
		})
	}
}
