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
	testFloat32PtrConsumerValue *float32
	testFloat32PtrConsumerError = errors.New("error")
)

type testFloat32PtrConsumerFactory func(t *testing.T) Float32PtrConsumer

func TestFloat32PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32PtrConsumer {
				return func(v *float32) error {
					require.Equal(t, testFloat32PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32PtrConsumer {
				return func(v *float32) error {
					require.Equal(t, testFloat32PtrConsumerValue, v)
					return testFloat32PtrConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testFloat32PtrConsumerValue)
			if err != nil {
				r.EqualError(err, testFloat32PtrConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestFloat32PtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32PtrConsumer {
				return func(v *float32) error {
					require.Equal(t, testFloat32PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32PtrConsumer {
				return func(v *float32) error {
					require.Equal(t, testFloat32PtrConsumerValue, v)
					return testFloat32PtrConsumerError
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

			err := c(testFloat32PtrConsumerValue)
			if err != nil {
				r.EqualError(err, testFloat32PtrConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestFloat32PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testFloat32PtrConsumerFactory
		cf2   testFloat32PtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float32PtrConsumer {
				return func(v *float32) error {
					calls++
					require.Equal(t, testFloat32PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32PtrConsumer {
				return func(v *float32) error {
					calls++
					require.Equal(t, testFloat32PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Float32PtrConsumer {
				return func(v *float32) error {
					calls++
					require.Equal(t, testFloat32PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testFloat32PtrConsumerError
				}
			},
			cf2: func(t *testing.T) Float32PtrConsumer {
				return func(v *float32) error {
					calls++
					require.Equal(t, testFloat32PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float32PtrConsumer {
				return func(v *float32) error {
					calls++
					require.Equal(t, testFloat32PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32PtrConsumer {
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
			err := cc(testFloat32PtrConsumerValue)
			if err != nil {
				r.EqualError(err, testFloat32PtrConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestFloat32PtrConsumer_ToSilentFloat32PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32PtrConsumer {
				return func(v *float32) error {
					require.Equal(t, testFloat32PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32PtrConsumer {
				return func(v *float32) error {
					require.Equal(t, testFloat32PtrConsumerValue, v)
					return testFloat32PtrConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentFloat32PtrConsumer()
			r.NotNil(sc)

			sc(testFloat32PtrConsumerValue)
		})
	}
}

func TestFloat32PtrConsumer_ToMustFloat32PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32PtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32PtrConsumer {
				return func(v *float32) error {
					require.Equal(t, testFloat32PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32PtrConsumer {
				return func(v *float32) error {
					require.Equal(t, testFloat32PtrConsumerValue, v)
					return testFloat32PtrConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustFloat32PtrConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testFloat32PtrConsumerError.Error(), func() {
					mc(testFloat32PtrConsumerValue)
				})
			} else {
				mc(testFloat32PtrConsumerValue)
			}
		})
	}
}

func TestSilentFloat32PtrConsumer(t *testing.T) {
	var sc SilentFloat32PtrConsumer = func(v *float32) {
		require.Equal(t, testFloat32PtrConsumerValue, v)
		return
	}
	sc(testFloat32PtrConsumerValue)
}

func TestSilentFloat32PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testFloat32PtrConsumerFactory
		cf2   testFloat32PtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float32PtrConsumer {
				return func(v *float32) error {
					calls++
					require.Equal(t, testFloat32PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32PtrConsumer {
				return func(v *float32) error {
					calls++
					require.Equal(t, testFloat32PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Float32PtrConsumer {
				return func(v *float32) error {
					calls++
					require.Equal(t, testFloat32PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testFloat32PtrConsumerError
				}
			},
			cf2: func(t *testing.T) Float32PtrConsumer {
				return func(v *float32) error {
					calls++
					require.Equal(t, testFloat32PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float32PtrConsumer {
				return func(v *float32) error {
					calls++
					require.Equal(t, testFloat32PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32PtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentFloat32PtrConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentFloat32PtrConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentFloat32PtrConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testFloat32PtrConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustFloat32PtrConsumer(t *testing.T) {
	var sc SilentFloat32PtrConsumer = func(v *float32) {
		require.Equal(t, testFloat32PtrConsumerValue, v)
		return
	}
	sc(testFloat32PtrConsumerValue)
}

func TestMustFloat32PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testFloat32PtrConsumerFactory
		cf2   testFloat32PtrConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float32PtrConsumer {
				return func(v *float32) error {
					calls++
					require.Equal(t, testFloat32PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32PtrConsumer {
				return func(v *float32) error {
					calls++
					require.Equal(t, testFloat32PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Float32PtrConsumer {
				return func(v *float32) error {
					calls++
					require.Equal(t, testFloat32PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testFloat32PtrConsumerError
				}
			},
			cf2: func(t *testing.T) Float32PtrConsumer {
				return func(v *float32) error {
					calls++
					require.Equal(t, testFloat32PtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float32PtrConsumer {
				return func(v *float32) error {
					calls++
					require.Equal(t, testFloat32PtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32PtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustFloat32PtrConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustFloat32PtrConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustFloat32PtrConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testFloat32PtrConsumerError.Error(), func() {
					cmc(testFloat32PtrConsumerValue)
				})
			} else {
				cmc(testFloat32PtrConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustFloat32PtrConsumer_ToSilentFloat32PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32PtrConsumer {
				return func(v *float32) error {
					require.Equal(t, testFloat32PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32PtrConsumer {
				return func(v *float32) error {
					require.Equal(t, testFloat32PtrConsumerValue, v)
					return testFloat32PtrConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustFloat32PtrConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentFloat32PtrConsumer()
			r.NotNil(sc)

			sc(testFloat32PtrConsumerValue)
		})
	}
}

func TestMustFloat32PtrConsumer_ToFloat32PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32PtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32PtrConsumer {
				return func(v *float32) error {
					require.Equal(t, testFloat32PtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32PtrConsumer {
				return func(v *float32) error {
					require.Equal(t, testFloat32PtrConsumerValue, v)
					return testFloat32PtrConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustFloat32PtrConsumer()
			r.NotNil(mc)

			c = mc.ToFloat32PtrConsumer()
			r.NotNil(c)

			err := c(testFloat32PtrConsumerValue)
			if tt.err {
				r.EqualError(err, testFloat32PtrConsumerError.Error())
			}
		})
	}
}
