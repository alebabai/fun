package fun

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	testSupplierValue interface{} = struct{}{}
	testSupplierError             = errors.New("error")
)

func testSupplier() (interface{}, error) {
	return testSupplierValue, nil
}

func testSupplierWithError() (interface{}, error) {
	return testSupplierValue, testSupplierError
}

func TestSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testSupplier,
		},
		{
			name:    "with error",
			s:       testSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, testSupplierError.Error())
				r.Empty(v)
			} else {
				r.Equal(testSupplierValue, v)
			}
		})
	}
}

func TestSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testSupplier,
		},
		{
			name:    "with error",
			s:       testSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(testSupplierValue, v)
			}
		})
	}
}

func TestSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testSupplier,
		},
		{
			name:    "with error",
			s:       testSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(testSupplierError.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(testSupplierValue, v)
			}
		})
	}
}

func TestSilentSupplier(t *testing.T) {
	var ss SilentSupplier = func() interface{} {
		return testSupplierValue
	}
	v := ss()
	require.Equal(t, testSupplierValue, v)
}

func TestMustSupplier(t *testing.T) {
	var ms MustSupplier = func() interface{} {
		return testSupplierValue
	}

	v := ms()
	require.Equal(t, testSupplierValue, v)
}

func TestMustSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testSupplier,
		},
		{
			name:    "with error",
			s:       testSupplierWithError,
			wantErr: true,
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
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(testSupplierValue, v)
			}
		})
	}
}

func TestMustSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Supplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testSupplier,
		},
		{
			name:    "with error",
			s:       testSupplierWithError,
			wantErr: true,
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
			if tt.wantErr {
				r.EqualError(err, testSupplierError.Error())
				r.Empty(v)
			} else {
				r.Equal(testSupplierValue, v)
			}
		})
	}
}
