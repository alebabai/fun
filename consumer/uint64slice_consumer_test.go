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
	valTestUint64SliceConsumer []uint64
	errTestUint64SliceConsumer = errors.New("error")
)

type testUint64SliceConsumerFactory func(t *testing.T) Uint64SliceConsumer

func TestUint64SliceConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUint64SliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					require.Equal(t, valTestUint64SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					require.Equal(t, valTestUint64SliceConsumer, v)
					return errTestUint64SliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestUint64SliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestUint64SliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint64SliceConsumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUint64SliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					require.Equal(t, valTestUint64SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					require.Equal(t, valTestUint64SliceConsumer, v)
					return errTestUint64SliceConsumer
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

			err := c(valTestUint64SliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestUint64SliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint64SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testUint64SliceConsumerFactory
		cf2     testUint64SliceConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, valTestUint64SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, valTestUint64SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, valTestUint64SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint64SliceConsumer
				}
			},
			cf2: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, valTestUint64SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, valTestUint64SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64SliceConsumer {
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
			err := cc(valTestUint64SliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestUint64SliceConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestUint64SliceConsumer_ToSilentUint64SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					require.Equal(t, valTestUint64SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					require.Equal(t, valTestUint64SliceConsumer, v)
					return errTestUint64SliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentUint64SliceConsumer()
			r.NotNil(sc)

			sc(valTestUint64SliceConsumer)
		})
	}
}

func TestUint64SliceConsumer_ToMustUint64SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64SliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					require.Equal(t, valTestUint64SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					require.Equal(t, valTestUint64SliceConsumer, v)
					return errTestUint64SliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustUint64SliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestUint64SliceConsumer.Error(), func() {
					mc(valTestUint64SliceConsumer)
				})
			} else {
				mc(valTestUint64SliceConsumer)
			}
		})
	}
}

func TestSilentUint64SliceConsumer(t *testing.T) {
	var sc SilentUint64SliceConsumer = func(v []uint64) {
		require.Equal(t, valTestUint64SliceConsumer, v)
		return
	}
	sc(valTestUint64SliceConsumer)
}

func TestSilentUint64SliceConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					require.Equal(t, valTestUint64SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					require.Equal(t, valTestUint64SliceConsumer, v)
					return errTestUint64SliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentUint64SliceConsumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestUint64SliceConsumer)
		})
	}
}

func TestSilentUint64SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint64SliceConsumerFactory
		cf2   testUint64SliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, valTestUint64SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, valTestUint64SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, valTestUint64SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint64SliceConsumer
				}
			},
			cf2: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, valTestUint64SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, valTestUint64SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64SliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentUint64SliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentUint64SliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentUint64SliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestUint64SliceConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint64SliceConsumer(t *testing.T) {
	var mc MustUint64SliceConsumer = func(v []uint64) {
		require.Equal(t, valTestUint64SliceConsumer, v)
		return
	}
	mc(valTestUint64SliceConsumer)
}

func TestMustUint64SliceConsumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUint64SliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					require.Equal(t, valTestUint64SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					require.Equal(t, valTestUint64SliceConsumer, v)
					return errTestUint64SliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustUint64SliceConsumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.wantErr {
				r.PanicsWithError(errTestUint64SliceConsumer.Error(), func() {
					mc(valTestUint64SliceConsumer)
				})
			} else {
				mc(valTestUint64SliceConsumer)
			}
		})
	}
}

func TestMustUint64SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testUint64SliceConsumerFactory
		cf2     testUint64SliceConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, valTestUint64SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, valTestUint64SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, valTestUint64SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint64SliceConsumer
				}
			},
			cf2: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, valTestUint64SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					calls++
					require.Equal(t, valTestUint64SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64SliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustUint64SliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustUint64SliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustUint64SliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.wantErr {
				r.PanicsWithError(errTestUint64SliceConsumer.Error(), func() {
					cmc(valTestUint64SliceConsumer)
				})
			} else {
				cmc(valTestUint64SliceConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint64SliceConsumer_ToSilentUint64SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					require.Equal(t, valTestUint64SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					require.Equal(t, valTestUint64SliceConsumer, v)
					return errTestUint64SliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint64SliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentUint64SliceConsumer()
			r.NotNil(sc)

			sc(valTestUint64SliceConsumer)
		})
	}
}

func TestMustUint64SliceConsumer_ToUint64SliceConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUint64SliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					require.Equal(t, valTestUint64SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint64SliceConsumer {
				return func(v []uint64) error {
					require.Equal(t, valTestUint64SliceConsumer, v)
					return errTestUint64SliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint64SliceConsumer()
			r.NotNil(mc)

			c = mc.ToUint64SliceConsumer()
			r.NotNil(c)

			err := c(valTestUint64SliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestUint64SliceConsumer.Error())
			}
		})
	}
}
