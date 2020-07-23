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
	valTestStringPtrSliceConsumer []*string
	errTestStringPtrSliceConsumer = errors.New("error")
)

type testStringPtrSliceConsumerFactory func(t *testing.T) StringPtrSliceConsumer

func TestStringPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testStringPtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					return errTestStringPtrSliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestStringPtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestStringPtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestStringPtrSliceConsumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testStringPtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					return errTestStringPtrSliceConsumer
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

			err := c(valTestStringPtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestStringPtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestStringPtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testStringPtrSliceConsumerFactory
		cf2     testStringPtrSliceConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestStringPtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringPtrSliceConsumer {
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
			err := cc(valTestStringPtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestStringPtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestStringPtrSliceConsumer_ToSilentStringPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringPtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					return errTestStringPtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentStringPtrSliceConsumer()
			r.NotNil(sc)

			sc(valTestStringPtrSliceConsumer)
		})
	}
}

func TestStringPtrSliceConsumer_ToMustStringPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringPtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					return errTestStringPtrSliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustStringPtrSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestStringPtrSliceConsumer.Error(), func() {
					mc(valTestStringPtrSliceConsumer)
				})
			} else {
				mc(valTestStringPtrSliceConsumer)
			}
		})
	}
}

func TestSilentStringPtrSliceConsumer(t *testing.T) {
	var sc SilentStringPtrSliceConsumer = func(v []*string) {
		require.Equal(t, valTestStringPtrSliceConsumer, v)
		return
	}
	sc(valTestStringPtrSliceConsumer)
}

func TestSilentStringPtrSliceConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringPtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					return errTestStringPtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentStringPtrSliceConsumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestStringPtrSliceConsumer)
		})
	}
}

func TestSilentStringPtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testStringPtrSliceConsumerFactory
		cf2   testStringPtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestStringPtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringPtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentStringPtrSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentStringPtrSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentStringPtrSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestStringPtrSliceConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustStringPtrSliceConsumer(t *testing.T) {
	var mc MustStringPtrSliceConsumer = func(v []*string) {
		require.Equal(t, valTestStringPtrSliceConsumer, v)
		return
	}
	mc(valTestStringPtrSliceConsumer)
}

func TestMustStringPtrSliceConsumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testStringPtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					return errTestStringPtrSliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustStringPtrSliceConsumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.wantErr {
				r.PanicsWithError(errTestStringPtrSliceConsumer.Error(), func() {
					mc(valTestStringPtrSliceConsumer)
				})
			} else {
				mc(valTestStringPtrSliceConsumer)
			}
		})
	}
}

func TestMustStringPtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testStringPtrSliceConsumerFactory
		cf2     testStringPtrSliceConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestStringPtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					calls++
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringPtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustStringPtrSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustStringPtrSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustStringPtrSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.wantErr {
				r.PanicsWithError(errTestStringPtrSliceConsumer.Error(), func() {
					cmc(valTestStringPtrSliceConsumer)
				})
			} else {
				cmc(valTestStringPtrSliceConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustStringPtrSliceConsumer_ToSilentStringPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringPtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					return errTestStringPtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustStringPtrSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentStringPtrSliceConsumer()
			r.NotNil(sc)

			sc(valTestStringPtrSliceConsumer)
		})
	}
}

func TestMustStringPtrSliceConsumer_ToStringPtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testStringPtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) StringPtrSliceConsumer {
				return func(v []*string) error {
					require.Equal(t, valTestStringPtrSliceConsumer, v)
					return errTestStringPtrSliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustStringPtrSliceConsumer()
			r.NotNil(mc)

			c = mc.ToStringPtrSliceConsumer()
			r.NotNil(c)

			err := c(valTestStringPtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestStringPtrSliceConsumer.Error())
			}
		})
	}
}
