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
	valTestBytePtrSliceConsumer []*byte
	errTestBytePtrSliceConsumer = errors.New("error")
)

type testBytePtrSliceConsumerFactory func(t *testing.T) BytePtrSliceConsumer

func TestBytePtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testBytePtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					return errTestBytePtrSliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestBytePtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestBytePtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestBytePtrSliceConsumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testBytePtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					return errTestBytePtrSliceConsumer
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

			err := c(valTestBytePtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestBytePtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestBytePtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testBytePtrSliceConsumerFactory
		cf2     testBytePtrSliceConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					calls++
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					calls++
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					calls++
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestBytePtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					calls++
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					calls++
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BytePtrSliceConsumer {
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
			err := cc(valTestBytePtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestBytePtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestBytePtrSliceConsumer_ToSilentBytePtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBytePtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					return errTestBytePtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentBytePtrSliceConsumer()
			r.NotNil(sc)

			sc(valTestBytePtrSliceConsumer)
		})
	}
}

func TestBytePtrSliceConsumer_ToMustBytePtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBytePtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					return errTestBytePtrSliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustBytePtrSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestBytePtrSliceConsumer.Error(), func() {
					mc(valTestBytePtrSliceConsumer)
				})
			} else {
				mc(valTestBytePtrSliceConsumer)
			}
		})
	}
}

func TestSilentBytePtrSliceConsumer(t *testing.T) {
	var sc SilentBytePtrSliceConsumer = func(v []*byte) {
		require.Equal(t, valTestBytePtrSliceConsumer, v)
		return
	}
	sc(valTestBytePtrSliceConsumer)
}

func TestSilentBytePtrSliceConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBytePtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					return errTestBytePtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentBytePtrSliceConsumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestBytePtrSliceConsumer)
		})
	}
}

func TestSilentBytePtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testBytePtrSliceConsumerFactory
		cf2   testBytePtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					calls++
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					calls++
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					calls++
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestBytePtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					calls++
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					calls++
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BytePtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentBytePtrSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentBytePtrSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentBytePtrSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestBytePtrSliceConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustBytePtrSliceConsumer(t *testing.T) {
	var mc MustBytePtrSliceConsumer = func(v []*byte) {
		require.Equal(t, valTestBytePtrSliceConsumer, v)
		return
	}
	mc(valTestBytePtrSliceConsumer)
}

func TestMustBytePtrSliceConsumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testBytePtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					return errTestBytePtrSliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustBytePtrSliceConsumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.wantErr {
				r.PanicsWithError(errTestBytePtrSliceConsumer.Error(), func() {
					mc(valTestBytePtrSliceConsumer)
				})
			} else {
				mc(valTestBytePtrSliceConsumer)
			}
		})
	}
}

func TestMustBytePtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testBytePtrSliceConsumerFactory
		cf2     testBytePtrSliceConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					calls++
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					calls++
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					calls++
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestBytePtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					calls++
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					calls++
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BytePtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustBytePtrSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustBytePtrSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustBytePtrSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.wantErr {
				r.PanicsWithError(errTestBytePtrSliceConsumer.Error(), func() {
					cmc(valTestBytePtrSliceConsumer)
				})
			} else {
				cmc(valTestBytePtrSliceConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustBytePtrSliceConsumer_ToSilentBytePtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBytePtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					return errTestBytePtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustBytePtrSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentBytePtrSliceConsumer()
			r.NotNil(sc)

			sc(valTestBytePtrSliceConsumer)
		})
	}
}

func TestMustBytePtrSliceConsumer_ToBytePtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testBytePtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BytePtrSliceConsumer {
				return func(v []*byte) error {
					require.Equal(t, valTestBytePtrSliceConsumer, v)
					return errTestBytePtrSliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustBytePtrSliceConsumer()
			r.NotNil(mc)

			c = mc.ToBytePtrSliceConsumer()
			r.NotNil(c)

			err := c(valTestBytePtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestBytePtrSliceConsumer.Error())
			}
		})
	}
}
