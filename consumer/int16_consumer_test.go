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
	valTestInt16Consumer int16
	errTestInt16Consumer = errors.New("error")
)

type testInt16ConsumerFactory func(t *testing.T) Int16Consumer

func TestInt16Consumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testInt16ConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					require.Equal(t, valTestInt16Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					require.Equal(t, valTestInt16Consumer, v)
					return errTestInt16Consumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestInt16Consumer)
			if tt.wantErr {
				r.EqualError(err, errTestInt16Consumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt16Consumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testInt16ConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					require.Equal(t, valTestInt16Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					require.Equal(t, valTestInt16Consumer, v)
					return errTestInt16Consumer
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

			err := c(valTestInt16Consumer)
			if tt.wantErr {
				r.EqualError(err, errTestInt16Consumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestInt16Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testInt16ConsumerFactory
		cf2     testInt16ConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, valTestInt16Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, valTestInt16Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, valTestInt16Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestInt16Consumer
				}
			},
			cf2: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, valTestInt16Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, valTestInt16Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16Consumer {
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
			err := cc(valTestInt16Consumer)
			if tt.wantErr {
				r.EqualError(err, errTestInt16Consumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestInt16Consumer_ToSilentInt16Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt16ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					require.Equal(t, valTestInt16Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					require.Equal(t, valTestInt16Consumer, v)
					return errTestInt16Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentInt16Consumer()
			r.NotNil(sc)

			sc(valTestInt16Consumer)
		})
	}
}

func TestInt16Consumer_ToMustInt16Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt16ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					require.Equal(t, valTestInt16Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					require.Equal(t, valTestInt16Consumer, v)
					return errTestInt16Consumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustInt16Consumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestInt16Consumer.Error(), func() {
					mc(valTestInt16Consumer)
				})
			} else {
				mc(valTestInt16Consumer)
			}
		})
	}
}

func TestSilentInt16Consumer(t *testing.T) {
	var sc SilentInt16Consumer = func(v int16) {
		require.Equal(t, valTestInt16Consumer, v)
		return
	}
	sc(valTestInt16Consumer)
}

func TestSilentInt16Consumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt16ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					require.Equal(t, valTestInt16Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					require.Equal(t, valTestInt16Consumer, v)
					return errTestInt16Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentInt16Consumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestInt16Consumer)
		})
	}
}

func TestSilentInt16Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testInt16ConsumerFactory
		cf2   testInt16ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, valTestInt16Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, valTestInt16Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, valTestInt16Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestInt16Consumer
				}
			},
			cf2: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, valTestInt16Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, valTestInt16Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentInt16Consumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentInt16Consumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentInt16Consumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestInt16Consumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt16Consumer(t *testing.T) {
	var mc MustInt16Consumer = func(v int16) {
		require.Equal(t, valTestInt16Consumer, v)
		return
	}
	mc(valTestInt16Consumer)
}

func TestMustInt16Consumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testInt16ConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					require.Equal(t, valTestInt16Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					require.Equal(t, valTestInt16Consumer, v)
					return errTestInt16Consumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustInt16Consumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.wantErr {
				r.PanicsWithError(errTestInt16Consumer.Error(), func() {
					mc(valTestInt16Consumer)
				})
			} else {
				mc(valTestInt16Consumer)
			}
		})
	}
}

func TestMustInt16Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testInt16ConsumerFactory
		cf2     testInt16ConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, valTestInt16Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, valTestInt16Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, valTestInt16Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestInt16Consumer
				}
			},
			cf2: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, valTestInt16Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					calls++
					require.Equal(t, valTestInt16Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Int16Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustInt16Consumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustInt16Consumer = nil
			if c2 != nil {
				mc2 = c2.ToMustInt16Consumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.wantErr {
				r.PanicsWithError(errTestInt16Consumer.Error(), func() {
					cmc(valTestInt16Consumer)
				})
			} else {
				cmc(valTestInt16Consumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustInt16Consumer_ToSilentInt16Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testInt16ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					require.Equal(t, valTestInt16Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					require.Equal(t, valTestInt16Consumer, v)
					return errTestInt16Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt16Consumer()
			r.NotNil(mc)

			sc := mc.ToSilentInt16Consumer()
			r.NotNil(sc)

			sc(valTestInt16Consumer)
		})
	}
}

func TestMustInt16Consumer_ToInt16Consumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testInt16ConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					require.Equal(t, valTestInt16Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Int16Consumer {
				return func(v int16) error {
					require.Equal(t, valTestInt16Consumer, v)
					return errTestInt16Consumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustInt16Consumer()
			r.NotNil(mc)

			c = mc.ToInt16Consumer()
			r.NotNil(c)

			err := c(valTestInt16Consumer)
			if tt.wantErr {
				r.EqualError(err, errTestInt16Consumer.Error())
			}
		})
	}
}
