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
	testStringPtrSliceConsumerValue []*string
	testStringPtrSliceConsumerError = errors.New("error")
)

type testStringPtrSliceConsumerFactory func(t *testing.T) StringPtrSliceConsumer

func TestStringPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringPtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					require.Equal(t, testStringPtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					require.Equal(t, testStringPtrSliceConsumerValue, v)
					return testStringPtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testStringPtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testStringPtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestStringPtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringPtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					require.Equal(t, testStringPtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					require.Equal(t, testStringPtrSliceConsumerValue, v)
					return testStringPtrSliceConsumerError
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

			err := c(testStringPtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testStringPtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestStringPtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testStringPtrSliceConsumerFactory
		cf2   testStringPtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, testStringPtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, testStringPtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, testStringPtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testStringPtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, testStringPtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, testStringPtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringPtrSliceConsumer {
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
			err := cc(testStringPtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testStringPtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestStringPtrSliceConsumer_ToSilentStringPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringPtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					require.Equal(t, testStringPtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					require.Equal(t, testStringPtrSliceConsumerValue, v)
					return testStringPtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentStringPtrSliceConsumer()
			r.NotNil(sc)

			sc(testStringPtrSliceConsumerValue)
		})
	}
}

func TestStringPtrSliceConsumer_ToMustStringPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringPtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					require.Equal(t, testStringPtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					require.Equal(t, testStringPtrSliceConsumerValue, v)
					return testStringPtrSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustStringPtrSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testStringPtrSliceConsumerError.Error(), func() {
					mc(testStringPtrSliceConsumerValue)
				})
			} else {
				mc(testStringPtrSliceConsumerValue)
			}
		})
	}
}

func TestSilentStringPtrSliceConsumer(t *testing.T) {
	var sc SilentStringPtrSliceConsumer = func(v []*string) {
		require.Equal(t, testStringPtrSliceConsumerValue, v)
		return
	}
	sc(testStringPtrSliceConsumerValue)
}

func TestSilentStringPtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testStringPtrSliceConsumerFactory
		cf2   testStringPtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, testStringPtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, testStringPtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, testStringPtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testStringPtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, testStringPtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, testStringPtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringPtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentStringPtrSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentStringPtrSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentStringPtrSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testStringPtrSliceConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustStringPtrSliceConsumer(t *testing.T) {
	var sc SilentStringPtrSliceConsumer = func(v []*string) {
		require.Equal(t, testStringPtrSliceConsumerValue, v)
		return
	}
	sc(testStringPtrSliceConsumerValue)
}

func TestMustStringPtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testStringPtrSliceConsumerFactory
		cf2   testStringPtrSliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, testStringPtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, testStringPtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, testStringPtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testStringPtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, testStringPtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, testStringPtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringPtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustStringPtrSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustStringPtrSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustStringPtrSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testStringPtrSliceConsumerError.Error(), func() {
					cmc(testStringPtrSliceConsumerValue)
				})
			} else {
				cmc(testStringPtrSliceConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustStringPtrSliceConsumer_ToSilentStringPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringPtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					require.Equal(t, testStringPtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					require.Equal(t, testStringPtrSliceConsumerValue, v)
					return testStringPtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustStringPtrSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentStringPtrSliceConsumer()
			r.NotNil(sc)

			sc(testStringPtrSliceConsumerValue)
		})
	}
}

func TestMustStringPtrSliceConsumer_ToStringPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringPtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					require.Equal(t, testStringPtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					require.Equal(t, testStringPtrSliceConsumerValue, v)
					return testStringPtrSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustStringPtrSliceConsumer()
			r.NotNil(mc)

			c = mc.ToStringPtrSliceConsumer()
			r.NotNil(c)

			err := c(testStringPtrSliceConsumerValue)
			if tt.err {
				r.EqualError(err, testStringPtrSliceConsumerError.Error())
			}
		})
	}
}
