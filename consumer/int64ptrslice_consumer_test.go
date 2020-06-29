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
	testInt64PtrSliceConsumerValue []*int64
	testInt64PtrSliceConsumerError = errors.New("error")
)

type testInt64PtrSliceConsumerFactory func(t *testing.T) Int64PtrSliceConsumer

func TestInt64PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64PtrSliceConsumer {
				return func(v []*int64) error {
					require.Equal(t, testInt64PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64PtrSliceConsumer {
				return func(v []*int64) error {
					require.Equal(t, testInt64PtrSliceConsumerValue, v)
					return testInt64PtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testInt64PtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testInt64PtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt64PtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64PtrSliceConsumer {
				return func(v []*int64) error {
					require.Equal(t, testInt64PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64PtrSliceConsumer {
				return func(v []*int64) error {
					require.Equal(t, testInt64PtrSliceConsumerValue, v)
					return testInt64PtrSliceConsumerError
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

			err := c(testInt64PtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testInt64PtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt64PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt64PtrSliceConsumerFactory
		cf2   testInt64PtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int64PtrSliceConsumer {
				return func(v []*int64) error {
					calls++
					require.Equal(t, testInt64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64PtrSliceConsumer {
				return func(v []*int64) error {
					calls++
					require.Equal(t, testInt64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int64PtrSliceConsumer {
				return func(v []*int64) error {
					calls++
					require.Equal(t, testInt64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt64PtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) Int64PtrSliceConsumer {
				return func(v []*int64) error {
					calls++
					require.Equal(t, testInt64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int64PtrSliceConsumer {
				return func(v []*int64) error {
					calls++
					require.Equal(t, testInt64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64PtrSliceConsumer {
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
			err := cc(testInt64PtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testInt64PtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestInt64PtrSliceConsumer_ToSilentInt64PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64PtrSliceConsumer {
				return func(v []*int64) error {
					require.Equal(t, testInt64PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64PtrSliceConsumer {
				return func(v []*int64) error {
					require.Equal(t, testInt64PtrSliceConsumerValue, v)
					return testInt64PtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentInt64PtrSliceConsumer()
			r.NotNil(sc)

			sc(testInt64PtrSliceConsumerValue)
		})
	}
}

func TestInt64PtrSliceConsumer_ToMustInt64PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64PtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64PtrSliceConsumer {
				return func(v []*int64) error {
					require.Equal(t, testInt64PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64PtrSliceConsumer {
				return func(v []*int64) error {
					require.Equal(t, testInt64PtrSliceConsumerValue, v)
					return testInt64PtrSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustInt64PtrSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testInt64PtrSliceConsumerError.Error(), func() {
					mc(testInt64PtrSliceConsumerValue)
				})
			} else {
				mc(testInt64PtrSliceConsumerValue)
			}
		})
	}
}

func TestSilentInt64PtrSliceConsumer(t *testing.T) {
	var sc SilentInt64PtrSliceConsumer = func(v []*int64) {
		require.Equal(t, testInt64PtrSliceConsumerValue, v)
		return
	}
	sc(testInt64PtrSliceConsumerValue)
}

func TestSilentInt64PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt64PtrSliceConsumerFactory
		cf2   testInt64PtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int64PtrSliceConsumer {
				return func(v []*int64) error {
					calls++
					require.Equal(t, testInt64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64PtrSliceConsumer {
				return func(v []*int64) error {
					calls++
					require.Equal(t, testInt64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int64PtrSliceConsumer {
				return func(v []*int64) error {
					calls++
					require.Equal(t, testInt64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt64PtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) Int64PtrSliceConsumer {
				return func(v []*int64) error {
					calls++
					require.Equal(t, testInt64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int64PtrSliceConsumer {
				return func(v []*int64) error {
					calls++
					require.Equal(t, testInt64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64PtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentInt64PtrSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentInt64PtrSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentInt64PtrSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testInt64PtrSliceConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt64PtrSliceConsumer(t *testing.T) {
	var sc SilentInt64PtrSliceConsumer = func(v []*int64) {
		require.Equal(t, testInt64PtrSliceConsumerValue, v)
		return
	}
	sc(testInt64PtrSliceConsumerValue)
}

func TestMustInt64PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt64PtrSliceConsumerFactory
		cf2   testInt64PtrSliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int64PtrSliceConsumer {
				return func(v []*int64) error {
					calls++
					require.Equal(t, testInt64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64PtrSliceConsumer {
				return func(v []*int64) error {
					calls++
					require.Equal(t, testInt64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int64PtrSliceConsumer {
				return func(v []*int64) error {
					calls++
					require.Equal(t, testInt64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt64PtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) Int64PtrSliceConsumer {
				return func(v []*int64) error {
					calls++
					require.Equal(t, testInt64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int64PtrSliceConsumer {
				return func(v []*int64) error {
					calls++
					require.Equal(t, testInt64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64PtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustInt64PtrSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustInt64PtrSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustInt64PtrSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testInt64PtrSliceConsumerError.Error(), func() {
					cmc(testInt64PtrSliceConsumerValue)
				})
			} else {
				cmc(testInt64PtrSliceConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt64PtrSliceConsumer_ToSilentInt64PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64PtrSliceConsumer {
				return func(v []*int64) error {
					require.Equal(t, testInt64PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64PtrSliceConsumer {
				return func(v []*int64) error {
					require.Equal(t, testInt64PtrSliceConsumerValue, v)
					return testInt64PtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt64PtrSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentInt64PtrSliceConsumer()
			r.NotNil(sc)

			sc(testInt64PtrSliceConsumerValue)
		})
	}
}

func TestMustInt64PtrSliceConsumer_ToInt64PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64PtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64PtrSliceConsumer {
				return func(v []*int64) error {
					require.Equal(t, testInt64PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64PtrSliceConsumer {
				return func(v []*int64) error {
					require.Equal(t, testInt64PtrSliceConsumerValue, v)
					return testInt64PtrSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt64PtrSliceConsumer()
			r.NotNil(mc)

			c = mc.ToInt64PtrSliceConsumer()
			r.NotNil(c)

			err := c(testInt64PtrSliceConsumerValue)
			if tt.err {
				r.EqualError(err, testInt64PtrSliceConsumerError.Error())
			}
		})
	}
}
