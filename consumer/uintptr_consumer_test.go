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
	valTestUintPtrConsumer *uint
	errTestUintPtrConsumer = errors.New("error")
)

type testUintPtrConsumerFactory func(t *testing.T) UintPtrConsumer

func TestUintPtrConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUintPtrConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					require.Equal(t, valTestUintPtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					require.Equal(t, valTestUintPtrConsumer, v)
					return errTestUintPtrConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestUintPtrConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestUintPtrConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUintPtrConsumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUintPtrConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					require.Equal(t, valTestUintPtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					require.Equal(t, valTestUintPtrConsumer, v)
					return errTestUintPtrConsumer
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

			err := c(valTestUintPtrConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestUintPtrConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUintPtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testUintPtrConsumerFactory
		cf2     testUintPtrConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					calls++
					require.Equal(t, valTestUintPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					calls++
					require.Equal(t, valTestUintPtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					calls++
					require.Equal(t, valTestUintPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUintPtrConsumer
				}
			},
			cf2: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					calls++
					require.Equal(t, valTestUintPtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					calls++
					require.Equal(t, valTestUintPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintPtrConsumer {
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
			err := cc(valTestUintPtrConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestUintPtrConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestUintPtrConsumer_ToSilentUintPtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintPtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					require.Equal(t, valTestUintPtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					require.Equal(t, valTestUintPtrConsumer, v)
					return errTestUintPtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentUintPtrConsumer()
			r.NotNil(sc)

			sc(valTestUintPtrConsumer)
		})
	}
}

func TestUintPtrConsumer_ToMustUintPtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintPtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					require.Equal(t, valTestUintPtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					require.Equal(t, valTestUintPtrConsumer, v)
					return errTestUintPtrConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustUintPtrConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestUintPtrConsumer.Error(), func() {
					mc(valTestUintPtrConsumer)
				})
			} else {
				mc(valTestUintPtrConsumer)
			}
		})
	}
}

func TestSilentUintPtrConsumer(t *testing.T) {
	var sc SilentUintPtrConsumer = func(v *uint) {
		require.Equal(t, valTestUintPtrConsumer, v)
		return
	}
	sc(valTestUintPtrConsumer)
}

func TestSilentUintPtrConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintPtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					require.Equal(t, valTestUintPtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					require.Equal(t, valTestUintPtrConsumer, v)
					return errTestUintPtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentUintPtrConsumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestUintPtrConsumer)
		})
	}
}

func TestSilentUintPtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUintPtrConsumerFactory
		cf2   testUintPtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					calls++
					require.Equal(t, valTestUintPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					calls++
					require.Equal(t, valTestUintPtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					calls++
					require.Equal(t, valTestUintPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUintPtrConsumer
				}
			},
			cf2: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					calls++
					require.Equal(t, valTestUintPtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					calls++
					require.Equal(t, valTestUintPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintPtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentUintPtrConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentUintPtrConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentUintPtrConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestUintPtrConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUintPtrConsumer(t *testing.T) {
	var mc MustUintPtrConsumer = func(v *uint) {
		require.Equal(t, valTestUintPtrConsumer, v)
		return
	}
	mc(valTestUintPtrConsumer)
}

func TestMustUintPtrConsumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUintPtrConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					require.Equal(t, valTestUintPtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					require.Equal(t, valTestUintPtrConsumer, v)
					return errTestUintPtrConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustUintPtrConsumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.wantErr {
				r.PanicsWithError(errTestUintPtrConsumer.Error(), func() {
					mc(valTestUintPtrConsumer)
				})
			} else {
				mc(valTestUintPtrConsumer)
			}
		})
	}
}

func TestMustUintPtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testUintPtrConsumerFactory
		cf2     testUintPtrConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					calls++
					require.Equal(t, valTestUintPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					calls++
					require.Equal(t, valTestUintPtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					calls++
					require.Equal(t, valTestUintPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUintPtrConsumer
				}
			},
			cf2: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					calls++
					require.Equal(t, valTestUintPtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					calls++
					require.Equal(t, valTestUintPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintPtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustUintPtrConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustUintPtrConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustUintPtrConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.wantErr {
				r.PanicsWithError(errTestUintPtrConsumer.Error(), func() {
					cmc(valTestUintPtrConsumer)
				})
			} else {
				cmc(valTestUintPtrConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUintPtrConsumer_ToSilentUintPtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintPtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					require.Equal(t, valTestUintPtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					require.Equal(t, valTestUintPtrConsumer, v)
					return errTestUintPtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUintPtrConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentUintPtrConsumer()
			r.NotNil(sc)

			sc(valTestUintPtrConsumer)
		})
	}
}

func TestMustUintPtrConsumer_ToUintPtrConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUintPtrConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					require.Equal(t, valTestUintPtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) UintPtrConsumer {
				return func(v *uint) error {
					require.Equal(t, valTestUintPtrConsumer, v)
					return errTestUintPtrConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUintPtrConsumer()
			r.NotNil(mc)

			c = mc.ToUintPtrConsumer()
			r.NotNil(c)

			err := c(valTestUintPtrConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestUintPtrConsumer.Error())
			}
		})
	}
}
