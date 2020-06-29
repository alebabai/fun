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
	testUintConsumerValue uint
	testUintConsumerError = errors.New("error")
)

type testUintConsumerFactory func(t *testing.T) UintConsumer

func TestUintConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					require.Equal(t, testUintConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					require.Equal(t, testUintConsumerValue, v)
					return testUintConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testUintConsumerValue)
			if err != nil {
				r.EqualError(err, testUintConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUintSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					require.Equal(t, testUintConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					require.Equal(t, testUintConsumerValue, v)
					return testUintConsumerError
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

			err := c(testUintConsumerValue)
			if err != nil {
				r.EqualError(err, testUintConsumerError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUintConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUintConsumerFactory
		cf2   testUintConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, testUintConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, testUintConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, testUintConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUintConsumerError
				}
			},
			cf2: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, testUintConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, testUintConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintConsumer {
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
			err := cc(testUintConsumerValue)
			if err != nil {
				r.EqualError(err, testUintConsumerError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestUintConsumer_ToSilentUintConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					require.Equal(t, testUintConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					require.Equal(t, testUintConsumerValue, v)
					return testUintConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentUintConsumer()
			r.NotNil(sc)

			sc(testUintConsumerValue)
		})
	}
}

func TestUintConsumer_ToMustUintConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					require.Equal(t, testUintConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					require.Equal(t, testUintConsumerValue, v)
					return testUintConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustUintConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testUintConsumerError.Error(), func() {
					mc(testUintConsumerValue)
				})
			} else {
				mc(testUintConsumerValue)
			}
		})
	}
}

func TestSilentUintConsumer(t *testing.T) {
	var sc SilentUintConsumer = func(v uint) {
		require.Equal(t, testUintConsumerValue, v)
		return
	}
	sc(testUintConsumerValue)
}

func TestSilentUintConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUintConsumerFactory
		cf2   testUintConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, testUintConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, testUintConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, testUintConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUintConsumerError
				}
			},
			cf2: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, testUintConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, testUintConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentUintConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			var sc2 SilentUintConsumer = nil
			if c2 != nil {
				sc2 = c2.ToSilentUintConsumer()
				r.NotNil(sc2)
			}

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			calls = 0
			csc(testUintConsumerValue)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUintConsumer(t *testing.T) {
	var sc SilentUintConsumer = func(v uint) {
		require.Equal(t, testUintConsumerValue, v)
		return
	}
	sc(testUintConsumerValue)
}

func TestMustUintConsumer_AndThen(t *testing.T) {
	var calls int
	tests := []struct {
		name  string
		cf1   testUintConsumerFactory
		cf2   testUintConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, testUintConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, testUintConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, testUintConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return testUintConsumerError
				}
			},
			cf2: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, testUintConsumerValue, v)
					require.Equal(t, calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
		{
			name: "nil after",
			cf1: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					calls++
					require.Equal(t, testUintConsumerValue, v)
					require.Equal(t, calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) UintConsumer {
				return nil
			},
			calls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustUintConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			var mc2 MustUintConsumer = nil
			if c2 != nil {
				mc2 = c2.ToMustUintConsumer()
				r.NotNil(mc2)
			}

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			calls = 0
			if tt.err {
				r.PanicsWithError(testUintConsumerError.Error(), func() {
					cmc(testUintConsumerValue)
				})
			} else {
				cmc(testUintConsumerValue)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustUintConsumer_ToSilentUintConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					require.Equal(t, testUintConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					require.Equal(t, testUintConsumerValue, v)
					return testUintConsumerError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUintConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentUintConsumer()
			r.NotNil(sc)

			sc(testUintConsumerValue)
		})
	}
}

func TestMustUintConsumer_ToUintConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   testUintConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					require.Equal(t, testUintConsumerValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) UintConsumer {
				return func(v uint) error {
					require.Equal(t, testUintConsumerValue, v)
					return testUintConsumerError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustUintConsumer()
			r.NotNil(mc)

			c = mc.ToUintConsumer()
			r.NotNil(c)

			err := c(testUintConsumerValue)
			if tt.err {
				r.EqualError(err, testUintConsumerError.Error())
			}
		})
	}
}
