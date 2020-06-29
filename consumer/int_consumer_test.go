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
	testIntConsumerValue int
	testIntConsumerError = errors.New("error")
)

type testIntConsumerFactory func(t *testing.T) IntConsumer

func TestIntConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntConsumer {
				return func(v int) error {
					require.Equal(t, testIntConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntConsumer {
				return func(v int) error {
					require.Equal(t, testIntConsumerValue, v)
					return testIntConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testIntConsumerValue)
			if err != nil {
				r.EqualError(err, testIntConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestIntSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntConsumer {
				return func(v int) error {
					require.Equal(t, testIntConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntConsumer {
				return func(v int) error {
					require.Equal(t, testIntConsumerValue, v)
					return testIntConsumerError
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

			err := c(testIntConsumerValue)
			if err != nil {
				r.EqualError(err, testIntConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestIntConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testIntConsumerFactory
		cf2   testIntConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) IntConsumer {
				return func(v int) error {
					calls++
					require.Equal(t, testIntConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntConsumer {
				return func(v int) error {
					calls++
					require.Equal(t, testIntConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) IntConsumer {
				return func(v int) error {
					calls++
					require.Equal(t, testIntConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testIntConsumerError
				}
			},
			cf2: func(t *testing.T) IntConsumer {
				return func(v int) error {
					calls++
					require.Equal(t, testIntConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) IntConsumer {
				return func(v int) error {
					calls++
					require.Equal(t, testIntConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntConsumer {
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
			err := cc(testIntConsumerValue)
			if err != nil {
				r.EqualError(err, testIntConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestIntConsumer_ToSilentIntConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntConsumer {
				return func(v int) error {
					require.Equal(t, testIntConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntConsumer {
				return func(v int) error {
					require.Equal(t, testIntConsumerValue, v)
					return testIntConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentIntConsumer()
			r.NotNil(sc)

			sc(testIntConsumerValue)
		})
	}
}

func TestIntConsumer_ToMustIntConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntConsumer {
				return func(v int) error {
					require.Equal(t, testIntConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntConsumer {
				return func(v int) error {
					require.Equal(t, testIntConsumerValue, v)
					return testIntConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustIntConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testIntConsumerError.Error(), func() {
					mc(testIntConsumerValue)
				})
			} else {
				mc(testIntConsumerValue)
			}
		})
	}
}

func TestSilentIntConsumer(t *testing.T) {
	var sc SilentIntConsumer = func(v int) {
		require.Equal(t, testIntConsumerValue, v)
		return
	}
	sc(testIntConsumerValue)
}

func TestSilentIntConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testIntConsumerFactory
		cf2   testIntConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) IntConsumer {
				return func(v int) error {
					calls++
					require.Equal(t, testIntConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntConsumer {
				return func(v int) error {
					calls++
					require.Equal(t, testIntConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) IntConsumer {
				return func(v int) error {
					calls++
					require.Equal(t, testIntConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testIntConsumerError
				}
			},
			cf2: func(t *testing.T) IntConsumer {
				return func(v int) error {
					calls++
					require.Equal(t, testIntConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) IntConsumer {
				return func(v int) error {
					calls++
					require.Equal(t, testIntConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentIntConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentIntConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentIntConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testIntConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustIntConsumer(t *testing.T) {
	var sc SilentIntConsumer = func(v int) {
		require.Equal(t, testIntConsumerValue, v)
		return
	}
	sc(testIntConsumerValue)
}

func TestMustIntConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testIntConsumerFactory
		cf2   testIntConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) IntConsumer {
				return func(v int) error {
					calls++
					require.Equal(t, testIntConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntConsumer {
				return func(v int) error {
					calls++
					require.Equal(t, testIntConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) IntConsumer {
				return func(v int) error {
					calls++
					require.Equal(t, testIntConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testIntConsumerError
				}
			},
			cf2: func(t *testing.T) IntConsumer {
				return func(v int) error {
					calls++
					require.Equal(t, testIntConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) IntConsumer {
				return func(v int) error {
					calls++
					require.Equal(t, testIntConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) IntConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustIntConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustIntConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustIntConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testIntConsumerError.Error(), func() {
					cmc(testIntConsumerValue)
				})
			} else {
				cmc(testIntConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustIntConsumer_ToSilentIntConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntConsumer {
				return func(v int) error {
					require.Equal(t, testIntConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntConsumer {
				return func(v int) error {
					require.Equal(t, testIntConsumerValue, v)
					return testIntConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustIntConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentIntConsumer()
			r.NotNil(sc)

			sc(testIntConsumerValue)
		})
	}
}

func TestMustIntConsumer_ToIntConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testIntConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) IntConsumer {
				return func(v int) error {
					require.Equal(t, testIntConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) IntConsumer {
				return func(v int) error {
					require.Equal(t, testIntConsumerValue, v)
					return testIntConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustIntConsumer()
			r.NotNil(mc)

			c = mc.ToIntConsumer()
			r.NotNil(c)

			err := c(testIntConsumerValue)
			if tt.err {
				r.EqualError(err, testIntConsumerError.Error())
			}
		})
	}
}
