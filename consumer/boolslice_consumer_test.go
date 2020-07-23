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
	valTestBoolSliceConsumer []bool
	errTestBoolSliceConsumer = errors.New("error")
)

type testBoolSliceConsumerFactory func(t *testing.T) BoolSliceConsumer

func TestBoolSliceConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testBoolSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					require.Equal(t, valTestBoolSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					require.Equal(t, valTestBoolSliceConsumer, v)
					return errTestBoolSliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestBoolSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestBoolSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestBoolSliceConsumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testBoolSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					require.Equal(t, valTestBoolSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					require.Equal(t, valTestBoolSliceConsumer, v)
					return errTestBoolSliceConsumer
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

			err := c(valTestBoolSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestBoolSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestBoolSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testBoolSliceConsumerFactory
		cf2     testBoolSliceConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, valTestBoolSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, valTestBoolSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, valTestBoolSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestBoolSliceConsumer
				}
			},
			cf2: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, valTestBoolSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, valTestBoolSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolSliceConsumer {
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
			err := cc(valTestBoolSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestBoolSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestBoolSliceConsumer_ToSilentBoolSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					require.Equal(t, valTestBoolSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					require.Equal(t, valTestBoolSliceConsumer, v)
					return errTestBoolSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentBoolSliceConsumer()
			r.NotNil(sc)

			sc(valTestBoolSliceConsumer)
		})
	}
}

func TestBoolSliceConsumer_ToMustBoolSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					require.Equal(t, valTestBoolSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					require.Equal(t, valTestBoolSliceConsumer, v)
					return errTestBoolSliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustBoolSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestBoolSliceConsumer.Error(), func() {
					mc(valTestBoolSliceConsumer)
				})
			} else {
				mc(valTestBoolSliceConsumer)
			}
		})
	}
}

func TestSilentBoolSliceConsumer(t *testing.T) {
	var sc SilentBoolSliceConsumer = func(v []bool) {
		require.Equal(t, valTestBoolSliceConsumer, v)
		return
	}
	sc(valTestBoolSliceConsumer)
}

func TestSilentBoolSliceConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					require.Equal(t, valTestBoolSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					require.Equal(t, valTestBoolSliceConsumer, v)
					return errTestBoolSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentBoolSliceConsumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestBoolSliceConsumer)
		})
	}
}

func TestSilentBoolSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testBoolSliceConsumerFactory
		cf2   testBoolSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, valTestBoolSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, valTestBoolSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, valTestBoolSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestBoolSliceConsumer
				}
			},
			cf2: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, valTestBoolSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, valTestBoolSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentBoolSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentBoolSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentBoolSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestBoolSliceConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustBoolSliceConsumer(t *testing.T) {
	var mc MustBoolSliceConsumer = func(v []bool) {
		require.Equal(t, valTestBoolSliceConsumer, v)
		return
	}
	mc(valTestBoolSliceConsumer)
}

func TestMustBoolSliceConsumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testBoolSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					require.Equal(t, valTestBoolSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					require.Equal(t, valTestBoolSliceConsumer, v)
					return errTestBoolSliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustBoolSliceConsumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.wantErr {
				r.PanicsWithError(errTestBoolSliceConsumer.Error(), func() {
					mc(valTestBoolSliceConsumer)
				})
			} else {
				mc(valTestBoolSliceConsumer)
			}
		})
	}
}

func TestMustBoolSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testBoolSliceConsumerFactory
		cf2     testBoolSliceConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, valTestBoolSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, valTestBoolSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, valTestBoolSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestBoolSliceConsumer
				}
			},
			cf2: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, valTestBoolSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					calls++
					require.Equal(t, valTestBoolSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustBoolSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustBoolSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustBoolSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.wantErr {
				r.PanicsWithError(errTestBoolSliceConsumer.Error(), func() {
					cmc(valTestBoolSliceConsumer)
				})
			} else {
				cmc(valTestBoolSliceConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustBoolSliceConsumer_ToSilentBoolSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					require.Equal(t, valTestBoolSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					require.Equal(t, valTestBoolSliceConsumer, v)
					return errTestBoolSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustBoolSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentBoolSliceConsumer()
			r.NotNil(sc)

			sc(valTestBoolSliceConsumer)
		})
	}
}

func TestMustBoolSliceConsumer_ToBoolSliceConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testBoolSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					require.Equal(t, valTestBoolSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BoolSliceConsumer {
				return func(v []bool) error {
					require.Equal(t, valTestBoolSliceConsumer, v)
					return errTestBoolSliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustBoolSliceConsumer()
			r.NotNil(mc)

			c = mc.ToBoolSliceConsumer()
			r.NotNil(c)

			err := c(valTestBoolSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestBoolSliceConsumer.Error())
			}
		})
	}
}
