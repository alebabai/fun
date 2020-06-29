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
	testInt16ConsumerValue int16
	testInt16ConsumerError = errors.New("error")
)

type testInt16ConsumerFactory func(t *testing.T) Int16Consumer

func TestInt16Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt16ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					require.Equal(t, testInt16ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					require.Equal(t, testInt16ConsumerValue, v)
					return testInt16ConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testInt16ConsumerValue)
			if err != nil {
				r.EqualError(err, testInt16ConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt16Supplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt16ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					require.Equal(t, testInt16ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					require.Equal(t, testInt16ConsumerValue, v)
					return testInt16ConsumerError
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

			err := c(testInt16ConsumerValue)
			if err != nil {
				r.EqualError(err, testInt16ConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt16Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt16ConsumerFactory
		cf2   testInt16ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, testInt16ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, testInt16ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, testInt16ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt16ConsumerError
				}
			},
			cf2: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, testInt16ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, testInt16ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16Consumer {
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
			err := cc(testInt16ConsumerValue)
			if err != nil {
				r.EqualError(err, testInt16ConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestInt16Consumer_ToSilentInt16Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt16ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					require.Equal(t, testInt16ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					require.Equal(t, testInt16ConsumerValue, v)
					return testInt16ConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentInt16Consumer()
			r.NotNil(sc)

			sc(testInt16ConsumerValue)
		})
	}
}

func TestInt16Consumer_ToMustInt16Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt16ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					require.Equal(t, testInt16ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					require.Equal(t, testInt16ConsumerValue, v)
					return testInt16ConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustInt16Consumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testInt16ConsumerError.Error(), func() {
					mc(testInt16ConsumerValue)
				})
			} else {
				mc(testInt16ConsumerValue)
			}
		})
	}
}

func TestSilentInt16Consumer(t *testing.T) {
	var sc SilentInt16Consumer = func(v int16) {
		require.Equal(t, testInt16ConsumerValue, v)
		return
	}
	sc(testInt16ConsumerValue)
}

func TestSilentInt16Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt16ConsumerFactory
		cf2   testInt16ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, testInt16ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, testInt16ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, testInt16ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt16ConsumerError
				}
			},
			cf2: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, testInt16ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, testInt16ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentInt16Consumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentInt16Consumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentInt16Consumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testInt16ConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt16Consumer(t *testing.T) {
	var sc SilentInt16Consumer = func(v int16) {
		require.Equal(t, testInt16ConsumerValue, v)
		return
	}
	sc(testInt16ConsumerValue)
}

func TestMustInt16Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt16ConsumerFactory
		cf2   testInt16ConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, testInt16ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, testInt16ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, testInt16ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt16ConsumerError
				}
			},
			cf2: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, testInt16ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, testInt16ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustInt16Consumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustInt16Consumer = nil
			if c2 != nil {
				mc2 = c2.ToMustInt16Consumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testInt16ConsumerError.Error(), func() {
					cmc(testInt16ConsumerValue)
				})
			} else {
				cmc(testInt16ConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt16Consumer_ToSilentInt16Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt16ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					require.Equal(t, testInt16ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					require.Equal(t, testInt16ConsumerValue, v)
					return testInt16ConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt16Consumer()
			r.NotNil(mc)

			sc := mc.ToSilentInt16Consumer()
			r.NotNil(sc)

			sc(testInt16ConsumerValue)
		})
	}
}

func TestMustInt16Consumer_ToInt16Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt16ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					require.Equal(t, testInt16ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					require.Equal(t, testInt16ConsumerValue, v)
					return testInt16ConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt16Consumer()
			r.NotNil(mc)

			c = mc.ToInt16Consumer()
			r.NotNil(c)

			err := c(testInt16ConsumerValue)
			if tt.err {
				r.EqualError(err, testInt16ConsumerError.Error())
			}
		})
	}
}
