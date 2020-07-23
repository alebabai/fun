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
	valTestInt32SliceConsumer []int32
	errTestInt32SliceConsumer = errors.New("error")
)

type testInt32SliceConsumerFactory func(t *testing.T) Int32SliceConsumer

func TestInt32SliceConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testInt32SliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					require.Equal(t, valTestInt32SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					require.Equal(t, valTestInt32SliceConsumer, v)
					return errTestInt32SliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestInt32SliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestInt32SliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt32SliceConsumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testInt32SliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					require.Equal(t, valTestInt32SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					require.Equal(t, valTestInt32SliceConsumer, v)
					return errTestInt32SliceConsumer
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

			err := c(valTestInt32SliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestInt32SliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt32SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testInt32SliceConsumerFactory
		cf2     testInt32SliceConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, valTestInt32SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, valTestInt32SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, valTestInt32SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestInt32SliceConsumer
				}
			},
			cf2: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, valTestInt32SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, valTestInt32SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32SliceConsumer {
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
			err := cc(valTestInt32SliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestInt32SliceConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestInt32SliceConsumer_ToSilentInt32SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					require.Equal(t, valTestInt32SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					require.Equal(t, valTestInt32SliceConsumer, v)
					return errTestInt32SliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentInt32SliceConsumer()
			r.NotNil(sc)

			sc(valTestInt32SliceConsumer)
		})
	}
}

func TestInt32SliceConsumer_ToMustInt32SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32SliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					require.Equal(t, valTestInt32SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					require.Equal(t, valTestInt32SliceConsumer, v)
					return errTestInt32SliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustInt32SliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestInt32SliceConsumer.Error(), func() {
					mc(valTestInt32SliceConsumer)
				})
			} else {
				mc(valTestInt32SliceConsumer)
			}
		})
	}
}

func TestSilentInt32SliceConsumer(t *testing.T) {
	var sc SilentInt32SliceConsumer = func(v []int32) {
		require.Equal(t, valTestInt32SliceConsumer, v)
		return
	}
	sc(valTestInt32SliceConsumer)
}

func TestSilentInt32SliceConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					require.Equal(t, valTestInt32SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					require.Equal(t, valTestInt32SliceConsumer, v)
					return errTestInt32SliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentInt32SliceConsumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestInt32SliceConsumer)
		})
	}
}

func TestSilentInt32SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt32SliceConsumerFactory
		cf2   testInt32SliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, valTestInt32SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, valTestInt32SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, valTestInt32SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestInt32SliceConsumer
				}
			},
			cf2: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, valTestInt32SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, valTestInt32SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32SliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentInt32SliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentInt32SliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentInt32SliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestInt32SliceConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt32SliceConsumer(t *testing.T) {
	var mc MustInt32SliceConsumer = func(v []int32) {
		require.Equal(t, valTestInt32SliceConsumer, v)
		return
	}
	mc(valTestInt32SliceConsumer)
}

func TestMustInt32SliceConsumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testInt32SliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					require.Equal(t, valTestInt32SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					require.Equal(t, valTestInt32SliceConsumer, v)
					return errTestInt32SliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustInt32SliceConsumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.wantErr {
				r.PanicsWithError(errTestInt32SliceConsumer.Error(), func() {
					mc(valTestInt32SliceConsumer)
				})
			} else {
				mc(valTestInt32SliceConsumer)
			}
		})
	}
}

func TestMustInt32SliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testInt32SliceConsumerFactory
		cf2     testInt32SliceConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, valTestInt32SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, valTestInt32SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, valTestInt32SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestInt32SliceConsumer
				}
			},
			cf2: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, valTestInt32SliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					calls++
					require.Equal(t, valTestInt32SliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int32SliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustInt32SliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustInt32SliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustInt32SliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.wantErr {
				r.PanicsWithError(errTestInt32SliceConsumer.Error(), func() {
					cmc(valTestInt32SliceConsumer)
				})
			} else {
				cmc(valTestInt32SliceConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt32SliceConsumer_ToSilentInt32SliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt32SliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					require.Equal(t, valTestInt32SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					require.Equal(t, valTestInt32SliceConsumer, v)
					return errTestInt32SliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt32SliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentInt32SliceConsumer()
			r.NotNil(sc)

			sc(valTestInt32SliceConsumer)
		})
	}
}

func TestMustInt32SliceConsumer_ToInt32SliceConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testInt32SliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					require.Equal(t, valTestInt32SliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Int32SliceConsumer {
				return func(v []int32) error {
					require.Equal(t, valTestInt32SliceConsumer, v)
					return errTestInt32SliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt32SliceConsumer()
			r.NotNil(mc)

			c = mc.ToInt32SliceConsumer()
			r.NotNil(c)

			err := c(valTestInt32SliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestInt32SliceConsumer.Error())
			}
		})
	}
}
