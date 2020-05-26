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

func supplier() (interface{}, error) {
	return testResult, nil
}

func supplierWithError() (interface{}, error) {
	return nil, testError
}

func TestSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Supplier
	}{
		{
			name: "ok",
			s:    supplier,
		},
		{
			name: "with_error",
			s:    supplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := tt.s()
			if err != nil {
				assert.EqualError(t, err, testError.Error())
			} else {
				assert.Equal(t, testResult, v)
			}
		})
	}
}

func TestSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Supplier
		v    interface{}
	}{
		{
			name: "ok",
			s:    supplier,
			v:    testResult,
		},
		{
			name: "with_error",
			s:    supplierWithError,
			v:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := tt.s.ToSilentSupplier()
			assert.NotNil(t, ss)

			v := ss()
			assert.Equal(t, tt.v, v)
		})
	}
}

func TestSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Supplier
		v    interface{}
		err  error
	}{
		{
			name: "ok",
			s:    supplier,
			v:    testResult,
			err:  nil,
		},
		{
			name: "with_error",
			s:    supplierWithError,
			v:    nil,
			err:  testError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := tt.s.ToMustSupplier()
			assert.NotNil(t, ms)

			if tt.err != nil {
				assert.PanicsWithError(t, testError.Error(), func() {
					v := ms()
					assert.Equal(t, tt.v, v)
				})
			} else {
				v := ms()
				assert.Equal(t, tt.v, v)
			}
		})
	}
}

func TestSilentSupplier(t *testing.T) {
	var ss SilentSupplier = func() interface{} {
		return testResult
	}
	v := ss()
	assert.Equal(t, testResult, v)
}

func TestMustSupplier(t *testing.T) {
	var ms MustSupplier = func() interface{} {
		return testResult
	}

	v := ms()
	assert.Equal(t, testResult, v)
}

func TestMustSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Supplier
		v    interface{}
	}{
		{
			name: "ok",
			s:    supplier,
			v:    testResult,
		},
		{
			name: "with_error",
			s:    supplierWithError,
			v:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := tt.s.ToMustSupplier()
			assert.NotNil(t, ms)

			ss := ms.ToSilentSupplier()
			assert.NotNil(t, ss)

			v := ss()
			assert.Equal(t, tt.v, v)
		})
	}
}

func TestMustSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name string
		s    Supplier
	}{
		{
			name: "ok",
			s:    supplier,
		},
		{
			name: "with_error",
			s:    supplierWithError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := tt.s.ToMustSupplier()
			assert.NotNil(t, ms)

			s := ms.ToSupplier()
			assert.NotNil(t, s)

			v, err := s()
			if err != nil {
				assert.EqualError(t, err, testError.Error())
			} else {
				assert.Equal(t, testResult, v)
			}
		})
	}
}
