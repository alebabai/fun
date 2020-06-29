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
	testInt8PtrSliceConsumerValue []*int8
	testInt8PtrSliceConsumerError = errors.New("error")
)

type testInt8PtrSliceConsumerFactory func(t *testing.T) Int8PtrSliceConsumer

func TestInt8PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					require.Equal(t, testInt8PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					require.Equal(t, testInt8PtrSliceConsumerValue, v)
					return testInt8PtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testInt8PtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testInt8PtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt8PtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					require.Equal(t, testInt8PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					require.Equal(t, testInt8PtrSliceConsumerValue, v)
					return testInt8PtrSliceConsumerError
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

			err := c(testInt8PtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testInt8PtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt8PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt8PtrSliceConsumerFactory
		cf2   testInt8PtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, testInt8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, testInt8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, testInt8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt8PtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, testInt8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, testInt8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8PtrSliceConsumer {
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
			err := cc(testInt8PtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testInt8PtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestInt8PtrSliceConsumer_ToSilentInt8PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					require.Equal(t, testInt8PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					require.Equal(t, testInt8PtrSliceConsumerValue, v)
					return testInt8PtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentInt8PtrSliceConsumer()
			r.NotNil(sc)

			sc(testInt8PtrSliceConsumerValue)
		})
	}
}

func TestInt8PtrSliceConsumer_ToMustInt8PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8PtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					require.Equal(t, testInt8PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					require.Equal(t, testInt8PtrSliceConsumerValue, v)
					return testInt8PtrSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustInt8PtrSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testInt8PtrSliceConsumerError.Error(), func() {
					mc(testInt8PtrSliceConsumerValue)
				})
			} else {
				mc(testInt8PtrSliceConsumerValue)
			}
		})
	}
}

func TestSilentInt8PtrSliceConsumer(t *testing.T) {
	var sc SilentInt8PtrSliceConsumer = func(v []*int8) {
		require.Equal(t, testInt8PtrSliceConsumerValue, v)
		return
	}
	sc(testInt8PtrSliceConsumerValue)
}

func TestSilentInt8PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt8PtrSliceConsumerFactory
		cf2   testInt8PtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, testInt8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, testInt8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, testInt8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt8PtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, testInt8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, testInt8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8PtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentInt8PtrSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentInt8PtrSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentInt8PtrSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testInt8PtrSliceConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt8PtrSliceConsumer(t *testing.T) {
	var sc SilentInt8PtrSliceConsumer = func(v []*int8) {
		require.Equal(t, testInt8PtrSliceConsumerValue, v)
		return
	}
	sc(testInt8PtrSliceConsumerValue)
}

func TestMustInt8PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt8PtrSliceConsumerFactory
		cf2   testInt8PtrSliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, testInt8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, testInt8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, testInt8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt8PtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, testInt8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					calls++
					require.Equal(t, testInt8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8PtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustInt8PtrSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustInt8PtrSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustInt8PtrSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testInt8PtrSliceConsumerError.Error(), func() {
					cmc(testInt8PtrSliceConsumerValue)
				})
			} else {
				cmc(testInt8PtrSliceConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt8PtrSliceConsumer_ToSilentInt8PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					require.Equal(t, testInt8PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					require.Equal(t, testInt8PtrSliceConsumerValue, v)
					return testInt8PtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt8PtrSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentInt8PtrSliceConsumer()
			r.NotNil(sc)

			sc(testInt8PtrSliceConsumerValue)
		})
	}
}

func TestMustInt8PtrSliceConsumer_ToInt8PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8PtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					require.Equal(t, testInt8PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8PtrSliceConsumer {
				return func(v []*int8) error {
					require.Equal(t, testInt8PtrSliceConsumerValue, v)
					return testInt8PtrSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt8PtrSliceConsumer()
			r.NotNil(mc)

			c = mc.ToInt8PtrSliceConsumer()
			r.NotNil(c)

			err := c(testInt8PtrSliceConsumerValue)
			if tt.err {
				r.EqualError(err, testInt8PtrSliceConsumerError.Error())
			}
		})
	}
}
