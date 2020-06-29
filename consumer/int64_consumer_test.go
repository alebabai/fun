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
	testInt64ConsumerValue int64
	testInt64ConsumerError = errors.New("error")
)

type testInt64ConsumerFactory func(t *testing.T) Int64Consumer

func TestInt64Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					require.Equal(t, testInt64ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					require.Equal(t, testInt64ConsumerValue, v)
					return testInt64ConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testInt64ConsumerValue)
			if err != nil {
				r.EqualError(err, testInt64ConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt64Supplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					require.Equal(t, testInt64ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					require.Equal(t, testInt64ConsumerValue, v)
					return testInt64ConsumerError
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

			err := c(testInt64ConsumerValue)
			if err != nil {
				r.EqualError(err, testInt64ConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt64Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt64ConsumerFactory
		cf2   testInt64ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, testInt64ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, testInt64ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, testInt64ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt64ConsumerError
				}
			},
			cf2: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, testInt64ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, testInt64ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64Consumer {
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
			err := cc(testInt64ConsumerValue)
			if err != nil {
				r.EqualError(err, testInt64ConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestInt64Consumer_ToSilentInt64Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					require.Equal(t, testInt64ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					require.Equal(t, testInt64ConsumerValue, v)
					return testInt64ConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentInt64Consumer()
			r.NotNil(sc)

			sc(testInt64ConsumerValue)
		})
	}
}

func TestInt64Consumer_ToMustInt64Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					require.Equal(t, testInt64ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					require.Equal(t, testInt64ConsumerValue, v)
					return testInt64ConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustInt64Consumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testInt64ConsumerError.Error(), func() {
					mc(testInt64ConsumerValue)
				})
			} else {
				mc(testInt64ConsumerValue)
			}
		})
	}
}

func TestSilentInt64Consumer(t *testing.T) {
	var sc SilentInt64Consumer = func(v int64) {
		require.Equal(t, testInt64ConsumerValue, v)
		return
	}
	sc(testInt64ConsumerValue)
}

func TestSilentInt64Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt64ConsumerFactory
		cf2   testInt64ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, testInt64ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, testInt64ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, testInt64ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt64ConsumerError
				}
			},
			cf2: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, testInt64ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, testInt64ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentInt64Consumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentInt64Consumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentInt64Consumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testInt64ConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt64Consumer(t *testing.T) {
	var sc SilentInt64Consumer = func(v int64) {
		require.Equal(t, testInt64ConsumerValue, v)
		return
	}
	sc(testInt64ConsumerValue)
}

func TestMustInt64Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt64ConsumerFactory
		cf2   testInt64ConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, testInt64ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, testInt64ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, testInt64ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt64ConsumerError
				}
			},
			cf2: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, testInt64ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					calls++
					require.Equal(t, testInt64ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int64Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustInt64Consumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustInt64Consumer = nil
			if c2 != nil {
				mc2 = c2.ToMustInt64Consumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testInt64ConsumerError.Error(), func() {
					cmc(testInt64ConsumerValue)
				})
			} else {
				cmc(testInt64ConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt64Consumer_ToSilentInt64Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					require.Equal(t, testInt64ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					require.Equal(t, testInt64ConsumerValue, v)
					return testInt64ConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt64Consumer()
			r.NotNil(mc)

			sc := mc.ToSilentInt64Consumer()
			r.NotNil(sc)

			sc(testInt64ConsumerValue)
		})
	}
}

func TestMustInt64Consumer_ToInt64Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt64ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					require.Equal(t, testInt64ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int64Consumer {
				return func(v int64) error {
					require.Equal(t, testInt64ConsumerValue, v)
					return testInt64ConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt64Consumer()
			r.NotNil(mc)

			c = mc.ToInt64Consumer()
			r.NotNil(c)

			err := c(testInt64ConsumerValue)
			if tt.err {
				r.EqualError(err, testInt64ConsumerError.Error())
			}
		})
	}
}
