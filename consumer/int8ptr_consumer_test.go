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
	testInt8PtrConsumerValue *int8
	testInt8PtrConsumerError = errors.New("error")
)

type testInt8PtrConsumerFactory func(t *testing.T) Int8PtrConsumer

func TestInt8PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8PtrConsumer {
				return func(v *int8) error {
					require.Equal(t, testInt8PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8PtrConsumer {
				return func(v *int8) error {
					require.Equal(t, testInt8PtrConsumerValue, v)
					return testInt8PtrConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testInt8PtrConsumerValue)
			if err != nil {
				r.EqualError(err, testInt8PtrConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt8PtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8PtrConsumer {
				return func(v *int8) error {
					require.Equal(t, testInt8PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8PtrConsumer {
				return func(v *int8) error {
					require.Equal(t, testInt8PtrConsumerValue, v)
					return testInt8PtrConsumerError
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

			err := c(testInt8PtrConsumerValue)
			if err != nil {
				r.EqualError(err, testInt8PtrConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt8PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt8PtrConsumerFactory
		cf2   testInt8PtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int8PtrConsumer {
				return func(v *int8) error {
					calls++
					require.Equal(t, testInt8PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8PtrConsumer {
				return func(v *int8) error {
					calls++
					require.Equal(t, testInt8PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int8PtrConsumer {
				return func(v *int8) error {
					calls++
					require.Equal(t, testInt8PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt8PtrConsumerError
				}
			},
			cf2: func(t *testing.T) Int8PtrConsumer {
				return func(v *int8) error {
					calls++
					require.Equal(t, testInt8PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int8PtrConsumer {
				return func(v *int8) error {
					calls++
					require.Equal(t, testInt8PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8PtrConsumer {
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
			err := cc(testInt8PtrConsumerValue)
			if err != nil {
				r.EqualError(err, testInt8PtrConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestInt8PtrConsumer_ToSilentInt8PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8PtrConsumer {
				return func(v *int8) error {
					require.Equal(t, testInt8PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8PtrConsumer {
				return func(v *int8) error {
					require.Equal(t, testInt8PtrConsumerValue, v)
					return testInt8PtrConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentInt8PtrConsumer()
			r.NotNil(sc)

			sc(testInt8PtrConsumerValue)
		})
	}
}

func TestInt8PtrConsumer_ToMustInt8PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8PtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8PtrConsumer {
				return func(v *int8) error {
					require.Equal(t, testInt8PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8PtrConsumer {
				return func(v *int8) error {
					require.Equal(t, testInt8PtrConsumerValue, v)
					return testInt8PtrConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustInt8PtrConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testInt8PtrConsumerError.Error(), func() {
					mc(testInt8PtrConsumerValue)
				})
			} else {
				mc(testInt8PtrConsumerValue)
			}
		})
	}
}

func TestSilentInt8PtrConsumer(t *testing.T) {
	var sc SilentInt8PtrConsumer = func(v *int8) {
		require.Equal(t, testInt8PtrConsumerValue, v)
		return
	}
	sc(testInt8PtrConsumerValue)
}

func TestSilentInt8PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt8PtrConsumerFactory
		cf2   testInt8PtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int8PtrConsumer {
				return func(v *int8) error {
					calls++
					require.Equal(t, testInt8PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8PtrConsumer {
				return func(v *int8) error {
					calls++
					require.Equal(t, testInt8PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int8PtrConsumer {
				return func(v *int8) error {
					calls++
					require.Equal(t, testInt8PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt8PtrConsumerError
				}
			},
			cf2: func(t *testing.T) Int8PtrConsumer {
				return func(v *int8) error {
					calls++
					require.Equal(t, testInt8PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int8PtrConsumer {
				return func(v *int8) error {
					calls++
					require.Equal(t, testInt8PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8PtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentInt8PtrConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentInt8PtrConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentInt8PtrConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testInt8PtrConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt8PtrConsumer(t *testing.T) {
	var sc SilentInt8PtrConsumer = func(v *int8) {
		require.Equal(t, testInt8PtrConsumerValue, v)
		return
	}
	sc(testInt8PtrConsumerValue)
}

func TestMustInt8PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt8PtrConsumerFactory
		cf2   testInt8PtrConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int8PtrConsumer {
				return func(v *int8) error {
					calls++
					require.Equal(t, testInt8PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8PtrConsumer {
				return func(v *int8) error {
					calls++
					require.Equal(t, testInt8PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int8PtrConsumer {
				return func(v *int8) error {
					calls++
					require.Equal(t, testInt8PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt8PtrConsumerError
				}
			},
			cf2: func(t *testing.T) Int8PtrConsumer {
				return func(v *int8) error {
					calls++
					require.Equal(t, testInt8PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int8PtrConsumer {
				return func(v *int8) error {
					calls++
					require.Equal(t, testInt8PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8PtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustInt8PtrConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustInt8PtrConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustInt8PtrConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testInt8PtrConsumerError.Error(), func() {
					cmc(testInt8PtrConsumerValue)
				})
			} else {
				cmc(testInt8PtrConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt8PtrConsumer_ToSilentInt8PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8PtrConsumer {
				return func(v *int8) error {
					require.Equal(t, testInt8PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8PtrConsumer {
				return func(v *int8) error {
					require.Equal(t, testInt8PtrConsumerValue, v)
					return testInt8PtrConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt8PtrConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentInt8PtrConsumer()
			r.NotNil(sc)

			sc(testInt8PtrConsumerValue)
		})
	}
}

func TestMustInt8PtrConsumer_ToInt8PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8PtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8PtrConsumer {
				return func(v *int8) error {
					require.Equal(t, testInt8PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8PtrConsumer {
				return func(v *int8) error {
					require.Equal(t, testInt8PtrConsumerValue, v)
					return testInt8PtrConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt8PtrConsumer()
			r.NotNil(mc)

			c = mc.ToInt8PtrConsumer()
			r.NotNil(c)

			err := c(testInt8PtrConsumerValue)
			if tt.err {
				r.EqualError(err, testInt8PtrConsumerError.Error())
			}
		})
	}
}
