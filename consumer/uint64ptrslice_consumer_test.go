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
	valTestUint64PtrSliceConsumer []*uint64
	errTestUint64PtrSliceConsumer = errors.New("error")
)

type testUint64PtrSliceConsumerFactory func(t *testing.T) Uint64PtrSliceConsumer

func TestUint64PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUint64PtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					return errTestUint64PtrSliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestUint64PtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestUint64PtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint64PtrSliceConsumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUint64PtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					return errTestUint64PtrSliceConsumer
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

			err := c(valTestUint64PtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestUint64PtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint64PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testUint64PtrSliceConsumerFactory
		cf2     testUint64PtrSliceConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint64PtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64PtrSliceConsumer {
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
			err := cc(valTestUint64PtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestUint64PtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestUint64PtrSliceConsumer_ToSilentUint64PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					return errTestUint64PtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentUint64PtrSliceConsumer()
			r.NotNil(sc)

			sc(valTestUint64PtrSliceConsumer)
		})
	}
}

func TestUint64PtrSliceConsumer_ToMustUint64PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64PtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					return errTestUint64PtrSliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustUint64PtrSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestUint64PtrSliceConsumer.Error(), func() {
					mc(valTestUint64PtrSliceConsumer)
				})
			} else {
				mc(valTestUint64PtrSliceConsumer)
			}
		})
	}
}

func TestSilentUint64PtrSliceConsumer(t *testing.T) {
	var sc SilentUint64PtrSliceConsumer = func(v []*uint64) {
		require.Equal(t, valTestUint64PtrSliceConsumer, v)
		return
	}
	sc(valTestUint64PtrSliceConsumer)
}

func TestSilentUint64PtrSliceConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					return errTestUint64PtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentUint64PtrSliceConsumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestUint64PtrSliceConsumer)
		})
	}
}

func TestSilentUint64PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint64PtrSliceConsumerFactory
		cf2   testUint64PtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint64PtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64PtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentUint64PtrSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentUint64PtrSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentUint64PtrSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestUint64PtrSliceConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint64PtrSliceConsumer(t *testing.T) {
	var mc MustUint64PtrSliceConsumer = func(v []*uint64) {
		require.Equal(t, valTestUint64PtrSliceConsumer, v)
		return
	}
	mc(valTestUint64PtrSliceConsumer)
}

func TestMustUint64PtrSliceConsumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUint64PtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					return errTestUint64PtrSliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustUint64PtrSliceConsumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.wantErr {
				r.PanicsWithError(errTestUint64PtrSliceConsumer.Error(), func() {
					mc(valTestUint64PtrSliceConsumer)
				})
			} else {
				mc(valTestUint64PtrSliceConsumer)
			}
		})
	}
}

func TestMustUint64PtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testUint64PtrSliceConsumerFactory
		cf2     testUint64PtrSliceConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint64PtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					calls++
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64PtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustUint64PtrSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustUint64PtrSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustUint64PtrSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.wantErr {
				r.PanicsWithError(errTestUint64PtrSliceConsumer.Error(), func() {
					cmc(valTestUint64PtrSliceConsumer)
				})
			} else {
				cmc(valTestUint64PtrSliceConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint64PtrSliceConsumer_ToSilentUint64PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64PtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					return errTestUint64PtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint64PtrSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentUint64PtrSliceConsumer()
			r.NotNil(sc)

			sc(valTestUint64PtrSliceConsumer)
		})
	}
}

func TestMustUint64PtrSliceConsumer_ToUint64PtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUint64PtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint64PtrSliceConsumer {
				return func(v []*uint64) error {
					require.Equal(t, valTestUint64PtrSliceConsumer, v)
					return errTestUint64PtrSliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint64PtrSliceConsumer()
			r.NotNil(mc)

			c = mc.ToUint64PtrSliceConsumer()
			r.NotNil(c)

			err := c(valTestUint64PtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestUint64PtrSliceConsumer.Error())
			}
		})
	}
}
