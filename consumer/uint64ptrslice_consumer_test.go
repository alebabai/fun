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
	testUint64PtrSliceConsumerValue []*uint64
	testUint64PtrSliceConsumerError = errors.New("error")
)

type testUint64PtrSliceConsumerFactory func(t *testing.T) Uint64PtrSliceConsumer

func TestUint64PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					require.Equal(t, testUint64PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					require.Equal(t, testUint64PtrSliceConsumerValue, v)
					return testUint64PtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testUint64PtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testUint64PtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint64PtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					require.Equal(t, testUint64PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					require.Equal(t, testUint64PtrSliceConsumerValue, v)
					return testUint64PtrSliceConsumerError
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

			err := c(testUint64PtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testUint64PtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint64PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint64PtrSliceConsumerFactory
		cf2   testUint64PtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, testUint64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, testUint64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, testUint64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUint64PtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, testUint64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, testUint64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64PtrSliceConsumer {
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
			err := cc(testUint64PtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testUint64PtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestUint64PtrSliceConsumer_ToSilentUint64PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					require.Equal(t, testUint64PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					require.Equal(t, testUint64PtrSliceConsumerValue, v)
					return testUint64PtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentUint64PtrSliceConsumer()
			r.NotNil(sc)

			sc(testUint64PtrSliceConsumerValue)
		})
	}
}

func TestUint64PtrSliceConsumer_ToMustUint64PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64PtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					require.Equal(t, testUint64PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					require.Equal(t, testUint64PtrSliceConsumerValue, v)
					return testUint64PtrSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustUint64PtrSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testUint64PtrSliceConsumerError.Error(), func() {
					mc(testUint64PtrSliceConsumerValue)
				})
			} else {
				mc(testUint64PtrSliceConsumerValue)
			}
		})
	}
}

func TestSilentUint64PtrSliceConsumer(t *testing.T) {
	var sc SilentUint64PtrSliceConsumer = func(v []*uint64) {
		require.Equal(t, testUint64PtrSliceConsumerValue, v)
		return
	}
	sc(testUint64PtrSliceConsumerValue)
}

func TestSilentUint64PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint64PtrSliceConsumerFactory
		cf2   testUint64PtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, testUint64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, testUint64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, testUint64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUint64PtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, testUint64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, testUint64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64PtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentUint64PtrSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentUint64PtrSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentUint64PtrSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testUint64PtrSliceConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint64PtrSliceConsumer(t *testing.T) {
	var sc SilentUint64PtrSliceConsumer = func(v []*uint64) {
		require.Equal(t, testUint64PtrSliceConsumerValue, v)
		return
	}
	sc(testUint64PtrSliceConsumerValue)
}

func TestMustUint64PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint64PtrSliceConsumerFactory
		cf2   testUint64PtrSliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, testUint64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, testUint64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, testUint64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUint64PtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, testUint64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, testUint64PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64PtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustUint64PtrSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustUint64PtrSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustUint64PtrSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testUint64PtrSliceConsumerError.Error(), func() {
					cmc(testUint64PtrSliceConsumerValue)
				})
			} else {
				cmc(testUint64PtrSliceConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint64PtrSliceConsumer_ToSilentUint64PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					require.Equal(t, testUint64PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					require.Equal(t, testUint64PtrSliceConsumerValue, v)
					return testUint64PtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint64PtrSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentUint64PtrSliceConsumer()
			r.NotNil(sc)

			sc(testUint64PtrSliceConsumerValue)
		})
	}
}

func TestMustUint64PtrSliceConsumer_ToUint64PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64PtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					require.Equal(t, testUint64PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					require.Equal(t, testUint64PtrSliceConsumerValue, v)
					return testUint64PtrSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint64PtrSliceConsumer()
			r.NotNil(mc)

			c = mc.ToUint64PtrSliceConsumer()
			r.NotNil(c)

			err := c(testUint64PtrSliceConsumerValue)
			if tt.err {
				r.EqualError(err, testUint64PtrSliceConsumerError.Error())
			}
		})
	}
}
