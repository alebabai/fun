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
	testInt32PtrConsumerValue *int32
	testInt32PtrConsumerError = errors.New("error")
)

type testInt32PtrConsumerFactory func(t *testing.T) Int32PtrConsumer

func TestInt32PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					require.Equal(t, testInt32PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					require.Equal(t, testInt32PtrConsumerValue, v)
					return testInt32PtrConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testInt32PtrConsumerValue)
			if err != nil {
				r.EqualError(err, testInt32PtrConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt32PtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					require.Equal(t, testInt32PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					require.Equal(t, testInt32PtrConsumerValue, v)
					return testInt32PtrConsumerError
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

			err := c(testInt32PtrConsumerValue)
			if err != nil {
				r.EqualError(err, testInt32PtrConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt32PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt32PtrConsumerFactory
		cf2   testInt32PtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, testInt32PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, testInt32PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, testInt32PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt32PtrConsumerError
				}
			},
			cf2: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, testInt32PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, testInt32PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32PtrConsumer {
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
			err := cc(testInt32PtrConsumerValue)
			if err != nil {
				r.EqualError(err, testInt32PtrConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestInt32PtrConsumer_ToSilentInt32PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					require.Equal(t, testInt32PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					require.Equal(t, testInt32PtrConsumerValue, v)
					return testInt32PtrConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentInt32PtrConsumer()
			r.NotNil(sc)

			sc(testInt32PtrConsumerValue)
		})
	}
}

func TestInt32PtrConsumer_ToMustInt32PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32PtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					require.Equal(t, testInt32PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					require.Equal(t, testInt32PtrConsumerValue, v)
					return testInt32PtrConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustInt32PtrConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testInt32PtrConsumerError.Error(), func() {
					mc(testInt32PtrConsumerValue)
				})
			} else {
				mc(testInt32PtrConsumerValue)
			}
		})
	}
}

func TestSilentInt32PtrConsumer(t *testing.T) {
	var sc SilentInt32PtrConsumer = func(v *int32) {
		require.Equal(t, testInt32PtrConsumerValue, v)
		return
	}
	sc(testInt32PtrConsumerValue)
}

func TestSilentInt32PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt32PtrConsumerFactory
		cf2   testInt32PtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, testInt32PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, testInt32PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, testInt32PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt32PtrConsumerError
				}
			},
			cf2: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, testInt32PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, testInt32PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32PtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentInt32PtrConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentInt32PtrConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentInt32PtrConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testInt32PtrConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt32PtrConsumer(t *testing.T) {
	var sc SilentInt32PtrConsumer = func(v *int32) {
		require.Equal(t, testInt32PtrConsumerValue, v)
		return
	}
	sc(testInt32PtrConsumerValue)
}

func TestMustInt32PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt32PtrConsumerFactory
		cf2   testInt32PtrConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, testInt32PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, testInt32PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, testInt32PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt32PtrConsumerError
				}
			},
			cf2: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, testInt32PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					calls++
					require.Equal(t, testInt32PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32PtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustInt32PtrConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustInt32PtrConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustInt32PtrConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testInt32PtrConsumerError.Error(), func() {
					cmc(testInt32PtrConsumerValue)
				})
			} else {
				cmc(testInt32PtrConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt32PtrConsumer_ToSilentInt32PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					require.Equal(t, testInt32PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					require.Equal(t, testInt32PtrConsumerValue, v)
					return testInt32PtrConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt32PtrConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentInt32PtrConsumer()
			r.NotNil(sc)

			sc(testInt32PtrConsumerValue)
		})
	}
}

func TestMustInt32PtrConsumer_ToInt32PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32PtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					require.Equal(t, testInt32PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32PtrConsumer {
				return func(v *int32) error {
					require.Equal(t, testInt32PtrConsumerValue, v)
					return testInt32PtrConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt32PtrConsumer()
			r.NotNil(mc)

			c = mc.ToInt32PtrConsumer()
			r.NotNil(c)

			err := c(testInt32PtrConsumerValue)
			if tt.err {
				r.EqualError(err, testInt32PtrConsumerError.Error())
			}
		})
	}
}
