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
	valTestFloat32PtrSliceConsumer []*float32
	errTestFloat32PtrSliceConsumer = errors.New("error")
)

type testFloat32PtrSliceConsumerFactory func(t *testing.T) Float32PtrSliceConsumer

func TestFloat32PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testFloat32PtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
					return errTestFloat32PtrSliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestFloat32PtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestFloat32PtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestFloat32PtrSliceConsumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testFloat32PtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
					return errTestFloat32PtrSliceConsumer
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

			err := c(valTestFloat32PtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestFloat32PtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestFloat32PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testFloat32PtrSliceConsumerFactory
		cf2     testFloat32PtrSliceConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					calls++
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					calls++
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					calls++
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestFloat32PtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					calls++
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					calls++
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
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
			err := cc(valTestFloat32PtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestFloat32PtrSliceConsumer.Error())
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
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
					return errTestFloat32PtrSliceConsumer
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

			sc(valTestFloat32PtrSliceConsumer)
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
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
					return errTestFloat32PtrSliceConsumer
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
				r.PanicsWithError(errTestFloat32PtrSliceConsumer.Error(), func() {
					mc(valTestFloat32PtrSliceConsumer)
				})
			} else {
				mc(valTestFloat32PtrSliceConsumer)
			}
		})
	}
}

func TestSilentFloat32PtrSliceConsumer(t *testing.T) {
	var sc SilentFloat32PtrSliceConsumer = func(v []*float32) {
		require.Equal(t, valTestFloat32PtrSliceConsumer, v)
		return
	}
	sc(valTestFloat32PtrSliceConsumer)
}

func TestSilentFloat32PtrSliceConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testFloat32PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
					return errTestFloat32PtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentFloat32PtrSliceConsumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestFloat32PtrSliceConsumer)
		})
	}
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
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					calls++
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					calls++
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestFloat32PtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					calls++
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
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
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
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
			csc(valTestFloat32PtrSliceConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustFloat32PtrSliceConsumer(t *testing.T) {
	var mc MustFloat32PtrSliceConsumer = func(v []*float32) {
		require.Equal(t, valTestFloat32PtrSliceConsumer, v)
		return
	}
	mc(valTestFloat32PtrSliceConsumer)
}

func TestMustFloat32PtrSliceConsumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testFloat32PtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
					return errTestFloat32PtrSliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustFloat32PtrSliceConsumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.wantErr {
				r.PanicsWithError(errTestFloat32PtrSliceConsumer.Error(), func() {
					mc(valTestFloat32PtrSliceConsumer)
				})
			} else {
				mc(valTestFloat32PtrSliceConsumer)
			}
		})
	}
}

func TestMustFloat32PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testFloat32PtrSliceConsumerFactory
		cf2     testFloat32PtrSliceConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					calls++
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					calls++
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					calls++
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestFloat32PtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					calls++
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					calls++
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
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
			if tt.wantErr {
				r.PanicsWithError(errTestFloat32PtrSliceConsumer.Error(), func() {
					cmc(valTestFloat32PtrSliceConsumer)
				})
			} else {
				cmc(valTestFloat32PtrSliceConsumer)
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
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
					return errTestFloat32PtrSliceConsumer
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

			sc(valTestFloat32PtrSliceConsumer)
		})
	}
}

func TestMustFloat32PtrSliceConsumer_ToFloat32PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testFloat32PtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Float32PtrSliceConsumer {
				return func(v []*float32) error {
					require.Equal(t, valTestFloat32PtrSliceConsumer, v)
					return errTestFloat32PtrSliceConsumer
				}
			},
			wantErr: true,
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

			err := c(valTestFloat32PtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestFloat32PtrSliceConsumer.Error())
			}
		})
	}
}
