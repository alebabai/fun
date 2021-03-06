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
	resTestUintSliceSupplier []uint
	errTestUintSliceSupplier = errors.New("error")
)

func testUintSliceSupplier() ([]uint, error) {
	return resTestUintSliceSupplier, nil
}

func testUintSliceSupplierWithError() ([]uint, error) {
	return resTestUintSliceSupplier, errTestUintSliceSupplier
}

func TestUintSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       UintSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUintSliceSupplier,
		},
		{
			name:    "with error",
			s:       testUintSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestUintSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUintSliceSupplier, v)
			}
		})
	}
}

func TestUintSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       UintSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUintSliceSupplier,
		},
		{
			name:    "with error",
			s:       testUintSliceSupplierWithError,
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
				r.EqualError(err, errTestUintSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUintSliceSupplier, v)
			}
		})
	}
}

func TestUintSliceSupplier_ToSilentUintSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       UintSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUintSliceSupplier,
		},
		{
			name:    "with error",
			s:       testUintSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUintSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUintSliceSupplier, v)
			}
		})
	}
}

func TestUintSliceSupplier_ToMustUintSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       UintSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUintSliceSupplier,
		},
		{
			name:    "with error",
			s:       testUintSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUintSliceSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestUintSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUintSliceSupplier, v)
			}
		})
	}
}

func TestSilentUintSliceSupplier(t *testing.T) {
	var ss SilentUintSliceSupplier = func() []uint {
		return resTestUintSliceSupplier
	}
	v := ss()
	require.Equal(t, resTestUintSliceSupplier, v)
}

func TestSilentUintSliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       UintSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUintSliceSupplier,
		},
		{
			name:    "with error",
			s:       testUintSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentUintSliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUintSliceSupplier, v)
			}
		})
	}
}

func TestMustUintSliceSupplier(t *testing.T) {
	var ms MustUintSliceSupplier = func() []uint {
		return resTestUintSliceSupplier
	}

	v := ms()
	require.Equal(t, resTestUintSliceSupplier, v)
}

func TestMustUintSliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       UintSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUintSliceSupplier,
		},
		{
			name:    "with error",
			s:       testUintSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustUintSliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestUintSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUintSliceSupplier, v)
			}
		})
	}
}

func TestMustUintSliceSupplier_ToSilentUintSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       UintSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUintSliceSupplier,
		},
		{
			name:    "with error",
			s:       testUintSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUintSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentUintSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUintSliceSupplier, v)
			}
		})
	}
}

func TestMustUintSliceSupplier_ToUintSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       UintSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUintSliceSupplier,
		},
		{
			name:    "with error",
			s:       testUintSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUintSliceSupplier()
			r.NotNil(ms)

			s := ms.ToUintSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestUintSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUintSliceSupplier, v)
			}
		})
	}
}
