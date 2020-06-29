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
	testInt16PtrSliceConsumerValue []*int16
	testInt16PtrSliceConsumerError = errors.New("error")
)

type testInt16PtrSliceConsumerFactory func(t *testing.T) Int16PtrSliceConsumer

func TestInt16PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt16PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16PtrSliceConsumer {
				return func(v []*int16) error {
					require.Equal(t, testInt16PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int16PtrSliceConsumer {
				return func(v []*int16) error {
					require.Equal(t, testInt16PtrSliceConsumerValue, v)
					return testInt16PtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testInt16PtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testInt16PtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt16PtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt16PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16PtrSliceConsumer {
				return func(v []*int16) error {
					require.Equal(t, testInt16PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int16PtrSliceConsumer {
				return func(v []*int16) error {
					require.Equal(t, testInt16PtrSliceConsumerValue, v)
					return testInt16PtrSliceConsumerError
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

			err := c(testInt16PtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testInt16PtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt16PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt16PtrSliceConsumerFactory
		cf2   testInt16PtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int16PtrSliceConsumer {
				return func(v []*int16) error {
					calls++
					require.Equal(t, testInt16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16PtrSliceConsumer {
				return func(v []*int16) error {
					calls++
					require.Equal(t, testInt16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int16PtrSliceConsumer {
				return func(v []*int16) error {
					calls++
					require.Equal(t, testInt16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt16PtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) Int16PtrSliceConsumer {
				return func(v []*int16) error {
					calls++
					require.Equal(t, testInt16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int16PtrSliceConsumer {
				return func(v []*int16) error {
					calls++
					require.Equal(t, testInt16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16PtrSliceConsumer {
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
			err := cc(testInt16PtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testInt16PtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestInt16PtrSliceConsumer_ToSilentInt16PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt16PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16PtrSliceConsumer {
				return func(v []*int16) error {
					require.Equal(t, testInt16PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int16PtrSliceConsumer {
				return func(v []*int16) error {
					require.Equal(t, testInt16PtrSliceConsumerValue, v)
					return testInt16PtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentInt16PtrSliceConsumer()
			r.NotNil(sc)

			sc(testInt16PtrSliceConsumerValue)
		})
	}
}

func TestInt16PtrSliceConsumer_ToMustInt16PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt16PtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16PtrSliceConsumer {
				return func(v []*int16) error {
					require.Equal(t, testInt16PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int16PtrSliceConsumer {
				return func(v []*int16) error {
					require.Equal(t, testInt16PtrSliceConsumerValue, v)
					return testInt16PtrSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustInt16PtrSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testInt16PtrSliceConsumerError.Error(), func() {
					mc(testInt16PtrSliceConsumerValue)
				})
			} else {
				mc(testInt16PtrSliceConsumerValue)
			}
		})
	}
}

func TestSilentInt16PtrSliceConsumer(t *testing.T) {
	var sc SilentInt16PtrSliceConsumer = func(v []*int16) {
		require.Equal(t, testInt16PtrSliceConsumerValue, v)
		return
	}
	sc(testInt16PtrSliceConsumerValue)
}

func TestSilentInt16PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt16PtrSliceConsumerFactory
		cf2   testInt16PtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int16PtrSliceConsumer {
				return func(v []*int16) error {
					calls++
					require.Equal(t, testInt16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16PtrSliceConsumer {
				return func(v []*int16) error {
					calls++
					require.Equal(t, testInt16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int16PtrSliceConsumer {
				return func(v []*int16) error {
					calls++
					require.Equal(t, testInt16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt16PtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) Int16PtrSliceConsumer {
				return func(v []*int16) error {
					calls++
					require.Equal(t, testInt16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int16PtrSliceConsumer {
				return func(v []*int16) error {
					calls++
					require.Equal(t, testInt16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16PtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentInt16PtrSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentInt16PtrSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentInt16PtrSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testInt16PtrSliceConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt16PtrSliceConsumer(t *testing.T) {
	var sc SilentInt16PtrSliceConsumer = func(v []*int16) {
		require.Equal(t, testInt16PtrSliceConsumerValue, v)
		return
	}
	sc(testInt16PtrSliceConsumerValue)
}

func TestMustInt16PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt16PtrSliceConsumerFactory
		cf2   testInt16PtrSliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int16PtrSliceConsumer {
				return func(v []*int16) error {
					calls++
					require.Equal(t, testInt16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16PtrSliceConsumer {
				return func(v []*int16) error {
					calls++
					require.Equal(t, testInt16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int16PtrSliceConsumer {
				return func(v []*int16) error {
					calls++
					require.Equal(t, testInt16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt16PtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) Int16PtrSliceConsumer {
				return func(v []*int16) error {
					calls++
					require.Equal(t, testInt16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int16PtrSliceConsumer {
				return func(v []*int16) error {
					calls++
					require.Equal(t, testInt16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16PtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustInt16PtrSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustInt16PtrSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustInt16PtrSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testInt16PtrSliceConsumerError.Error(), func() {
					cmc(testInt16PtrSliceConsumerValue)
				})
			} else {
				cmc(testInt16PtrSliceConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt16PtrSliceConsumer_ToSilentInt16PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt16PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16PtrSliceConsumer {
				return func(v []*int16) error {
					require.Equal(t, testInt16PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int16PtrSliceConsumer {
				return func(v []*int16) error {
					require.Equal(t, testInt16PtrSliceConsumerValue, v)
					return testInt16PtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt16PtrSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentInt16PtrSliceConsumer()
			r.NotNil(sc)

			sc(testInt16PtrSliceConsumerValue)
		})
	}
}

func TestMustInt16PtrSliceConsumer_ToInt16PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt16PtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16PtrSliceConsumer {
				return func(v []*int16) error {
					require.Equal(t, testInt16PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int16PtrSliceConsumer {
				return func(v []*int16) error {
					require.Equal(t, testInt16PtrSliceConsumerValue, v)
					return testInt16PtrSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt16PtrSliceConsumer()
			r.NotNil(mc)

			c = mc.ToInt16PtrSliceConsumer()
			r.NotNil(c)

			err := c(testInt16PtrSliceConsumerValue)
			if tt.err {
				r.EqualError(err, testInt16PtrSliceConsumerError.Error())
			}
		})
	}
}
