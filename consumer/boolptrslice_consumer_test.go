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
	valTestBoolPtrSliceConsumer []*bool
	errTestBoolPtrSliceConsumer = errors.New("error")
)

type testBoolPtrSliceConsumerFactory func(t *testing.T) BoolPtrSliceConsumer

func TestBoolPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testBoolPtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					return errTestBoolPtrSliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestBoolPtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestBoolPtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestBoolPtrSliceConsumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testBoolPtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					return errTestBoolPtrSliceConsumer
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

			err := c(valTestBoolPtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestBoolPtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestBoolPtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testBoolPtrSliceConsumerFactory
		cf2     testBoolPtrSliceConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestBoolPtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolPtrSliceConsumer {
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
			err := cc(valTestBoolPtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestBoolPtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestBoolPtrSliceConsumer_ToSilentBoolPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolPtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					return errTestBoolPtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentBoolPtrSliceConsumer()
			r.NotNil(sc)

			sc(valTestBoolPtrSliceConsumer)
		})
	}
}

func TestBoolPtrSliceConsumer_ToMustBoolPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolPtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					return errTestBoolPtrSliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustBoolPtrSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestBoolPtrSliceConsumer.Error(), func() {
					mc(valTestBoolPtrSliceConsumer)
				})
			} else {
				mc(valTestBoolPtrSliceConsumer)
			}
		})
	}
}

func TestSilentBoolPtrSliceConsumer(t *testing.T) {
	var sc SilentBoolPtrSliceConsumer = func(v []*bool) {
		require.Equal(t, valTestBoolPtrSliceConsumer, v)
		return
	}
	sc(valTestBoolPtrSliceConsumer)
}

func TestSilentBoolPtrSliceConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolPtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					return errTestBoolPtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentBoolPtrSliceConsumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestBoolPtrSliceConsumer)
		})
	}
}

func TestSilentBoolPtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testBoolPtrSliceConsumerFactory
		cf2   testBoolPtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestBoolPtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolPtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentBoolPtrSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentBoolPtrSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentBoolPtrSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestBoolPtrSliceConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustBoolPtrSliceConsumer(t *testing.T) {
	var mc MustBoolPtrSliceConsumer = func(v []*bool) {
		require.Equal(t, valTestBoolPtrSliceConsumer, v)
		return
	}
	mc(valTestBoolPtrSliceConsumer)
}

func TestMustBoolPtrSliceConsumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testBoolPtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					return errTestBoolPtrSliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustBoolPtrSliceConsumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.wantErr {
				r.PanicsWithError(errTestBoolPtrSliceConsumer.Error(), func() {
					mc(valTestBoolPtrSliceConsumer)
				})
			} else {
				mc(valTestBoolPtrSliceConsumer)
			}
		})
	}
}

func TestMustBoolPtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testBoolPtrSliceConsumerFactory
		cf2     testBoolPtrSliceConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestBoolPtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					calls++
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolPtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustBoolPtrSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustBoolPtrSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustBoolPtrSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.wantErr {
				r.PanicsWithError(errTestBoolPtrSliceConsumer.Error(), func() {
					cmc(valTestBoolPtrSliceConsumer)
				})
			} else {
				cmc(valTestBoolPtrSliceConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustBoolPtrSliceConsumer_ToSilentBoolPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolPtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					return errTestBoolPtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustBoolPtrSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentBoolPtrSliceConsumer()
			r.NotNil(sc)

			sc(valTestBoolPtrSliceConsumer)
		})
	}
}

func TestMustBoolPtrSliceConsumer_ToBoolPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testBoolPtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) BoolPtrSliceConsumer {
				return func(v []*bool) error {
					require.Equal(t, valTestBoolPtrSliceConsumer, v)
					return errTestBoolPtrSliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustBoolPtrSliceConsumer()
			r.NotNil(mc)

			c = mc.ToBoolPtrSliceConsumer()
			r.NotNil(c)

			err := c(valTestBoolPtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestBoolPtrSliceConsumer.Error())
			}
		})
	}
}
