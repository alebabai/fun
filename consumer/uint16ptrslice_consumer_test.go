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
	valTestUint16PtrSliceConsumer []*uint16
	errTestUint16PtrSliceConsumer = errors.New("error")
)

type testUint16PtrSliceConsumerFactory func(t *testing.T) Uint16PtrSliceConsumer

func TestUint16PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUint16PtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
					return errTestUint16PtrSliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestUint16PtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestUint16PtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint16PtrSliceConsumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUint16PtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
					return errTestUint16PtrSliceConsumer
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

			err := c(valTestUint16PtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestUint16PtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint16PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testUint16PtrSliceConsumerFactory
		cf2     testUint16PtrSliceConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					calls++
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					calls++
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					calls++
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint16PtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					calls++
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					calls++
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
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
			err := cc(valTestUint16PtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestUint16PtrSliceConsumer.Error())
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
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
					return errTestUint16PtrSliceConsumer
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

			sc(valTestUint16PtrSliceConsumer)
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
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
					return errTestUint16PtrSliceConsumer
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
				r.PanicsWithError(errTestUint16PtrSliceConsumer.Error(), func() {
					mc(valTestUint16PtrSliceConsumer)
				})
			} else {
				mc(valTestUint16PtrSliceConsumer)
			}
		})
	}
}

func TestSilentUint16PtrSliceConsumer(t *testing.T) {
	var sc SilentUint16PtrSliceConsumer = func(v []*uint16) {
		require.Equal(t, valTestUint16PtrSliceConsumer, v)
		return
	}
	sc(valTestUint16PtrSliceConsumer)
}

func TestSilentUint16PtrSliceConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint16PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
					return errTestUint16PtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentUint16PtrSliceConsumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestUint16PtrSliceConsumer)
		})
	}
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
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					calls++
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					calls++
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint16PtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					calls++
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
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
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
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
			csc(valTestUint16PtrSliceConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint16PtrSliceConsumer(t *testing.T) {
	var mc MustUint16PtrSliceConsumer = func(v []*uint16) {
		require.Equal(t, valTestUint16PtrSliceConsumer, v)
		return
	}
	mc(valTestUint16PtrSliceConsumer)
}

func TestMustUint16PtrSliceConsumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUint16PtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
					return errTestUint16PtrSliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustUint16PtrSliceConsumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.wantErr {
				r.PanicsWithError(errTestUint16PtrSliceConsumer.Error(), func() {
					mc(valTestUint16PtrSliceConsumer)
				})
			} else {
				mc(valTestUint16PtrSliceConsumer)
			}
		})
	}
}

func TestMustUint16PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testUint16PtrSliceConsumerFactory
		cf2     testUint16PtrSliceConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					calls++
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					calls++
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					calls++
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint16PtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					calls++
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					calls++
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
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
			if tt.wantErr {
				r.PanicsWithError(errTestUint16PtrSliceConsumer.Error(), func() {
					cmc(valTestUint16PtrSliceConsumer)
				})
			} else {
				cmc(valTestUint16PtrSliceConsumer)
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
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
					return errTestUint16PtrSliceConsumer
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

			sc(valTestUint16PtrSliceConsumer)
		})
	}
}

func TestMustUint16PtrSliceConsumer_ToUint16PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUint16PtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint16PtrSliceConsumer {
				return func(v []*uint16) error {
					require.Equal(t, valTestUint16PtrSliceConsumer, v)
					return errTestUint16PtrSliceConsumer
				}
			},
			wantErr: true,
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

			err := c(valTestUint16PtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestUint16PtrSliceConsumer.Error())
			}
		})
	}
}
