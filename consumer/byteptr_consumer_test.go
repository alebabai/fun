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
	valTestBytePtrConsumer *byte
	errTestBytePtrConsumer = errors.New("error")
)

type testBytePtrConsumerFactory func(t *testing.T) BytePtrConsumer

func TestBytePtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBytePtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					require.Equal(t, valTestBytePtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					require.Equal(t, valTestBytePtrConsumer, v)
					return errTestBytePtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestBytePtrConsumer)
			if err != nil {
				r.EqualError(err, errTestBytePtrConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestBytePtrConsumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBytePtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					require.Equal(t, valTestBytePtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					require.Equal(t, valTestBytePtrConsumer, v)
					return errTestBytePtrConsumer
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

			err := c(valTestBytePtrConsumer)
			if err != nil {
				r.EqualError(err, errTestBytePtrConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestBytePtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testBytePtrConsumerFactory
		cf2   testBytePtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					calls++
					require.Equal(t, valTestBytePtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					calls++
					require.Equal(t, valTestBytePtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					calls++
					require.Equal(t, valTestBytePtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestBytePtrConsumer
				}
			},
			cf2: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					calls++
					require.Equal(t, valTestBytePtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					calls++
					require.Equal(t, valTestBytePtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BytePtrConsumer {
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
			err := cc(valTestBytePtrConsumer)
			if err != nil {
				r.EqualError(err, errTestBytePtrConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestBytePtrConsumer_ToSilentBytePtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBytePtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					require.Equal(t, valTestBytePtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					require.Equal(t, valTestBytePtrConsumer, v)
					return errTestBytePtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentBytePtrConsumer()
			r.NotNil(sc)

			sc(valTestBytePtrConsumer)
		})
	}
}

func TestBytePtrConsumer_ToMustBytePtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBytePtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					require.Equal(t, valTestBytePtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					require.Equal(t, valTestBytePtrConsumer, v)
					return errTestBytePtrConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustBytePtrConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestBytePtrConsumer.Error(), func() {
					mc(valTestBytePtrConsumer)
				})
			} else {
				mc(valTestBytePtrConsumer)
			}
		})
	}
}

func TestSilentBytePtrConsumer(t *testing.T) {
	var sc SilentBytePtrConsumer = func(v *byte) {
		require.Equal(t, valTestBytePtrConsumer, v)
		return
	}
	sc(valTestBytePtrConsumer)
}

func TestSilentBytePtrConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBytePtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					require.Equal(t, valTestBytePtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					require.Equal(t, valTestBytePtrConsumer, v)
					return errTestBytePtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentBytePtrConsumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestBytePtrConsumer)
		})
	}
}

func TestSilentBytePtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testBytePtrConsumerFactory
		cf2   testBytePtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					calls++
					require.Equal(t, valTestBytePtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					calls++
					require.Equal(t, valTestBytePtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					calls++
					require.Equal(t, valTestBytePtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestBytePtrConsumer
				}
			},
			cf2: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					calls++
					require.Equal(t, valTestBytePtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					calls++
					require.Equal(t, valTestBytePtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BytePtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentBytePtrConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentBytePtrConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentBytePtrConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestBytePtrConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustBytePtrConsumer(t *testing.T) {
	var mc MustBytePtrConsumer = func(v *byte) {
		require.Equal(t, valTestBytePtrConsumer, v)
		return
	}
	mc(valTestBytePtrConsumer)
}

func TestMustBytePtrConsumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBytePtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					require.Equal(t, valTestBytePtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					require.Equal(t, valTestBytePtrConsumer, v)
					return errTestBytePtrConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustBytePtrConsumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestBytePtrConsumer.Error(), func() {
					mc(valTestBytePtrConsumer)
				})
			} else {
				mc(valTestBytePtrConsumer)
			}
		})
	}
}

func TestMustBytePtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testBytePtrConsumerFactory
		cf2   testBytePtrConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					calls++
					require.Equal(t, valTestBytePtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					calls++
					require.Equal(t, valTestBytePtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					calls++
					require.Equal(t, valTestBytePtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestBytePtrConsumer
				}
			},
			cf2: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					calls++
					require.Equal(t, valTestBytePtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					calls++
					require.Equal(t, valTestBytePtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BytePtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustBytePtrConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustBytePtrConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustBytePtrConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(errTestBytePtrConsumer.Error(), func() {
					cmc(valTestBytePtrConsumer)
				})
			} else {
				cmc(valTestBytePtrConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustBytePtrConsumer_ToSilentBytePtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBytePtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					require.Equal(t, valTestBytePtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					require.Equal(t, valTestBytePtrConsumer, v)
					return errTestBytePtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustBytePtrConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentBytePtrConsumer()
			r.NotNil(sc)

			sc(valTestBytePtrConsumer)
		})
	}
}

func TestMustBytePtrConsumer_ToBytePtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBytePtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					require.Equal(t, valTestBytePtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BytePtrConsumer {
				return func(v *byte) error {
					require.Equal(t, valTestBytePtrConsumer, v)
					return errTestBytePtrConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustBytePtrConsumer()
			r.NotNil(mc)

			c = mc.ToBytePtrConsumer()
			r.NotNil(c)

			err := c(valTestBytePtrConsumer)
			if tt.err {
				r.EqualError(err, errTestBytePtrConsumer.Error())
			}
		})
	}
}
