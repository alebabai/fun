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
	resTestFloat64SliceSupplier []float64
	errTestFloat64SliceSupplier = errors.New("error")
)

func testFloat64SliceSupplier() ([]float64, error) {
	return resTestFloat64SliceSupplier, nil
}

func testFloat64SliceSupplierWithError() ([]float64, error) {
	return resTestFloat64SliceSupplier, errTestFloat64SliceSupplier
}

func TestFloat64SliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float64SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat64SliceSupplier,
		},
		{
			name:    "with error",
			s:       testFloat64SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestFloat64SliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestFloat64SliceSupplier, v)
			}
		})
	}
}

func TestFloat64SliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float64SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat64SliceSupplier,
		},
		{
			name:    "with error",
			s:       testFloat64SliceSupplierWithError,
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
				r.EqualError(err, errTestFloat64SliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestFloat64SliceSupplier, v)
			}
		})
	}
}

func TestFloat64SliceSupplier_ToSilentFloat64SliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float64SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat64SliceSupplier,
		},
		{
			name:    "with error",
			s:       testFloat64SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentFloat64SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestFloat64SliceSupplier, v)
			}
		})
	}
}

func TestFloat64SliceSupplier_ToMustFloat64SliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float64SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat64SliceSupplier,
		},
		{
			name:    "with error",
			s:       testFloat64SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat64SliceSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestFloat64SliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestFloat64SliceSupplier, v)
			}
		})
	}
}

func TestSilentFloat64SliceSupplier(t *testing.T) {
	var ss SilentFloat64SliceSupplier = func() []float64 {
		return resTestFloat64SliceSupplier
	}
	v := ss()
	require.Equal(t, resTestFloat64SliceSupplier, v)
}

func TestSilentFloat64SliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float64SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat64SliceSupplier,
		},
		{
			name:    "with error",
			s:       testFloat64SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentFloat64SliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestFloat64SliceSupplier, v)
			}
		})
	}
}

func TestMustFloat64SliceSupplier(t *testing.T) {
	var ms MustFloat64SliceSupplier = func() []float64 {
		return resTestFloat64SliceSupplier
	}

	v := ms()
	require.Equal(t, resTestFloat64SliceSupplier, v)
}

func TestMustFloat64SliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float64SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat64SliceSupplier,
		},
		{
			name:    "with error",
			s:       testFloat64SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustFloat64SliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestFloat64SliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestFloat64SliceSupplier, v)
			}
		})
	}
}

func TestMustFloat64SliceSupplier_ToSilentFloat64SliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float64SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat64SliceSupplier,
		},
		{
			name:    "with error",
			s:       testFloat64SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat64SliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentFloat64SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestFloat64SliceSupplier, v)
			}
		})
	}
}

func TestMustFloat64SliceSupplier_ToFloat64SliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Float64SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testFloat64SliceSupplier,
		},
		{
			name:    "with error",
			s:       testFloat64SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustFloat64SliceSupplier()
			r.NotNil(ms)

			s := ms.ToFloat64SliceSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestFloat64SliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestFloat64SliceSupplier, v)
			}
		})
	}
}
