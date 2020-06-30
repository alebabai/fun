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
	valTestFloat32SliceConsumer []float32
	errTestFloat32SliceConsumer = errors.New("error")
)

type testFloat32SliceConsumerFactory func(t *testing.T) Float32SliceConsumer

func TestFloat32SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32SliceConsumer {
				return func(v []float32) error {
					require.Equal(t, valTestFloat32SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32SliceConsumer {
				return func(v []float32) error {
					require.Equal(t, valTestFloat32SliceConsumer, v)
					return errTestFloat32SliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestFloat32SliceConsumer)
			if err != nil {
				r.EqualError(err, errTestFloat32SliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestFloat32SliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32SliceConsumer {
				return func(v []float32) error {
					require.Equal(t, valTestFloat32SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32SliceConsumer {
				return func(v []float32) error {
					require.Equal(t, valTestFloat32SliceConsumer, v)
					return errTestFloat32SliceConsumer
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

			err := c(valTestFloat32SliceConsumer)
			if err != nil {
				r.EqualError(err, errTestFloat32SliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestFloat32SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testFloat32SliceConsumerFactory
		cf2   testFloat32SliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float32SliceConsumer {
				return func(v []float32) error {
					calls++
					require.Equal(t, valTestFloat32SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32SliceConsumer {
				return func(v []float32) error {
					calls++
					require.Equal(t, valTestFloat32SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Float32SliceConsumer {
				return func(v []float32) error {
					calls++
					require.Equal(t, valTestFloat32SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestFloat32SliceConsumer
				}
			},
			cf2: func(t *testing.T) Float32SliceConsumer {
				return func(v []float32) error {
					calls++
					require.Equal(t, valTestFloat32SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float32SliceConsumer {
				return func(v []float32) error {
					calls++
					require.Equal(t, valTestFloat32SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32SliceConsumer {
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
			err := cc(valTestFloat32SliceConsumer)
			if err != nil {
				r.EqualError(err, errTestFloat32SliceConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestFloat32SliceConsumer_ToSilentFloat32SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32SliceConsumer {
				return func(v []float32) error {
					require.Equal(t, valTestFloat32SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32SliceConsumer {
				return func(v []float32) error {
					require.Equal(t, valTestFloat32SliceConsumer, v)
					return errTestFloat32SliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentFloat32SliceConsumer()
			r.NotNil(sc)

			sc(valTestFloat32SliceConsumer)
		})
	}
}

func TestFloat32SliceConsumer_ToMustFloat32SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32SliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32SliceConsumer {
				return func(v []float32) error {
					require.Equal(t, valTestFloat32SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32SliceConsumer {
				return func(v []float32) error {
					require.Equal(t, valTestFloat32SliceConsumer, v)
					return errTestFloat32SliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustFloat32SliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestFloat32SliceConsumer.Error(), func() {
					mc(valTestFloat32SliceConsumer)
				})
			} else {
				mc(valTestFloat32SliceConsumer)
			}
		})
	}
}

func TestSilentFloat32SliceConsumer(t *testing.T) {
	var sc SilentFloat32SliceConsumer = func(v []float32) {
		require.Equal(t, valTestFloat32SliceConsumer, v)
		return
	}
	sc(valTestFloat32SliceConsumer)
}

func TestSilentFloat32SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testFloat32SliceConsumerFactory
		cf2   testFloat32SliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float32SliceConsumer {
				return func(v []float32) error {
					calls++
					require.Equal(t, valTestFloat32SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32SliceConsumer {
				return func(v []float32) error {
					calls++
					require.Equal(t, valTestFloat32SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Float32SliceConsumer {
				return func(v []float32) error {
					calls++
					require.Equal(t, valTestFloat32SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestFloat32SliceConsumer
				}
			},
			cf2: func(t *testing.T) Float32SliceConsumer {
				return func(v []float32) error {
					calls++
					require.Equal(t, valTestFloat32SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float32SliceConsumer {
				return func(v []float32) error {
					calls++
					require.Equal(t, valTestFloat32SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32SliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentFloat32SliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentFloat32SliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentFloat32SliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestFloat32SliceConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustFloat32SliceConsumer(t *testing.T) {
	var sc SilentFloat32SliceConsumer = func(v []float32) {
		require.Equal(t, valTestFloat32SliceConsumer, v)
		return
	}
	sc(valTestFloat32SliceConsumer)
}

func TestMustFloat32SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testFloat32SliceConsumerFactory
		cf2   testFloat32SliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float32SliceConsumer {
				return func(v []float32) error {
					calls++
					require.Equal(t, valTestFloat32SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32SliceConsumer {
				return func(v []float32) error {
					calls++
					require.Equal(t, valTestFloat32SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Float32SliceConsumer {
				return func(v []float32) error {
					calls++
					require.Equal(t, valTestFloat32SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestFloat32SliceConsumer
				}
			},
			cf2: func(t *testing.T) Float32SliceConsumer {
				return func(v []float32) error {
					calls++
					require.Equal(t, valTestFloat32SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float32SliceConsumer {
				return func(v []float32) error {
					calls++
					require.Equal(t, valTestFloat32SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32SliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustFloat32SliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustFloat32SliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustFloat32SliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(errTestFloat32SliceConsumer.Error(), func() {
					cmc(valTestFloat32SliceConsumer)
				})
			} else {
				cmc(valTestFloat32SliceConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustFloat32SliceConsumer_ToSilentFloat32SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32SliceConsumer {
				return func(v []float32) error {
					require.Equal(t, valTestFloat32SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32SliceConsumer {
				return func(v []float32) error {
					require.Equal(t, valTestFloat32SliceConsumer, v)
					return errTestFloat32SliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustFloat32SliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentFloat32SliceConsumer()
			r.NotNil(sc)

			sc(valTestFloat32SliceConsumer)
		})
	}
}

func TestMustFloat32SliceConsumer_ToFloat32SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32SliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32SliceConsumer {
				return func(v []float32) error {
					require.Equal(t, valTestFloat32SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Float32SliceConsumer {
				return func(v []float32) error {
					require.Equal(t, valTestFloat32SliceConsumer, v)
					return errTestFloat32SliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustFloat32SliceConsumer()
			r.NotNil(mc)

			c = mc.ToFloat32SliceConsumer()
			r.NotNil(c)

			err := c(valTestFloat32SliceConsumer)
			if tt.err {
				r.EqualError(err, errTestFloat32SliceConsumer.Error())
			}
		})
	}
}
