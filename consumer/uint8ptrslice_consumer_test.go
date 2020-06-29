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
	testUint8PtrSliceConsumerValue []*uint8
	testUint8PtrSliceConsumerError = errors.New("error")
)

type testUint8PtrSliceConsumerFactory func(t *testing.T) Uint8PtrSliceConsumer

func TestUint8PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint8PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint8PtrSliceConsumer {
				return func(v []*uint8) error {
					require.Equal(t, testUint8PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint8PtrSliceConsumer {
				return func(v []*uint8) error {
					require.Equal(t, testUint8PtrSliceConsumerValue, v)
					return testUint8PtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testUint8PtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testUint8PtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint8PtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint8PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint8PtrSliceConsumer {
				return func(v []*uint8) error {
					require.Equal(t, testUint8PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint8PtrSliceConsumer {
				return func(v []*uint8) error {
					require.Equal(t, testUint8PtrSliceConsumerValue, v)
					return testUint8PtrSliceConsumerError
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

			err := c(testUint8PtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testUint8PtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint8PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint8PtrSliceConsumerFactory
		cf2   testUint8PtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint8PtrSliceConsumer {
				return func(v []*uint8) error {
					calls++
					require.Equal(t, testUint8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint8PtrSliceConsumer {
				return func(v []*uint8) error {
					calls++
					require.Equal(t, testUint8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint8PtrSliceConsumer {
				return func(v []*uint8) error {
					calls++
					require.Equal(t, testUint8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUint8PtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) Uint8PtrSliceConsumer {
				return func(v []*uint8) error {
					calls++
					require.Equal(t, testUint8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint8PtrSliceConsumer {
				return func(v []*uint8) error {
					calls++
					require.Equal(t, testUint8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint8PtrSliceConsumer {
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
			err := cc(testUint8PtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testUint8PtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestUint8PtrSliceConsumer_ToSilentUint8PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint8PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint8PtrSliceConsumer {
				return func(v []*uint8) error {
					require.Equal(t, testUint8PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint8PtrSliceConsumer {
				return func(v []*uint8) error {
					require.Equal(t, testUint8PtrSliceConsumerValue, v)
					return testUint8PtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentUint8PtrSliceConsumer()
			r.NotNil(sc)

			sc(testUint8PtrSliceConsumerValue)
		})
	}
}

func TestUint8PtrSliceConsumer_ToMustUint8PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint8PtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint8PtrSliceConsumer {
				return func(v []*uint8) error {
					require.Equal(t, testUint8PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint8PtrSliceConsumer {
				return func(v []*uint8) error {
					require.Equal(t, testUint8PtrSliceConsumerValue, v)
					return testUint8PtrSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustUint8PtrSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testUint8PtrSliceConsumerError.Error(), func() {
					mc(testUint8PtrSliceConsumerValue)
				})
			} else {
				mc(testUint8PtrSliceConsumerValue)
			}
		})
	}
}

func TestSilentUint8PtrSliceConsumer(t *testing.T) {
	var sc SilentUint8PtrSliceConsumer = func(v []*uint8) {
		require.Equal(t, testUint8PtrSliceConsumerValue, v)
		return
	}
	sc(testUint8PtrSliceConsumerValue)
}

func TestSilentUint8PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint8PtrSliceConsumerFactory
		cf2   testUint8PtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint8PtrSliceConsumer {
				return func(v []*uint8) error {
					calls++
					require.Equal(t, testUint8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint8PtrSliceConsumer {
				return func(v []*uint8) error {
					calls++
					require.Equal(t, testUint8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint8PtrSliceConsumer {
				return func(v []*uint8) error {
					calls++
					require.Equal(t, testUint8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUint8PtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) Uint8PtrSliceConsumer {
				return func(v []*uint8) error {
					calls++
					require.Equal(t, testUint8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint8PtrSliceConsumer {
				return func(v []*uint8) error {
					calls++
					require.Equal(t, testUint8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint8PtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentUint8PtrSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentUint8PtrSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentUint8PtrSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testUint8PtrSliceConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint8PtrSliceConsumer(t *testing.T) {
	var sc SilentUint8PtrSliceConsumer = func(v []*uint8) {
		require.Equal(t, testUint8PtrSliceConsumerValue, v)
		return
	}
	sc(testUint8PtrSliceConsumerValue)
}

func TestMustUint8PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint8PtrSliceConsumerFactory
		cf2   testUint8PtrSliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint8PtrSliceConsumer {
				return func(v []*uint8) error {
					calls++
					require.Equal(t, testUint8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint8PtrSliceConsumer {
				return func(v []*uint8) error {
					calls++
					require.Equal(t, testUint8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint8PtrSliceConsumer {
				return func(v []*uint8) error {
					calls++
					require.Equal(t, testUint8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUint8PtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) Uint8PtrSliceConsumer {
				return func(v []*uint8) error {
					calls++
					require.Equal(t, testUint8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint8PtrSliceConsumer {
				return func(v []*uint8) error {
					calls++
					require.Equal(t, testUint8PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint8PtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustUint8PtrSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustUint8PtrSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustUint8PtrSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testUint8PtrSliceConsumerError.Error(), func() {
					cmc(testUint8PtrSliceConsumerValue)
				})
			} else {
				cmc(testUint8PtrSliceConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint8PtrSliceConsumer_ToSilentUint8PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint8PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint8PtrSliceConsumer {
				return func(v []*uint8) error {
					require.Equal(t, testUint8PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint8PtrSliceConsumer {
				return func(v []*uint8) error {
					require.Equal(t, testUint8PtrSliceConsumerValue, v)
					return testUint8PtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint8PtrSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentUint8PtrSliceConsumer()
			r.NotNil(sc)

			sc(testUint8PtrSliceConsumerValue)
		})
	}
}

func TestMustUint8PtrSliceConsumer_ToUint8PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint8PtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint8PtrSliceConsumer {
				return func(v []*uint8) error {
					require.Equal(t, testUint8PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint8PtrSliceConsumer {
				return func(v []*uint8) error {
					require.Equal(t, testUint8PtrSliceConsumerValue, v)
					return testUint8PtrSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint8PtrSliceConsumer()
			r.NotNil(mc)

			c = mc.ToUint8PtrSliceConsumer()
			r.NotNil(c)

			err := c(testUint8PtrSliceConsumerValue)
			if tt.err {
				r.EqualError(err, testUint8PtrSliceConsumerError.Error())
			}
		})
	}
}
