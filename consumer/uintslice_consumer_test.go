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
	testUintSliceConsumerValue []uint
	testUintSliceConsumerError = errors.New("error")
)

type testUintSliceConsumerFactory func(t *testing.T) UintSliceConsumer

func TestUintSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					require.Equal(t, testUintSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					require.Equal(t, testUintSliceConsumerValue, v)
					return testUintSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testUintSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testUintSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUintSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					require.Equal(t, testUintSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					require.Equal(t, testUintSliceConsumerValue, v)
					return testUintSliceConsumerError
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

			err := c(testUintSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testUintSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUintSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUintSliceConsumerFactory
		cf2   testUintSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, testUintSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, testUintSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, testUintSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUintSliceConsumerError
				}
			},
			cf2: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, testUintSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, testUintSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintSliceConsumer {
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
			err := cc(testUintSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testUintSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestUintSliceConsumer_ToSilentUintSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					require.Equal(t, testUintSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					require.Equal(t, testUintSliceConsumerValue, v)
					return testUintSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentUintSliceConsumer()
			r.NotNil(sc)

			sc(testUintSliceConsumerValue)
		})
	}
}

func TestUintSliceConsumer_ToMustUintSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					require.Equal(t, testUintSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					require.Equal(t, testUintSliceConsumerValue, v)
					return testUintSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustUintSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testUintSliceConsumerError.Error(), func() {
					mc(testUintSliceConsumerValue)
				})
			} else {
				mc(testUintSliceConsumerValue)
			}
		})
	}
}

func TestSilentUintSliceConsumer(t *testing.T) {
	var sc SilentUintSliceConsumer = func(v []uint) {
		require.Equal(t, testUintSliceConsumerValue, v)
		return
	}
	sc(testUintSliceConsumerValue)
}

func TestSilentUintSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUintSliceConsumerFactory
		cf2   testUintSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, testUintSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, testUintSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, testUintSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUintSliceConsumerError
				}
			},
			cf2: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, testUintSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, testUintSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentUintSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentUintSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentUintSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testUintSliceConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUintSliceConsumer(t *testing.T) {
	var sc SilentUintSliceConsumer = func(v []uint) {
		require.Equal(t, testUintSliceConsumerValue, v)
		return
	}
	sc(testUintSliceConsumerValue)
}

func TestMustUintSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUintSliceConsumerFactory
		cf2   testUintSliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, testUintSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, testUintSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, testUintSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUintSliceConsumerError
				}
			},
			cf2: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, testUintSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, testUintSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustUintSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustUintSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustUintSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testUintSliceConsumerError.Error(), func() {
					cmc(testUintSliceConsumerValue)
				})
			} else {
				cmc(testUintSliceConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUintSliceConsumer_ToSilentUintSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					require.Equal(t, testUintSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					require.Equal(t, testUintSliceConsumerValue, v)
					return testUintSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUintSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentUintSliceConsumer()
			r.NotNil(sc)

			sc(testUintSliceConsumerValue)
		})
	}
}

func TestMustUintSliceConsumer_ToUintSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					require.Equal(t, testUintSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					require.Equal(t, testUintSliceConsumerValue, v)
					return testUintSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUintSliceConsumer()
			r.NotNil(mc)

			c = mc.ToUintSliceConsumer()
			r.NotNil(c)

			err := c(testUintSliceConsumerValue)
			if tt.err {
				r.EqualError(err, testUintSliceConsumerError.Error())
			}
		})
	}
}
