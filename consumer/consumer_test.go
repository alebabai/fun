package consumer

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testValue = "value"
	testError = errors.New("error")
)

func TestConsumer(t *testing.T) {
	var c Consumer = func(v interface{}) error {
		assert.Equal(t, testValue, v)
		return nil
	}

	err := c(testValue)
	assert.NoError(t, err)
}

func TestConsumer_AndThen(t *testing.T) {
	var calls int
	var c1 Consumer = func(v interface{}) error {
		calls++
		assert.Equal(t, testValue, v)
		assert.Equal(t, calls, 1, "should be called first and only once")
		return nil
	}
	var c2 Consumer = func(v interface{}) error {
		calls++
		assert.Equal(t, testValue, v)
		assert.Equal(t, calls, 2, "should be called second and only once")
		return nil
	}
	cc := c1.AndThen(c2)

	err := cc(testValue)
	assert.NoError(t, err)
	assert.Equal(t, 2, calls)
}

func TestConsumer_ToSilentConsumer(t *testing.T) {
	var c Consumer = func(v interface{}) error {
		assert.Equal(t, testValue, v)
		return nil
	}

	sc := c.ToSilentConsumer()
	assert.NotNil(t, sc)

	sc(testValue)
}

func TestConsumer_ToMustConsumer(t *testing.T) {
	var c Consumer = func(v interface{}) error {
		assert.Equal(t, testValue, v)
		return nil
	}

	mc := c.ToMustConsumer()
	assert.NotNil(t, mc)

	mc(testValue)
}

func TestConsumerWithError(t *testing.T) {
	var c Consumer = func(v interface{}) error {
		assert.Equal(t, testValue, v)
		return testError
	}

	err := c(testValue)
	assert.Error(t, err, testError)
}

func TestConsumer_AndThenWithError(t *testing.T) {
	var calls int
	// c1 returns an error
	var c1 Consumer = func(v interface{}) error {
		calls++
		assert.Equal(t, testValue, v)
		assert.Equal(t, calls, 1, "should be called first and only once")
		return testError
	}
	// c2 just basic consumer
	var c2 Consumer = func(v interface{}) error {
		calls++
		assert.Equal(t, testValue, v)
		assert.Equal(t, calls, 2, "should be called second and only once")
		return nil
	}
	cc := c1.AndThen(c2)

	err := cc(testValue)
	assert.Error(t, err, testError)
	assert.Equal(t, 1, calls)
}

func TestConsumer_ToSilentConsumerWithError(t *testing.T) {
	var c Consumer = func(v interface{}) error {
		assert.Equal(t, testValue, v)
		return testError
	}

	sc := c.ToSilentConsumer()
	assert.NotNil(t, sc)

	sc(testValue)
}

func TestConsumer_ToMustConsumerWithError(t *testing.T) {
	var c Consumer = func(v interface{}) error {
		assert.Equal(t, testValue, v)
		return testError
	}

	mc := c.ToMustConsumer()
	assert.NotNil(t, mc)

	assert.PanicsWithError(t, testError.Error(), func() {
		mc(testValue)
	})
}

func TestSilentConsumer(t *testing.T) {
	var sc SilentConsumer = func(v interface{}) {
		assert.Equal(t, testValue, v)
		return
	}
	sc(testValue)
}

func TestMustConsumer(t *testing.T) {
	var sc SilentConsumer = func(v interface{}) {
		assert.Equal(t, testValue, v)
		return
	}
	sc(testValue)
}

func TestMustConsumer_ToSilentConsumer(t *testing.T) {
	var c Consumer = func(v interface{}) error {
		assert.Equal(t, testValue, v)
		return testError
	}

	mc := c.ToMustConsumer()
	assert.NotNil(t, mc)

	sc := mc.ToSilentConsumer()
	assert.NotNil(t, sc)

	sc(testValue)
}

func TestMustConsumer_ToConsumer(t *testing.T) {
	var c Consumer = func(v interface{}) error {
		assert.Equal(t, testValue, v)
		return testError
	}

	mc := c.ToMustConsumer()
	assert.NotNil(t, mc)

	c = mc.ToConsumer()
	assert.NotNil(t, c)

	err := c(testValue)
	assert.Error(t, err, testError)
}