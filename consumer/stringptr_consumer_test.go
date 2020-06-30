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
	valTestStringPtrConsumer *string
	errTestStringPtrConsumer = errors.New("error")
)

type testStringPtrConsumerFactory func(t *testing.T) StringPtrConsumer

func TestStringPtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringPtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringPtrConsumer {
				return func(v *string) error {
					require.Equal(t, valTestStringPtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) StringPtrConsumer {
				return func(v *string) error {
					require.Equal(t, valTestStringPtrConsumer, v)
					return errTestStringPtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(valTestStringPtrConsumer)
			if err != nil {
				r.EqualError(err, errTestStringPtrConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestStringPtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringPtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringPtrConsumer {
				return func(v *string) error {
					require.Equal(t, valTestStringPtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) StringPtrConsumer {
				return func(v *string) error {
					require.Equal(t, valTestStringPtrConsumer, v)
					return errTestStringPtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tc := tt.cf(t)
			c := tc.ToConsumer()
			r.NotNil(c)

			err := c(valTestStringPtrConsumer)
			if err != nil {
				r.EqualError(err, errTestStringPtrConsumer.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestStringPtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testStringPtrConsumerFactory
		cf2   testStringPtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) StringPtrConsumer {
				return func(v *string) error {
					calls++
					require.Equal(t, valTestStringPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringPtrConsumer {
				return func(v *string) error {
					calls++
					require.Equal(t, valTestStringPtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) StringPtrConsumer {
				return func(v *string) error {
					calls++
					require.Equal(t, valTestStringPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestStringPtrConsumer
				}
			},
			cf2: func(t *testing.T) StringPtrConsumer {
				return func(v *string) error {
					calls++
					require.Equal(t, valTestStringPtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) StringPtrConsumer {
				return func(v *string) error {
					calls++
					require.Equal(t, valTestStringPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringPtrConsumer {
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
			err := cc(valTestStringPtrConsumer)
			if err != nil {
				r.EqualError(err, errTestStringPtrConsumer.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestStringPtrConsumer_ToSilentStringPtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringPtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringPtrConsumer {
				return func(v *string) error {
					require.Equal(t, valTestStringPtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) StringPtrConsumer {
				return func(v *string) error {
					require.Equal(t, valTestStringPtrConsumer, v)
					return errTestStringPtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentStringPtrConsumer()
			r.NotNil(sc)

			sc(valTestStringPtrConsumer)
		})
	}
}

func TestStringPtrConsumer_ToMustStringPtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringPtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringPtrConsumer {
				return func(v *string) error {
					require.Equal(t, valTestStringPtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) StringPtrConsumer {
				return func(v *string) error {
					require.Equal(t, valTestStringPtrConsumer, v)
					return errTestStringPtrConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustStringPtrConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(errTestStringPtrConsumer.Error(), func() {
					mc(valTestStringPtrConsumer)
				})
			} else {
				mc(valTestStringPtrConsumer)
			}
		})
	}
}

func TestSilentStringPtrConsumer(t *testing.T) {
	var sc SilentStringPtrConsumer = func(v *string) {
		require.Equal(t, valTestStringPtrConsumer, v)
		return
	}
	sc(valTestStringPtrConsumer)
}

func TestSilentStringPtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testStringPtrConsumerFactory
		cf2   testStringPtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) StringPtrConsumer {
				return func(v *string) error {
					calls++
					require.Equal(t, valTestStringPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringPtrConsumer {
				return func(v *string) error {
					calls++
					require.Equal(t, valTestStringPtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) StringPtrConsumer {
				return func(v *string) error {
					calls++
					require.Equal(t, valTestStringPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestStringPtrConsumer
				}
			},
			cf2: func(t *testing.T) StringPtrConsumer {
				return func(v *string) error {
					calls++
					require.Equal(t, valTestStringPtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) StringPtrConsumer {
				return func(v *string) error {
					calls++
					require.Equal(t, valTestStringPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringPtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentStringPtrConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentStringPtrConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentStringPtrConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(valTestStringPtrConsumer)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustStringPtrConsumer(t *testing.T) {
	var sc SilentStringPtrConsumer = func(v *string) {
		require.Equal(t, valTestStringPtrConsumer, v)
		return
	}
	sc(valTestStringPtrConsumer)
}

func TestMustStringPtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testStringPtrConsumerFactory
		cf2   testStringPtrConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) StringPtrConsumer {
				return func(v *string) error {
					calls++
					require.Equal(t, valTestStringPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringPtrConsumer {
				return func(v *string) error {
					calls++
					require.Equal(t, valTestStringPtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) StringPtrConsumer {
				return func(v *string) error {
					calls++
					require.Equal(t, valTestStringPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return errTestStringPtrConsumer
				}
			},
			cf2: func(t *testing.T) StringPtrConsumer {
				return func(v *string) error {
					calls++
					require.Equal(t, valTestStringPtrConsumer, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) StringPtrConsumer {
				return func(v *string) error {
					calls++
					require.Equal(t, valTestStringPtrConsumer, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) StringPtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustStringPtrConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustStringPtrConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustStringPtrConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(errTestStringPtrConsumer.Error(), func() {
					cmc(valTestStringPtrConsumer)
				})
			} else {
				cmc(valTestStringPtrConsumer)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustStringPtrConsumer_ToSilentStringPtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringPtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringPtrConsumer {
				return func(v *string) error {
					require.Equal(t, valTestStringPtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) StringPtrConsumer {
				return func(v *string) error {
					require.Equal(t, valTestStringPtrConsumer, v)
					return errTestStringPtrConsumer
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustStringPtrConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentStringPtrConsumer()
			r.NotNil(sc)

			sc(valTestStringPtrConsumer)
		})
	}
}

func TestMustStringPtrConsumer_ToStringPtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testStringPtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) StringPtrConsumer {
				return func(v *string) error {
					require.Equal(t, valTestStringPtrConsumer, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) StringPtrConsumer {
				return func(v *string) error {
					require.Equal(t, valTestStringPtrConsumer, v)
					return errTestStringPtrConsumer
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustStringPtrConsumer()
			r.NotNil(mc)

			c = mc.ToStringPtrConsumer()
			r.NotNil(c)

			err := c(valTestStringPtrConsumer)
			if tt.err {
				r.EqualError(err, errTestStringPtrConsumer.Error())
			}
		})
	}
}
