package fun

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	valTestSupplier interface{} = struct{}{}
	errTestSupplier             = errors.New("error")
)

func testSupplier() (interface{}, error) {
	return valTestSupplier, nil
}

func testSupplierWithError() (interface{}, error) {
	return valTestSupplier, errTestSupplier
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
				r.EqualError(err, errTestSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(valTestSupplier, v)
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
				r.Equal(valTestSupplier, v)
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
				r.PanicsWithError(errTestSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(valTestSupplier, v)
			}
		})
	}
}

func TestSilentSupplier(t *testing.T) {
	var ss SilentSupplier = func() interface{} {
		return valTestSupplier
	}
	v := ss()
	require.Equal(t, valTestSupplier, v)
}

func TestMustSupplier(t *testing.T) {
	var ms MustSupplier = func() interface{} {
		return valTestSupplier
	}

	v := ms()
	require.Equal(t, valTestSupplier, v)
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
				r.Equal(valTestSupplier, v)
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
				r.EqualError(err, errTestSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(valTestSupplier, v)
			}
		})
	}
}
