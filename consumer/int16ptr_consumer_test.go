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
	valTestInt16PtrConsumer *int16
	errTestInt16PtrConsumer = errors.New("error")
)

type testInt16PtrConsumerFactory func(t *testing.T) Int16PtrConsumer

func TestInt16PtrConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testInt16PtrConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					require.Equal(t, valTestInt16PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					require.Equal(t, valTestInt16PtrConsumer, v)
					return errTestInt16PtrConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestInt16PtrConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestInt16PtrConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt16PtrConsumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testInt16PtrConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					require.Equal(t, valTestInt16PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					require.Equal(t, valTestInt16PtrConsumer, v)
					return errTestInt16PtrConsumer
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

			err := c(valTestInt16PtrConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestInt16PtrConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt16PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testInt16PtrConsumerFactory
		cf2     testInt16PtrConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					calls++
					require.Equal(t, valTestInt16PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					calls++
					require.Equal(t, valTestInt16PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					calls++
					require.Equal(t, valTestInt16PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestInt16PtrConsumer
				}
			},
			cf2: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					calls++
					require.Equal(t, valTestInt16PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					calls++
					require.Equal(t, valTestInt16PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16PtrConsumer {
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
			err := cc(valTestInt16PtrConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestInt16PtrConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestInt16PtrConsumer_ToSilentInt16PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt16PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					require.Equal(t, valTestInt16PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					require.Equal(t, valTestInt16PtrConsumer, v)
					return errTestInt16PtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentInt16PtrConsumer()
			r.NotNil(sc)

			sc(valTestInt16PtrConsumer)
		})
	}
}

func TestInt16PtrConsumer_ToMustInt16PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt16PtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					require.Equal(t, valTestInt16PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					require.Equal(t, valTestInt16PtrConsumer, v)
					return errTestInt16PtrConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustInt16PtrConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestInt16PtrConsumer.Error(), func() {
					mc(valTestInt16PtrConsumer)
				})
			} else {
				mc(valTestInt16PtrConsumer)
			}
		})
	}
}

func TestSilentInt16PtrConsumer(t *testing.T) {
	var sc SilentInt16PtrConsumer = func(v *int16) {
		require.Equal(t, valTestInt16PtrConsumer, v)
		return
	}
	sc(valTestInt16PtrConsumer)
}

func TestSilentInt16PtrConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt16PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					require.Equal(t, valTestInt16PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					require.Equal(t, valTestInt16PtrConsumer, v)
					return errTestInt16PtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentInt16PtrConsumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestInt16PtrConsumer)
		})
	}
}

func TestSilentInt16PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt16PtrConsumerFactory
		cf2   testInt16PtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					calls++
					require.Equal(t, valTestInt16PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					calls++
					require.Equal(t, valTestInt16PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					calls++
					require.Equal(t, valTestInt16PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestInt16PtrConsumer
				}
			},
			cf2: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					calls++
					require.Equal(t, valTestInt16PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					calls++
					require.Equal(t, valTestInt16PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16PtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentInt16PtrConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentInt16PtrConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentInt16PtrConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestInt16PtrConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt16PtrConsumer(t *testing.T) {
	var mc MustInt16PtrConsumer = func(v *int16) {
		require.Equal(t, valTestInt16PtrConsumer, v)
		return
	}
	mc(valTestInt16PtrConsumer)
}

func TestMustInt16PtrConsumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testInt16PtrConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					require.Equal(t, valTestInt16PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					require.Equal(t, valTestInt16PtrConsumer, v)
					return errTestInt16PtrConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustInt16PtrConsumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.wantErr {
				r.PanicsWithError(errTestInt16PtrConsumer.Error(), func() {
					mc(valTestInt16PtrConsumer)
				})
			} else {
				mc(valTestInt16PtrConsumer)
			}
		})
	}
}

func TestMustInt16PtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testInt16PtrConsumerFactory
		cf2     testInt16PtrConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					calls++
					require.Equal(t, valTestInt16PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					calls++
					require.Equal(t, valTestInt16PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					calls++
					require.Equal(t, valTestInt16PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestInt16PtrConsumer
				}
			},
			cf2: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					calls++
					require.Equal(t, valTestInt16PtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					calls++
					require.Equal(t, valTestInt16PtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16PtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustInt16PtrConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustInt16PtrConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustInt16PtrConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.wantErr {
				r.PanicsWithError(errTestInt16PtrConsumer.Error(), func() {
					cmc(valTestInt16PtrConsumer)
				})
			} else {
				cmc(valTestInt16PtrConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt16PtrConsumer_ToSilentInt16PtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt16PtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					require.Equal(t, valTestInt16PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					require.Equal(t, valTestInt16PtrConsumer, v)
					return errTestInt16PtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt16PtrConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentInt16PtrConsumer()
			r.NotNil(sc)

			sc(valTestInt16PtrConsumer)
		})
	}
}

func TestMustInt16PtrConsumer_ToInt16PtrConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testInt16PtrConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					require.Equal(t, valTestInt16PtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Int16PtrConsumer {
				return func(v *int16) error {
					require.Equal(t, valTestInt16PtrConsumer, v)
					return errTestInt16PtrConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt16PtrConsumer()
			r.NotNil(mc)

			c = mc.ToInt16PtrConsumer()
			r.NotNil(c)

			err := c(valTestInt16PtrConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestInt16PtrConsumer.Error())
			}
		})
	}
}
