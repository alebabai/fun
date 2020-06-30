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
	valTestByteSliceConsumer []byte
	errTestByteSliceConsumer = errors.New("error")
)

type testByteSliceConsumerFactory func(t *testing.T) ByteSliceConsumer

func TestByteSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testByteSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					require.Equal(t, valTestByteSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					require.Equal(t, valTestByteSliceConsumer, v)
					return errTestByteSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestByteSliceConsumer)
			if err != nil {
				r.EqualError(err, errTestByteSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestByteSliceConsumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testByteSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					require.Equal(t, valTestByteSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					require.Equal(t, valTestByteSliceConsumer, v)
					return errTestByteSliceConsumer
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

			err := c(valTestByteSliceConsumer)
			if err != nil {
				r.EqualError(err, errTestByteSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestByteSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testByteSliceConsumerFactory
		cf2   testByteSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					calls++
					require.Equal(t, valTestByteSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					calls++
					require.Equal(t, valTestByteSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					calls++
					require.Equal(t, valTestByteSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestByteSliceConsumer
				}
			},
			cf2: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					calls++
					require.Equal(t, valTestByteSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					calls++
					require.Equal(t, valTestByteSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) ByteSliceConsumer {
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
			err := cc(valTestByteSliceConsumer)
			if err != nil {
				r.EqualError(err, errTestByteSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestByteSliceConsumer_ToSilentByteSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testByteSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					require.Equal(t, valTestByteSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					require.Equal(t, valTestByteSliceConsumer, v)
					return errTestByteSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentByteSliceConsumer()
			r.NotNil(sc)

			sc(valTestByteSliceConsumer)
		})
	}
}

func TestByteSliceConsumer_ToMustByteSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testByteSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					require.Equal(t, valTestByteSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					require.Equal(t, valTestByteSliceConsumer, v)
					return errTestByteSliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustByteSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestByteSliceConsumer.Error(), func() {
					mc(valTestByteSliceConsumer)
				})
			} else {
				mc(valTestByteSliceConsumer)
			}
		})
	}
}

func TestSilentByteSliceConsumer(t *testing.T) {
	var sc SilentByteSliceConsumer = func(v []byte) {
		require.Equal(t, valTestByteSliceConsumer, v)
		return
	}
	sc(valTestByteSliceConsumer)
}

func TestSilentByteSliceConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testByteSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					require.Equal(t, valTestByteSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					require.Equal(t, valTestByteSliceConsumer, v)
					return errTestByteSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentByteSliceConsumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestByteSliceConsumer)
		})
	}
}

func TestSilentByteSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testByteSliceConsumerFactory
		cf2   testByteSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					calls++
					require.Equal(t, valTestByteSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					calls++
					require.Equal(t, valTestByteSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					calls++
					require.Equal(t, valTestByteSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestByteSliceConsumer
				}
			},
			cf2: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					calls++
					require.Equal(t, valTestByteSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					calls++
					require.Equal(t, valTestByteSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) ByteSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentByteSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentByteSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentByteSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestByteSliceConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustByteSliceConsumer(t *testing.T) {
	var mc MustByteSliceConsumer = func(v []byte) {
		require.Equal(t, valTestByteSliceConsumer, v)
		return
	}
	mc(valTestByteSliceConsumer)
}

func TestMustByteSliceConsumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testByteSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					require.Equal(t, valTestByteSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					require.Equal(t, valTestByteSliceConsumer, v)
					return errTestByteSliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustByteSliceConsumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestByteSliceConsumer.Error(), func() {
					mc(valTestByteSliceConsumer)
				})
			} else {
				mc(valTestByteSliceConsumer)
			}
		})
	}
}

func TestMustByteSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testByteSliceConsumerFactory
		cf2   testByteSliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					calls++
					require.Equal(t, valTestByteSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					calls++
					require.Equal(t, valTestByteSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					calls++
					require.Equal(t, valTestByteSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestByteSliceConsumer
				}
			},
			cf2: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					calls++
					require.Equal(t, valTestByteSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					calls++
					require.Equal(t, valTestByteSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) ByteSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustByteSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustByteSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustByteSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(errTestByteSliceConsumer.Error(), func() {
					cmc(valTestByteSliceConsumer)
				})
			} else {
				cmc(valTestByteSliceConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustByteSliceConsumer_ToSilentByteSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testByteSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					require.Equal(t, valTestByteSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					require.Equal(t, valTestByteSliceConsumer, v)
					return errTestByteSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustByteSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentByteSliceConsumer()
			r.NotNil(sc)

			sc(valTestByteSliceConsumer)
		})
	}
}

func TestMustByteSliceConsumer_ToByteSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testByteSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					require.Equal(t, valTestByteSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) ByteSliceConsumer {
				return func(v []byte) error {
					require.Equal(t, valTestByteSliceConsumer, v)
					return errTestByteSliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustByteSliceConsumer()
			r.NotNil(mc)

			c = mc.ToByteSliceConsumer()
			r.NotNil(c)

			err := c(valTestByteSliceConsumer)
			if tt.err {
				r.EqualError(err, errTestByteSliceConsumer.Error())
			}
		})
	}
}
