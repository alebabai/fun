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
	testFloat32PtrSliceConsumerValue []*float32
	testFloat32PtrSliceConsumerError = errors.New("error")
)

type testFloat32PtrSliceConsumerFactory func(t *testing.T) Float32PtrSliceConsumer

func TestFloat32PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					require.Equal(t, testFloat32PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					require.Equal(t, testFloat32PtrSliceConsumerValue, v)
					return testFloat32PtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testFloat32PtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testFloat32PtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestFloat32PtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					require.Equal(t, testFloat32PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					require.Equal(t, testFloat32PtrSliceConsumerValue, v)
					return testFloat32PtrSliceConsumerError
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

			err := c(testFloat32PtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testFloat32PtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestFloat32PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testFloat32PtrSliceConsumerFactory
		cf2   testFloat32PtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					calls++
					require.Equal(t, testFloat32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					calls++
					require.Equal(t, testFloat32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					calls++
					require.Equal(t, testFloat32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testFloat32PtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					calls++
					require.Equal(t, testFloat32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					calls++
					require.Equal(t, testFloat32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32PtrSliceConsumer {
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
			err := cc(testFloat32PtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testFloat32PtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestFloat32PtrSliceConsumer_ToSilentFloat32PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					require.Equal(t, testFloat32PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					require.Equal(t, testFloat32PtrSliceConsumerValue, v)
					return testFloat32PtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentFloat32PtrSliceConsumer()
			r.NotNil(sc)

			sc(testFloat32PtrSliceConsumerValue)
		})
	}
}

func TestFloat32PtrSliceConsumer_ToMustFloat32PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32PtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					require.Equal(t, testFloat32PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					require.Equal(t, testFloat32PtrSliceConsumerValue, v)
					return testFloat32PtrSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustFloat32PtrSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testFloat32PtrSliceConsumerError.Error(), func() {
					mc(testFloat32PtrSliceConsumerValue)
				})
			} else {
				mc(testFloat32PtrSliceConsumerValue)
			}
		})
	}
}

func TestSilentFloat32PtrSliceConsumer(t *testing.T) {
	var sc SilentFloat32PtrSliceConsumer = func(v []*float32) {
		require.Equal(t, testFloat32PtrSliceConsumerValue, v)
		return
	}
	sc(testFloat32PtrSliceConsumerValue)
}

func TestSilentFloat32PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testFloat32PtrSliceConsumerFactory
		cf2   testFloat32PtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					calls++
					require.Equal(t, testFloat32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					calls++
					require.Equal(t, testFloat32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					calls++
					require.Equal(t, testFloat32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testFloat32PtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					calls++
					require.Equal(t, testFloat32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					calls++
					require.Equal(t, testFloat32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32PtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentFloat32PtrSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentFloat32PtrSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentFloat32PtrSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testFloat32PtrSliceConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustFloat32PtrSliceConsumer(t *testing.T) {
	var sc SilentFloat32PtrSliceConsumer = func(v []*float32) {
		require.Equal(t, testFloat32PtrSliceConsumerValue, v)
		return
	}
	sc(testFloat32PtrSliceConsumerValue)
}

func TestMustFloat32PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testFloat32PtrSliceConsumerFactory
		cf2   testFloat32PtrSliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					calls++
					require.Equal(t, testFloat32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					calls++
					require.Equal(t, testFloat32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					calls++
					require.Equal(t, testFloat32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testFloat32PtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					calls++
					require.Equal(t, testFloat32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					calls++
					require.Equal(t, testFloat32PtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32PtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustFloat32PtrSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustFloat32PtrSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustFloat32PtrSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testFloat32PtrSliceConsumerError.Error(), func() {
					cmc(testFloat32PtrSliceConsumerValue)
				})
			} else {
				cmc(testFloat32PtrSliceConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustFloat32PtrSliceConsumer_ToSilentFloat32PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					require.Equal(t, testFloat32PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					require.Equal(t, testFloat32PtrSliceConsumerValue, v)
					return testFloat32PtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustFloat32PtrSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentFloat32PtrSliceConsumer()
			r.NotNil(sc)

			sc(testFloat32PtrSliceConsumerValue)
		})
	}
}

func TestMustFloat32PtrSliceConsumer_ToFloat32PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32PtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					require.Equal(t, testFloat32PtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					require.Equal(t, testFloat32PtrSliceConsumerValue, v)
					return testFloat32PtrSliceConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustFloat32PtrSliceConsumer()
			r.NotNil(mc)

			c = mc.ToFloat32PtrSliceConsumer()
			r.NotNil(c)

			err := c(testFloat32PtrSliceConsumerValue)
			if tt.err {
				r.EqualError(err, testFloat32PtrSliceConsumerError.Error())
			}
		})
	}
}
