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
	resTestInt64PtrSliceSupplier []*int64
	errTestInt64PtrSliceSupplier = errors.New("error")
)

func testInt64PtrSliceSupplier() ([]*int64, error) {
	return resTestInt64PtrSliceSupplier, nil
}

func testInt64PtrSliceSupplierWithError() ([]*int64, error) {
	return resTestInt64PtrSliceSupplier, errTestInt64PtrSliceSupplier
}

func TestInt64PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int64PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt64PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt64PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestInt64PtrSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt64PtrSliceSupplier, v)
			}
		})
	}
}

func TestInt64PtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int64PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt64PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt64PtrSliceSupplierWithError,
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
				r.EqualError(err, errTestInt64PtrSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt64PtrSliceSupplier, v)
			}
		})
	}
}

func TestInt64PtrSliceSupplier_ToSilentInt64PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int64PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt64PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt64PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentInt64PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt64PtrSliceSupplier, v)
			}
		})
	}
}

func TestInt64PtrSliceSupplier_ToMustInt64PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int64PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt64PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt64PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt64PtrSliceSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestInt64PtrSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt64PtrSliceSupplier, v)
			}
		})
	}
}

func TestSilentInt64PtrSliceSupplier(t *testing.T) {
	var ss SilentInt64PtrSliceSupplier = func() []*int64 {
		return resTestInt64PtrSliceSupplier
	}
	v := ss()
	require.Equal(t, resTestInt64PtrSliceSupplier, v)
}

func TestSilentInt64PtrSliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int64PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt64PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt64PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentInt64PtrSliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt64PtrSliceSupplier, v)
			}
		})
	}
}

func TestMustInt64PtrSliceSupplier(t *testing.T) {
	var ms MustInt64PtrSliceSupplier = func() []*int64 {
		return resTestInt64PtrSliceSupplier
	}

	v := ms()
	require.Equal(t, resTestInt64PtrSliceSupplier, v)
}

func TestMustInt64PtrSliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int64PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt64PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt64PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustInt64PtrSliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestInt64PtrSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt64PtrSliceSupplier, v)
			}
		})
	}
}

func TestMustInt64PtrSliceSupplier_ToSilentInt64PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int64PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt64PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt64PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt64PtrSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentInt64PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt64PtrSliceSupplier, v)
			}
		})
	}
}

func TestMustInt64PtrSliceSupplier_ToInt64PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int64PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt64PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt64PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt64PtrSliceSupplier()
			r.NotNil(ms)

			s := ms.ToInt64PtrSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestInt64PtrSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt64PtrSliceSupplier, v)
			}
		})
	}
}
