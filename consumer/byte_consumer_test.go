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
	testByteConsumerValue byte
	testByteConsumerError = errors.New("error")
)

type testByteConsumerFactory func(t *testing.T) ByteConsumer

func TestByteConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testByteConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					require.Equal(t, testByteConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					require.Equal(t, testByteConsumerValue, v)
					return testByteConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testByteConsumerValue)
			if err != nil {
				r.EqualError(err, testByteConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestByteSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testByteConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					require.Equal(t, testByteConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					require.Equal(t, testByteConsumerValue, v)
					return testByteConsumerError
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

			err := c(testByteConsumerValue)
			if err != nil {
				r.EqualError(err, testByteConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestByteConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testByteConsumerFactory
		cf2   testByteConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, testByteConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, testByteConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, testByteConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testByteConsumerError
				}
			},
			cf2: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, testByteConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, testByteConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) ByteConsumer {
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
			err := cc(testByteConsumerValue)
			if err != nil {
				r.EqualError(err, testByteConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestByteConsumer_ToSilentByteConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testByteConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					require.Equal(t, testByteConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					require.Equal(t, testByteConsumerValue, v)
					return testByteConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentByteConsumer()
			r.NotNil(sc)

			sc(testByteConsumerValue)
		})
	}
}

func TestByteConsumer_ToMustByteConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testByteConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					require.Equal(t, testByteConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					require.Equal(t, testByteConsumerValue, v)
					return testByteConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustByteConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testByteConsumerError.Error(), func() {
					mc(testByteConsumerValue)
				})
			} else {
				mc(testByteConsumerValue)
			}
		})
	}
}

func TestSilentByteConsumer(t *testing.T) {
	var sc SilentByteConsumer = func(v byte) {
		require.Equal(t, testByteConsumerValue, v)
		return
	}
	sc(testByteConsumerValue)
}

func TestSilentByteConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testByteConsumerFactory
		cf2   testByteConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, testByteConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, testByteConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, testByteConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testByteConsumerError
				}
			},
			cf2: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, testByteConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, testByteConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) ByteConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentByteConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentByteConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentByteConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testByteConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustByteConsumer(t *testing.T) {
	var sc SilentByteConsumer = func(v byte) {
		require.Equal(t, testByteConsumerValue, v)
		return
	}
	sc(testByteConsumerValue)
}

func TestMustByteConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testByteConsumerFactory
		cf2   testByteConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, testByteConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, testByteConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, testByteConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testByteConsumerError
				}
			},
			cf2: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, testByteConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, testByteConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) ByteConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustByteConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustByteConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustByteConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testByteConsumerError.Error(), func() {
					cmc(testByteConsumerValue)
				})
			} else {
				cmc(testByteConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustByteConsumer_ToSilentByteConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testByteConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					require.Equal(t, testByteConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					require.Equal(t, testByteConsumerValue, v)
					return testByteConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustByteConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentByteConsumer()
			r.NotNil(sc)

			sc(testByteConsumerValue)
		})
	}
}

func TestMustByteConsumer_ToByteConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testByteConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					require.Equal(t, testByteConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					require.Equal(t, testByteConsumerValue, v)
					return testByteConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustByteConsumer()
			r.NotNil(mc)

			c = mc.ToByteConsumer()
			r.NotNil(c)

			err := c(testByteConsumerValue)
			if tt.err {
				r.EqualError(err, testByteConsumerError.Error())
			}
		})
	}
}
