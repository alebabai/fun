package supplier

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testResult = "result"
	testError  = errors.New("error")
)

func TestSupplier(t *testing.T) {
	var s Supplier = func() (interface{}, error) {
		return testResult, nil
	}

	v, err := s()
	assert.NoError(t, err)
	assert.Equal(t, testResult, v)
}

func TestSupplier_ToSilentSupplier(t *testing.T) {
	var s Supplier = func() (interface{}, error) {
		return testResult, nil
	}

	ms := s.ToSilentSupplier()
	assert.NotNil(t, ms)

	v := ms()
	assert.Equal(t, testResult, v)
}

func TestSupplier_ToMustSupplier(t *testing.T) {
	var s Supplier = func() (interface{}, error) {
		return testResult, nil
	}

	ms := s.ToMustSupplier()
	assert.NotNil(t, ms)

	v := ms()
	assert.Equal(t, testResult, v)
}

func TestSupplierWithError(t *testing.T) {
	var s Supplier = func() (interface{}, error) {
		return nil, testError
	}

	v, err := s()
	assert.Errorf(t, err, testError.Error())
	assert.Nil(t, v)
}

func TestSupplier_ToSilentSupplierWithError(t *testing.T) {
	var s Supplier = func() (interface{}, error) {
		return nil, testError
	}

	ms := s.ToSilentSupplier()
	assert.NotNil(t, ms)

	v := ms()
	assert.Nil(t, v)
}

func TestSupplier_ToMustSupplierWithError(t *testing.T) {
	var s Supplier = func() (interface{}, error) {
		return nil, testError
	}

	ms := s.ToMustSupplier()
	assert.NotNil(t, ms)

	assert.PanicsWithError(t, testError.Error(), func() {
		ms()
	})
}

func TestSilentSupplier(t *testing.T) {
	var s SilentSupplier = func() interface{} {
		return testResult
	}
	v := s()
	assert.Equal(t, testResult, v)
}

func TestMustSupplier(t *testing.T) {
	var s SilentSupplier = func() interface{} {
		return testResult
	}

	v := s()
	assert.Equal(t, testResult, v)
}

func TestMustSupplier_ToSilentSupplier(t *testing.T) {
	var s Supplier = func() (interface{}, error) {
		return nil, testError
	}

	ms := s.ToMustSupplier()
	assert.NotNil(t, ms)

	ss := ms.ToSilentSupplier()
	assert.NotNil(t, ss)

	v := ss()
	assert.Nil(t, v)
}

func TestMustSupplier_ToSupplier(t *testing.T) {
	var s Supplier = func() (interface{}, error) {
		return nil, testError
	}

	ms := s.ToMustSupplier()
	assert.NotNil(t, ms)

	s = ms.ToSupplier()
	assert.NotNil(t, s)

	v, err := s()
	assert.Nil(t, v)
	assert.Errorf(t, err, testError.Error())
}
