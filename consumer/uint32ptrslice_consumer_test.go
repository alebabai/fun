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
	valTestUint32PtrSliceConsumer []*uint32
	errTestUint32PtrSliceConsumer = errors.New("error")
)

type testUint32PtrSliceConsumerFactory func(t *testing.T) Uint32PtrSliceConsumer

func TestUint32PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint32PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					return errTestUint32PtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestUint32PtrSliceConsumer)
			if err != nil {
				r.EqualError(err, errTestUint32PtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint32PtrSliceConsumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint32PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					return errTestUint32PtrSliceConsumer
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

			err := c(valTestUint32PtrSliceConsumer)
			if err != nil {
				r.EqualError(err, errTestUint32PtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint32PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint32PtrSliceConsumerFactory
		cf2   testUint32PtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					calls++
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					calls++
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					calls++
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint32PtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					calls++
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					calls++
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint32PtrSliceConsumer {
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
			err := cc(valTestUint32PtrSliceConsumer)
			if err != nil {
				r.EqualError(err, errTestUint32PtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestUint32PtrSliceConsumer_ToSilentUint32PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint32PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					return errTestUint32PtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentUint32PtrSliceConsumer()
			r.NotNil(sc)

			sc(valTestUint32PtrSliceConsumer)
		})
	}
}

func TestUint32PtrSliceConsumer_ToMustUint32PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint32PtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					return errTestUint32PtrSliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustUint32PtrSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestUint32PtrSliceConsumer.Error(), func() {
					mc(valTestUint32PtrSliceConsumer)
				})
			} else {
				mc(valTestUint32PtrSliceConsumer)
			}
		})
	}
}

func TestSilentUint32PtrSliceConsumer(t *testing.T) {
	var sc SilentUint32PtrSliceConsumer = func(v []*uint32) {
		require.Equal(t, valTestUint32PtrSliceConsumer, v)
		return
	}
	sc(valTestUint32PtrSliceConsumer)
}

func TestSilentUint32PtrSliceConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint32PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					return errTestUint32PtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentUint32PtrSliceConsumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestUint32PtrSliceConsumer)
		})
	}
}

func TestSilentUint32PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint32PtrSliceConsumerFactory
		cf2   testUint32PtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					calls++
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					calls++
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					calls++
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint32PtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					calls++
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					calls++
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint32PtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentUint32PtrSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentUint32PtrSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentUint32PtrSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestUint32PtrSliceConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint32PtrSliceConsumer(t *testing.T) {
	var mc MustUint32PtrSliceConsumer = func(v []*uint32) {
		require.Equal(t, valTestUint32PtrSliceConsumer, v)
		return
	}
	mc(valTestUint32PtrSliceConsumer)
}

func TestMustUint32PtrSliceConsumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint32PtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					return errTestUint32PtrSliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustUint32PtrSliceConsumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestUint32PtrSliceConsumer.Error(), func() {
					mc(valTestUint32PtrSliceConsumer)
				})
			} else {
				mc(valTestUint32PtrSliceConsumer)
			}
		})
	}
}

func TestMustUint32PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint32PtrSliceConsumerFactory
		cf2   testUint32PtrSliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					calls++
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					calls++
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					calls++
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint32PtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					calls++
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					calls++
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint32PtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustUint32PtrSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustUint32PtrSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustUint32PtrSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(errTestUint32PtrSliceConsumer.Error(), func() {
					cmc(valTestUint32PtrSliceConsumer)
				})
			} else {
				cmc(valTestUint32PtrSliceConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint32PtrSliceConsumer_ToSilentUint32PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint32PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					return errTestUint32PtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint32PtrSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentUint32PtrSliceConsumer()
			r.NotNil(sc)

			sc(valTestUint32PtrSliceConsumer)
		})
	}
}

func TestMustUint32PtrSliceConsumer_ToUint32PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint32PtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint32PtrSliceConsumer {
				return func(v []*uint32) error {
					require.Equal(t, valTestUint32PtrSliceConsumer, v)
					return errTestUint32PtrSliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint32PtrSliceConsumer()
			r.NotNil(mc)

			c = mc.ToUint32PtrSliceConsumer()
			r.NotNil(c)

			err := c(valTestUint32PtrSliceConsumer)
			if tt.err {
				r.EqualError(err, errTestUint32PtrSliceConsumer.Error())
			}
		})
	}
}
