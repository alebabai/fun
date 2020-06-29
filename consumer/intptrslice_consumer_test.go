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
	testIntPtrSliceConsumerValue []*int
	testIntPtrSliceConsumerError = errors.New("error")
)

type testIntPtrSliceConsumerFactory func(t *testing.T) IntPtrSliceConsumer

func TestIntPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntPtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntPtrSliceConsumer {
				return func(v []*int) error {
					require.Equal(t, testIntPtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntPtrSliceConsumer {
				return func(v []*int) error {
					require.Equal(t, testIntPtrSliceConsumerValue, v)
					return testIntPtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testIntPtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testIntPtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestIntPtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntPtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntPtrSliceConsumer {
				return func(v []*int) error {
					require.Equal(t, testIntPtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntPtrSliceConsumer {
				return func(v []*int) error {
					require.Equal(t, testIntPtrSliceConsumerValue, v)
					return testIntPtrSliceConsumerError
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

			err := c(testIntPtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testIntPtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestIntPtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testIntPtrSliceConsumerFactory
		cf2   testIntPtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) IntPtrSliceConsumer {
				return func(v []*int) error {
					calls++
					require.Equal(t, testIntPtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntPtrSliceConsumer {
				return func(v []*int) error {
					calls++
					require.Equal(t, testIntPtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) IntPtrSliceConsumer {
				return func(v []*int) error {
					calls++
					require.Equal(t, testIntPtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testIntPtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) IntPtrSliceConsumer {
				return func(v []*int) error {
					calls++
					require.Equal(t, testIntPtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) IntPtrSliceConsumer {
				return func(v []*int) error {
					calls++
					require.Equal(t, testIntPtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntPtrSliceConsumer {
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
			err := cc(testIntPtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testIntPtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestIntPtrSliceConsumer_ToSilentIntPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntPtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntPtrSliceConsumer {
				return func(v []*int) error {
					require.Equal(t, testIntPtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntPtrSliceConsumer {
				return func(v []*int) error {
					require.Equal(t, testIntPtrSliceConsumerValue, v)
					return testIntPtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentIntPtrSliceConsumer()
			r.NotNil(sc)

			sc(testIntPtrSliceConsumerValue)
		})
	}
}

func TestIntPtrSliceConsumer_ToMustIntPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntPtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntPtrSliceConsumer {
				return func(v []*int) error {
					require.Equal(t, testIntPtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntPtrSliceConsumer {
				return func(v []*int) error {
					require.Equal(t, testIntPtrSliceConsumerValue, v)
					return testIntPtrSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustIntPtrSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testIntPtrSliceConsumerError.Error(), func() {
					mc(testIntPtrSliceConsumerValue)
				})
			} else {
				mc(testIntPtrSliceConsumerValue)
			}
		})
	}
}

func TestSilentIntPtrSliceConsumer(t *testing.T) {
	var sc SilentIntPtrSliceConsumer = func(v []*int) {
		require.Equal(t, testIntPtrSliceConsumerValue, v)
		return
	}
	sc(testIntPtrSliceConsumerValue)
}

func TestSilentIntPtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testIntPtrSliceConsumerFactory
		cf2   testIntPtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) IntPtrSliceConsumer {
				return func(v []*int) error {
					calls++
					require.Equal(t, testIntPtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntPtrSliceConsumer {
				return func(v []*int) error {
					calls++
					require.Equal(t, testIntPtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) IntPtrSliceConsumer {
				return func(v []*int) error {
					calls++
					require.Equal(t, testIntPtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testIntPtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) IntPtrSliceConsumer {
				return func(v []*int) error {
					calls++
					require.Equal(t, testIntPtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) IntPtrSliceConsumer {
				return func(v []*int) error {
					calls++
					require.Equal(t, testIntPtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntPtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentIntPtrSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentIntPtrSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentIntPtrSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testIntPtrSliceConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustIntPtrSliceConsumer(t *testing.T) {
	var sc SilentIntPtrSliceConsumer = func(v []*int) {
		require.Equal(t, testIntPtrSliceConsumerValue, v)
		return
	}
	sc(testIntPtrSliceConsumerValue)
}

func TestMustIntPtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testIntPtrSliceConsumerFactory
		cf2   testIntPtrSliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) IntPtrSliceConsumer {
				return func(v []*int) error {
					calls++
					require.Equal(t, testIntPtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntPtrSliceConsumer {
				return func(v []*int) error {
					calls++
					require.Equal(t, testIntPtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) IntPtrSliceConsumer {
				return func(v []*int) error {
					calls++
					require.Equal(t, testIntPtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testIntPtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) IntPtrSliceConsumer {
				return func(v []*int) error {
					calls++
					require.Equal(t, testIntPtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) IntPtrSliceConsumer {
				return func(v []*int) error {
					calls++
					require.Equal(t, testIntPtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntPtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustIntPtrSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustIntPtrSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustIntPtrSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testIntPtrSliceConsumerError.Error(), func() {
					cmc(testIntPtrSliceConsumerValue)
				})
			} else {
				cmc(testIntPtrSliceConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustIntPtrSliceConsumer_ToSilentIntPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntPtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntPtrSliceConsumer {
				return func(v []*int) error {
					require.Equal(t, testIntPtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntPtrSliceConsumer {
				return func(v []*int) error {
					require.Equal(t, testIntPtrSliceConsumerValue, v)
					return testIntPtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustIntPtrSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentIntPtrSliceConsumer()
			r.NotNil(sc)

			sc(testIntPtrSliceConsumerValue)
		})
	}
}

func TestMustIntPtrSliceConsumer_ToIntPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntPtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntPtrSliceConsumer {
				return func(v []*int) error {
					require.Equal(t, testIntPtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntPtrSliceConsumer {
				return func(v []*int) error {
					require.Equal(t, testIntPtrSliceConsumerValue, v)
					return testIntPtrSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustIntPtrSliceConsumer()
			r.NotNil(mc)

			c = mc.ToIntPtrSliceConsumer()
			r.NotNil(c)

			err := c(testIntPtrSliceConsumerValue)
			if tt.err {
				r.EqualError(err, testIntPtrSliceConsumerError.Error())
			}
		})
	}
}
