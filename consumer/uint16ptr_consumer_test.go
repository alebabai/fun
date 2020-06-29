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
	testUint16PtrConsumerValue *uint16
	testUint16PtrConsumerError = errors.New("error")
)

type testUint16PtrConsumerFactory func(t *testing.T) Uint16PtrConsumer

func TestUint16PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint16PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16PtrConsumer {
				return func(v *uint16) error {
					require.Equal(t, testUint16PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint16PtrConsumer {
				return func(v *uint16) error {
					require.Equal(t, testUint16PtrConsumerValue, v)
					return testUint16PtrConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testUint16PtrConsumerValue)
			if err != nil {
				r.EqualError(err, testUint16PtrConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint16PtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint16PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16PtrConsumer {
				return func(v *uint16) error {
					require.Equal(t, testUint16PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint16PtrConsumer {
				return func(v *uint16) error {
					require.Equal(t, testUint16PtrConsumerValue, v)
					return testUint16PtrConsumerError
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

			err := c(testUint16PtrConsumerValue)
			if err != nil {
				r.EqualError(err, testUint16PtrConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint16PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint16PtrConsumerFactory
		cf2   testUint16PtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint16PtrConsumer {
				return func(v *uint16) error {
					calls++
					require.Equal(t, testUint16PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16PtrConsumer {
				return func(v *uint16) error {
					calls++
					require.Equal(t, testUint16PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint16PtrConsumer {
				return func(v *uint16) error {
					calls++
					require.Equal(t, testUint16PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUint16PtrConsumerError
				}
			},
			cf2: func(t *testing.T) Uint16PtrConsumer {
				return func(v *uint16) error {
					calls++
					require.Equal(t, testUint16PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint16PtrConsumer {
				return func(v *uint16) error {
					calls++
					require.Equal(t, testUint16PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16PtrConsumer {
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
			err := cc(testUint16PtrConsumerValue)
			if err != nil {
				r.EqualError(err, testUint16PtrConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestUint16PtrConsumer_ToSilentUint16PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint16PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16PtrConsumer {
				return func(v *uint16) error {
					require.Equal(t, testUint16PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint16PtrConsumer {
				return func(v *uint16) error {
					require.Equal(t, testUint16PtrConsumerValue, v)
					return testUint16PtrConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentUint16PtrConsumer()
			r.NotNil(sc)

			sc(testUint16PtrConsumerValue)
		})
	}
}

func TestUint16PtrConsumer_ToMustUint16PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint16PtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16PtrConsumer {
				return func(v *uint16) error {
					require.Equal(t, testUint16PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint16PtrConsumer {
				return func(v *uint16) error {
					require.Equal(t, testUint16PtrConsumerValue, v)
					return testUint16PtrConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustUint16PtrConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testUint16PtrConsumerError.Error(), func() {
					mc(testUint16PtrConsumerValue)
				})
			} else {
				mc(testUint16PtrConsumerValue)
			}
		})
	}
}

func TestSilentUint16PtrConsumer(t *testing.T) {
	var sc SilentUint16PtrConsumer = func(v *uint16) {
		require.Equal(t, testUint16PtrConsumerValue, v)
		return
	}
	sc(testUint16PtrConsumerValue)
}

func TestSilentUint16PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint16PtrConsumerFactory
		cf2   testUint16PtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint16PtrConsumer {
				return func(v *uint16) error {
					calls++
					require.Equal(t, testUint16PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16PtrConsumer {
				return func(v *uint16) error {
					calls++
					require.Equal(t, testUint16PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint16PtrConsumer {
				return func(v *uint16) error {
					calls++
					require.Equal(t, testUint16PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUint16PtrConsumerError
				}
			},
			cf2: func(t *testing.T) Uint16PtrConsumer {
				return func(v *uint16) error {
					calls++
					require.Equal(t, testUint16PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint16PtrConsumer {
				return func(v *uint16) error {
					calls++
					require.Equal(t, testUint16PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16PtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentUint16PtrConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentUint16PtrConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentUint16PtrConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testUint16PtrConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint16PtrConsumer(t *testing.T) {
	var sc SilentUint16PtrConsumer = func(v *uint16) {
		require.Equal(t, testUint16PtrConsumerValue, v)
		return
	}
	sc(testUint16PtrConsumerValue)
}

func TestMustUint16PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint16PtrConsumerFactory
		cf2   testUint16PtrConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint16PtrConsumer {
				return func(v *uint16) error {
					calls++
					require.Equal(t, testUint16PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16PtrConsumer {
				return func(v *uint16) error {
					calls++
					require.Equal(t, testUint16PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Uint16PtrConsumer {
				return func(v *uint16) error {
					calls++
					require.Equal(t, testUint16PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUint16PtrConsumerError
				}
			},
			cf2: func(t *testing.T) Uint16PtrConsumer {
				return func(v *uint16) error {
					calls++
					require.Equal(t, testUint16PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint16PtrConsumer {
				return func(v *uint16) error {
					calls++
					require.Equal(t, testUint16PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16PtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustUint16PtrConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustUint16PtrConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustUint16PtrConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testUint16PtrConsumerError.Error(), func() {
					cmc(testUint16PtrConsumerValue)
				})
			} else {
				cmc(testUint16PtrConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint16PtrConsumer_ToSilentUint16PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint16PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16PtrConsumer {
				return func(v *uint16) error {
					require.Equal(t, testUint16PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint16PtrConsumer {
				return func(v *uint16) error {
					require.Equal(t, testUint16PtrConsumerValue, v)
					return testUint16PtrConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint16PtrConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentUint16PtrConsumer()
			r.NotNil(sc)

			sc(testUint16PtrConsumerValue)
		})
	}
}

func TestMustUint16PtrConsumer_ToUint16PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint16PtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16PtrConsumer {
				return func(v *uint16) error {
					require.Equal(t, testUint16PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Uint16PtrConsumer {
				return func(v *uint16) error {
					require.Equal(t, testUint16PtrConsumerValue, v)
					return testUint16PtrConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint16PtrConsumer()
			r.NotNil(mc)

			c = mc.ToUint16PtrConsumer()
			r.NotNil(c)

			err := c(testUint16PtrConsumerValue)
			if tt.err {
				r.EqualError(err, testUint16PtrConsumerError.Error())
			}
		})
	}
}
