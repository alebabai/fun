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
	testFloat64SliceConsumerValue []float64
	testFloat64SliceConsumerError = errors.New("error")
)

type testFloat64SliceConsumerFactory func(t *testing.T) Float64SliceConsumer

func TestFloat64SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat64SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64SliceConsumer {
				return func(v []float64) error {
					require.Equal(t, testFloat64SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float64SliceConsumer {
				return func(v []float64) error {
					require.Equal(t, testFloat64SliceConsumerValue, v)
					return testFloat64SliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testFloat64SliceConsumerValue)
			if err != nil {
				r.EqualError(err, testFloat64SliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestFloat64SliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat64SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64SliceConsumer {
				return func(v []float64) error {
					require.Equal(t, testFloat64SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float64SliceConsumer {
				return func(v []float64) error {
					require.Equal(t, testFloat64SliceConsumerValue, v)
					return testFloat64SliceConsumerError
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

			err := c(testFloat64SliceConsumerValue)
			if err != nil {
				r.EqualError(err, testFloat64SliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestFloat64SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testFloat64SliceConsumerFactory
		cf2   testFloat64SliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float64SliceConsumer {
				return func(v []float64) error {
					calls++
					require.Equal(t, testFloat64SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64SliceConsumer {
				return func(v []float64) error {
					calls++
					require.Equal(t, testFloat64SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Float64SliceConsumer {
				return func(v []float64) error {
					calls++
					require.Equal(t, testFloat64SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testFloat64SliceConsumerError
				}
			},
			cf2: func(t *testing.T) Float64SliceConsumer {
				return func(v []float64) error {
					calls++
					require.Equal(t, testFloat64SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float64SliceConsumer {
				return func(v []float64) error {
					calls++
					require.Equal(t, testFloat64SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64SliceConsumer {
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
			err := cc(testFloat64SliceConsumerValue)
			if err != nil {
				r.EqualError(err, testFloat64SliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestFloat64SliceConsumer_ToSilentFloat64SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat64SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64SliceConsumer {
				return func(v []float64) error {
					require.Equal(t, testFloat64SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float64SliceConsumer {
				return func(v []float64) error {
					require.Equal(t, testFloat64SliceConsumerValue, v)
					return testFloat64SliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentFloat64SliceConsumer()
			r.NotNil(sc)

			sc(testFloat64SliceConsumerValue)
		})
	}
}

func TestFloat64SliceConsumer_ToMustFloat64SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat64SliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64SliceConsumer {
				return func(v []float64) error {
					require.Equal(t, testFloat64SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float64SliceConsumer {
				return func(v []float64) error {
					require.Equal(t, testFloat64SliceConsumerValue, v)
					return testFloat64SliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustFloat64SliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testFloat64SliceConsumerError.Error(), func() {
					mc(testFloat64SliceConsumerValue)
				})
			} else {
				mc(testFloat64SliceConsumerValue)
			}
		})
	}
}

func TestSilentFloat64SliceConsumer(t *testing.T) {
	var sc SilentFloat64SliceConsumer = func(v []float64) {
		require.Equal(t, testFloat64SliceConsumerValue, v)
		return
	}
	sc(testFloat64SliceConsumerValue)
}

func TestSilentFloat64SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testFloat64SliceConsumerFactory
		cf2   testFloat64SliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float64SliceConsumer {
				return func(v []float64) error {
					calls++
					require.Equal(t, testFloat64SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64SliceConsumer {
				return func(v []float64) error {
					calls++
					require.Equal(t, testFloat64SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Float64SliceConsumer {
				return func(v []float64) error {
					calls++
					require.Equal(t, testFloat64SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testFloat64SliceConsumerError
				}
			},
			cf2: func(t *testing.T) Float64SliceConsumer {
				return func(v []float64) error {
					calls++
					require.Equal(t, testFloat64SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float64SliceConsumer {
				return func(v []float64) error {
					calls++
					require.Equal(t, testFloat64SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64SliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentFloat64SliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentFloat64SliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentFloat64SliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testFloat64SliceConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustFloat64SliceConsumer(t *testing.T) {
	var sc SilentFloat64SliceConsumer = func(v []float64) {
		require.Equal(t, testFloat64SliceConsumerValue, v)
		return
	}
	sc(testFloat64SliceConsumerValue)
}

func TestMustFloat64SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testFloat64SliceConsumerFactory
		cf2   testFloat64SliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float64SliceConsumer {
				return func(v []float64) error {
					calls++
					require.Equal(t, testFloat64SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64SliceConsumer {
				return func(v []float64) error {
					calls++
					require.Equal(t, testFloat64SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Float64SliceConsumer {
				return func(v []float64) error {
					calls++
					require.Equal(t, testFloat64SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testFloat64SliceConsumerError
				}
			},
			cf2: func(t *testing.T) Float64SliceConsumer {
				return func(v []float64) error {
					calls++
					require.Equal(t, testFloat64SliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float64SliceConsumer {
				return func(v []float64) error {
					calls++
					require.Equal(t, testFloat64SliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64SliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustFloat64SliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustFloat64SliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustFloat64SliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testFloat64SliceConsumerError.Error(), func() {
					cmc(testFloat64SliceConsumerValue)
				})
			} else {
				cmc(testFloat64SliceConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustFloat64SliceConsumer_ToSilentFloat64SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat64SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64SliceConsumer {
				return func(v []float64) error {
					require.Equal(t, testFloat64SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float64SliceConsumer {
				return func(v []float64) error {
					require.Equal(t, testFloat64SliceConsumerValue, v)
					return testFloat64SliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustFloat64SliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentFloat64SliceConsumer()
			r.NotNil(sc)

			sc(testFloat64SliceConsumerValue)
		})
	}
}

func TestMustFloat64SliceConsumer_ToFloat64SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat64SliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64SliceConsumer {
				return func(v []float64) error {
					require.Equal(t, testFloat64SliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float64SliceConsumer {
				return func(v []float64) error {
					require.Equal(t, testFloat64SliceConsumerValue, v)
					return testFloat64SliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustFloat64SliceConsumer()
			r.NotNil(mc)

			c = mc.ToFloat64SliceConsumer()
			r.NotNil(c)

			err := c(testFloat64SliceConsumerValue)
			if tt.err {
				r.EqualError(err, testFloat64SliceConsumerError.Error())
			}
		})
	}
}
