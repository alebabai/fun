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
	testBoolPtrSliceConsumerValue []*bool
	testBoolPtrSliceConsumerError = errors.New("error")
)

type testBoolPtrSliceConsumerFactory func(t *testing.T) BoolPtrSliceConsumer

func TestBoolPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolPtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					require.Equal(t, testBoolPtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					require.Equal(t, testBoolPtrSliceConsumerValue, v)
					return testBoolPtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testBoolPtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testBoolPtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestBoolPtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolPtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					require.Equal(t, testBoolPtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					require.Equal(t, testBoolPtrSliceConsumerValue, v)
					return testBoolPtrSliceConsumerError
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

			err := c(testBoolPtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testBoolPtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestBoolPtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testBoolPtrSliceConsumerFactory
		cf2   testBoolPtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, testBoolPtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, testBoolPtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, testBoolPtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testBoolPtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, testBoolPtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, testBoolPtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolPtrSliceConsumer {
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
			err := cc(testBoolPtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testBoolPtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestBoolPtrSliceConsumer_ToSilentBoolPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolPtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					require.Equal(t, testBoolPtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					require.Equal(t, testBoolPtrSliceConsumerValue, v)
					return testBoolPtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentBoolPtrSliceConsumer()
			r.NotNil(sc)

			sc(testBoolPtrSliceConsumerValue)
		})
	}
}

func TestBoolPtrSliceConsumer_ToMustBoolPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolPtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					require.Equal(t, testBoolPtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					require.Equal(t, testBoolPtrSliceConsumerValue, v)
					return testBoolPtrSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustBoolPtrSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testBoolPtrSliceConsumerError.Error(), func() {
					mc(testBoolPtrSliceConsumerValue)
				})
			} else {
				mc(testBoolPtrSliceConsumerValue)
			}
		})
	}
}

func TestSilentBoolPtrSliceConsumer(t *testing.T) {
	var sc SilentBoolPtrSliceConsumer = func(v []*bool) {
		require.Equal(t, testBoolPtrSliceConsumerValue, v)
		return
	}
	sc(testBoolPtrSliceConsumerValue)
}

func TestSilentBoolPtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testBoolPtrSliceConsumerFactory
		cf2   testBoolPtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, testBoolPtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, testBoolPtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, testBoolPtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testBoolPtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, testBoolPtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, testBoolPtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolPtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentBoolPtrSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentBoolPtrSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentBoolPtrSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testBoolPtrSliceConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustBoolPtrSliceConsumer(t *testing.T) {
	var sc SilentBoolPtrSliceConsumer = func(v []*bool) {
		require.Equal(t, testBoolPtrSliceConsumerValue, v)
		return
	}
	sc(testBoolPtrSliceConsumerValue)
}

func TestMustBoolPtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testBoolPtrSliceConsumerFactory
		cf2   testBoolPtrSliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, testBoolPtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, testBoolPtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, testBoolPtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testBoolPtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, testBoolPtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, testBoolPtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolPtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustBoolPtrSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustBoolPtrSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustBoolPtrSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testBoolPtrSliceConsumerError.Error(), func() {
					cmc(testBoolPtrSliceConsumerValue)
				})
			} else {
				cmc(testBoolPtrSliceConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustBoolPtrSliceConsumer_ToSilentBoolPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolPtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					require.Equal(t, testBoolPtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					require.Equal(t, testBoolPtrSliceConsumerValue, v)
					return testBoolPtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustBoolPtrSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentBoolPtrSliceConsumer()
			r.NotNil(sc)

			sc(testBoolPtrSliceConsumerValue)
		})
	}
}

func TestMustBoolPtrSliceConsumer_ToBoolPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolPtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					require.Equal(t, testBoolPtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					require.Equal(t, testBoolPtrSliceConsumerValue, v)
					return testBoolPtrSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustBoolPtrSliceConsumer()
			r.NotNil(mc)

			c = mc.ToBoolPtrSliceConsumer()
			r.NotNil(c)

			err := c(testBoolPtrSliceConsumerValue)
			if tt.err {
				r.EqualError(err, testBoolPtrSliceConsumerError.Error())
			}
		})
	}
}
