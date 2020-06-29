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
	testFloat32ConsumerValue float32
	testFloat32ConsumerError = errors.New("error")
)

type testFloat32ConsumerFactory func(t *testing.T) Float32Consumer

func TestFloat32Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					require.Equal(t, testFloat32ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					require.Equal(t, testFloat32ConsumerValue, v)
					return testFloat32ConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testFloat32ConsumerValue)
			if err != nil {
				r.EqualError(err, testFloat32ConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestFloat32Supplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					require.Equal(t, testFloat32ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					require.Equal(t, testFloat32ConsumerValue, v)
					return testFloat32ConsumerError
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

			err := c(testFloat32ConsumerValue)
			if err != nil {
				r.EqualError(err, testFloat32ConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestFloat32Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testFloat32ConsumerFactory
		cf2   testFloat32ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, testFloat32ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, testFloat32ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, testFloat32ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testFloat32ConsumerError
				}
			},
			cf2: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, testFloat32ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, testFloat32ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32Consumer {
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
			err := cc(testFloat32ConsumerValue)
			if err != nil {
				r.EqualError(err, testFloat32ConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestFloat32Consumer_ToSilentFloat32Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					require.Equal(t, testFloat32ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					require.Equal(t, testFloat32ConsumerValue, v)
					return testFloat32ConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentFloat32Consumer()
			r.NotNil(sc)

			sc(testFloat32ConsumerValue)
		})
	}
}

func TestFloat32Consumer_ToMustFloat32Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					require.Equal(t, testFloat32ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					require.Equal(t, testFloat32ConsumerValue, v)
					return testFloat32ConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustFloat32Consumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testFloat32ConsumerError.Error(), func() {
					mc(testFloat32ConsumerValue)
				})
			} else {
				mc(testFloat32ConsumerValue)
			}
		})
	}
}

func TestSilentFloat32Consumer(t *testing.T) {
	var sc SilentFloat32Consumer = func(v float32) {
		require.Equal(t, testFloat32ConsumerValue, v)
		return
	}
	sc(testFloat32ConsumerValue)
}

func TestSilentFloat32Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testFloat32ConsumerFactory
		cf2   testFloat32ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, testFloat32ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, testFloat32ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, testFloat32ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testFloat32ConsumerError
				}
			},
			cf2: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, testFloat32ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, testFloat32ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentFloat32Consumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentFloat32Consumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentFloat32Consumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testFloat32ConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustFloat32Consumer(t *testing.T) {
	var sc SilentFloat32Consumer = func(v float32) {
		require.Equal(t, testFloat32ConsumerValue, v)
		return
	}
	sc(testFloat32ConsumerValue)
}

func TestMustFloat32Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testFloat32ConsumerFactory
		cf2   testFloat32ConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, testFloat32ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, testFloat32ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, testFloat32ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testFloat32ConsumerError
				}
			},
			cf2: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, testFloat32ConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					calls++
					require.Equal(t, testFloat32ConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustFloat32Consumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustFloat32Consumer = nil
			if c2 != nil {
				mc2 = c2.ToMustFloat32Consumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testFloat32ConsumerError.Error(), func() {
					cmc(testFloat32ConsumerValue)
				})
			} else {
				cmc(testFloat32ConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustFloat32Consumer_ToSilentFloat32Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					require.Equal(t, testFloat32ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					require.Equal(t, testFloat32ConsumerValue, v)
					return testFloat32ConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustFloat32Consumer()
			r.NotNil(mc)

			sc := mc.ToSilentFloat32Consumer()
			r.NotNil(sc)

			sc(testFloat32ConsumerValue)
		})
	}
}

func TestMustFloat32Consumer_ToFloat32Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					require.Equal(t, testFloat32ConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32Consumer {
				return func(v float32) error {
					require.Equal(t, testFloat32ConsumerValue, v)
					return testFloat32ConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustFloat32Consumer()
			r.NotNil(mc)

			c = mc.ToFloat32Consumer()
			r.NotNil(c)

			err := c(testFloat32ConsumerValue)
			if tt.err {
				r.EqualError(err, testFloat32ConsumerError.Error())
			}
		})
	}
}
