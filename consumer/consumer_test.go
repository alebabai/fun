package consumer

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	testValue = "value"
	testError = errors.New("error")
)

type ConsumerFactory func(t *testing.T) Consumer

func TestConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					require.Equal(t, testValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					require.Equal(t, testValue, v)
					return testError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			err := c(testValue)
			if err != nil {
				r.EqualError(err, testError.Error())
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestConsumer_AndThen(t *testing.T) {
	tests := []struct {
		name  string
		cf1   ConsumerFactory
		cf2   ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					calls := v.(*int)
					*calls++
					require.Equal(t, *calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					calls := v.(*int)
					*calls++
					require.Equal(t, *calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					calls := v.(*int)
					*calls++
					require.Equal(t, *calls, 1, "should be called first and only once")
					return testError
				}
			},
			cf2: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					calls := v.(*int)
					*calls++
					require.Equal(t, *calls, 2, "should be called second and only once")
					return nil
				}
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

			var calls int
			err := cc(&calls)
			if err != nil {
				r.EqualError(err, testError.Error())
			} else {
				r.NoError(err)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					require.Equal(t, testValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					require.Equal(t, testValue, v)
					return testError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			sc := c.ToSilentConsumer()
			r.NotNil(sc)

			sc(testValue)
		})
	}
}

func TestConsumer_ToMustConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					require.Equal(t, testValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					require.Equal(t, testValue, v)
					return testError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)
			mc := c.ToMustConsumer()
			r.NotNil(mc)

			if tt.err {
				r.PanicsWithError(testError.Error(), func() {
					mc(testValue)
				})
			} else {
				mc(testValue)
			}
		})
	}
}

func TestSilentConsumer(t *testing.T) {
	var sc SilentConsumer = func(v interface{}) {
		require.Equal(t, testValue, v)
		return
	}
	sc(testValue)
}

func TestSilentConsumer_AndThen(t *testing.T) {
	tests := []struct {
		name  string
		cf1   ConsumerFactory
		cf2   ConsumerFactory
		calls int
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					calls := v.(*int)
					*calls++
					require.Equal(t, *calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					calls := v.(*int)
					*calls++
					require.Equal(t, *calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					calls := v.(*int)
					*calls++
					require.Equal(t, *calls, 1, "should be called first and only once")
					return testError
				}
			},
			cf2: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					calls := v.(*int)
					*calls++
					require.Equal(t, *calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			sc1 := c1.ToSilentConsumer()
			r.NotNil(sc1)

			c2 := tt.cf2(t)
			sc2 := c2.ToSilentConsumer()
			r.NotNil(sc2)

			csc := sc1.AndThen(sc2)
			r.NotNil(csc)

			var calls int
			csc(&calls)
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustConsumer(t *testing.T) {
	var sc SilentConsumer = func(v interface{}) {
		require.Equal(t, testValue, v)
		return
	}
	sc(testValue)
}

func TestMustConsumer_AndThen(t *testing.T) {
	tests := []struct {
		name  string
		cf1   ConsumerFactory
		cf2   ConsumerFactory
		calls int
		err   bool
	}{
		{
			name: "ok",
			cf1: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					calls := v.(*int)
					*calls++
					require.Equal(t, *calls, 1, "should be called first and only once")
					return nil
				}
			},
			cf2: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					calls := v.(*int)
					*calls++
					require.Equal(t, *calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 2,
		},
		{
			name: "with_error",
			cf1: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					calls := v.(*int)
					*calls++
					require.Equal(t, *calls, 1, "should be called first and only once")
					return testError
				}
			},
			cf2: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					calls := v.(*int)
					*calls++
					require.Equal(t, *calls, 2, "should be called second and only once")
					return nil
				}
			},
			calls: 1,
			err:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c1 := tt.cf1(t)
			mc1 := c1.ToMustConsumer()
			r.NotNil(mc1)

			c2 := tt.cf2(t)
			mc2 := c2.ToMustConsumer()
			r.NotNil(mc2)

			cmc := mc1.AndThen(mc2)
			r.NotNil(cmc)

			var calls int
			if tt.err {
				r.PanicsWithError(testError.Error(), func() {
					cmc(&calls)
				})
			} else {
				cmc(&calls)
			}
			r.Equal(tt.calls, calls)
		})
	}
}

func TestMustConsumer_ToSilentConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   ConsumerFactory
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					require.Equal(t, testValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					require.Equal(t, testValue, v)
					return testError
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustConsumer()
			r.NotNil(mc)

			sc := mc.ToSilentConsumer()
			r.NotNil(sc)

			sc(testValue)
		})
	}
}

func TestMustConsumer_ToConsumer(t *testing.T) {
	tests := []struct {
		name string
		cf   ConsumerFactory
		err  bool
	}{
		{
			name: "ok",
			cf: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					require.Equal(t, testValue, v)
					return nil
				}
			},
		},
		{
			name: "with_error",
			cf: func(t *testing.T) Consumer {
				return func(v interface{}) error {
					require.Equal(t, testValue, v)
					return testError
				}
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			c := tt.cf(t)

			mc := c.ToMustConsumer()
			r.NotNil(mc)

			c = mc.ToConsumer()
			r.NotNil(c)

			err := c(testValue)
			if tt.err {
				r.EqualError(err, testError.Error())
			}
		})
	}
}
