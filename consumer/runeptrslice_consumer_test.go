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
	valTestRunePtrSliceConsumer []*rune
	errTestRunePtrSliceConsumer = errors.New("error")
)

type testRunePtrSliceConsumerFactory func(t *testing.T) RunePtrSliceConsumer

func TestRunePtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testRunePtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					return errTestRunePtrSliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestRunePtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestRunePtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestRunePtrSliceConsumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testRunePtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					return errTestRunePtrSliceConsumer
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

			err := c(valTestRunePtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestRunePtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestRunePtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testRunePtrSliceConsumerFactory
		cf2     testRunePtrSliceConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					calls++
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					calls++
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					calls++
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestRunePtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					calls++
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					calls++
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RunePtrSliceConsumer {
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
			err := cc(valTestRunePtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestRunePtrSliceConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestRunePtrSliceConsumer_ToSilentRunePtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testRunePtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					return errTestRunePtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentRunePtrSliceConsumer()
			r.NotNil(sc)

			sc(valTestRunePtrSliceConsumer)
		})
	}
}

func TestRunePtrSliceConsumer_ToMustRunePtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testRunePtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					return errTestRunePtrSliceConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustRunePtrSliceConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestRunePtrSliceConsumer.Error(), func() {
					mc(valTestRunePtrSliceConsumer)
				})
			} else {
				mc(valTestRunePtrSliceConsumer)
			}
		})
	}
}

func TestSilentRunePtrSliceConsumer(t *testing.T) {
	var sc SilentRunePtrSliceConsumer = func(v []*rune) {
		require.Equal(t, valTestRunePtrSliceConsumer, v)
		return
	}
	sc(valTestRunePtrSliceConsumer)
}

func TestSilentRunePtrSliceConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testRunePtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					return errTestRunePtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentRunePtrSliceConsumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestRunePtrSliceConsumer)
		})
	}
}

func TestSilentRunePtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testRunePtrSliceConsumerFactory
		cf2   testRunePtrSliceConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					calls++
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					calls++
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					calls++
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestRunePtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					calls++
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					calls++
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RunePtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentRunePtrSliceConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentRunePtrSliceConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentRunePtrSliceConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestRunePtrSliceConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustRunePtrSliceConsumer(t *testing.T) {
	var mc MustRunePtrSliceConsumer = func(v []*rune) {
		require.Equal(t, valTestRunePtrSliceConsumer, v)
		return
	}
	mc(valTestRunePtrSliceConsumer)
}

func TestMustRunePtrSliceConsumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testRunePtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					return errTestRunePtrSliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustRunePtrSliceConsumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.wantErr {
				r.PanicsWithError(errTestRunePtrSliceConsumer.Error(), func() {
					mc(valTestRunePtrSliceConsumer)
				})
			} else {
				mc(valTestRunePtrSliceConsumer)
			}
		})
	}
}

func TestMustRunePtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testRunePtrSliceConsumerFactory
		cf2     testRunePtrSliceConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					calls++
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					calls++
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					calls++
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestRunePtrSliceConsumer
				}
			},
			cf2: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					calls++
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					calls++
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RunePtrSliceConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustRunePtrSliceConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustRunePtrSliceConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustRunePtrSliceConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.wantErr {
				r.PanicsWithError(errTestRunePtrSliceConsumer.Error(), func() {
					cmc(valTestRunePtrSliceConsumer)
				})
			} else {
				cmc(valTestRunePtrSliceConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustRunePtrSliceConsumer_ToSilentRunePtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testRunePtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					return errTestRunePtrSliceConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustRunePtrSliceConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentRunePtrSliceConsumer()
			r.NotNil(sc)

			sc(valTestRunePtrSliceConsumer)
		})
	}
}

func TestMustRunePtrSliceConsumer_ToRunePtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testRunePtrSliceConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					require.Equal(t, valTestRunePtrSliceConsumer, v)
					return errTestRunePtrSliceConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustRunePtrSliceConsumer()
			r.NotNil(mc)

			c = mc.ToRunePtrSliceConsumer()
			r.NotNil(c)

			err := c(valTestRunePtrSliceConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestRunePtrSliceConsumer.Error())
			}
		})
	}
}
