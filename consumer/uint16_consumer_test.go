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
	testUint16ConsumerValue uint16
	testUint16ConsumerError = errors.New("error")
)

type testUint16ConsumerFactory func(t *testing.T) Uint16Consumer

func TestUint16Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint16ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					require.Equal(t, testUint16ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					require.Equal(t, testUint16ConsumerValue, v)
					return testUint16ConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testUint16ConsumerValue)
			if err != nil {
				r.EqualError(err, testUint16ConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint16Supplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint16ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					require.Equal(t, testUint16ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					require.Equal(t, testUint16ConsumerValue, v)
					return testUint16ConsumerError
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

			err := c(testUint16ConsumerValue)
			if err != nil {
				r.EqualError(err, testUint16ConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint16Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint16ConsumerFactory
		cf2   testUint16ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, testUint16ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, testUint16ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, testUint16ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUint16ConsumerError
				}
			},
			cf2: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, testUint16ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, testUint16ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16Consumer {
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
			err := cc(testUint16ConsumerValue)
			if err != nil {
				r.EqualError(err, testUint16ConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestUint16Consumer_ToSilentUint16Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint16ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					require.Equal(t, testUint16ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					require.Equal(t, testUint16ConsumerValue, v)
					return testUint16ConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentUint16Consumer()
			r.NotNil(sc)

			sc(testUint16ConsumerValue)
		})
	}
}

func TestUint16Consumer_ToMustUint16Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint16ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					require.Equal(t, testUint16ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					require.Equal(t, testUint16ConsumerValue, v)
					return testUint16ConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustUint16Consumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testUint16ConsumerError.Error(), func() {
					mc(testUint16ConsumerValue)
				})
			} else {
				mc(testUint16ConsumerValue)
			}
		})
	}
}

func TestSilentUint16Consumer(t *testing.T) {
	var sc SilentUint16Consumer = func(v uint16) {
		require.Equal(t, testUint16ConsumerValue, v)
		return
	}
	sc(testUint16ConsumerValue)
}

func TestSilentUint16Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint16ConsumerFactory
		cf2   testUint16ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, testUint16ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, testUint16ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, testUint16ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUint16ConsumerError
				}
			},
			cf2: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, testUint16ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, testUint16ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentUint16Consumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentUint16Consumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentUint16Consumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testUint16ConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint16Consumer(t *testing.T) {
	var sc SilentUint16Consumer = func(v uint16) {
		require.Equal(t, testUint16ConsumerValue, v)
		return
	}
	sc(testUint16ConsumerValue)
}

func TestMustUint16Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint16ConsumerFactory
		cf2   testUint16ConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, testUint16ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, testUint16ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, testUint16ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUint16ConsumerError
				}
			},
			cf2: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, testUint16ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, testUint16ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustUint16Consumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustUint16Consumer = nil
			if c2 != nil {
				mc2 = c2.ToMustUint16Consumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testUint16ConsumerError.Error(), func() {
					cmc(testUint16ConsumerValue)
				})
			} else {
				cmc(testUint16ConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint16Consumer_ToSilentUint16Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint16ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					require.Equal(t, testUint16ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					require.Equal(t, testUint16ConsumerValue, v)
					return testUint16ConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint16Consumer()
			r.NotNil(mc)

			sc := mc.ToSilentUint16Consumer()
			r.NotNil(sc)

			sc(testUint16ConsumerValue)
		})
	}
}

func TestMustUint16Consumer_ToUint16Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint16ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					require.Equal(t, testUint16ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					require.Equal(t, testUint16ConsumerValue, v)
					return testUint16ConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint16Consumer()
			r.NotNil(mc)

			c = mc.ToUint16Consumer()
			r.NotNil(c)

			err := c(testUint16ConsumerValue)
			if tt.err {
				r.EqualError(err, testUint16ConsumerError.Error())
			}
		})
	}
}
