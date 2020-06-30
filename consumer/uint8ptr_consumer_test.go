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
	valTestUint8PtrConsumer *uint8
	errTestUint8PtrConsumer = errors.New("error")
)

type testUint8PtrConsumerFactory func(t *testing.T) Uint8PtrConsumer

func TestUint8PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint8PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint8PtrConsumer {
				return func(v *uint8) error {
					require.Equal(t, valTestUint8PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint8PtrConsumer {
				return func(v *uint8) error {
					require.Equal(t, valTestUint8PtrConsumer, v)
					return errTestUint8PtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestUint8PtrConsumer)
			if err != nil {
				r.EqualError(err, errTestUint8PtrConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint8PtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint8PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint8PtrConsumer {
				return func(v *uint8) error {
					require.Equal(t, valTestUint8PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint8PtrConsumer {
				return func(v *uint8) error {
					require.Equal(t, valTestUint8PtrConsumer, v)
					return errTestUint8PtrConsumer
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

			err := c(valTestUint8PtrConsumer)
			if err != nil {
				r.EqualError(err, errTestUint8PtrConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint8PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint8PtrConsumerFactory
		cf2   testUint8PtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint8PtrConsumer {
				return func(v *uint8) error {
					calls++
					require.Equal(t, valTestUint8PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint8PtrConsumer {
				return func(v *uint8) error {
					calls++
					require.Equal(t, valTestUint8PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint8PtrConsumer {
				return func(v *uint8) error {
					calls++
					require.Equal(t, valTestUint8PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint8PtrConsumer
				}
			},
			cf2: func(t *testing.T) Uint8PtrConsumer {
				return func(v *uint8) error {
					calls++
					require.Equal(t, valTestUint8PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint8PtrConsumer {
				return func(v *uint8) error {
					calls++
					require.Equal(t, valTestUint8PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint8PtrConsumer {
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
			err := cc(valTestUint8PtrConsumer)
			if err != nil {
				r.EqualError(err, errTestUint8PtrConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestUint8PtrConsumer_ToSilentUint8PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint8PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint8PtrConsumer {
				return func(v *uint8) error {
					require.Equal(t, valTestUint8PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint8PtrConsumer {
				return func(v *uint8) error {
					require.Equal(t, valTestUint8PtrConsumer, v)
					return errTestUint8PtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentUint8PtrConsumer()
			r.NotNil(sc)

			sc(valTestUint8PtrConsumer)
		})
	}
}

func TestUint8PtrConsumer_ToMustUint8PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint8PtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint8PtrConsumer {
				return func(v *uint8) error {
					require.Equal(t, valTestUint8PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint8PtrConsumer {
				return func(v *uint8) error {
					require.Equal(t, valTestUint8PtrConsumer, v)
					return errTestUint8PtrConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustUint8PtrConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestUint8PtrConsumer.Error(), func() {
					mc(valTestUint8PtrConsumer)
				})
			} else {
				mc(valTestUint8PtrConsumer)
			}
		})
	}
}

func TestSilentUint8PtrConsumer(t *testing.T) {
	var sc SilentUint8PtrConsumer = func(v *uint8) {
		require.Equal(t, valTestUint8PtrConsumer, v)
		return
	}
	sc(valTestUint8PtrConsumer)
}

func TestSilentUint8PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint8PtrConsumerFactory
		cf2   testUint8PtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint8PtrConsumer {
				return func(v *uint8) error {
					calls++
					require.Equal(t, valTestUint8PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint8PtrConsumer {
				return func(v *uint8) error {
					calls++
					require.Equal(t, valTestUint8PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint8PtrConsumer {
				return func(v *uint8) error {
					calls++
					require.Equal(t, valTestUint8PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint8PtrConsumer
				}
			},
			cf2: func(t *testing.T) Uint8PtrConsumer {
				return func(v *uint8) error {
					calls++
					require.Equal(t, valTestUint8PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint8PtrConsumer {
				return func(v *uint8) error {
					calls++
					require.Equal(t, valTestUint8PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint8PtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentUint8PtrConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentUint8PtrConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentUint8PtrConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestUint8PtrConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint8PtrConsumer(t *testing.T) {
	var sc SilentUint8PtrConsumer = func(v *uint8) {
		require.Equal(t, valTestUint8PtrConsumer, v)
		return
	}
	sc(valTestUint8PtrConsumer)
}

func TestMustUint8PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint8PtrConsumerFactory
		cf2   testUint8PtrConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint8PtrConsumer {
				return func(v *uint8) error {
					calls++
					require.Equal(t, valTestUint8PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint8PtrConsumer {
				return func(v *uint8) error {
					calls++
					require.Equal(t, valTestUint8PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint8PtrConsumer {
				return func(v *uint8) error {
					calls++
					require.Equal(t, valTestUint8PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint8PtrConsumer
				}
			},
			cf2: func(t *testing.T) Uint8PtrConsumer {
				return func(v *uint8) error {
					calls++
					require.Equal(t, valTestUint8PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint8PtrConsumer {
				return func(v *uint8) error {
					calls++
					require.Equal(t, valTestUint8PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint8PtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustUint8PtrConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustUint8PtrConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustUint8PtrConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(errTestUint8PtrConsumer.Error(), func() {
					cmc(valTestUint8PtrConsumer)
				})
			} else {
				cmc(valTestUint8PtrConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint8PtrConsumer_ToSilentUint8PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint8PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint8PtrConsumer {
				return func(v *uint8) error {
					require.Equal(t, valTestUint8PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint8PtrConsumer {
				return func(v *uint8) error {
					require.Equal(t, valTestUint8PtrConsumer, v)
					return errTestUint8PtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint8PtrConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentUint8PtrConsumer()
			r.NotNil(sc)

			sc(valTestUint8PtrConsumer)
		})
	}
}

func TestMustUint8PtrConsumer_ToUint8PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint8PtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint8PtrConsumer {
				return func(v *uint8) error {
					require.Equal(t, valTestUint8PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint8PtrConsumer {
				return func(v *uint8) error {
					require.Equal(t, valTestUint8PtrConsumer, v)
					return errTestUint8PtrConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint8PtrConsumer()
			r.NotNil(mc)

			c = mc.ToUint8PtrConsumer()
			r.NotNil(c)

			err := c(valTestUint8PtrConsumer)
			if tt.err {
				r.EqualError(err, errTestUint8PtrConsumer.Error())
			}
		})
	}
}
