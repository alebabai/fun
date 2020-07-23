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
	resTestFloat64PtrSliceSupplier []*float64
	errTestFloat64PtrSliceSupplier = errors.New("error")
)

func testFloat64PtrSliceSupplier() ([]*float64, error) {
	return resTestFloat64PtrSliceSupplier, nil
}

func testFloat64PtrSliceSupplierWithError() ([]*float64, error) {
	return resTestFloat64PtrSliceSupplier, errTestFloat64PtrSliceSupplier
}

func TestFloat64PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float64PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat64PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testFloat64PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestFloat64PtrSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestFloat64PtrSliceSupplier, v)
			}
		})
	}
}

func TestFloat64PtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float64PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat64PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testFloat64PtrSliceSupplierWithError,
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
				r.EqualError(err, errTestFloat64PtrSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestFloat64PtrSliceSupplier, v)
			}
		})
	}
}

func TestFloat64PtrSliceSupplier_ToSilentFloat64PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float64PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat64PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testFloat64PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentFloat64PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestFloat64PtrSliceSupplier, v)
			}
		})
	}
}

func TestFloat64PtrSliceSupplier_ToMustFloat64PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float64PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat64PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testFloat64PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat64PtrSliceSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestFloat64PtrSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestFloat64PtrSliceSupplier, v)
			}
		})
	}
}

func TestSilentFloat64PtrSliceSupplier(t *testing.T) {
	var ss SilentFloat64PtrSliceSupplier = func() []*float64 {
		return resTestFloat64PtrSliceSupplier
	}
	v := ss()
	require.Equal(t, resTestFloat64PtrSliceSupplier, v)
}

func TestSilentFloat64PtrSliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float64PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat64PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testFloat64PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentFloat64PtrSliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestFloat64PtrSliceSupplier, v)
			}
		})
	}
}

func TestMustFloat64PtrSliceSupplier(t *testing.T) {
	var ms MustFloat64PtrSliceSupplier = func() []*float64 {
		return resTestFloat64PtrSliceSupplier
	}

	v := ms()
	require.Equal(t, resTestFloat64PtrSliceSupplier, v)
}

func TestMustFloat64PtrSliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float64PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat64PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testFloat64PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustFloat64PtrSliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestFloat64PtrSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestFloat64PtrSliceSupplier, v)
			}
		})
	}
}

func TestMustFloat64PtrSliceSupplier_ToSilentFloat64PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float64PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat64PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testFloat64PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat64PtrSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentFloat64PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestFloat64PtrSliceSupplier, v)
			}
		})
	}
}

func TestMustFloat64PtrSliceSupplier_ToFloat64PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float64PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat64PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testFloat64PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat64PtrSliceSupplier()
			r.NotNil(ms)

			s := ms.ToFloat64PtrSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestFloat64PtrSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestFloat64PtrSliceSupplier, v)
			}
		})
	}
}
