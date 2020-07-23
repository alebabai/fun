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
	valTestUint64Consumer uint64
	errTestUint64Consumer = errors.New("error")
)

type testUint64ConsumerFactory func(t *testing.T) Uint64Consumer

func TestUint64Consumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUint64ConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					require.Equal(t, valTestUint64Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					require.Equal(t, valTestUint64Consumer, v)
					return errTestUint64Consumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestUint64Consumer)
			if tt.wantErr {
				r.EqualError(err, errTestUint64Consumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint64Consumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUint64ConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					require.Equal(t, valTestUint64Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					require.Equal(t, valTestUint64Consumer, v)
					return errTestUint64Consumer
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

			err := c(valTestUint64Consumer)
			if tt.wantErr {
				r.EqualError(err, errTestUint64Consumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUint64Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testUint64ConsumerFactory
		cf2     testUint64ConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, valTestUint64Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, valTestUint64Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, valTestUint64Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint64Consumer
				}
			},
			cf2: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, valTestUint64Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, valTestUint64Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64Consumer {
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
			err := cc(valTestUint64Consumer)
			if tt.wantErr {
				r.EqualError(err, errTestUint64Consumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestUint64Consumer_ToSilentUint64Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					require.Equal(t, valTestUint64Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					require.Equal(t, valTestUint64Consumer, v)
					return errTestUint64Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentUint64Consumer()
			r.NotNil(sc)

			sc(valTestUint64Consumer)
		})
	}
}

func TestUint64Consumer_ToMustUint64Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					require.Equal(t, valTestUint64Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					require.Equal(t, valTestUint64Consumer, v)
					return errTestUint64Consumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustUint64Consumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestUint64Consumer.Error(), func() {
					mc(valTestUint64Consumer)
				})
			} else {
				mc(valTestUint64Consumer)
			}
		})
	}
}

func TestSilentUint64Consumer(t *testing.T) {
	var sc SilentUint64Consumer = func(v uint64) {
		require.Equal(t, valTestUint64Consumer, v)
		return
	}
	sc(valTestUint64Consumer)
}

func TestSilentUint64Consumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					require.Equal(t, valTestUint64Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					require.Equal(t, valTestUint64Consumer, v)
					return errTestUint64Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentUint64Consumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestUint64Consumer)
		})
	}
}

func TestSilentUint64Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUint64ConsumerFactory
		cf2   testUint64ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, valTestUint64Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, valTestUint64Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, valTestUint64Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint64Consumer
				}
			},
			cf2: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, valTestUint64Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, valTestUint64Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentUint64Consumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentUint64Consumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentUint64Consumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestUint64Consumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint64Consumer(t *testing.T) {
	var mc MustUint64Consumer = func(v uint64) {
		require.Equal(t, valTestUint64Consumer, v)
		return
	}
	mc(valTestUint64Consumer)
}

func TestMustUint64Consumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUint64ConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					require.Equal(t, valTestUint64Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					require.Equal(t, valTestUint64Consumer, v)
					return errTestUint64Consumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustUint64Consumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.wantErr {
				r.PanicsWithError(errTestUint64Consumer.Error(), func() {
					mc(valTestUint64Consumer)
				})
			} else {
				mc(valTestUint64Consumer)
			}
		})
	}
}

func TestMustUint64Consumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testUint64ConsumerFactory
		cf2     testUint64ConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, valTestUint64Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, valTestUint64Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, valTestUint64Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestUint64Consumer
				}
			},
			cf2: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, valTestUint64Consumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					calls++
					require.Equal(t, valTestUint64Consumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Uint64Consumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustUint64Consumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustUint64Consumer = nil
			if c2 != nil {
				mc2 = c2.ToMustUint64Consumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.wantErr {
				r.PanicsWithError(errTestUint64Consumer.Error(), func() {
					cmc(valTestUint64Consumer)
				})
			} else {
				cmc(valTestUint64Consumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUint64Consumer_ToSilentUint64Consumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUint64ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					require.Equal(t, valTestUint64Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					require.Equal(t, valTestUint64Consumer, v)
					return errTestUint64Consumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint64Consumer()
			r.NotNil(mc)

			sc := mc.ToSilentUint64Consumer()
			r.NotNil(sc)

			sc(valTestUint64Consumer)
		})
	}
}

func TestMustUint64Consumer_ToUint64Consumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testUint64ConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					require.Equal(t, valTestUint64Consumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) Uint64Consumer {
				return func(v uint64) error {
					require.Equal(t, valTestUint64Consumer, v)
					return errTestUint64Consumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUint64Consumer()
			r.NotNil(mc)

			c = mc.ToUint64Consumer()
			r.NotNil(c)

			err := c(valTestUint64Consumer)
			if tt.wantErr {
				r.EqualError(err, errTestUint64Consumer.Error())
			}
		})
	}
}
