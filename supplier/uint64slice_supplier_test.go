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
	resTestUint64SliceSupplier []uint64
	errTestUint64SliceSupplier = errors.New("error")
)

func testUint64SliceSupplier() ([]uint64, error) {
	return resTestUint64SliceSupplier, nil
}

func testUint64SliceSupplierWithError() ([]uint64, error) {
	return resTestUint64SliceSupplier, errTestUint64SliceSupplier
}

func TestUint64SliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint64SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint64SliceSupplier,
		},
		{
			name:    "with error",
			s:       testUint64SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestUint64SliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUint64SliceSupplier, v)
			}
		})
	}
}

func TestUint64SliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint64SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint64SliceSupplier,
		},
		{
			name:    "with error",
			s:       testUint64SliceSupplierWithError,
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
				r.EqualError(err, errTestUint64SliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUint64SliceSupplier, v)
			}
		})
	}
}

func TestUint64SliceSupplier_ToSilentUint64SliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint64SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint64SliceSupplier,
		},
		{
			name:    "with error",
			s:       testUint64SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUint64SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUint64SliceSupplier, v)
			}
		})
	}
}

func TestUint64SliceSupplier_ToMustUint64SliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint64SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint64SliceSupplier,
		},
		{
			name:    "with error",
			s:       testUint64SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint64SliceSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestUint64SliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint64SliceSupplier, v)
			}
		})
	}
}

func TestSilentUint64SliceSupplier(t *testing.T) {
	var ss SilentUint64SliceSupplier = func() []uint64 {
		return resTestUint64SliceSupplier
	}
	v := ss()
	require.Equal(t, resTestUint64SliceSupplier, v)
}

func TestSilentUint64SliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint64SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint64SliceSupplier,
		},
		{
			name:    "with error",
			s:       testUint64SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentUint64SliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUint64SliceSupplier, v)
			}
		})
	}
}

func TestMustUint64SliceSupplier(t *testing.T) {
	var ms MustUint64SliceSupplier = func() []uint64 {
		return resTestUint64SliceSupplier
	}

	v := ms()
	require.Equal(t, resTestUint64SliceSupplier, v)
}

func TestMustUint64SliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint64SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint64SliceSupplier,
		},
		{
			name:    "with error",
			s:       testUint64SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustUint64SliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestUint64SliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint64SliceSupplier, v)
			}
		})
	}
}

func TestMustUint64SliceSupplier_ToSilentUint64SliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint64SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint64SliceSupplier,
		},
		{
			name:    "with error",
			s:       testUint64SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint64SliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentUint64SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUint64SliceSupplier, v)
			}
		})
	}
}

func TestMustUint64SliceSupplier_ToUint64SliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint64SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint64SliceSupplier,
		},
		{
			name:    "with error",
			s:       testUint64SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint64SliceSupplier()
			r.NotNil(ms)

			s := ms.ToUint64SliceSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestUint64SliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUint64SliceSupplier, v)
			}
		})
	}
}
