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
	valTestUint8SliceConsumer []uint8
	errTestUint8SliceConsumer = errors.New("error")
)

type testUint8SliceConsumerFactory func(t *testing.T) Uint8SliceConsumer

func TestUint8SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint8SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint8SliceConsumer {
				return func(v []uint8) error {
					require.Equal(t, valTestUint8SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint8SliceConsumer {
				return func(v []uint8) error {
					require.Equal(t, valTestUint8SliceConsumer, v)
					return errTestUint8SliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestUint8SliceConsumer)
			if err != nil {
				r.EqualError(err, errTestUint8SliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint8SliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint8SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint8SliceConsumer {
				return func(v []uint8) error {
					require.Equal(t, valTestUint8SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint8SliceConsumer {
				return func(v []uint8) error {
					require.Equal(t, valTestUint8SliceConsumer, v)
					return errTestUint8SliceConsumer
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

			err := c(valTestUint8SliceConsumer)
			if err != nil {
				r.EqualError(err, errTestUint8SliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint8SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint8SliceConsumerFactory
		cf2   testUint8SliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint8SliceConsumer {
				return func(v []uint8) error {
					calls++
					require.Equal(t, valTestUint8SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint8SliceConsumer {
				return func(v []uint8) error {
					calls++
					require.Equal(t, valTestUint8SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint8SliceConsumer {
				return func(v []uint8) error {
					calls++
					require.Equal(t, valTestUint8SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint8SliceConsumer
				}
			},
			cf2: func(t *testing.T) Uint8SliceConsumer {
				return func(v []uint8) error {
					calls++
					require.Equal(t, valTestUint8SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint8SliceConsumer {
				return func(v []uint8) error {
					calls++
					require.Equal(t, valTestUint8SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint8SliceConsumer {
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
			err := cc(valTestUint8SliceConsumer)
			if err != nil {
				r.EqualError(err, errTestUint8SliceConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestUint8SliceConsumer_ToSilentUint8SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint8SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint8SliceConsumer {
				return func(v []uint8) error {
					require.Equal(t, valTestUint8SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint8SliceConsumer {
				return func(v []uint8) error {
					require.Equal(t, valTestUint8SliceConsumer, v)
					return errTestUint8SliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentUint8SliceConsumer()
			r.NotNil(sc)

			sc(valTestUint8SliceConsumer)
		})
	}
}

func TestUint8SliceConsumer_ToMustUint8SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint8SliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint8SliceConsumer {
				return func(v []uint8) error {
					require.Equal(t, valTestUint8SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint8SliceConsumer {
				return func(v []uint8) error {
					require.Equal(t, valTestUint8SliceConsumer, v)
					return errTestUint8SliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustUint8SliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestUint8SliceConsumer.Error(), func() {
					mc(valTestUint8SliceConsumer)
				})
			} else {
				mc(valTestUint8SliceConsumer)
			}
		})
	}
}

func TestSilentUint8SliceConsumer(t *testing.T) {
	var sc SilentUint8SliceConsumer = func(v []uint8) {
		require.Equal(t, valTestUint8SliceConsumer, v)
		return
	}
	sc(valTestUint8SliceConsumer)
}

func TestSilentUint8SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint8SliceConsumerFactory
		cf2   testUint8SliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint8SliceConsumer {
				return func(v []uint8) error {
					calls++
					require.Equal(t, valTestUint8SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint8SliceConsumer {
				return func(v []uint8) error {
					calls++
					require.Equal(t, valTestUint8SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint8SliceConsumer {
				return func(v []uint8) error {
					calls++
					require.Equal(t, valTestUint8SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint8SliceConsumer
				}
			},
			cf2: func(t *testing.T) Uint8SliceConsumer {
				return func(v []uint8) error {
					calls++
					require.Equal(t, valTestUint8SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint8SliceConsumer {
				return func(v []uint8) error {
					calls++
					require.Equal(t, valTestUint8SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint8SliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentUint8SliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentUint8SliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentUint8SliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestUint8SliceConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint8SliceConsumer(t *testing.T) {
	var sc SilentUint8SliceConsumer = func(v []uint8) {
		require.Equal(t, valTestUint8SliceConsumer, v)
		return
	}
	sc(valTestUint8SliceConsumer)
}

func TestMustUint8SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint8SliceConsumerFactory
		cf2   testUint8SliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint8SliceConsumer {
				return func(v []uint8) error {
					calls++
					require.Equal(t, valTestUint8SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint8SliceConsumer {
				return func(v []uint8) error {
					calls++
					require.Equal(t, valTestUint8SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint8SliceConsumer {
				return func(v []uint8) error {
					calls++
					require.Equal(t, valTestUint8SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint8SliceConsumer
				}
			},
			cf2: func(t *testing.T) Uint8SliceConsumer {
				return func(v []uint8) error {
					calls++
					require.Equal(t, valTestUint8SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint8SliceConsumer {
				return func(v []uint8) error {
					calls++
					require.Equal(t, valTestUint8SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint8SliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustUint8SliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustUint8SliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustUint8SliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(errTestUint8SliceConsumer.Error(), func() {
					cmc(valTestUint8SliceConsumer)
				})
			} else {
				cmc(valTestUint8SliceConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint8SliceConsumer_ToSilentUint8SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint8SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint8SliceConsumer {
				return func(v []uint8) error {
					require.Equal(t, valTestUint8SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint8SliceConsumer {
				return func(v []uint8) error {
					require.Equal(t, valTestUint8SliceConsumer, v)
					return errTestUint8SliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint8SliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentUint8SliceConsumer()
			r.NotNil(sc)

			sc(valTestUint8SliceConsumer)
		})
	}
}

func TestMustUint8SliceConsumer_ToUint8SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint8SliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint8SliceConsumer {
				return func(v []uint8) error {
					require.Equal(t, valTestUint8SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint8SliceConsumer {
				return func(v []uint8) error {
					require.Equal(t, valTestUint8SliceConsumer, v)
					return errTestUint8SliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint8SliceConsumer()
			r.NotNil(mc)

			c = mc.ToUint8SliceConsumer()
			r.NotNil(c)

			err := c(valTestUint8SliceConsumer)
			if tt.err {
				r.EqualError(err, errTestUint8SliceConsumer.Error())
			}
		})
	}
}
