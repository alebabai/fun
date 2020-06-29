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
	testInt32SliceConsumerValue []int32
	testInt32SliceConsumerError = errors.New("error")
)

type testInt32SliceConsumerFactory func(t *testing.T) Int32SliceConsumer

func TestInt32SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					require.Equal(t, testInt32SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					require.Equal(t, testInt32SliceConsumerValue, v)
					return testInt32SliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testInt32SliceConsumerValue)
			if err != nil {
				r.EqualError(err, testInt32SliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt32SliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					require.Equal(t, testInt32SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					require.Equal(t, testInt32SliceConsumerValue, v)
					return testInt32SliceConsumerError
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

			err := c(testInt32SliceConsumerValue)
			if err != nil {
				r.EqualError(err, testInt32SliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt32SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt32SliceConsumerFactory
		cf2   testInt32SliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, testInt32SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, testInt32SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, testInt32SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt32SliceConsumerError
				}
			},
			cf2: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, testInt32SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, testInt32SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32SliceConsumer {
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
			err := cc(testInt32SliceConsumerValue)
			if err != nil {
				r.EqualError(err, testInt32SliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestInt32SliceConsumer_ToSilentInt32SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					require.Equal(t, testInt32SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					require.Equal(t, testInt32SliceConsumerValue, v)
					return testInt32SliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentInt32SliceConsumer()
			r.NotNil(sc)

			sc(testInt32SliceConsumerValue)
		})
	}
}

func TestInt32SliceConsumer_ToMustInt32SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32SliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					require.Equal(t, testInt32SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					require.Equal(t, testInt32SliceConsumerValue, v)
					return testInt32SliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustInt32SliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testInt32SliceConsumerError.Error(), func() {
					mc(testInt32SliceConsumerValue)
				})
			} else {
				mc(testInt32SliceConsumerValue)
			}
		})
	}
}

func TestSilentInt32SliceConsumer(t *testing.T) {
	var sc SilentInt32SliceConsumer = func(v []int32) {
		require.Equal(t, testInt32SliceConsumerValue, v)
		return
	}
	sc(testInt32SliceConsumerValue)
}

func TestSilentInt32SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt32SliceConsumerFactory
		cf2   testInt32SliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, testInt32SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, testInt32SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, testInt32SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt32SliceConsumerError
				}
			},
			cf2: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, testInt32SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, testInt32SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32SliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentInt32SliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentInt32SliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentInt32SliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testInt32SliceConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt32SliceConsumer(t *testing.T) {
	var sc SilentInt32SliceConsumer = func(v []int32) {
		require.Equal(t, testInt32SliceConsumerValue, v)
		return
	}
	sc(testInt32SliceConsumerValue)
}

func TestMustInt32SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt32SliceConsumerFactory
		cf2   testInt32SliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, testInt32SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, testInt32SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, testInt32SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt32SliceConsumerError
				}
			},
			cf2: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, testInt32SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, testInt32SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32SliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustInt32SliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustInt32SliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustInt32SliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testInt32SliceConsumerError.Error(), func() {
					cmc(testInt32SliceConsumerValue)
				})
			} else {
				cmc(testInt32SliceConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt32SliceConsumer_ToSilentInt32SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					require.Equal(t, testInt32SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					require.Equal(t, testInt32SliceConsumerValue, v)
					return testInt32SliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt32SliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentInt32SliceConsumer()
			r.NotNil(sc)

			sc(testInt32SliceConsumerValue)
		})
	}
}

func TestMustInt32SliceConsumer_ToInt32SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32SliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					require.Equal(t, testInt32SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					require.Equal(t, testInt32SliceConsumerValue, v)
					return testInt32SliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt32SliceConsumer()
			r.NotNil(mc)

			c = mc.ToInt32SliceConsumer()
			r.NotNil(c)

			err := c(testInt32SliceConsumerValue)
			if tt.err {
				r.EqualError(err, testInt32SliceConsumerError.Error())
			}
		})
	}
}
