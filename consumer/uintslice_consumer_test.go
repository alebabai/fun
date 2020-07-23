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
	valTestUintSliceConsumer []uint
	errTestUintSliceConsumer = errors.New("error")
)

type testUintSliceConsumerFactory func(t *testing.T) UintSliceConsumer

func TestUintSliceConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUintSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					require.Equal(t, valTestUintSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					require.Equal(t, valTestUintSliceConsumer, v)
					return errTestUintSliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestUintSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestUintSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUintSliceConsumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUintSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					require.Equal(t, valTestUintSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					require.Equal(t, valTestUintSliceConsumer, v)
					return errTestUintSliceConsumer
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

			err := c(valTestUintSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestUintSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUintSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testUintSliceConsumerFactory
		cf2     testUintSliceConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, valTestUintSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, valTestUintSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, valTestUintSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUintSliceConsumer
				}
			},
			cf2: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, valTestUintSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, valTestUintSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintSliceConsumer {
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
			err := cc(valTestUintSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestUintSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestUintSliceConsumer_ToSilentUintSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					require.Equal(t, valTestUintSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					require.Equal(t, valTestUintSliceConsumer, v)
					return errTestUintSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentUintSliceConsumer()
			r.NotNil(sc)

			sc(valTestUintSliceConsumer)
		})
	}
}

func TestUintSliceConsumer_ToMustUintSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					require.Equal(t, valTestUintSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					require.Equal(t, valTestUintSliceConsumer, v)
					return errTestUintSliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustUintSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestUintSliceConsumer.Error(), func() {
					mc(valTestUintSliceConsumer)
				})
			} else {
				mc(valTestUintSliceConsumer)
			}
		})
	}
}

func TestSilentUintSliceConsumer(t *testing.T) {
	var sc SilentUintSliceConsumer = func(v []uint) {
		require.Equal(t, valTestUintSliceConsumer, v)
		return
	}
	sc(valTestUintSliceConsumer)
}

func TestSilentUintSliceConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					require.Equal(t, valTestUintSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					require.Equal(t, valTestUintSliceConsumer, v)
					return errTestUintSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentUintSliceConsumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestUintSliceConsumer)
		})
	}
}

func TestSilentUintSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUintSliceConsumerFactory
		cf2   testUintSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, valTestUintSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, valTestUintSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, valTestUintSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUintSliceConsumer
				}
			},
			cf2: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, valTestUintSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, valTestUintSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentUintSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentUintSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentUintSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestUintSliceConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUintSliceConsumer(t *testing.T) {
	var mc MustUintSliceConsumer = func(v []uint) {
		require.Equal(t, valTestUintSliceConsumer, v)
		return
	}
	mc(valTestUintSliceConsumer)
}

func TestMustUintSliceConsumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUintSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					require.Equal(t, valTestUintSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					require.Equal(t, valTestUintSliceConsumer, v)
					return errTestUintSliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustUintSliceConsumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.wantErr {
				r.PanicsWithError(errTestUintSliceConsumer.Error(), func() {
					mc(valTestUintSliceConsumer)
				})
			} else {
				mc(valTestUintSliceConsumer)
			}
		})
	}
}

func TestMustUintSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testUintSliceConsumerFactory
		cf2     testUintSliceConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, valTestUintSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, valTestUintSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, valTestUintSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUintSliceConsumer
				}
			},
			cf2: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, valTestUintSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					calls++
					require.Equal(t, valTestUintSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustUintSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustUintSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustUintSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.wantErr {
				r.PanicsWithError(errTestUintSliceConsumer.Error(), func() {
					cmc(valTestUintSliceConsumer)
				})
			} else {
				cmc(valTestUintSliceConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUintSliceConsumer_ToSilentUintSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					require.Equal(t, valTestUintSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					require.Equal(t, valTestUintSliceConsumer, v)
					return errTestUintSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUintSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentUintSliceConsumer()
			r.NotNil(sc)

			sc(valTestUintSliceConsumer)
		})
	}
}

func TestMustUintSliceConsumer_ToUintSliceConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUintSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					require.Equal(t, valTestUintSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) UintSliceConsumer {
				return func(v []uint) error {
					require.Equal(t, valTestUintSliceConsumer, v)
					return errTestUintSliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUintSliceConsumer()
			r.NotNil(mc)

			c = mc.ToUintSliceConsumer()
			r.NotNil(c)

			err := c(valTestUintSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestUintSliceConsumer.Error())
			}
		})
	}
}
