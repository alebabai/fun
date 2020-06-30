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
	valTestUint16SliceConsumer []uint16
	errTestUint16SliceConsumer = errors.New("error")
)

type testUint16SliceConsumerFactory func(t *testing.T) Uint16SliceConsumer

func TestUint16SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint16SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16SliceConsumer {
				return func(v []uint16) error {
					require.Equal(t, valTestUint16SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint16SliceConsumer {
				return func(v []uint16) error {
					require.Equal(t, valTestUint16SliceConsumer, v)
					return errTestUint16SliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestUint16SliceConsumer)
			if err != nil {
				r.EqualError(err, errTestUint16SliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint16SliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint16SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16SliceConsumer {
				return func(v []uint16) error {
					require.Equal(t, valTestUint16SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint16SliceConsumer {
				return func(v []uint16) error {
					require.Equal(t, valTestUint16SliceConsumer, v)
					return errTestUint16SliceConsumer
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

			err := c(valTestUint16SliceConsumer)
			if err != nil {
				r.EqualError(err, errTestUint16SliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint16SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint16SliceConsumerFactory
		cf2   testUint16SliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint16SliceConsumer {
				return func(v []uint16) error {
					calls++
					require.Equal(t, valTestUint16SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16SliceConsumer {
				return func(v []uint16) error {
					calls++
					require.Equal(t, valTestUint16SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint16SliceConsumer {
				return func(v []uint16) error {
					calls++
					require.Equal(t, valTestUint16SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint16SliceConsumer
				}
			},
			cf2: func(t *testing.T) Uint16SliceConsumer {
				return func(v []uint16) error {
					calls++
					require.Equal(t, valTestUint16SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint16SliceConsumer {
				return func(v []uint16) error {
					calls++
					require.Equal(t, valTestUint16SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16SliceConsumer {
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
			err := cc(valTestUint16SliceConsumer)
			if err != nil {
				r.EqualError(err, errTestUint16SliceConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestUint16SliceConsumer_ToSilentUint16SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint16SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16SliceConsumer {
				return func(v []uint16) error {
					require.Equal(t, valTestUint16SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint16SliceConsumer {
				return func(v []uint16) error {
					require.Equal(t, valTestUint16SliceConsumer, v)
					return errTestUint16SliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentUint16SliceConsumer()
			r.NotNil(sc)

			sc(valTestUint16SliceConsumer)
		})
	}
}

func TestUint16SliceConsumer_ToMustUint16SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint16SliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16SliceConsumer {
				return func(v []uint16) error {
					require.Equal(t, valTestUint16SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint16SliceConsumer {
				return func(v []uint16) error {
					require.Equal(t, valTestUint16SliceConsumer, v)
					return errTestUint16SliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustUint16SliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestUint16SliceConsumer.Error(), func() {
					mc(valTestUint16SliceConsumer)
				})
			} else {
				mc(valTestUint16SliceConsumer)
			}
		})
	}
}

func TestSilentUint16SliceConsumer(t *testing.T) {
	var sc SilentUint16SliceConsumer = func(v []uint16) {
		require.Equal(t, valTestUint16SliceConsumer, v)
		return
	}
	sc(valTestUint16SliceConsumer)
}

func TestSilentUint16SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint16SliceConsumerFactory
		cf2   testUint16SliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint16SliceConsumer {
				return func(v []uint16) error {
					calls++
					require.Equal(t, valTestUint16SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16SliceConsumer {
				return func(v []uint16) error {
					calls++
					require.Equal(t, valTestUint16SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint16SliceConsumer {
				return func(v []uint16) error {
					calls++
					require.Equal(t, valTestUint16SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint16SliceConsumer
				}
			},
			cf2: func(t *testing.T) Uint16SliceConsumer {
				return func(v []uint16) error {
					calls++
					require.Equal(t, valTestUint16SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint16SliceConsumer {
				return func(v []uint16) error {
					calls++
					require.Equal(t, valTestUint16SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16SliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentUint16SliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentUint16SliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentUint16SliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestUint16SliceConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint16SliceConsumer(t *testing.T) {
	var sc SilentUint16SliceConsumer = func(v []uint16) {
		require.Equal(t, valTestUint16SliceConsumer, v)
		return
	}
	sc(valTestUint16SliceConsumer)
}

func TestMustUint16SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint16SliceConsumerFactory
		cf2   testUint16SliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint16SliceConsumer {
				return func(v []uint16) error {
					calls++
					require.Equal(t, valTestUint16SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16SliceConsumer {
				return func(v []uint16) error {
					calls++
					require.Equal(t, valTestUint16SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint16SliceConsumer {
				return func(v []uint16) error {
					calls++
					require.Equal(t, valTestUint16SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint16SliceConsumer
				}
			},
			cf2: func(t *testing.T) Uint16SliceConsumer {
				return func(v []uint16) error {
					calls++
					require.Equal(t, valTestUint16SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint16SliceConsumer {
				return func(v []uint16) error {
					calls++
					require.Equal(t, valTestUint16SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16SliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustUint16SliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustUint16SliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustUint16SliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(errTestUint16SliceConsumer.Error(), func() {
					cmc(valTestUint16SliceConsumer)
				})
			} else {
				cmc(valTestUint16SliceConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint16SliceConsumer_ToSilentUint16SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint16SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16SliceConsumer {
				return func(v []uint16) error {
					require.Equal(t, valTestUint16SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint16SliceConsumer {
				return func(v []uint16) error {
					require.Equal(t, valTestUint16SliceConsumer, v)
					return errTestUint16SliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint16SliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentUint16SliceConsumer()
			r.NotNil(sc)

			sc(valTestUint16SliceConsumer)
		})
	}
}

func TestMustUint16SliceConsumer_ToUint16SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint16SliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16SliceConsumer {
				return func(v []uint16) error {
					require.Equal(t, valTestUint16SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint16SliceConsumer {
				return func(v []uint16) error {
					require.Equal(t, valTestUint16SliceConsumer, v)
					return errTestUint16SliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint16SliceConsumer()
			r.NotNil(mc)

			c = mc.ToUint16SliceConsumer()
			r.NotNil(c)

			err := c(valTestUint16SliceConsumer)
			if tt.err {
				r.EqualError(err, errTestUint16SliceConsumer.Error())
			}
		})
	}
}
