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
	testUint32SliceConsumerValue []uint32
	testUint32SliceConsumerError = errors.New("error")
)

type testUint32SliceConsumerFactory func(t *testing.T) Uint32SliceConsumer

func TestUint32SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint32SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32SliceConsumer {
				return func(v []uint32) error {
					require.Equal(t, testUint32SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint32SliceConsumer {
				return func(v []uint32) error {
					require.Equal(t, testUint32SliceConsumerValue, v)
					return testUint32SliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testUint32SliceConsumerValue)
			if err != nil {
				r.EqualError(err, testUint32SliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint32SliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint32SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32SliceConsumer {
				return func(v []uint32) error {
					require.Equal(t, testUint32SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint32SliceConsumer {
				return func(v []uint32) error {
					require.Equal(t, testUint32SliceConsumerValue, v)
					return testUint32SliceConsumerError
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

			err := c(testUint32SliceConsumerValue)
			if err != nil {
				r.EqualError(err, testUint32SliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint32SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint32SliceConsumerFactory
		cf2   testUint32SliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint32SliceConsumer {
				return func(v []uint32) error {
					calls++
					require.Equal(t, testUint32SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint32SliceConsumer {
				return func(v []uint32) error {
					calls++
					require.Equal(t, testUint32SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint32SliceConsumer {
				return func(v []uint32) error {
					calls++
					require.Equal(t, testUint32SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUint32SliceConsumerError
				}
			},
			cf2: func(t *testing.T) Uint32SliceConsumer {
				return func(v []uint32) error {
					calls++
					require.Equal(t, testUint32SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint32SliceConsumer {
				return func(v []uint32) error {
					calls++
					require.Equal(t, testUint32SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint32SliceConsumer {
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
			err := cc(testUint32SliceConsumerValue)
			if err != nil {
				r.EqualError(err, testUint32SliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestUint32SliceConsumer_ToSilentUint32SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint32SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32SliceConsumer {
				return func(v []uint32) error {
					require.Equal(t, testUint32SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint32SliceConsumer {
				return func(v []uint32) error {
					require.Equal(t, testUint32SliceConsumerValue, v)
					return testUint32SliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentUint32SliceConsumer()
			r.NotNil(sc)

			sc(testUint32SliceConsumerValue)
		})
	}
}

func TestUint32SliceConsumer_ToMustUint32SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint32SliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32SliceConsumer {
				return func(v []uint32) error {
					require.Equal(t, testUint32SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint32SliceConsumer {
				return func(v []uint32) error {
					require.Equal(t, testUint32SliceConsumerValue, v)
					return testUint32SliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustUint32SliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testUint32SliceConsumerError.Error(), func() {
					mc(testUint32SliceConsumerValue)
				})
			} else {
				mc(testUint32SliceConsumerValue)
			}
		})
	}
}

func TestSilentUint32SliceConsumer(t *testing.T) {
	var sc SilentUint32SliceConsumer = func(v []uint32) {
		require.Equal(t, testUint32SliceConsumerValue, v)
		return
	}
	sc(testUint32SliceConsumerValue)
}

func TestSilentUint32SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint32SliceConsumerFactory
		cf2   testUint32SliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint32SliceConsumer {
				return func(v []uint32) error {
					calls++
					require.Equal(t, testUint32SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint32SliceConsumer {
				return func(v []uint32) error {
					calls++
					require.Equal(t, testUint32SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint32SliceConsumer {
				return func(v []uint32) error {
					calls++
					require.Equal(t, testUint32SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUint32SliceConsumerError
				}
			},
			cf2: func(t *testing.T) Uint32SliceConsumer {
				return func(v []uint32) error {
					calls++
					require.Equal(t, testUint32SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint32SliceConsumer {
				return func(v []uint32) error {
					calls++
					require.Equal(t, testUint32SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint32SliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentUint32SliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentUint32SliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentUint32SliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testUint32SliceConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint32SliceConsumer(t *testing.T) {
	var sc SilentUint32SliceConsumer = func(v []uint32) {
		require.Equal(t, testUint32SliceConsumerValue, v)
		return
	}
	sc(testUint32SliceConsumerValue)
}

func TestMustUint32SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint32SliceConsumerFactory
		cf2   testUint32SliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint32SliceConsumer {
				return func(v []uint32) error {
					calls++
					require.Equal(t, testUint32SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint32SliceConsumer {
				return func(v []uint32) error {
					calls++
					require.Equal(t, testUint32SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint32SliceConsumer {
				return func(v []uint32) error {
					calls++
					require.Equal(t, testUint32SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUint32SliceConsumerError
				}
			},
			cf2: func(t *testing.T) Uint32SliceConsumer {
				return func(v []uint32) error {
					calls++
					require.Equal(t, testUint32SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint32SliceConsumer {
				return func(v []uint32) error {
					calls++
					require.Equal(t, testUint32SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint32SliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustUint32SliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustUint32SliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustUint32SliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testUint32SliceConsumerError.Error(), func() {
					cmc(testUint32SliceConsumerValue)
				})
			} else {
				cmc(testUint32SliceConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint32SliceConsumer_ToSilentUint32SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint32SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32SliceConsumer {
				return func(v []uint32) error {
					require.Equal(t, testUint32SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint32SliceConsumer {
				return func(v []uint32) error {
					require.Equal(t, testUint32SliceConsumerValue, v)
					return testUint32SliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint32SliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentUint32SliceConsumer()
			r.NotNil(sc)

			sc(testUint32SliceConsumerValue)
		})
	}
}

func TestMustUint32SliceConsumer_ToUint32SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint32SliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32SliceConsumer {
				return func(v []uint32) error {
					require.Equal(t, testUint32SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint32SliceConsumer {
				return func(v []uint32) error {
					require.Equal(t, testUint32SliceConsumerValue, v)
					return testUint32SliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint32SliceConsumer()
			r.NotNil(mc)

			c = mc.ToUint32SliceConsumer()
			r.NotNil(c)

			err := c(testUint32SliceConsumerValue)
			if tt.err {
				r.EqualError(err, testUint32SliceConsumerError.Error())
			}
		})
	}
}
