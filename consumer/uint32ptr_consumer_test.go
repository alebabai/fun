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
	testUint32PtrConsumerValue *uint32
	testUint32PtrConsumerError = errors.New("error")
)

type testUint32PtrConsumerFactory func(t *testing.T) Uint32PtrConsumer

func TestUint32PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint32PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					require.Equal(t, testUint32PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					require.Equal(t, testUint32PtrConsumerValue, v)
					return testUint32PtrConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testUint32PtrConsumerValue)
			if err != nil {
				r.EqualError(err, testUint32PtrConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint32PtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint32PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					require.Equal(t, testUint32PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					require.Equal(t, testUint32PtrConsumerValue, v)
					return testUint32PtrConsumerError
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

			err := c(testUint32PtrConsumerValue)
			if err != nil {
				r.EqualError(err, testUint32PtrConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint32PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint32PtrConsumerFactory
		cf2   testUint32PtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					calls++
					require.Equal(t, testUint32PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					calls++
					require.Equal(t, testUint32PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					calls++
					require.Equal(t, testUint32PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUint32PtrConsumerError
				}
			},
			cf2: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					calls++
					require.Equal(t, testUint32PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					calls++
					require.Equal(t, testUint32PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint32PtrConsumer {
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
			err := cc(testUint32PtrConsumerValue)
			if err != nil {
				r.EqualError(err, testUint32PtrConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestUint32PtrConsumer_ToSilentUint32PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint32PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					require.Equal(t, testUint32PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					require.Equal(t, testUint32PtrConsumerValue, v)
					return testUint32PtrConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentUint32PtrConsumer()
			r.NotNil(sc)

			sc(testUint32PtrConsumerValue)
		})
	}
}

func TestUint32PtrConsumer_ToMustUint32PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint32PtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					require.Equal(t, testUint32PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					require.Equal(t, testUint32PtrConsumerValue, v)
					return testUint32PtrConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustUint32PtrConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testUint32PtrConsumerError.Error(), func() {
					mc(testUint32PtrConsumerValue)
				})
			} else {
				mc(testUint32PtrConsumerValue)
			}
		})
	}
}

func TestSilentUint32PtrConsumer(t *testing.T) {
	var sc SilentUint32PtrConsumer = func(v *uint32) {
		require.Equal(t, testUint32PtrConsumerValue, v)
		return
	}
	sc(testUint32PtrConsumerValue)
}

func TestSilentUint32PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint32PtrConsumerFactory
		cf2   testUint32PtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					calls++
					require.Equal(t, testUint32PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					calls++
					require.Equal(t, testUint32PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					calls++
					require.Equal(t, testUint32PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUint32PtrConsumerError
				}
			},
			cf2: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					calls++
					require.Equal(t, testUint32PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					calls++
					require.Equal(t, testUint32PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint32PtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentUint32PtrConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentUint32PtrConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentUint32PtrConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testUint32PtrConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint32PtrConsumer(t *testing.T) {
	var sc SilentUint32PtrConsumer = func(v *uint32) {
		require.Equal(t, testUint32PtrConsumerValue, v)
		return
	}
	sc(testUint32PtrConsumerValue)
}

func TestMustUint32PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint32PtrConsumerFactory
		cf2   testUint32PtrConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					calls++
					require.Equal(t, testUint32PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					calls++
					require.Equal(t, testUint32PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					calls++
					require.Equal(t, testUint32PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUint32PtrConsumerError
				}
			},
			cf2: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					calls++
					require.Equal(t, testUint32PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					calls++
					require.Equal(t, testUint32PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint32PtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustUint32PtrConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustUint32PtrConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustUint32PtrConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testUint32PtrConsumerError.Error(), func() {
					cmc(testUint32PtrConsumerValue)
				})
			} else {
				cmc(testUint32PtrConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint32PtrConsumer_ToSilentUint32PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint32PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					require.Equal(t, testUint32PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					require.Equal(t, testUint32PtrConsumerValue, v)
					return testUint32PtrConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint32PtrConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentUint32PtrConsumer()
			r.NotNil(sc)

			sc(testUint32PtrConsumerValue)
		})
	}
}

func TestMustUint32PtrConsumer_ToUint32PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint32PtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					require.Equal(t, testUint32PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					require.Equal(t, testUint32PtrConsumerValue, v)
					return testUint32PtrConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint32PtrConsumer()
			r.NotNil(mc)

			c = mc.ToUint32PtrConsumer()
			r.NotNil(c)

			err := c(testUint32PtrConsumerValue)
			if tt.err {
				r.EqualError(err, testUint32PtrConsumerError.Error())
			}
		})
	}
}
