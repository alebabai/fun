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
	valTestUint64PtrConsumer *uint64
	errTestUint64PtrConsumer = errors.New("error")
)

type testUint64PtrConsumerFactory func(t *testing.T) Uint64PtrConsumer

func TestUint64PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64PtrConsumer {
				return func(v *uint64) error {
					require.Equal(t, valTestUint64PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint64PtrConsumer {
				return func(v *uint64) error {
					require.Equal(t, valTestUint64PtrConsumer, v)
					return errTestUint64PtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestUint64PtrConsumer)
			if err != nil {
				r.EqualError(err, errTestUint64PtrConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint64PtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64PtrConsumer {
				return func(v *uint64) error {
					require.Equal(t, valTestUint64PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint64PtrConsumer {
				return func(v *uint64) error {
					require.Equal(t, valTestUint64PtrConsumer, v)
					return errTestUint64PtrConsumer
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

			err := c(valTestUint64PtrConsumer)
			if err != nil {
				r.EqualError(err, errTestUint64PtrConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint64PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint64PtrConsumerFactory
		cf2   testUint64PtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint64PtrConsumer {
				return func(v *uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64PtrConsumer {
				return func(v *uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint64PtrConsumer {
				return func(v *uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint64PtrConsumer
				}
			},
			cf2: func(t *testing.T) Uint64PtrConsumer {
				return func(v *uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint64PtrConsumer {
				return func(v *uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64PtrConsumer {
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
			err := cc(valTestUint64PtrConsumer)
			if err != nil {
				r.EqualError(err, errTestUint64PtrConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestUint64PtrConsumer_ToSilentUint64PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64PtrConsumer {
				return func(v *uint64) error {
					require.Equal(t, valTestUint64PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint64PtrConsumer {
				return func(v *uint64) error {
					require.Equal(t, valTestUint64PtrConsumer, v)
					return errTestUint64PtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentUint64PtrConsumer()
			r.NotNil(sc)

			sc(valTestUint64PtrConsumer)
		})
	}
}

func TestUint64PtrConsumer_ToMustUint64PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64PtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64PtrConsumer {
				return func(v *uint64) error {
					require.Equal(t, valTestUint64PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint64PtrConsumer {
				return func(v *uint64) error {
					require.Equal(t, valTestUint64PtrConsumer, v)
					return errTestUint64PtrConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustUint64PtrConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestUint64PtrConsumer.Error(), func() {
					mc(valTestUint64PtrConsumer)
				})
			} else {
				mc(valTestUint64PtrConsumer)
			}
		})
	}
}

func TestSilentUint64PtrConsumer(t *testing.T) {
	var sc SilentUint64PtrConsumer = func(v *uint64) {
		require.Equal(t, valTestUint64PtrConsumer, v)
		return
	}
	sc(valTestUint64PtrConsumer)
}

func TestSilentUint64PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint64PtrConsumerFactory
		cf2   testUint64PtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint64PtrConsumer {
				return func(v *uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64PtrConsumer {
				return func(v *uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint64PtrConsumer {
				return func(v *uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint64PtrConsumer
				}
			},
			cf2: func(t *testing.T) Uint64PtrConsumer {
				return func(v *uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint64PtrConsumer {
				return func(v *uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64PtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentUint64PtrConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentUint64PtrConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentUint64PtrConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestUint64PtrConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint64PtrConsumer(t *testing.T) {
	var sc SilentUint64PtrConsumer = func(v *uint64) {
		require.Equal(t, valTestUint64PtrConsumer, v)
		return
	}
	sc(valTestUint64PtrConsumer)
}

func TestMustUint64PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint64PtrConsumerFactory
		cf2   testUint64PtrConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint64PtrConsumer {
				return func(v *uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64PtrConsumer {
				return func(v *uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint64PtrConsumer {
				return func(v *uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint64PtrConsumer
				}
			},
			cf2: func(t *testing.T) Uint64PtrConsumer {
				return func(v *uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint64PtrConsumer {
				return func(v *uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64PtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustUint64PtrConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustUint64PtrConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustUint64PtrConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(errTestUint64PtrConsumer.Error(), func() {
					cmc(valTestUint64PtrConsumer)
				})
			} else {
				cmc(valTestUint64PtrConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint64PtrConsumer_ToSilentUint64PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64PtrConsumer {
				return func(v *uint64) error {
					require.Equal(t, valTestUint64PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint64PtrConsumer {
				return func(v *uint64) error {
					require.Equal(t, valTestUint64PtrConsumer, v)
					return errTestUint64PtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint64PtrConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentUint64PtrConsumer()
			r.NotNil(sc)

			sc(valTestUint64PtrConsumer)
		})
	}
}

func TestMustUint64PtrConsumer_ToUint64PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64PtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64PtrConsumer {
				return func(v *uint64) error {
					require.Equal(t, valTestUint64PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint64PtrConsumer {
				return func(v *uint64) error {
					require.Equal(t, valTestUint64PtrConsumer, v)
					return errTestUint64PtrConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint64PtrConsumer()
			r.NotNil(mc)

			c = mc.ToUint64PtrConsumer()
			r.NotNil(c)

			err := c(valTestUint64PtrConsumer)
			if tt.err {
				r.EqualError(err, errTestUint64PtrConsumer.Error())
			}
		})
	}
}
