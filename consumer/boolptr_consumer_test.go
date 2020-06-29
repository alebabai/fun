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
	testBoolPtrConsumerValue *bool
	testBoolPtrConsumerError = errors.New("error")
)

type testBoolPtrConsumerFactory func(t *testing.T) BoolPtrConsumer

func TestBoolPtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolPtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					require.Equal(t, testBoolPtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					require.Equal(t, testBoolPtrConsumerValue, v)
					return testBoolPtrConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testBoolPtrConsumerValue)
			if err != nil {
				r.EqualError(err, testBoolPtrConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestBoolPtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolPtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					require.Equal(t, testBoolPtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					require.Equal(t, testBoolPtrConsumerValue, v)
					return testBoolPtrConsumerError
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

			err := c(testBoolPtrConsumerValue)
			if err != nil {
				r.EqualError(err, testBoolPtrConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestBoolPtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testBoolPtrConsumerFactory
		cf2   testBoolPtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, testBoolPtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, testBoolPtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, testBoolPtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testBoolPtrConsumerError
				}
			},
			cf2: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, testBoolPtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, testBoolPtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolPtrConsumer {
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
			err := cc(testBoolPtrConsumerValue)
			if err != nil {
				r.EqualError(err, testBoolPtrConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestBoolPtrConsumer_ToSilentBoolPtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolPtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					require.Equal(t, testBoolPtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					require.Equal(t, testBoolPtrConsumerValue, v)
					return testBoolPtrConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentBoolPtrConsumer()
			r.NotNil(sc)

			sc(testBoolPtrConsumerValue)
		})
	}
}

func TestBoolPtrConsumer_ToMustBoolPtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolPtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					require.Equal(t, testBoolPtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					require.Equal(t, testBoolPtrConsumerValue, v)
					return testBoolPtrConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustBoolPtrConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testBoolPtrConsumerError.Error(), func() {
					mc(testBoolPtrConsumerValue)
				})
			} else {
				mc(testBoolPtrConsumerValue)
			}
		})
	}
}

func TestSilentBoolPtrConsumer(t *testing.T) {
	var sc SilentBoolPtrConsumer = func(v *bool) {
		require.Equal(t, testBoolPtrConsumerValue, v)
		return
	}
	sc(testBoolPtrConsumerValue)
}

func TestSilentBoolPtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testBoolPtrConsumerFactory
		cf2   testBoolPtrConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, testBoolPtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, testBoolPtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, testBoolPtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testBoolPtrConsumerError
				}
			},
			cf2: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, testBoolPtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, testBoolPtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolPtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentBoolPtrConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentBoolPtrConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentBoolPtrConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testBoolPtrConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustBoolPtrConsumer(t *testing.T) {
	var sc SilentBoolPtrConsumer = func(v *bool) {
		require.Equal(t, testBoolPtrConsumerValue, v)
		return
	}
	sc(testBoolPtrConsumerValue)
}

func TestMustBoolPtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testBoolPtrConsumerFactory
		cf2   testBoolPtrConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, testBoolPtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, testBoolPtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, testBoolPtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testBoolPtrConsumerError
				}
			},
			cf2: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, testBoolPtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					calls++
					require.Equal(t, testBoolPtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) BoolPtrConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustBoolPtrConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustBoolPtrConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustBoolPtrConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testBoolPtrConsumerError.Error(), func() {
					cmc(testBoolPtrConsumerValue)
				})
			} else {
				cmc(testBoolPtrConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustBoolPtrConsumer_ToSilentBoolPtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolPtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					require.Equal(t, testBoolPtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					require.Equal(t, testBoolPtrConsumerValue, v)
					return testBoolPtrConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustBoolPtrConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentBoolPtrConsumer()
			r.NotNil(sc)

			sc(testBoolPtrConsumerValue)
		})
	}
}

func TestMustBoolPtrConsumer_ToBoolPtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testBoolPtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					require.Equal(t, testBoolPtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) BoolPtrConsumer {
				return func(v *bool) error {
					require.Equal(t, testBoolPtrConsumerValue, v)
					return testBoolPtrConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustBoolPtrConsumer()
			r.NotNil(mc)

			c = mc.ToBoolPtrConsumer()
			r.NotNil(c)

			err := c(testBoolPtrConsumerValue)
			if tt.err {
				r.EqualError(err, testBoolPtrConsumerError.Error())
			}
		})
	}
}
