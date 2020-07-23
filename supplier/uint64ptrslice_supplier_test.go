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
	resTestUint64PtrSliceSupplier []*uint64
	errTestUint64PtrSliceSupplier = errors.New("error")
)

func testUint64PtrSliceSupplier() ([]*uint64, error) {
	return resTestUint64PtrSliceSupplier, nil
}

func testUint64PtrSliceSupplierWithError() ([]*uint64, error) {
	return resTestUint64PtrSliceSupplier, errTestUint64PtrSliceSupplier
}

func TestUint64PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint64PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint64PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testUint64PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestUint64PtrSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUint64PtrSliceSupplier, v)
			}
		})
	}
}

func TestUint64PtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint64PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint64PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testUint64PtrSliceSupplierWithError,
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
				r.EqualError(err, errTestUint64PtrSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUint64PtrSliceSupplier, v)
			}
		})
	}
}

func TestUint64PtrSliceSupplier_ToSilentUint64PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint64PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint64PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testUint64PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUint64PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUint64PtrSliceSupplier, v)
			}
		})
	}
}

func TestUint64PtrSliceSupplier_ToMustUint64PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint64PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint64PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testUint64PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint64PtrSliceSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestUint64PtrSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint64PtrSliceSupplier, v)
			}
		})
	}
}

func TestSilentUint64PtrSliceSupplier(t *testing.T) {
	var ss SilentUint64PtrSliceSupplier = func() []*uint64 {
		return resTestUint64PtrSliceSupplier
	}
	v := ss()
	require.Equal(t, resTestUint64PtrSliceSupplier, v)
}

func TestSilentUint64PtrSliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint64PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint64PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testUint64PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentUint64PtrSliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUint64PtrSliceSupplier, v)
			}
		})
	}
}

func TestMustUint64PtrSliceSupplier(t *testing.T) {
	var ms MustUint64PtrSliceSupplier = func() []*uint64 {
		return resTestUint64PtrSliceSupplier
	}

	v := ms()
	require.Equal(t, resTestUint64PtrSliceSupplier, v)
}

func TestMustUint64PtrSliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint64PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint64PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testUint64PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustUint64PtrSliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestUint64PtrSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint64PtrSliceSupplier, v)
			}
		})
	}
}

func TestMustUint64PtrSliceSupplier_ToSilentUint64PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint64PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint64PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testUint64PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint64PtrSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentUint64PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUint64PtrSliceSupplier, v)
			}
		})
	}
}

func TestMustUint64PtrSliceSupplier_ToUint64PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint64PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint64PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testUint64PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint64PtrSliceSupplier()
			r.NotNil(ms)

			s := ms.ToUint64PtrSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestUint64PtrSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUint64PtrSliceSupplier, v)
			}
		})
	}
}
