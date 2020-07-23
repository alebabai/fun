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
	resTestInt64PtrSupplier *int64
	errTestInt64PtrSupplier = errors.New("error")
)

func testInt64PtrSupplier() (*int64, error) {
	return resTestInt64PtrSupplier, nil
}

func testInt64PtrSupplierWithError() (*int64, error) {
	return resTestInt64PtrSupplier, errTestInt64PtrSupplier
}

func TestInt64PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int64PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt64PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt64PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestInt64PtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt64PtrSupplier, v)
			}
		})
	}
}

func TestInt64PtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int64PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt64PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt64PtrSupplierWithError,
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
				r.EqualError(err, errTestInt64PtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt64PtrSupplier, v)
			}
		})
	}
}

func TestInt64PtrSupplier_ToSilentInt64PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int64PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt64PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt64PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentInt64PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt64PtrSupplier, v)
			}
		})
	}
}

func TestInt64PtrSupplier_ToMustInt64PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int64PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt64PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt64PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt64PtrSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestInt64PtrSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt64PtrSupplier, v)
			}
		})
	}
}

func TestSilentInt64PtrSupplier(t *testing.T) {
	var ss SilentInt64PtrSupplier = func() *int64 {
		return resTestInt64PtrSupplier
	}
	v := ss()
	require.Equal(t, resTestInt64PtrSupplier, v)
}

func TestSilentInt64PtrSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int64PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt64PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt64PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentInt64PtrSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt64PtrSupplier, v)
			}
		})
	}
}

func TestMustInt64PtrSupplier(t *testing.T) {
	var ms MustInt64PtrSupplier = func() *int64 {
		return resTestInt64PtrSupplier
	}

	v := ms()
	require.Equal(t, resTestInt64PtrSupplier, v)
}

func TestMustInt64PtrSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int64PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt64PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt64PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustInt64PtrSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestInt64PtrSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt64PtrSupplier, v)
			}
		})
	}
}

func TestMustInt64PtrSupplier_ToSilentInt64PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int64PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt64PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt64PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt64PtrSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentInt64PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt64PtrSupplier, v)
			}
		})
	}
}

func TestMustInt64PtrSupplier_ToInt64PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int64PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt64PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt64PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt64PtrSupplier()
			r.NotNil(ms)

			s := ms.ToInt64PtrSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestInt64PtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt64PtrSupplier, v)
			}
		})
	}
}
