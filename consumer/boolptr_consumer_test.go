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
	valTestBoolPtrConsumer *bool
	errTestBoolPtrConsumer = errors.New("error")
)

type testBoolPtrConsumerFactory func(t *testing.T) BoolPtrConsumer

func TestBoolPtrConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testBoolPtrConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					require.Equal(t, valTestBoolPtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					require.Equal(t, valTestBoolPtrConsumer, v)
					return errTestBoolPtrConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestBoolPtrConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestBoolPtrConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestBoolPtrConsumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testBoolPtrConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					require.Equal(t, valTestBoolPtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					require.Equal(t, valTestBoolPtrConsumer, v)
					return errTestBoolPtrConsumer
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

			err := c(valTestBoolPtrConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestBoolPtrConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestBoolPtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testBoolPtrConsumerFactory
		cf2     testBoolPtrConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, valTestBoolPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, valTestBoolPtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, valTestBoolPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestBoolPtrConsumer
				}
			},
			cf2: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, valTestBoolPtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, valTestBoolPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolPtrConsumer {
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
			err := cc(valTestBoolPtrConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestBoolPtrConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestBoolPtrConsumer_ToSilentBoolPtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolPtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					require.Equal(t, valTestBoolPtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					require.Equal(t, valTestBoolPtrConsumer, v)
					return errTestBoolPtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentBoolPtrConsumer()
			r.NotNil(sc)

			sc(valTestBoolPtrConsumer)
		})
	}
}

func TestBoolPtrConsumer_ToMustBoolPtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolPtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					require.Equal(t, valTestBoolPtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					require.Equal(t, valTestBoolPtrConsumer, v)
					return errTestBoolPtrConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustBoolPtrConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestBoolPtrConsumer.Error(), func() {
					mc(valTestBoolPtrConsumer)
				})
			} else {
				mc(valTestBoolPtrConsumer)
			}
		})
	}
}

func TestSilentBoolPtrConsumer(t *testing.T) {
	var sc SilentBoolPtrConsumer = func(v *bool) {
		require.Equal(t, valTestBoolPtrConsumer, v)
		return
	}
	sc(valTestBoolPtrConsumer)
}

func TestSilentBoolPtrConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolPtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					require.Equal(t, valTestBoolPtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					require.Equal(t, valTestBoolPtrConsumer, v)
					return errTestBoolPtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentBoolPtrConsumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestBoolPtrConsumer)
		})
	}
}

func TestSilentBoolPtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testBoolPtrConsumerFactory
		cf2   testBoolPtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, valTestBoolPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, valTestBoolPtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, valTestBoolPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestBoolPtrConsumer
				}
			},
			cf2: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, valTestBoolPtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, valTestBoolPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolPtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentBoolPtrConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentBoolPtrConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentBoolPtrConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestBoolPtrConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustBoolPtrConsumer(t *testing.T) {
	var mc MustBoolPtrConsumer = func(v *bool) {
		require.Equal(t, valTestBoolPtrConsumer, v)
		return
	}
	mc(valTestBoolPtrConsumer)
}

func TestMustBoolPtrConsumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testBoolPtrConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					require.Equal(t, valTestBoolPtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					require.Equal(t, valTestBoolPtrConsumer, v)
					return errTestBoolPtrConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustBoolPtrConsumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.wantErr {
				r.PanicsWithError(errTestBoolPtrConsumer.Error(), func() {
					mc(valTestBoolPtrConsumer)
				})
			} else {
				mc(valTestBoolPtrConsumer)
			}
		})
	}
}

func TestMustBoolPtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testBoolPtrConsumerFactory
		cf2     testBoolPtrConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, valTestBoolPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, valTestBoolPtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, valTestBoolPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestBoolPtrConsumer
				}
			},
			cf2: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, valTestBoolPtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, valTestBoolPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolPtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustBoolPtrConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustBoolPtrConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustBoolPtrConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.wantErr {
				r.PanicsWithError(errTestBoolPtrConsumer.Error(), func() {
					cmc(valTestBoolPtrConsumer)
				})
			} else {
				cmc(valTestBoolPtrConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustBoolPtrConsumer_ToSilentBoolPtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolPtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					require.Equal(t, valTestBoolPtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					require.Equal(t, valTestBoolPtrConsumer, v)
					return errTestBoolPtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustBoolPtrConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentBoolPtrConsumer()
			r.NotNil(sc)

			sc(valTestBoolPtrConsumer)
		})
	}
}

func TestMustBoolPtrConsumer_ToBoolPtrConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testBoolPtrConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					require.Equal(t, valTestBoolPtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					require.Equal(t, valTestBoolPtrConsumer, v)
					return errTestBoolPtrConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustBoolPtrConsumer()
			r.NotNil(mc)

			c = mc.ToBoolPtrConsumer()
			r.NotNil(c)

			err := c(valTestBoolPtrConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestBoolPtrConsumer.Error())
			}
		})
	}
}
