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
	valTestFloat64PtrSliceConsumer []*float64
	errTestFloat64PtrSliceConsumer = errors.New("error")
)

type testFloat64PtrSliceConsumerFactory func(t *testing.T) Float64PtrSliceConsumer

func TestFloat64PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testFloat64PtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					return errTestFloat64PtrSliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestFloat64PtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestFloat64PtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestFloat64PtrSliceConsumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testFloat64PtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					return errTestFloat64PtrSliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			c := tc.ToConsumer()
			r.NotNil(c)

			err := c(valTestFloat64PtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestFloat64PtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestFloat64PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testFloat64PtrSliceConsumerFactory
		cf2     testFloat64PtrSliceConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestFloat64PtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64PtrSliceConsumer {
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
			err := cc(valTestFloat64PtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestFloat64PtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestFloat64PtrSliceConsumer_ToSilentFloat64PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat64PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					return errTestFloat64PtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentFloat64PtrSliceConsumer()
			r.NotNil(sc)

			sc(valTestFloat64PtrSliceConsumer)
		})
	}
}

func TestFloat64PtrSliceConsumer_ToMustFloat64PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat64PtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					return errTestFloat64PtrSliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustFloat64PtrSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestFloat64PtrSliceConsumer.Error(), func() {
					mc(valTestFloat64PtrSliceConsumer)
				})
			} else {
				mc(valTestFloat64PtrSliceConsumer)
			}
		})
	}
}

func TestSilentFloat64PtrSliceConsumer(t *testing.T) {
	var sc SilentFloat64PtrSliceConsumer = func(v []*float64) {
		require.Equal(t, valTestFloat64PtrSliceConsumer, v)
		return
	}
	sc(valTestFloat64PtrSliceConsumer)
}

func TestSilentFloat64PtrSliceConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat64PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					return errTestFloat64PtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentFloat64PtrSliceConsumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestFloat64PtrSliceConsumer)
		})
	}
}

func TestSilentFloat64PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testFloat64PtrSliceConsumerFactory
		cf2   testFloat64PtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestFloat64PtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64PtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentFloat64PtrSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentFloat64PtrSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentFloat64PtrSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestFloat64PtrSliceConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustFloat64PtrSliceConsumer(t *testing.T) {
	var mc MustFloat64PtrSliceConsumer = func(v []*float64) {
		require.Equal(t, valTestFloat64PtrSliceConsumer, v)
		return
	}
	mc(valTestFloat64PtrSliceConsumer)
}

func TestMustFloat64PtrSliceConsumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testFloat64PtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					return errTestFloat64PtrSliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustFloat64PtrSliceConsumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.wantErr {
				r.PanicsWithError(errTestFloat64PtrSliceConsumer.Error(), func() {
					mc(valTestFloat64PtrSliceConsumer)
				})
			} else {
				mc(valTestFloat64PtrSliceConsumer)
			}
		})
	}
}

func TestMustFloat64PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testFloat64PtrSliceConsumerFactory
		cf2     testFloat64PtrSliceConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestFloat64PtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					calls++
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float64PtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustFloat64PtrSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustFloat64PtrSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustFloat64PtrSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.wantErr {
				r.PanicsWithError(errTestFloat64PtrSliceConsumer.Error(), func() {
					cmc(valTestFloat64PtrSliceConsumer)
				})
			} else {
				cmc(valTestFloat64PtrSliceConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustFloat64PtrSliceConsumer_ToSilentFloat64PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat64PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					return errTestFloat64PtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustFloat64PtrSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentFloat64PtrSliceConsumer()
			r.NotNil(sc)

			sc(valTestFloat64PtrSliceConsumer)
		})
	}
}

func TestMustFloat64PtrSliceConsumer_ToFloat64PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testFloat64PtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Float64PtrSliceConsumer {
				return func(v []*float64) error {
					require.Equal(t, valTestFloat64PtrSliceConsumer, v)
					return errTestFloat64PtrSliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustFloat64PtrSliceConsumer()
			r.NotNil(mc)

			c = mc.ToFloat64PtrSliceConsumer()
			r.NotNil(c)

			err := c(valTestFloat64PtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestFloat64PtrSliceConsumer.Error())
			}
		})
	}
}
