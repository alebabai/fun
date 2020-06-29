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
	testIntSliceConsumerValue []int
	testIntSliceConsumerError = errors.New("error")
)

type testIntSliceConsumerFactory func(t *testing.T) IntSliceConsumer

func TestIntSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					require.Equal(t, testIntSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					require.Equal(t, testIntSliceConsumerValue, v)
					return testIntSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testIntSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testIntSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestIntSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					require.Equal(t, testIntSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					require.Equal(t, testIntSliceConsumerValue, v)
					return testIntSliceConsumerError
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

			err := c(testIntSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testIntSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestIntSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testIntSliceConsumerFactory
		cf2   testIntSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, testIntSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, testIntSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, testIntSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testIntSliceConsumerError
				}
			},
			cf2: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, testIntSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, testIntSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntSliceConsumer {
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
			err := cc(testIntSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testIntSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestIntSliceConsumer_ToSilentIntSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					require.Equal(t, testIntSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					require.Equal(t, testIntSliceConsumerValue, v)
					return testIntSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentIntSliceConsumer()
			r.NotNil(sc)

			sc(testIntSliceConsumerValue)
		})
	}
}

func TestIntSliceConsumer_ToMustIntSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					require.Equal(t, testIntSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					require.Equal(t, testIntSliceConsumerValue, v)
					return testIntSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustIntSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testIntSliceConsumerError.Error(), func() {
					mc(testIntSliceConsumerValue)
				})
			} else {
				mc(testIntSliceConsumerValue)
			}
		})
	}
}

func TestSilentIntSliceConsumer(t *testing.T) {
	var sc SilentIntSliceConsumer = func(v []int) {
		require.Equal(t, testIntSliceConsumerValue, v)
		return
	}
	sc(testIntSliceConsumerValue)
}

func TestSilentIntSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testIntSliceConsumerFactory
		cf2   testIntSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, testIntSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, testIntSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, testIntSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testIntSliceConsumerError
				}
			},
			cf2: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, testIntSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, testIntSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentIntSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentIntSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentIntSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testIntSliceConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustIntSliceConsumer(t *testing.T) {
	var sc SilentIntSliceConsumer = func(v []int) {
		require.Equal(t, testIntSliceConsumerValue, v)
		return
	}
	sc(testIntSliceConsumerValue)
}

func TestMustIntSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testIntSliceConsumerFactory
		cf2   testIntSliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, testIntSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, testIntSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, testIntSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testIntSliceConsumerError
				}
			},
			cf2: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, testIntSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					calls++
					require.Equal(t, testIntSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustIntSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustIntSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustIntSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testIntSliceConsumerError.Error(), func() {
					cmc(testIntSliceConsumerValue)
				})
			} else {
				cmc(testIntSliceConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustIntSliceConsumer_ToSilentIntSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					require.Equal(t, testIntSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					require.Equal(t, testIntSliceConsumerValue, v)
					return testIntSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustIntSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentIntSliceConsumer()
			r.NotNil(sc)

			sc(testIntSliceConsumerValue)
		})
	}
}

func TestMustIntSliceConsumer_ToIntSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					require.Equal(t, testIntSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntSliceConsumer {
				return func(v []int) error {
					require.Equal(t, testIntSliceConsumerValue, v)
					return testIntSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustIntSliceConsumer()
			r.NotNil(mc)

			c = mc.ToIntSliceConsumer()
			r.NotNil(c)

			err := c(testIntSliceConsumerValue)
			if tt.err {
				r.EqualError(err, testIntSliceConsumerError.Error())
			}
		})
	}
}
