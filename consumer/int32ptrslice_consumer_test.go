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
	testInt32PtrSliceConsumerValue []*int32
	testInt32PtrSliceConsumerError = errors.New("error")
)

type testInt32PtrSliceConsumerFactory func(t *testing.T) Int32PtrSliceConsumer

func TestInt32PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32PtrSliceConsumer {
				return func(v []*int32) error {
					require.Equal(t, testInt32PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32PtrSliceConsumer {
				return func(v []*int32) error {
					require.Equal(t, testInt32PtrSliceConsumerValue, v)
					return testInt32PtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testInt32PtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testInt32PtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt32PtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32PtrSliceConsumer {
				return func(v []*int32) error {
					require.Equal(t, testInt32PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32PtrSliceConsumer {
				return func(v []*int32) error {
					require.Equal(t, testInt32PtrSliceConsumerValue, v)
					return testInt32PtrSliceConsumerError
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

			err := c(testInt32PtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testInt32PtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt32PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt32PtrSliceConsumerFactory
		cf2   testInt32PtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int32PtrSliceConsumer {
				return func(v []*int32) error {
					calls++
					require.Equal(t, testInt32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32PtrSliceConsumer {
				return func(v []*int32) error {
					calls++
					require.Equal(t, testInt32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int32PtrSliceConsumer {
				return func(v []*int32) error {
					calls++
					require.Equal(t, testInt32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt32PtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) Int32PtrSliceConsumer {
				return func(v []*int32) error {
					calls++
					require.Equal(t, testInt32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int32PtrSliceConsumer {
				return func(v []*int32) error {
					calls++
					require.Equal(t, testInt32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32PtrSliceConsumer {
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
			err := cc(testInt32PtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testInt32PtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestInt32PtrSliceConsumer_ToSilentInt32PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32PtrSliceConsumer {
				return func(v []*int32) error {
					require.Equal(t, testInt32PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32PtrSliceConsumer {
				return func(v []*int32) error {
					require.Equal(t, testInt32PtrSliceConsumerValue, v)
					return testInt32PtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentInt32PtrSliceConsumer()
			r.NotNil(sc)

			sc(testInt32PtrSliceConsumerValue)
		})
	}
}

func TestInt32PtrSliceConsumer_ToMustInt32PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32PtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32PtrSliceConsumer {
				return func(v []*int32) error {
					require.Equal(t, testInt32PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32PtrSliceConsumer {
				return func(v []*int32) error {
					require.Equal(t, testInt32PtrSliceConsumerValue, v)
					return testInt32PtrSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustInt32PtrSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testInt32PtrSliceConsumerError.Error(), func() {
					mc(testInt32PtrSliceConsumerValue)
				})
			} else {
				mc(testInt32PtrSliceConsumerValue)
			}
		})
	}
}

func TestSilentInt32PtrSliceConsumer(t *testing.T) {
	var sc SilentInt32PtrSliceConsumer = func(v []*int32) {
		require.Equal(t, testInt32PtrSliceConsumerValue, v)
		return
	}
	sc(testInt32PtrSliceConsumerValue)
}

func TestSilentInt32PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt32PtrSliceConsumerFactory
		cf2   testInt32PtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int32PtrSliceConsumer {
				return func(v []*int32) error {
					calls++
					require.Equal(t, testInt32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32PtrSliceConsumer {
				return func(v []*int32) error {
					calls++
					require.Equal(t, testInt32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int32PtrSliceConsumer {
				return func(v []*int32) error {
					calls++
					require.Equal(t, testInt32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt32PtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) Int32PtrSliceConsumer {
				return func(v []*int32) error {
					calls++
					require.Equal(t, testInt32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int32PtrSliceConsumer {
				return func(v []*int32) error {
					calls++
					require.Equal(t, testInt32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32PtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentInt32PtrSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentInt32PtrSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentInt32PtrSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testInt32PtrSliceConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt32PtrSliceConsumer(t *testing.T) {
	var sc SilentInt32PtrSliceConsumer = func(v []*int32) {
		require.Equal(t, testInt32PtrSliceConsumerValue, v)
		return
	}
	sc(testInt32PtrSliceConsumerValue)
}

func TestMustInt32PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt32PtrSliceConsumerFactory
		cf2   testInt32PtrSliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int32PtrSliceConsumer {
				return func(v []*int32) error {
					calls++
					require.Equal(t, testInt32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32PtrSliceConsumer {
				return func(v []*int32) error {
					calls++
					require.Equal(t, testInt32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int32PtrSliceConsumer {
				return func(v []*int32) error {
					calls++
					require.Equal(t, testInt32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt32PtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) Int32PtrSliceConsumer {
				return func(v []*int32) error {
					calls++
					require.Equal(t, testInt32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int32PtrSliceConsumer {
				return func(v []*int32) error {
					calls++
					require.Equal(t, testInt32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32PtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustInt32PtrSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustInt32PtrSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustInt32PtrSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testInt32PtrSliceConsumerError.Error(), func() {
					cmc(testInt32PtrSliceConsumerValue)
				})
			} else {
				cmc(testInt32PtrSliceConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt32PtrSliceConsumer_ToSilentInt32PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32PtrSliceConsumer {
				return func(v []*int32) error {
					require.Equal(t, testInt32PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32PtrSliceConsumer {
				return func(v []*int32) error {
					require.Equal(t, testInt32PtrSliceConsumerValue, v)
					return testInt32PtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt32PtrSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentInt32PtrSliceConsumer()
			r.NotNil(sc)

			sc(testInt32PtrSliceConsumerValue)
		})
	}
}

func TestMustInt32PtrSliceConsumer_ToInt32PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32PtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32PtrSliceConsumer {
				return func(v []*int32) error {
					require.Equal(t, testInt32PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32PtrSliceConsumer {
				return func(v []*int32) error {
					require.Equal(t, testInt32PtrSliceConsumerValue, v)
					return testInt32PtrSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt32PtrSliceConsumer()
			r.NotNil(mc)

			c = mc.ToInt32PtrSliceConsumer()
			r.NotNil(c)

			err := c(testInt32PtrSliceConsumerValue)
			if tt.err {
				r.EqualError(err, testInt32PtrSliceConsumerError.Error())
			}
		})
	}
}
