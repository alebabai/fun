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
	valTestByteConsumer byte
	errTestByteConsumer = errors.New("error")
)

type testByteConsumerFactory func(t *testing.T) ByteConsumer

func TestByteConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testByteConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					require.Equal(t, valTestByteConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					require.Equal(t, valTestByteConsumer, v)
					return errTestByteConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestByteConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestByteConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestByteConsumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testByteConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					require.Equal(t, valTestByteConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					require.Equal(t, valTestByteConsumer, v)
					return errTestByteConsumer
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

			err := c(valTestByteConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestByteConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestByteConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testByteConsumerFactory
		cf2     testByteConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, valTestByteConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, valTestByteConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, valTestByteConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestByteConsumer
				}
			},
			cf2: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, valTestByteConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, valTestByteConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) ByteConsumer {
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
			err := cc(valTestByteConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestByteConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestByteConsumer_ToSilentByteConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testByteConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					require.Equal(t, valTestByteConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					require.Equal(t, valTestByteConsumer, v)
					return errTestByteConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentByteConsumer()
			r.NotNil(sc)

			sc(valTestByteConsumer)
		})
	}
}

func TestByteConsumer_ToMustByteConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testByteConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					require.Equal(t, valTestByteConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					require.Equal(t, valTestByteConsumer, v)
					return errTestByteConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustByteConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestByteConsumer.Error(), func() {
					mc(valTestByteConsumer)
				})
			} else {
				mc(valTestByteConsumer)
			}
		})
	}
}

func TestSilentByteConsumer(t *testing.T) {
	var sc SilentByteConsumer = func(v byte) {
		require.Equal(t, valTestByteConsumer, v)
		return
	}
	sc(valTestByteConsumer)
}

func TestSilentByteConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testByteConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					require.Equal(t, valTestByteConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					require.Equal(t, valTestByteConsumer, v)
					return errTestByteConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentByteConsumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestByteConsumer)
		})
	}
}

func TestSilentByteConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testByteConsumerFactory
		cf2   testByteConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, valTestByteConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, valTestByteConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, valTestByteConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestByteConsumer
				}
			},
			cf2: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, valTestByteConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, valTestByteConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) ByteConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentByteConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentByteConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentByteConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestByteConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustByteConsumer(t *testing.T) {
	var mc MustByteConsumer = func(v byte) {
		require.Equal(t, valTestByteConsumer, v)
		return
	}
	mc(valTestByteConsumer)
}

func TestMustByteConsumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testByteConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					require.Equal(t, valTestByteConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					require.Equal(t, valTestByteConsumer, v)
					return errTestByteConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustByteConsumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.wantErr {
				r.PanicsWithError(errTestByteConsumer.Error(), func() {
					mc(valTestByteConsumer)
				})
			} else {
				mc(valTestByteConsumer)
			}
		})
	}
}

func TestMustByteConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testByteConsumerFactory
		cf2     testByteConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, valTestByteConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, valTestByteConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, valTestByteConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestByteConsumer
				}
			},
			cf2: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, valTestByteConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					calls++
					require.Equal(t, valTestByteConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) ByteConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustByteConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustByteConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustByteConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.wantErr {
				r.PanicsWithError(errTestByteConsumer.Error(), func() {
					cmc(valTestByteConsumer)
				})
			} else {
				cmc(valTestByteConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustByteConsumer_ToSilentByteConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testByteConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					require.Equal(t, valTestByteConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					require.Equal(t, valTestByteConsumer, v)
					return errTestByteConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustByteConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentByteConsumer()
			r.NotNil(sc)

			sc(valTestByteConsumer)
		})
	}
}

func TestMustByteConsumer_ToByteConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testByteConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					require.Equal(t, valTestByteConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) ByteConsumer {
				return func(v byte) error {
					require.Equal(t, valTestByteConsumer, v)
					return errTestByteConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustByteConsumer()
			r.NotNil(mc)

			c = mc.ToByteConsumer()
			r.NotNil(c)

			err := c(valTestByteConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestByteConsumer.Error())
			}
		})
	}
}
