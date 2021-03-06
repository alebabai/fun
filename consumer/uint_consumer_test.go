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
	valTestUintConsumer uint
	errTestUintConsumer = errors.New("error")
)

type testUintConsumerFactory func(t *testing.T) UintConsumer

func TestUintConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUintConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					require.Equal(t, valTestUintConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					require.Equal(t, valTestUintConsumer, v)
					return errTestUintConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestUintConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestUintConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUintConsumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUintConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					require.Equal(t, valTestUintConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					require.Equal(t, valTestUintConsumer, v)
					return errTestUintConsumer
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

			err := c(valTestUintConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestUintConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUintConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testUintConsumerFactory
		cf2     testUintConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, valTestUintConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, valTestUintConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, valTestUintConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUintConsumer
				}
			},
			cf2: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, valTestUintConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, valTestUintConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintConsumer {
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
			err := cc(valTestUintConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestUintConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestUintConsumer_ToSilentUintConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					require.Equal(t, valTestUintConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					require.Equal(t, valTestUintConsumer, v)
					return errTestUintConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentUintConsumer()
			r.NotNil(sc)

			sc(valTestUintConsumer)
		})
	}
}

func TestUintConsumer_ToMustUintConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					require.Equal(t, valTestUintConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					require.Equal(t, valTestUintConsumer, v)
					return errTestUintConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustUintConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestUintConsumer.Error(), func() {
					mc(valTestUintConsumer)
				})
			} else {
				mc(valTestUintConsumer)
			}
		})
	}
}

func TestSilentUintConsumer(t *testing.T) {
	var sc SilentUintConsumer = func(v uint) {
		require.Equal(t, valTestUintConsumer, v)
		return
	}
	sc(valTestUintConsumer)
}

func TestSilentUintConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					require.Equal(t, valTestUintConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					require.Equal(t, valTestUintConsumer, v)
					return errTestUintConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentUintConsumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestUintConsumer)
		})
	}
}

func TestSilentUintConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUintConsumerFactory
		cf2   testUintConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, valTestUintConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, valTestUintConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, valTestUintConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUintConsumer
				}
			},
			cf2: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, valTestUintConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, valTestUintConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentUintConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentUintConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentUintConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestUintConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUintConsumer(t *testing.T) {
	var mc MustUintConsumer = func(v uint) {
		require.Equal(t, valTestUintConsumer, v)
		return
	}
	mc(valTestUintConsumer)
}

func TestMustUintConsumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUintConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					require.Equal(t, valTestUintConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					require.Equal(t, valTestUintConsumer, v)
					return errTestUintConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustUintConsumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.wantErr {
				r.PanicsWithError(errTestUintConsumer.Error(), func() {
					mc(valTestUintConsumer)
				})
			} else {
				mc(valTestUintConsumer)
			}
		})
	}
}

func TestMustUintConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testUintConsumerFactory
		cf2     testUintConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, valTestUintConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, valTestUintConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, valTestUintConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUintConsumer
				}
			},
			cf2: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, valTestUintConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, valTestUintConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustUintConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustUintConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustUintConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.wantErr {
				r.PanicsWithError(errTestUintConsumer.Error(), func() {
					cmc(valTestUintConsumer)
				})
			} else {
				cmc(valTestUintConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUintConsumer_ToSilentUintConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					require.Equal(t, valTestUintConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					require.Equal(t, valTestUintConsumer, v)
					return errTestUintConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUintConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentUintConsumer()
			r.NotNil(sc)

			sc(valTestUintConsumer)
		})
	}
}

func TestMustUintConsumer_ToUintConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUintConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					require.Equal(t, valTestUintConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					require.Equal(t, valTestUintConsumer, v)
					return errTestUintConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUintConsumer()
			r.NotNil(mc)

			c = mc.ToUintConsumer()
			r.NotNil(c)

			err := c(valTestUintConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestUintConsumer.Error())
			}
		})
	}
}
