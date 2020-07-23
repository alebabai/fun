// CODE GENERATED AUTOMATICALLY
// SOURCE: supplier_test.go.tmpl
// DO NOT EDIT

package supplier

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	resTestFloat64PtrSupplier *float64
	errTestFloat64PtrSupplier = errors.New("error")
)

func testFloat64PtrSupplier() (*float64, error) {
	return resTestFloat64PtrSupplier, nil
}

func testFloat64PtrSupplierWithError() (*float64, error) {
	return resTestFloat64PtrSupplier, errTestFloat64PtrSupplier
}

func TestFloat64PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float64PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat64PtrSupplier,
		},
		{
			name:    "with error",
			s:       testFloat64PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestFloat64PtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestFloat64PtrSupplier, v)
			}
		})
	}
}

func TestFloat64PtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float64PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat64PtrSupplier,
		},
		{
			name:    "with error",
			s:       testFloat64PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			s := tt.s.ToSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestFloat64PtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestFloat64PtrSupplier, v)
			}
		})
	}
}

func TestFloat64PtrSupplier_ToSilentFloat64PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float64PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat64PtrSupplier,
		},
		{
			name:    "with error",
			s:       testFloat64PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentFloat64PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestFloat64PtrSupplier, v)
			}
		})
	}
}

func TestFloat64PtrSupplier_ToMustFloat64PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float64PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat64PtrSupplier,
		},
		{
			name:    "with error",
			s:       testFloat64PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat64PtrSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestFloat64PtrSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestFloat64PtrSupplier, v)
			}
		})
	}
}

func TestSilentFloat64PtrSupplier(t *testing.T) {
	var ss SilentFloat64PtrSupplier = func() *float64 {
		return resTestFloat64PtrSupplier
	}
	v := ss()
	require.Equal(t, resTestFloat64PtrSupplier, v)
}

func TestSilentFloat64PtrSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float64PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat64PtrSupplier,
		},
		{
			name:    "with error",
			s:       testFloat64PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentFloat64PtrSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestFloat64PtrSupplier, v)
			}
		})
	}
}

func TestMustFloat64PtrSupplier(t *testing.T) {
	var ms MustFloat64PtrSupplier = func() *float64 {
		return resTestFloat64PtrSupplier
	}

	v := ms()
	require.Equal(t, resTestFloat64PtrSupplier, v)
}

func TestMustFloat64PtrSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float64PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat64PtrSupplier,
		},
		{
			name:    "with error",
			s:       testFloat64PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustFloat64PtrSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestFloat64PtrSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestFloat64PtrSupplier, v)
			}
		})
	}
}

func TestMustFloat64PtrSupplier_ToSilentFloat64PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float64PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat64PtrSupplier,
		},
		{
			name:    "with error",
			s:       testFloat64PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat64PtrSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentFloat64PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestFloat64PtrSupplier, v)
			}
		})
	}
}

func TestMustFloat64PtrSupplier_ToFloat64PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float64PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat64PtrSupplier,
		},
		{
			name:    "with error",
			s:       testFloat64PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat64PtrSupplier()
			r.NotNil(ms)

			s := ms.ToFloat64PtrSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestFloat64PtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestFloat64PtrSupplier, v)
			}
		})
	}
}
