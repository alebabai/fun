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
	testRunePtrConsumerValue *rune
	testRunePtrConsumerError = errors.New("error")
)

type testRunePtrConsumerFactory func(t *testing.T) RunePtrConsumer

func TestRunePtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testRunePtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					require.Equal(t, testRunePtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					require.Equal(t, testRunePtrConsumerValue, v)
					return testRunePtrConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testRunePtrConsumerValue)
			if err != nil {
				r.EqualError(err, testRunePtrConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestRunePtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testRunePtrConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					require.Equal(t, testRunePtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					require.Equal(t, testRunePtrConsumerValue, v)
					return testRunePtrConsumerError
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

			err := c(testRunePtrConsumerValue)
			if err != nil {
				r.EqualError(err, testRunePtrConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestRunePtrConsumer_AndThen(t *testing.T) {
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
					require.Equal(t, testRunePtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					calls++
					require.Equal(t, testRunePtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					calls++
					require.Equal(t, testRunePtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testRunePtrConsumerError
				}
			},
			cf2: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					calls++
					require.Equal(t, testRunePtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					calls++
					require.Equal(t, testRunePtrConsumerValue, v)
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
			err := cc(testRunePtrConsumerValue)
			if err != nil {
				r.EqualError(err, testRunePtrConsumerError.Error())
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
					require.Equal(t, testRunePtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					require.Equal(t, testRunePtrConsumerValue, v)
					return testRunePtrConsumerError
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

			sc(testRunePtrConsumerValue)
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
					require.Equal(t, testRunePtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					require.Equal(t, testRunePtrConsumerValue, v)
					return testRunePtrConsumerError
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
				r.PanicsWithError(testRunePtrConsumerError.Error(), func() {
					mc(testRunePtrConsumerValue)
				})
			} else {
				mc(testRunePtrConsumerValue)
			}
		})
	}
}

func TestSilentRunePtrConsumer(t *testing.T) {
	var sc SilentRunePtrConsumer = func(v *rune) {
		require.Equal(t, testRunePtrConsumerValue, v)
		return
	}
	sc(testRunePtrConsumerValue)
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
					require.Equal(t, testRunePtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					calls++
					require.Equal(t, testRunePtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					calls++
					require.Equal(t, testRunePtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testRunePtrConsumerError
				}
			},
			cf2: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					calls++
					require.Equal(t, testRunePtrConsumerValue, v)
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
					require.Equal(t, testRunePtrConsumerValue, v)
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
			csc(testRunePtrConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustRunePtrConsumer(t *testing.T) {
	var sc SilentRunePtrConsumer = func(v *rune) {
		require.Equal(t, testRunePtrConsumerValue, v)
		return
	}
	sc(testRunePtrConsumerValue)
}

func TestMustRunePtrConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testRunePtrConsumerFactory
		cf2   testRunePtrConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					calls++
					require.Equal(t, testRunePtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					calls++
					require.Equal(t, testRunePtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					calls++
					require.Equal(t, testRunePtrConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testRunePtrConsumerError
				}
			},
			cf2: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					calls++
					require.Equal(t, testRunePtrConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					calls++
					require.Equal(t, testRunePtrConsumerValue, v)
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
			if tt.err {
				r.PanicsWithError(testRunePtrConsumerError.Error(), func() {
					cmc(testRunePtrConsumerValue)
				})
			} else {
				cmc(testRunePtrConsumerValue)
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
					require.Equal(t, testRunePtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					require.Equal(t, testRunePtrConsumerValue, v)
					return testRunePtrConsumerError
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

			sc(testRunePtrConsumerValue)
		})
	}
}

func TestMustRunePtrConsumer_ToRunePtrConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testRunePtrConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					require.Equal(t, testRunePtrConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) RunePtrConsumer {
				return func(v *rune) error {
					require.Equal(t, testRunePtrConsumerValue, v)
					return testRunePtrConsumerError
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

			c = mc.ToRunePtrConsumer()
			r.NotNil(c)

			err := c(testRunePtrConsumerValue)
			if tt.err {
				r.EqualError(err, testRunePtrConsumerError.Error())
			}
		})
	}
}
