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
	testUint64ConsumerValue uint64
	testUint64ConsumerError = errors.New("error")
)

type testUint64ConsumerFactory func(t *testing.T) Uint64Consumer

func TestUint64Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					require.Equal(t, testUint64ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					require.Equal(t, testUint64ConsumerValue, v)
					return testUint64ConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testUint64ConsumerValue)
			if err != nil {
				r.EqualError(err, testUint64ConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint64Supplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					require.Equal(t, testUint64ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					require.Equal(t, testUint64ConsumerValue, v)
					return testUint64ConsumerError
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

			err := c(testUint64ConsumerValue)
			if err != nil {
				r.EqualError(err, testUint64ConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint64Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint64ConsumerFactory
		cf2   testUint64ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, testUint64ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, testUint64ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, testUint64ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUint64ConsumerError
				}
			},
			cf2: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, testUint64ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, testUint64ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64Consumer {
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
			err := cc(testUint64ConsumerValue)
			if err != nil {
				r.EqualError(err, testUint64ConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestUint64Consumer_ToSilentUint64Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					require.Equal(t, testUint64ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					require.Equal(t, testUint64ConsumerValue, v)
					return testUint64ConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentUint64Consumer()
			r.NotNil(sc)

			sc(testUint64ConsumerValue)
		})
	}
}

func TestUint64Consumer_ToMustUint64Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					require.Equal(t, testUint64ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					require.Equal(t, testUint64ConsumerValue, v)
					return testUint64ConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustUint64Consumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testUint64ConsumerError.Error(), func() {
					mc(testUint64ConsumerValue)
				})
			} else {
				mc(testUint64ConsumerValue)
			}
		})
	}
}

func TestSilentUint64Consumer(t *testing.T) {
	var sc SilentUint64Consumer = func(v uint64) {
		require.Equal(t, testUint64ConsumerValue, v)
		return
	}
	sc(testUint64ConsumerValue)
}

func TestSilentUint64Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint64ConsumerFactory
		cf2   testUint64ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, testUint64ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, testUint64ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, testUint64ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUint64ConsumerError
				}
			},
			cf2: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, testUint64ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, testUint64ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentUint64Consumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentUint64Consumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentUint64Consumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testUint64ConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint64Consumer(t *testing.T) {
	var sc SilentUint64Consumer = func(v uint64) {
		require.Equal(t, testUint64ConsumerValue, v)
		return
	}
	sc(testUint64ConsumerValue)
}

func TestMustUint64Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint64ConsumerFactory
		cf2   testUint64ConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, testUint64ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, testUint64ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, testUint64ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUint64ConsumerError
				}
			},
			cf2: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, testUint64ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, testUint64ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustUint64Consumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustUint64Consumer = nil
			if c2 != nil {
				mc2 = c2.ToMustUint64Consumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testUint64ConsumerError.Error(), func() {
					cmc(testUint64ConsumerValue)
				})
			} else {
				cmc(testUint64ConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint64Consumer_ToSilentUint64Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					require.Equal(t, testUint64ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					require.Equal(t, testUint64ConsumerValue, v)
					return testUint64ConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint64Consumer()
			r.NotNil(mc)

			sc := mc.ToSilentUint64Consumer()
			r.NotNil(sc)

			sc(testUint64ConsumerValue)
		})
	}
}

func TestMustUint64Consumer_ToUint64Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					require.Equal(t, testUint64ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					require.Equal(t, testUint64ConsumerValue, v)
					return testUint64ConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint64Consumer()
			r.NotNil(mc)

			c = mc.ToUint64Consumer()
			r.NotNil(c)

			err := c(testUint64ConsumerValue)
			if tt.err {
				r.EqualError(err, testUint64ConsumerError.Error())
			}
		})
	}
}
