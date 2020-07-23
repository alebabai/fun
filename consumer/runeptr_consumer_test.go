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
	valTestRunePtrConsumer *rune
	errTestRunePtrConsumer = errors.New("error")
)

type testRunePtrConsumerFactory func(t *testing.T) RunePtrConsumer

func TestRunePtrConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testRunePtrConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					require.Equal(t, valTestRunePtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					require.Equal(t, valTestRunePtrConsumer, v)
					return errTestRunePtrConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestRunePtrConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestRunePtrConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestRunePtrConsumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testRunePtrConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					require.Equal(t, valTestRunePtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					require.Equal(t, valTestRunePtrConsumer, v)
					return errTestRunePtrConsumer
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

			err := c(valTestRunePtrConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestRunePtrConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestRunePtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testRunePtrConsumerFactory
		cf2     testRunePtrConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					calls++
					require.Equal(t, valTestRunePtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					calls++
					require.Equal(t, valTestRunePtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					calls++
					require.Equal(t, valTestRunePtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestRunePtrConsumer
				}
			},
			cf2: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					calls++
					require.Equal(t, valTestRunePtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					calls++
					require.Equal(t, valTestRunePtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RunePtrConsumer {
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
			err := cc(valTestRunePtrConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestRunePtrConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestRunePtrConsumer_ToSilentRunePtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testRunePtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					require.Equal(t, valTestRunePtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					require.Equal(t, valTestRunePtrConsumer, v)
					return errTestRunePtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentRunePtrConsumer()
			r.NotNil(sc)

			sc(valTestRunePtrConsumer)
		})
	}
}

func TestRunePtrConsumer_ToMustRunePtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testRunePtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					require.Equal(t, valTestRunePtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					require.Equal(t, valTestRunePtrConsumer, v)
					return errTestRunePtrConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustRunePtrConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestRunePtrConsumer.Error(), func() {
					mc(valTestRunePtrConsumer)
				})
			} else {
				mc(valTestRunePtrConsumer)
			}
		})
	}
}

func TestSilentRunePtrConsumer(t *testing.T) {
	var sc SilentRunePtrConsumer = func(v *rune) {
		require.Equal(t, valTestRunePtrConsumer, v)
		return
	}
	sc(valTestRunePtrConsumer)
}

func TestSilentRunePtrConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testRunePtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					require.Equal(t, valTestRunePtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					require.Equal(t, valTestRunePtrConsumer, v)
					return errTestRunePtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tsc := tc.ToSilentRunePtrConsumer()
			r.NotNil(tsc)

			sc := tsc.ToSilentConsumer()
			r.NotNil(sc)

			sc(valTestRunePtrConsumer)
		})
	}
}

func TestSilentRunePtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testRunePtrConsumerFactory
		cf2   testRunePtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					calls++
					require.Equal(t, valTestRunePtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					calls++
					require.Equal(t, valTestRunePtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					calls++
					require.Equal(t, valTestRunePtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestRunePtrConsumer
				}
			},
			cf2: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					calls++
					require.Equal(t, valTestRunePtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					calls++
					require.Equal(t, valTestRunePtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RunePtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentRunePtrConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentRunePtrConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentRunePtrConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestRunePtrConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustRunePtrConsumer(t *testing.T) {
	var mc MustRunePtrConsumer = func(v *rune) {
		require.Equal(t, valTestRunePtrConsumer, v)
		return
	}
	mc(valTestRunePtrConsumer)
}

func TestMustRunePtrConsumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testRunePtrConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					require.Equal(t, valTestRunePtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					require.Equal(t, valTestRunePtrConsumer, v)
					return errTestRunePtrConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			tmc := tc.ToMustRunePtrConsumer()
			r.NotNil(tmc)

			mc := tmc.ToMustConsumer()
			r.NotNil(mc)

			if tt.wantErr {
				r.PanicsWithError(errTestRunePtrConsumer.Error(), func() {
					mc(valTestRunePtrConsumer)
				})
			} else {
				mc(valTestRunePtrConsumer)
			}
		})
	}
}

func TestMustRunePtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name    string
		cf1     testRunePtrConsumerFactory
		cf2     testRunePtrConsumerFactory
		calls   int
		wantErr bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					calls++
					require.Equal(t, valTestRunePtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					calls++
					require.Equal(t, valTestRunePtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with error",
			cf1: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					calls++
					require.Equal(t, valTestRunePtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestRunePtrConsumer
				}
			},
			cf2: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					calls++
					require.Equal(t, valTestRunePtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls:   1,
			wantErr: true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					calls++
					require.Equal(t, valTestRunePtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RunePtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustRunePtrConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustRunePtrConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustRunePtrConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.wantErr {
				r.PanicsWithError(errTestRunePtrConsumer.Error(), func() {
					cmc(valTestRunePtrConsumer)
				})
			} else {
				cmc(valTestRunePtrConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustRunePtrConsumer_ToSilentRunePtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testRunePtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					require.Equal(t, valTestRunePtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					require.Equal(t, valTestRunePtrConsumer, v)
					return errTestRunePtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustRunePtrConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentRunePtrConsumer()
			r.NotNil(sc)

			sc(valTestRunePtrConsumer)
		})
	}
}

func TestMustRunePtrConsumer_ToRunePtrConsumer(t *testing.T) {
	tests := []struct {
		name    string
		cf      testRunePtrConsumerFactory
		wantErr bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					require.Equal(t, valTestRunePtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with error",
			cf: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					require.Equal(t, valTestRunePtrConsumer, v)
					return errTestRunePtrConsumer
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustRunePtrConsumer()
			r.NotNil(mc)

			c = mc.ToRunePtrConsumer()
			r.NotNil(c)

			err := c(valTestRunePtrConsumer)
			if tt.wantErr {
				r.EqualError(err, errTestRunePtrConsumer.Error())
			}
		})
	}
}
