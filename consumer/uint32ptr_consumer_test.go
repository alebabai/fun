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
	valTestUint32PtrConsumer *uint32
	errTestUint32PtrConsumer = errors.New("error")
)

type testUint32PtrConsumerFactory func(t *testing.T) Uint32PtrConsumer

func TestUint32PtrConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUint32PtrConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					require.Equal(t, valTestUint32PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					require.Equal(t, valTestUint32PtrConsumer, v)
					return errTestUint32PtrConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestUint32PtrConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestUint32PtrConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint32PtrConsumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUint32PtrConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					require.Equal(t, valTestUint32PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					require.Equal(t, valTestUint32PtrConsumer, v)
					return errTestUint32PtrConsumer
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

			err := c(valTestUint32PtrConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestUint32PtrConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint32PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testUint32PtrConsumerFactory
		cf2     testUint32PtrConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					calls++
					require.Equal(t, valTestUint32PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					calls++
					require.Equal(t, valTestUint32PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					calls++
					require.Equal(t, valTestUint32PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint32PtrConsumer
				}
			},
			cf2: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					calls++
					require.Equal(t, valTestUint32PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					calls++
					require.Equal(t, valTestUint32PtrConsumer, v)
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
			err := cc(valTestUint32PtrConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestUint32PtrConsumer.Error())
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
					require.Equal(t, valTestUint32PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					require.Equal(t, valTestUint32PtrConsumer, v)
					return errTestUint32PtrConsumer
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

			sc(valTestUint32PtrConsumer)
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
					require.Equal(t, valTestUint32PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					require.Equal(t, valTestUint32PtrConsumer, v)
					return errTestUint32PtrConsumer
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
				r.PanicsWithError(errTestUint32PtrConsumer.Error(), func() {
					mc(valTestUint32PtrConsumer)
				})
			} else {
				mc(valTestUint32PtrConsumer)
			}
		})
	}
}

func TestSilentUint32PtrConsumer(t *testing.T) {
	var sc SilentUint32PtrConsumer = func(v *uint32) {
		require.Equal(t, valTestUint32PtrConsumer, v)
		return
	}
	sc(valTestUint32PtrConsumer)
}

func TestSilentUint32PtrConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint32PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					require.Equal(t, valTestUint32PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					require.Equal(t, valTestUint32PtrConsumer, v)
					return errTestUint32PtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentUint32PtrConsumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestUint32PtrConsumer)
		})
	}
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
					require.Equal(t, valTestUint32PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					calls++
					require.Equal(t, valTestUint32PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					calls++
					require.Equal(t, valTestUint32PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint32PtrConsumer
				}
			},
			cf2: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					calls++
					require.Equal(t, valTestUint32PtrConsumer, v)
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
					require.Equal(t, valTestUint32PtrConsumer, v)
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
			csc(valTestUint32PtrConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint32PtrConsumer(t *testing.T) {
	var mc MustUint32PtrConsumer = func(v *uint32) {
		require.Equal(t, valTestUint32PtrConsumer, v)
		return
	}
	mc(valTestUint32PtrConsumer)
}

func TestMustUint32PtrConsumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUint32PtrConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					require.Equal(t, valTestUint32PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					require.Equal(t, valTestUint32PtrConsumer, v)
					return errTestUint32PtrConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustUint32PtrConsumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.wantErr {
				r.PanicsWithError(errTestUint32PtrConsumer.Error(), func() {
					mc(valTestUint32PtrConsumer)
				})
			} else {
				mc(valTestUint32PtrConsumer)
			}
		})
	}
}

func TestMustUint32PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testUint32PtrConsumerFactory
		cf2     testUint32PtrConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					calls++
					require.Equal(t, valTestUint32PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					calls++
					require.Equal(t, valTestUint32PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					calls++
					require.Equal(t, valTestUint32PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint32PtrConsumer
				}
			},
			cf2: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					calls++
					require.Equal(t, valTestUint32PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					calls++
					require.Equal(t, valTestUint32PtrConsumer, v)
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
			if tt.wantErr {
				r.PanicsWithError(errTestUint32PtrConsumer.Error(), func() {
					cmc(valTestUint32PtrConsumer)
				})
			} else {
				cmc(valTestUint32PtrConsumer)
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
					require.Equal(t, valTestUint32PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					require.Equal(t, valTestUint32PtrConsumer, v)
					return errTestUint32PtrConsumer
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

			sc(valTestUint32PtrConsumer)
		})
	}
}

func TestMustUint32PtrConsumer_ToUint32PtrConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUint32PtrConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					require.Equal(t, valTestUint32PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint32PtrConsumer {
				return func(v *uint32) error {
					require.Equal(t, valTestUint32PtrConsumer, v)
					return errTestUint32PtrConsumer
				}
			},
			wantErr: true,
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

			err := c(valTestUint32PtrConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestUint32PtrConsumer.Error())
			}
		})
	}
}
