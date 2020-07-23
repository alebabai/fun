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
	resTestUintPtrSliceSupplier []*uint
	errTestUintPtrSliceSupplier = errors.New("error")
)

func testUintPtrSliceSupplier() ([]*uint, error) {
	return resTestUintPtrSliceSupplier, nil
}

func testUintPtrSliceSupplierWithError() ([]*uint, error) {
	return resTestUintPtrSliceSupplier, errTestUintPtrSliceSupplier
}

func TestUintPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       UintPtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUintPtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testUintPtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestUintPtrSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUintPtrSliceSupplier, v)
			}
		})
	}
}

func TestUintPtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       UintPtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUintPtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testUintPtrSliceSupplierWithError,
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
				r.EqualError(err, errTestUintPtrSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUintPtrSliceSupplier, v)
			}
		})
	}
}

func TestUintPtrSliceSupplier_ToSilentUintPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       UintPtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUintPtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testUintPtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUintPtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUintPtrSliceSupplier, v)
			}
		})
	}
}

func TestUintPtrSliceSupplier_ToMustUintPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       UintPtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUintPtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testUintPtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUintPtrSliceSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestUintPtrSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUintPtrSliceSupplier, v)
			}
		})
	}
}

func TestSilentUintPtrSliceSupplier(t *testing.T) {
	var ss SilentUintPtrSliceSupplier = func() []*uint {
		return resTestUintPtrSliceSupplier
	}
	v := ss()
	require.Equal(t, resTestUintPtrSliceSupplier, v)
}

func TestSilentUintPtrSliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       UintPtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUintPtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testUintPtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentUintPtrSliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUintPtrSliceSupplier, v)
			}
		})
	}
}

func TestMustUintPtrSliceSupplier(t *testing.T) {
	var ms MustUintPtrSliceSupplier = func() []*uint {
		return resTestUintPtrSliceSupplier
	}

	v := ms()
	require.Equal(t, resTestUintPtrSliceSupplier, v)
}

func TestMustUintPtrSliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       UintPtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUintPtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testUintPtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustUintPtrSliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestUintPtrSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUintPtrSliceSupplier, v)
			}
		})
	}
}

func TestMustUintPtrSliceSupplier_ToSilentUintPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       UintPtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUintPtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testUintPtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUintPtrSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentUintPtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUintPtrSliceSupplier, v)
			}
		})
	}
}

func TestMustUintPtrSliceSupplier_ToUintPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       UintPtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUintPtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testUintPtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUintPtrSliceSupplier()
			r.NotNil(ms)

			s := ms.ToUintPtrSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestUintPtrSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUintPtrSliceSupplier, v)
			}
		})
	}
}
