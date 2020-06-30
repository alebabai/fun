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
	valTestUintPtrSliceConsumer []*uint
	errTestUintPtrSliceConsumer = errors.New("error")
)

type testUintPtrSliceConsumerFactory func(t *testing.T) UintPtrSliceConsumer

func TestUintPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintPtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintPtrSliceConsumer {
				return func(v []*uint) error {
					require.Equal(t, valTestUintPtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) UintPtrSliceConsumer {
				return func(v []*uint) error {
					require.Equal(t, valTestUintPtrSliceConsumer, v)
					return errTestUintPtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestUintPtrSliceConsumer)
			if err != nil {
				r.EqualError(err, errTestUintPtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUintPtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintPtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintPtrSliceConsumer {
				return func(v []*uint) error {
					require.Equal(t, valTestUintPtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) UintPtrSliceConsumer {
				return func(v []*uint) error {
					require.Equal(t, valTestUintPtrSliceConsumer, v)
					return errTestUintPtrSliceConsumer
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

			err := c(valTestUintPtrSliceConsumer)
			if err != nil {
				r.EqualError(err, errTestUintPtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUintPtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUintPtrSliceConsumerFactory
		cf2   testUintPtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) UintPtrSliceConsumer {
				return func(v []*uint) error {
					calls++
					require.Equal(t, valTestUintPtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintPtrSliceConsumer {
				return func(v []*uint) error {
					calls++
					require.Equal(t, valTestUintPtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) UintPtrSliceConsumer {
				return func(v []*uint) error {
					calls++
					require.Equal(t, valTestUintPtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUintPtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) UintPtrSliceConsumer {
				return func(v []*uint) error {
					calls++
					require.Equal(t, valTestUintPtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) UintPtrSliceConsumer {
				return func(v []*uint) error {
					calls++
					require.Equal(t, valTestUintPtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintPtrSliceConsumer {
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
			err := cc(valTestUintPtrSliceConsumer)
			if err != nil {
				r.EqualError(err, errTestUintPtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestUintPtrSliceConsumer_ToSilentUintPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintPtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintPtrSliceConsumer {
				return func(v []*uint) error {
					require.Equal(t, valTestUintPtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) UintPtrSliceConsumer {
				return func(v []*uint) error {
					require.Equal(t, valTestUintPtrSliceConsumer, v)
					return errTestUintPtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentUintPtrSliceConsumer()
			r.NotNil(sc)

			sc(valTestUintPtrSliceConsumer)
		})
	}
}

func TestUintPtrSliceConsumer_ToMustUintPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintPtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintPtrSliceConsumer {
				return func(v []*uint) error {
					require.Equal(t, valTestUintPtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) UintPtrSliceConsumer {
				return func(v []*uint) error {
					require.Equal(t, valTestUintPtrSliceConsumer, v)
					return errTestUintPtrSliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustUintPtrSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestUintPtrSliceConsumer.Error(), func() {
					mc(valTestUintPtrSliceConsumer)
				})
			} else {
				mc(valTestUintPtrSliceConsumer)
			}
		})
	}
}

func TestSilentUintPtrSliceConsumer(t *testing.T) {
	var sc SilentUintPtrSliceConsumer = func(v []*uint) {
		require.Equal(t, valTestUintPtrSliceConsumer, v)
		return
	}
	sc(valTestUintPtrSliceConsumer)
}

func TestSilentUintPtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUintPtrSliceConsumerFactory
		cf2   testUintPtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) UintPtrSliceConsumer {
				return func(v []*uint) error {
					calls++
					require.Equal(t, valTestUintPtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintPtrSliceConsumer {
				return func(v []*uint) error {
					calls++
					require.Equal(t, valTestUintPtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) UintPtrSliceConsumer {
				return func(v []*uint) error {
					calls++
					require.Equal(t, valTestUintPtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUintPtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) UintPtrSliceConsumer {
				return func(v []*uint) error {
					calls++
					require.Equal(t, valTestUintPtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) UintPtrSliceConsumer {
				return func(v []*uint) error {
					calls++
					require.Equal(t, valTestUintPtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintPtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentUintPtrSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentUintPtrSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentUintPtrSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestUintPtrSliceConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUintPtrSliceConsumer(t *testing.T) {
	var sc SilentUintPtrSliceConsumer = func(v []*uint) {
		require.Equal(t, valTestUintPtrSliceConsumer, v)
		return
	}
	sc(valTestUintPtrSliceConsumer)
}

func TestMustUintPtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUintPtrSliceConsumerFactory
		cf2   testUintPtrSliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) UintPtrSliceConsumer {
				return func(v []*uint) error {
					calls++
					require.Equal(t, valTestUintPtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintPtrSliceConsumer {
				return func(v []*uint) error {
					calls++
					require.Equal(t, valTestUintPtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) UintPtrSliceConsumer {
				return func(v []*uint) error {
					calls++
					require.Equal(t, valTestUintPtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUintPtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) UintPtrSliceConsumer {
				return func(v []*uint) error {
					calls++
					require.Equal(t, valTestUintPtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) UintPtrSliceConsumer {
				return func(v []*uint) error {
					calls++
					require.Equal(t, valTestUintPtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintPtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustUintPtrSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustUintPtrSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustUintPtrSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(errTestUintPtrSliceConsumer.Error(), func() {
					cmc(valTestUintPtrSliceConsumer)
				})
			} else {
				cmc(valTestUintPtrSliceConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUintPtrSliceConsumer_ToSilentUintPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintPtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintPtrSliceConsumer {
				return func(v []*uint) error {
					require.Equal(t, valTestUintPtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) UintPtrSliceConsumer {
				return func(v []*uint) error {
					require.Equal(t, valTestUintPtrSliceConsumer, v)
					return errTestUintPtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUintPtrSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentUintPtrSliceConsumer()
			r.NotNil(sc)

			sc(valTestUintPtrSliceConsumer)
		})
	}
}

func TestMustUintPtrSliceConsumer_ToUintPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintPtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintPtrSliceConsumer {
				return func(v []*uint) error {
					require.Equal(t, valTestUintPtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) UintPtrSliceConsumer {
				return func(v []*uint) error {
					require.Equal(t, valTestUintPtrSliceConsumer, v)
					return errTestUintPtrSliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUintPtrSliceConsumer()
			r.NotNil(mc)

			c = mc.ToUintPtrSliceConsumer()
			r.NotNil(c)

			err := c(valTestUintPtrSliceConsumer)
			if tt.err {
				r.EqualError(err, errTestUintPtrSliceConsumer.Error())
			}
		})
	}
}
