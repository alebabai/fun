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
	testRunePtrSliceConsumerValue []*rune
	testRunePtrSliceConsumerError = errors.New("error")
)

type testRunePtrSliceConsumerFactory func(t *testing.T) RunePtrSliceConsumer

func TestRunePtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testRunePtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					require.Equal(t, testRunePtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					require.Equal(t, testRunePtrSliceConsumerValue, v)
					return testRunePtrSliceConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testRunePtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testRunePtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestRunePtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testRunePtrSliceConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					require.Equal(t, testRunePtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					require.Equal(t, testRunePtrSliceConsumerValue, v)
					return testRunePtrSliceConsumerError
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

			err := c(testRunePtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testRunePtrSliceConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestRunePtrSliceConsumer_AndThen(t *testing.T) {
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
					require.Equal(t, testRunePtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					calls++
					require.Equal(t, testRunePtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					calls++
					require.Equal(t, testRunePtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testRunePtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					calls++
					require.Equal(t, testRunePtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					calls++
					require.Equal(t, testRunePtrSliceConsumerValue, v)
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
			err := cc(testRunePtrSliceConsumerValue)
			if err != nil {
				r.EqualError(err, testRunePtrSliceConsumerError.Error())
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
					require.Equal(t, testRunePtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					require.Equal(t, testRunePtrSliceConsumerValue, v)
					return testRunePtrSliceConsumerError
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

			sc(testRunePtrSliceConsumerValue)
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
					require.Equal(t, testRunePtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					require.Equal(t, testRunePtrSliceConsumerValue, v)
					return testRunePtrSliceConsumerError
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
				r.PanicsWithError(testRunePtrSliceConsumerError.Error(), func() {
					mc(testRunePtrSliceConsumerValue)
				})
			} else {
				mc(testRunePtrSliceConsumerValue)
			}
		})
	}
}

func TestSilentRunePtrSliceConsumer(t *testing.T) {
	var sc SilentRunePtrSliceConsumer = func(v []*rune) {
		require.Equal(t, testRunePtrSliceConsumerValue, v)
		return
	}
	sc(testRunePtrSliceConsumerValue)
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
					require.Equal(t, testRunePtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					calls++
					require.Equal(t, testRunePtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					calls++
					require.Equal(t, testRunePtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testRunePtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					calls++
					require.Equal(t, testRunePtrSliceConsumerValue, v)
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
					require.Equal(t, testRunePtrSliceConsumerValue, v)
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
			csc(testRunePtrSliceConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustRunePtrSliceConsumer(t *testing.T) {
	var sc SilentRunePtrSliceConsumer = func(v []*rune) {
		require.Equal(t, testRunePtrSliceConsumerValue, v)
		return
	}
	sc(testRunePtrSliceConsumerValue)
}

func TestMustRunePtrSliceConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testRunePtrSliceConsumerFactory
		cf2   testRunePtrSliceConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					calls++
					require.Equal(t, testRunePtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					calls++
					require.Equal(t, testRunePtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					calls++
					require.Equal(t, testRunePtrSliceConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testRunePtrSliceConsumerError
				}
			},
			cf2: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					calls++
					require.Equal(t, testRunePtrSliceConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					calls++
					require.Equal(t, testRunePtrSliceConsumerValue, v)
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
			if tt.err {
				r.PanicsWithError(testRunePtrSliceConsumerError.Error(), func() {
					cmc(testRunePtrSliceConsumerValue)
				})
			} else {
				cmc(testRunePtrSliceConsumerValue)
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
					require.Equal(t, testRunePtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					require.Equal(t, testRunePtrSliceConsumerValue, v)
					return testRunePtrSliceConsumerError
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

			sc(testRunePtrSliceConsumerValue)
		})
	}
}

func TestMustRunePtrSliceConsumer_ToRunePtrSliceConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testRunePtrSliceConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					require.Equal(t, testRunePtrSliceConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) RunePtrSliceConsumer {
				return func(v []*rune) error {
					require.Equal(t, testRunePtrSliceConsumerValue, v)
					return testRunePtrSliceConsumerError
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

			c = mc.ToRunePtrSliceConsumer()
			r.NotNil(c)

			err := c(testRunePtrSliceConsumerValue)
			if tt.err {
				r.EqualError(err, testRunePtrSliceConsumerError.Error())
			}
		})
	}
}
