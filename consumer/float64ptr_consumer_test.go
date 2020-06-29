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
	testFloat64PtrConsumerValue *float64
	testFloat64PtrConsumerError = errors.New("error")
)

type testFloat64PtrConsumerFactory func(t *testing.T) Float64PtrConsumer

func TestFloat64PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat64PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64PtrConsumer {
				return func(v *float64) error {
					require.Equal(t, testFloat64PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float64PtrConsumer {
				return func(v *float64) error {
					require.Equal(t, testFloat64PtrConsumerValue, v)
					return testFloat64PtrConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testFloat64PtrConsumerValue)
			if err != nil {
				r.EqualError(err, testFloat64PtrConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestFloat64PtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat64PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64PtrConsumer {
				return func(v *float64) error {
					require.Equal(t, testFloat64PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float64PtrConsumer {
				return func(v *float64) error {
					require.Equal(t, testFloat64PtrConsumerValue, v)
					return testFloat64PtrConsumerError
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

			err := c(testFloat64PtrConsumerValue)
			if err != nil {
				r.EqualError(err, testFloat64PtrConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestFloat64PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testFloat64PtrConsumerFactory
		cf2   testFloat64PtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float64PtrConsumer {
				return func(v *float64) error {
					calls++
					require.Equal(t, testFloat64PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64PtrConsumer {
				return func(v *float64) error {
					calls++
					require.Equal(t, testFloat64PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Float64PtrConsumer {
				return func(v *float64) error {
					calls++
					require.Equal(t, testFloat64PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testFloat64PtrConsumerError
				}
			},
			cf2: func(t *testing.T) Float64PtrConsumer {
				return func(v *float64) error {
					calls++
					require.Equal(t, testFloat64PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float64PtrConsumer {
				return func(v *float64) error {
					calls++
					require.Equal(t, testFloat64PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64PtrConsumer {
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
			err := cc(testFloat64PtrConsumerValue)
			if err != nil {
				r.EqualError(err, testFloat64PtrConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestFloat64PtrConsumer_ToSilentFloat64PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat64PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64PtrConsumer {
				return func(v *float64) error {
					require.Equal(t, testFloat64PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float64PtrConsumer {
				return func(v *float64) error {
					require.Equal(t, testFloat64PtrConsumerValue, v)
					return testFloat64PtrConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentFloat64PtrConsumer()
			r.NotNil(sc)

			sc(testFloat64PtrConsumerValue)
		})
	}
}

func TestFloat64PtrConsumer_ToMustFloat64PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat64PtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64PtrConsumer {
				return func(v *float64) error {
					require.Equal(t, testFloat64PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float64PtrConsumer {
				return func(v *float64) error {
					require.Equal(t, testFloat64PtrConsumerValue, v)
					return testFloat64PtrConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustFloat64PtrConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testFloat64PtrConsumerError.Error(), func() {
					mc(testFloat64PtrConsumerValue)
				})
			} else {
				mc(testFloat64PtrConsumerValue)
			}
		})
	}
}

func TestSilentFloat64PtrConsumer(t *testing.T) {
	var sc SilentFloat64PtrConsumer = func(v *float64) {
		require.Equal(t, testFloat64PtrConsumerValue, v)
		return
	}
	sc(testFloat64PtrConsumerValue)
}

func TestSilentFloat64PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testFloat64PtrConsumerFactory
		cf2   testFloat64PtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float64PtrConsumer {
				return func(v *float64) error {
					calls++
					require.Equal(t, testFloat64PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64PtrConsumer {
				return func(v *float64) error {
					calls++
					require.Equal(t, testFloat64PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Float64PtrConsumer {
				return func(v *float64) error {
					calls++
					require.Equal(t, testFloat64PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testFloat64PtrConsumerError
				}
			},
			cf2: func(t *testing.T) Float64PtrConsumer {
				return func(v *float64) error {
					calls++
					require.Equal(t, testFloat64PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float64PtrConsumer {
				return func(v *float64) error {
					calls++
					require.Equal(t, testFloat64PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64PtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentFloat64PtrConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentFloat64PtrConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentFloat64PtrConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testFloat64PtrConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustFloat64PtrConsumer(t *testing.T) {
	var sc SilentFloat64PtrConsumer = func(v *float64) {
		require.Equal(t, testFloat64PtrConsumerValue, v)
		return
	}
	sc(testFloat64PtrConsumerValue)
}

func TestMustFloat64PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testFloat64PtrConsumerFactory
		cf2   testFloat64PtrConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float64PtrConsumer {
				return func(v *float64) error {
					calls++
					require.Equal(t, testFloat64PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64PtrConsumer {
				return func(v *float64) error {
					calls++
					require.Equal(t, testFloat64PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Float64PtrConsumer {
				return func(v *float64) error {
					calls++
					require.Equal(t, testFloat64PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testFloat64PtrConsumerError
				}
			},
			cf2: func(t *testing.T) Float64PtrConsumer {
				return func(v *float64) error {
					calls++
					require.Equal(t, testFloat64PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float64PtrConsumer {
				return func(v *float64) error {
					calls++
					require.Equal(t, testFloat64PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64PtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustFloat64PtrConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustFloat64PtrConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustFloat64PtrConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testFloat64PtrConsumerError.Error(), func() {
					cmc(testFloat64PtrConsumerValue)
				})
			} else {
				cmc(testFloat64PtrConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustFloat64PtrConsumer_ToSilentFloat64PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat64PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64PtrConsumer {
				return func(v *float64) error {
					require.Equal(t, testFloat64PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float64PtrConsumer {
				return func(v *float64) error {
					require.Equal(t, testFloat64PtrConsumerValue, v)
					return testFloat64PtrConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustFloat64PtrConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentFloat64PtrConsumer()
			r.NotNil(sc)

			sc(testFloat64PtrConsumerValue)
		})
	}
}

func TestMustFloat64PtrConsumer_ToFloat64PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat64PtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64PtrConsumer {
				return func(v *float64) error {
					require.Equal(t, testFloat64PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float64PtrConsumer {
				return func(v *float64) error {
					require.Equal(t, testFloat64PtrConsumerValue, v)
					return testFloat64PtrConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustFloat64PtrConsumer()
			r.NotNil(mc)

			c = mc.ToFloat64PtrConsumer()
			r.NotNil(c)

			err := c(testFloat64PtrConsumerValue)
			if tt.err {
				r.EqualError(err, testFloat64PtrConsumerError.Error())
			}
		})
	}
}
