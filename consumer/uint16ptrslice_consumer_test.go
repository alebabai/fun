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
	testUint16PtrSliceConsumerValue []*uint16
	testUint16PtrSliceConsumerError = errors.New("error")
)

type testUint16PtrSliceConsumerFactory func(t *testing.T) Uint16PtrSliceConsumer

func TestUint16PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint16PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					require.Equal(t, testUint16PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					require.Equal(t, testUint16PtrSliceConsumerValue, v)
					return testUint16PtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testUint16PtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testUint16PtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint16PtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint16PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					require.Equal(t, testUint16PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					require.Equal(t, testUint16PtrSliceConsumerValue, v)
					return testUint16PtrSliceConsumerError
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

			err := c(testUint16PtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testUint16PtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint16PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint16PtrSliceConsumerFactory
		cf2   testUint16PtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					calls++
					require.Equal(t, testUint16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					calls++
					require.Equal(t, testUint16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					calls++
					require.Equal(t, testUint16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUint16PtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					calls++
					require.Equal(t, testUint16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					calls++
					require.Equal(t, testUint16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16PtrSliceConsumer {
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
			err := cc(testUint16PtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testUint16PtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestUint16PtrSliceConsumer_ToSilentUint16PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint16PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					require.Equal(t, testUint16PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					require.Equal(t, testUint16PtrSliceConsumerValue, v)
					return testUint16PtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentUint16PtrSliceConsumer()
			r.NotNil(sc)

			sc(testUint16PtrSliceConsumerValue)
		})
	}
}

func TestUint16PtrSliceConsumer_ToMustUint16PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint16PtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					require.Equal(t, testUint16PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					require.Equal(t, testUint16PtrSliceConsumerValue, v)
					return testUint16PtrSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustUint16PtrSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testUint16PtrSliceConsumerError.Error(), func() {
					mc(testUint16PtrSliceConsumerValue)
				})
			} else {
				mc(testUint16PtrSliceConsumerValue)
			}
		})
	}
}

func TestSilentUint16PtrSliceConsumer(t *testing.T) {
	var sc SilentUint16PtrSliceConsumer = func(v []*uint16) {
		require.Equal(t, testUint16PtrSliceConsumerValue, v)
		return
	}
	sc(testUint16PtrSliceConsumerValue)
}

func TestSilentUint16PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint16PtrSliceConsumerFactory
		cf2   testUint16PtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					calls++
					require.Equal(t, testUint16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					calls++
					require.Equal(t, testUint16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					calls++
					require.Equal(t, testUint16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUint16PtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					calls++
					require.Equal(t, testUint16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					calls++
					require.Equal(t, testUint16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16PtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentUint16PtrSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentUint16PtrSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentUint16PtrSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testUint16PtrSliceConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint16PtrSliceConsumer(t *testing.T) {
	var sc SilentUint16PtrSliceConsumer = func(v []*uint16) {
		require.Equal(t, testUint16PtrSliceConsumerValue, v)
		return
	}
	sc(testUint16PtrSliceConsumerValue)
}

func TestMustUint16PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint16PtrSliceConsumerFactory
		cf2   testUint16PtrSliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					calls++
					require.Equal(t, testUint16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					calls++
					require.Equal(t, testUint16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					calls++
					require.Equal(t, testUint16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUint16PtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					calls++
					require.Equal(t, testUint16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					calls++
					require.Equal(t, testUint16PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16PtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustUint16PtrSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustUint16PtrSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustUint16PtrSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testUint16PtrSliceConsumerError.Error(), func() {
					cmc(testUint16PtrSliceConsumerValue)
				})
			} else {
				cmc(testUint16PtrSliceConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint16PtrSliceConsumer_ToSilentUint16PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint16PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					require.Equal(t, testUint16PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					require.Equal(t, testUint16PtrSliceConsumerValue, v)
					return testUint16PtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint16PtrSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentUint16PtrSliceConsumer()
			r.NotNil(sc)

			sc(testUint16PtrSliceConsumerValue)
		})
	}
}

func TestMustUint16PtrSliceConsumer_ToUint16PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint16PtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					require.Equal(t, testUint16PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					require.Equal(t, testUint16PtrSliceConsumerValue, v)
					return testUint16PtrSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint16PtrSliceConsumer()
			r.NotNil(mc)

			c = mc.ToUint16PtrSliceConsumer()
			r.NotNil(c)

			err := c(testUint16PtrSliceConsumerValue)
			if tt.err {
				r.EqualError(err, testUint16PtrSliceConsumerError.Error())
			}
		})
	}
}
