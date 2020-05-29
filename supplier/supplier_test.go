package supplier

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
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
			r := require.New(t)

			v, err := tt.s()
			if err != nil {
				r.EqualError(err, testError.Error())
			} else {
				r.Equal(testResult, v)
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
			r := require.New(t)

			ss := tt.s.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			r.Equal(tt.v, v)
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
			r := require.New(t)

			ms := tt.s.ToMustSupplier()
			r.NotNil(ms)

			if tt.err != nil {
				r.PanicsWithError(testError.Error(), func() {
					v := ms()
					r.Equal(tt.v, v)
				})
			} else {
				v := ms()
				r.Equal(tt.v, v)
			}
		})
	}
}

func TestSilentSupplier(t *testing.T) {
	var ss SilentSupplier = func() interface{} {
		return testResult
	}
	v := ss()
	require.Equal(t, testResult, v)
}

func TestMustSupplier(t *testing.T) {
	var ms MustSupplier = func() interface{} {
		return testResult
	}

	v := ms()
	require.Equal(t, testResult, v)
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
			r := require.New(t)

			ms := tt.s.ToMustSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			r.Equal(tt.v, v)
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
			r := require.New(t)

			ms := tt.s.ToMustSupplier()
			r.NotNil(ms)

			s := ms.ToSupplier()
			r.NotNil(s)

			v, err := s()
			if err != nil {
				r.EqualError(err, testError.Error())
			} else {
				r.Equal(testResult, v)
			}
		})
	}
}
