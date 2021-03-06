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
	valTestUint16Consumer uint16
	errTestUint16Consumer = errors.New("error")
)

type testUint16ConsumerFactory func(t *testing.T) Uint16Consumer

func TestUint16Consumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUint16ConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					require.Equal(t, valTestUint16Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					require.Equal(t, valTestUint16Consumer, v)
					return errTestUint16Consumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestUint16Consumer)
			if tt.wantErr {
				r.EqualError(err, errTestUint16Consumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint16Consumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUint16ConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					require.Equal(t, valTestUint16Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					require.Equal(t, valTestUint16Consumer, v)
					return errTestUint16Consumer
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

			err := c(valTestUint16Consumer)
			if tt.wantErr {
				r.EqualError(err, errTestUint16Consumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint16Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testUint16ConsumerFactory
		cf2     testUint16ConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, valTestUint16Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, valTestUint16Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, valTestUint16Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint16Consumer
				}
			},
			cf2: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, valTestUint16Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, valTestUint16Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16Consumer {
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
			err := cc(valTestUint16Consumer)
			if tt.wantErr {
				r.EqualError(err, errTestUint16Consumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestUint16Consumer_ToSilentUint16Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint16ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					require.Equal(t, valTestUint16Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					require.Equal(t, valTestUint16Consumer, v)
					return errTestUint16Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentUint16Consumer()
			r.NotNil(sc)

			sc(valTestUint16Consumer)
		})
	}
}

func TestUint16Consumer_ToMustUint16Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint16ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					require.Equal(t, valTestUint16Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					require.Equal(t, valTestUint16Consumer, v)
					return errTestUint16Consumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustUint16Consumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestUint16Consumer.Error(), func() {
					mc(valTestUint16Consumer)
				})
			} else {
				mc(valTestUint16Consumer)
			}
		})
	}
}

func TestSilentUint16Consumer(t *testing.T) {
	var sc SilentUint16Consumer = func(v uint16) {
		require.Equal(t, valTestUint16Consumer, v)
		return
	}
	sc(valTestUint16Consumer)
}

func TestSilentUint16Consumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint16ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					require.Equal(t, valTestUint16Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					require.Equal(t, valTestUint16Consumer, v)
					return errTestUint16Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentUint16Consumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestUint16Consumer)
		})
	}
}

func TestSilentUint16Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint16ConsumerFactory
		cf2   testUint16ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, valTestUint16Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, valTestUint16Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, valTestUint16Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint16Consumer
				}
			},
			cf2: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, valTestUint16Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, valTestUint16Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentUint16Consumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentUint16Consumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentUint16Consumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestUint16Consumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint16Consumer(t *testing.T) {
	var mc MustUint16Consumer = func(v uint16) {
		require.Equal(t, valTestUint16Consumer, v)
		return
	}
	mc(valTestUint16Consumer)
}

func TestMustUint16Consumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUint16ConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					require.Equal(t, valTestUint16Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					require.Equal(t, valTestUint16Consumer, v)
					return errTestUint16Consumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustUint16Consumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.wantErr {
				r.PanicsWithError(errTestUint16Consumer.Error(), func() {
					mc(valTestUint16Consumer)
				})
			} else {
				mc(valTestUint16Consumer)
			}
		})
	}
}

func TestMustUint16Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testUint16ConsumerFactory
		cf2     testUint16ConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, valTestUint16Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, valTestUint16Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, valTestUint16Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint16Consumer
				}
			},
			cf2: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, valTestUint16Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					calls++
					require.Equal(t, valTestUint16Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint16Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustUint16Consumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustUint16Consumer = nil
			if c2 != nil {
				mc2 = c2.ToMustUint16Consumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.wantErr {
				r.PanicsWithError(errTestUint16Consumer.Error(), func() {
					cmc(valTestUint16Consumer)
				})
			} else {
				cmc(valTestUint16Consumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint16Consumer_ToSilentUint16Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint16ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					require.Equal(t, valTestUint16Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					require.Equal(t, valTestUint16Consumer, v)
					return errTestUint16Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint16Consumer()
			r.NotNil(mc)

			sc := mc.ToSilentUint16Consumer()
			r.NotNil(sc)

			sc(valTestUint16Consumer)
		})
	}
}

func TestMustUint16Consumer_ToUint16Consumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUint16ConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					require.Equal(t, valTestUint16Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint16Consumer {
				return func(v uint16) error {
					require.Equal(t, valTestUint16Consumer, v)
					return errTestUint16Consumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint16Consumer()
			r.NotNil(mc)

			c = mc.ToUint16Consumer()
			r.NotNil(c)

			err := c(valTestUint16Consumer)
			if tt.wantErr {
				r.EqualError(err, errTestUint16Consumer.Error())
			}
		})
	}
}
