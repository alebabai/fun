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
	testInt32ConsumerValue int32
	testInt32ConsumerError = errors.New("error")
)

type testInt32ConsumerFactory func(t *testing.T) Int32Consumer

func TestInt32Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					require.Equal(t, testInt32ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					require.Equal(t, testInt32ConsumerValue, v)
					return testInt32ConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testInt32ConsumerValue)
			if err != nil {
				r.EqualError(err, testInt32ConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt32Supplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					require.Equal(t, testInt32ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					require.Equal(t, testInt32ConsumerValue, v)
					return testInt32ConsumerError
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

			err := c(testInt32ConsumerValue)
			if err != nil {
				r.EqualError(err, testInt32ConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt32Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt32ConsumerFactory
		cf2   testInt32ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, testInt32ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, testInt32ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, testInt32ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt32ConsumerError
				}
			},
			cf2: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, testInt32ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, testInt32ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32Consumer {
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
			err := cc(testInt32ConsumerValue)
			if err != nil {
				r.EqualError(err, testInt32ConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestInt32Consumer_ToSilentInt32Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					require.Equal(t, testInt32ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					require.Equal(t, testInt32ConsumerValue, v)
					return testInt32ConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentInt32Consumer()
			r.NotNil(sc)

			sc(testInt32ConsumerValue)
		})
	}
}

func TestInt32Consumer_ToMustInt32Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					require.Equal(t, testInt32ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					require.Equal(t, testInt32ConsumerValue, v)
					return testInt32ConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustInt32Consumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testInt32ConsumerError.Error(), func() {
					mc(testInt32ConsumerValue)
				})
			} else {
				mc(testInt32ConsumerValue)
			}
		})
	}
}

func TestSilentInt32Consumer(t *testing.T) {
	var sc SilentInt32Consumer = func(v int32) {
		require.Equal(t, testInt32ConsumerValue, v)
		return
	}
	sc(testInt32ConsumerValue)
}

func TestSilentInt32Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt32ConsumerFactory
		cf2   testInt32ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, testInt32ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, testInt32ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, testInt32ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt32ConsumerError
				}
			},
			cf2: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, testInt32ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, testInt32ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentInt32Consumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentInt32Consumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentInt32Consumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testInt32ConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt32Consumer(t *testing.T) {
	var sc SilentInt32Consumer = func(v int32) {
		require.Equal(t, testInt32ConsumerValue, v)
		return
	}
	sc(testInt32ConsumerValue)
}

func TestMustInt32Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt32ConsumerFactory
		cf2   testInt32ConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, testInt32ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, testInt32ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, testInt32ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testInt32ConsumerError
				}
			},
			cf2: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, testInt32ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					calls++
					require.Equal(t, testInt32ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustInt32Consumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustInt32Consumer = nil
			if c2 != nil {
				mc2 = c2.ToMustInt32Consumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testInt32ConsumerError.Error(), func() {
					cmc(testInt32ConsumerValue)
				})
			} else {
				cmc(testInt32ConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt32Consumer_ToSilentInt32Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					require.Equal(t, testInt32ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					require.Equal(t, testInt32ConsumerValue, v)
					return testInt32ConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt32Consumer()
			r.NotNil(mc)

			sc := mc.ToSilentInt32Consumer()
			r.NotNil(sc)

			sc(testInt32ConsumerValue)
		})
	}
}

func TestMustInt32Consumer_ToInt32Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					require.Equal(t, testInt32ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Int32Consumer {
				return func(v int32) error {
					require.Equal(t, testInt32ConsumerValue, v)
					return testInt32ConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt32Consumer()
			r.NotNil(mc)

			c = mc.ToInt32Consumer()
			r.NotNil(c)

			err := c(testInt32ConsumerValue)
			if tt.err {
				r.EqualError(err, testInt32ConsumerError.Error())
			}
		})
	}
}
