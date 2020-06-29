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
	testFloat64PtrSliceConsumerValue []*float64
	testFloat64PtrSliceConsumerError = errors.New("error")
)

type testFloat64PtrSliceConsumerFactory func(t *testing.T) Float64PtrSliceConsumer

func TestFloat64PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat64PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					require.Equal(t, testFloat64PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					require.Equal(t, testFloat64PtrSliceConsumerValue, v)
					return testFloat64PtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testFloat64PtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testFloat64PtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestFloat64PtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat64PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					require.Equal(t, testFloat64PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					require.Equal(t, testFloat64PtrSliceConsumerValue, v)
					return testFloat64PtrSliceConsumerError
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

			err := c(testFloat64PtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testFloat64PtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestFloat64PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testFloat64PtrSliceConsumerFactory
		cf2   testFloat64PtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, testFloat64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, testFloat64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, testFloat64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testFloat64PtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, testFloat64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, testFloat64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64PtrSliceConsumer {
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
			err := cc(testFloat64PtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testFloat64PtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestFloat64PtrSliceConsumer_ToSilentFloat64PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat64PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					require.Equal(t, testFloat64PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					require.Equal(t, testFloat64PtrSliceConsumerValue, v)
					return testFloat64PtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentFloat64PtrSliceConsumer()
			r.NotNil(sc)

			sc(testFloat64PtrSliceConsumerValue)
		})
	}
}

func TestFloat64PtrSliceConsumer_ToMustFloat64PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat64PtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					require.Equal(t, testFloat64PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					require.Equal(t, testFloat64PtrSliceConsumerValue, v)
					return testFloat64PtrSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustFloat64PtrSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testFloat64PtrSliceConsumerError.Error(), func() {
					mc(testFloat64PtrSliceConsumerValue)
				})
			} else {
				mc(testFloat64PtrSliceConsumerValue)
			}
		})
	}
}

func TestSilentFloat64PtrSliceConsumer(t *testing.T) {
	var sc SilentFloat64PtrSliceConsumer = func(v []*float64) {
		require.Equal(t, testFloat64PtrSliceConsumerValue, v)
		return
	}
	sc(testFloat64PtrSliceConsumerValue)
}

func TestSilentFloat64PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testFloat64PtrSliceConsumerFactory
		cf2   testFloat64PtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, testFloat64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, testFloat64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, testFloat64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testFloat64PtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, testFloat64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, testFloat64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64PtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentFloat64PtrSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentFloat64PtrSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentFloat64PtrSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testFloat64PtrSliceConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustFloat64PtrSliceConsumer(t *testing.T) {
	var sc SilentFloat64PtrSliceConsumer = func(v []*float64) {
		require.Equal(t, testFloat64PtrSliceConsumerValue, v)
		return
	}
	sc(testFloat64PtrSliceConsumerValue)
}

func TestMustFloat64PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testFloat64PtrSliceConsumerFactory
		cf2   testFloat64PtrSliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, testFloat64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, testFloat64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, testFloat64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testFloat64PtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, testFloat64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, testFloat64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64PtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustFloat64PtrSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustFloat64PtrSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustFloat64PtrSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testFloat64PtrSliceConsumerError.Error(), func() {
					cmc(testFloat64PtrSliceConsumerValue)
				})
			} else {
				cmc(testFloat64PtrSliceConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustFloat64PtrSliceConsumer_ToSilentFloat64PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat64PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					require.Equal(t, testFloat64PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					require.Equal(t, testFloat64PtrSliceConsumerValue, v)
					return testFloat64PtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustFloat64PtrSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentFloat64PtrSliceConsumer()
			r.NotNil(sc)

			sc(testFloat64PtrSliceConsumerValue)
		})
	}
}

func TestMustFloat64PtrSliceConsumer_ToFloat64PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat64PtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					require.Equal(t, testFloat64PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					require.Equal(t, testFloat64PtrSliceConsumerValue, v)
					return testFloat64PtrSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustFloat64PtrSliceConsumer()
			r.NotNil(mc)

			c = mc.ToFloat64PtrSliceConsumer()
			r.NotNil(c)

			err := c(testFloat64PtrSliceConsumerValue)
			if tt.err {
				r.EqualError(err, testFloat64PtrSliceConsumerError.Error())
			}
		})
	}
}
