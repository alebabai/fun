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
	testUint64SliceConsumerValue []uint64
	testUint64SliceConsumerError = errors.New("error")
)

type testUint64SliceConsumerFactory func(t *testing.T) Uint64SliceConsumer

func TestUint64SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					require.Equal(t, testUint64SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					require.Equal(t, testUint64SliceConsumerValue, v)
					return testUint64SliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testUint64SliceConsumerValue)
			if err != nil {
				r.EqualError(err, testUint64SliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint64SliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					require.Equal(t, testUint64SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					require.Equal(t, testUint64SliceConsumerValue, v)
					return testUint64SliceConsumerError
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

			err := c(testUint64SliceConsumerValue)
			if err != nil {
				r.EqualError(err, testUint64SliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint64SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint64SliceConsumerFactory
		cf2   testUint64SliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, testUint64SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, testUint64SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, testUint64SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUint64SliceConsumerError
				}
			},
			cf2: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, testUint64SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, testUint64SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64SliceConsumer {
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
			err := cc(testUint64SliceConsumerValue)
			if err != nil {
				r.EqualError(err, testUint64SliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestUint64SliceConsumer_ToSilentUint64SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					require.Equal(t, testUint64SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					require.Equal(t, testUint64SliceConsumerValue, v)
					return testUint64SliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentUint64SliceConsumer()
			r.NotNil(sc)

			sc(testUint64SliceConsumerValue)
		})
	}
}

func TestUint64SliceConsumer_ToMustUint64SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64SliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					require.Equal(t, testUint64SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					require.Equal(t, testUint64SliceConsumerValue, v)
					return testUint64SliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustUint64SliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testUint64SliceConsumerError.Error(), func() {
					mc(testUint64SliceConsumerValue)
				})
			} else {
				mc(testUint64SliceConsumerValue)
			}
		})
	}
}

func TestSilentUint64SliceConsumer(t *testing.T) {
	var sc SilentUint64SliceConsumer = func(v []uint64) {
		require.Equal(t, testUint64SliceConsumerValue, v)
		return
	}
	sc(testUint64SliceConsumerValue)
}

func TestSilentUint64SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint64SliceConsumerFactory
		cf2   testUint64SliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, testUint64SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, testUint64SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, testUint64SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUint64SliceConsumerError
				}
			},
			cf2: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, testUint64SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, testUint64SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64SliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentUint64SliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentUint64SliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentUint64SliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testUint64SliceConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint64SliceConsumer(t *testing.T) {
	var sc SilentUint64SliceConsumer = func(v []uint64) {
		require.Equal(t, testUint64SliceConsumerValue, v)
		return
	}
	sc(testUint64SliceConsumerValue)
}

func TestMustUint64SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint64SliceConsumerFactory
		cf2   testUint64SliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, testUint64SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, testUint64SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, testUint64SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUint64SliceConsumerError
				}
			},
			cf2: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, testUint64SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, testUint64SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64SliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustUint64SliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustUint64SliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustUint64SliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testUint64SliceConsumerError.Error(), func() {
					cmc(testUint64SliceConsumerValue)
				})
			} else {
				cmc(testUint64SliceConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint64SliceConsumer_ToSilentUint64SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					require.Equal(t, testUint64SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					require.Equal(t, testUint64SliceConsumerValue, v)
					return testUint64SliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint64SliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentUint64SliceConsumer()
			r.NotNil(sc)

			sc(testUint64SliceConsumerValue)
		})
	}
}

func TestMustUint64SliceConsumer_ToUint64SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64SliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					require.Equal(t, testUint64SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					require.Equal(t, testUint64SliceConsumerValue, v)
					return testUint64SliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint64SliceConsumer()
			r.NotNil(mc)

			c = mc.ToUint64SliceConsumer()
			r.NotNil(c)

			err := c(testUint64SliceConsumerValue)
			if tt.err {
				r.EqualError(err, testUint64SliceConsumerError.Error())
			}
		})
	}
}
