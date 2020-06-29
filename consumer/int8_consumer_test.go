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
	testInt8ConsumerValue int8
	testInt8ConsumerError = errors.New("error")
)

type testInt8ConsumerFactory func(t *testing.T) Int8Consumer

func TestInt8Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8Consumer {
				return func(v int8) error {
					require.Equal(t, testInt8ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8Consumer {
				return func(v int8) error {
					require.Equal(t, testInt8ConsumerValue, v)
					return testInt8ConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testInt8ConsumerValue)
			if err != nil {
				r.EqualError(err, testInt8ConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt8Supplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8Consumer {
				return func(v int8) error {
					require.Equal(t, testInt8ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8Consumer {
				return func(v int8) error {
					require.Equal(t, testInt8ConsumerValue, v)
					return testInt8ConsumerError
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

			err := c(testInt8ConsumerValue)
			if err != nil {
				r.EqualError(err, testInt8ConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt8Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt8ConsumerFactory
		cf2   testInt8ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int8Consumer {
				return func(v int8) error {
					calls++
					require.Equal(t, testInt8ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8Consumer {
				return func(v int8) error {
					calls++
					require.Equal(t, testInt8ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int8Consumer {
				return func(v int8) error {
					calls++
					require.Equal(t, testInt8ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt8ConsumerError
				}
			},
			cf2: func(t *testing.T) Int8Consumer {
				return func(v int8) error {
					calls++
					require.Equal(t, testInt8ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int8Consumer {
				return func(v int8) error {
					calls++
					require.Equal(t, testInt8ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8Consumer {
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
			err := cc(testInt8ConsumerValue)
			if err != nil {
				r.EqualError(err, testInt8ConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestInt8Consumer_ToSilentInt8Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8Consumer {
				return func(v int8) error {
					require.Equal(t, testInt8ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8Consumer {
				return func(v int8) error {
					require.Equal(t, testInt8ConsumerValue, v)
					return testInt8ConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentInt8Consumer()
			r.NotNil(sc)

			sc(testInt8ConsumerValue)
		})
	}
}

func TestInt8Consumer_ToMustInt8Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8Consumer {
				return func(v int8) error {
					require.Equal(t, testInt8ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8Consumer {
				return func(v int8) error {
					require.Equal(t, testInt8ConsumerValue, v)
					return testInt8ConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustInt8Consumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testInt8ConsumerError.Error(), func() {
					mc(testInt8ConsumerValue)
				})
			} else {
				mc(testInt8ConsumerValue)
			}
		})
	}
}

func TestSilentInt8Consumer(t *testing.T) {
	var sc SilentInt8Consumer = func(v int8) {
		require.Equal(t, testInt8ConsumerValue, v)
		return
	}
	sc(testInt8ConsumerValue)
}

func TestSilentInt8Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt8ConsumerFactory
		cf2   testInt8ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int8Consumer {
				return func(v int8) error {
					calls++
					require.Equal(t, testInt8ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8Consumer {
				return func(v int8) error {
					calls++
					require.Equal(t, testInt8ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int8Consumer {
				return func(v int8) error {
					calls++
					require.Equal(t, testInt8ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt8ConsumerError
				}
			},
			cf2: func(t *testing.T) Int8Consumer {
				return func(v int8) error {
					calls++
					require.Equal(t, testInt8ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int8Consumer {
				return func(v int8) error {
					calls++
					require.Equal(t, testInt8ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentInt8Consumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentInt8Consumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentInt8Consumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testInt8ConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt8Consumer(t *testing.T) {
	var sc SilentInt8Consumer = func(v int8) {
		require.Equal(t, testInt8ConsumerValue, v)
		return
	}
	sc(testInt8ConsumerValue)
}

func TestMustInt8Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt8ConsumerFactory
		cf2   testInt8ConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int8Consumer {
				return func(v int8) error {
					calls++
					require.Equal(t, testInt8ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8Consumer {
				return func(v int8) error {
					calls++
					require.Equal(t, testInt8ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int8Consumer {
				return func(v int8) error {
					calls++
					require.Equal(t, testInt8ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt8ConsumerError
				}
			},
			cf2: func(t *testing.T) Int8Consumer {
				return func(v int8) error {
					calls++
					require.Equal(t, testInt8ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int8Consumer {
				return func(v int8) error {
					calls++
					require.Equal(t, testInt8ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int8Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustInt8Consumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustInt8Consumer = nil
			if c2 != nil {
				mc2 = c2.ToMustInt8Consumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testInt8ConsumerError.Error(), func() {
					cmc(testInt8ConsumerValue)
				})
			} else {
				cmc(testInt8ConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt8Consumer_ToSilentInt8Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8Consumer {
				return func(v int8) error {
					require.Equal(t, testInt8ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8Consumer {
				return func(v int8) error {
					require.Equal(t, testInt8ConsumerValue, v)
					return testInt8ConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt8Consumer()
			r.NotNil(mc)

			sc := mc.ToSilentInt8Consumer()
			r.NotNil(sc)

			sc(testInt8ConsumerValue)
		})
	}
}

func TestMustInt8Consumer_ToInt8Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt8ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int8Consumer {
				return func(v int8) error {
					require.Equal(t, testInt8ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int8Consumer {
				return func(v int8) error {
					require.Equal(t, testInt8ConsumerValue, v)
					return testInt8ConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt8Consumer()
			r.NotNil(mc)

			c = mc.ToInt8Consumer()
			r.NotNil(c)

			err := c(testInt8ConsumerValue)
			if tt.err {
				r.EqualError(err, testInt8ConsumerError.Error())
			}
		})
	}
}
