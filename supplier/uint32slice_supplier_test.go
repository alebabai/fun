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
	resTestUint32SliceSupplier []uint32
	errTestUint32SliceSupplier = errors.New("error")
)

func testUint32SliceSupplier() ([]uint32, error) {
	return resTestUint32SliceSupplier, nil
}

func testUint32SliceSupplierWithError() ([]uint32, error) {
	return resTestUint32SliceSupplier, errTestUint32SliceSupplier
}

func TestUint32SliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint32SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint32SliceSupplier,
		},
		{
			name:    "with error",
			s:       testUint32SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestUint32SliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUint32SliceSupplier, v)
			}
		})
	}
}

func TestUint32SliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint32SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint32SliceSupplier,
		},
		{
			name:    "with error",
			s:       testUint32SliceSupplierWithError,
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
				r.EqualError(err, errTestUint32SliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUint32SliceSupplier, v)
			}
		})
	}
}

func TestUint32SliceSupplier_ToSilentUint32SliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint32SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint32SliceSupplier,
		},
		{
			name:    "with error",
			s:       testUint32SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUint32SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUint32SliceSupplier, v)
			}
		})
	}
}

func TestUint32SliceSupplier_ToMustUint32SliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint32SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint32SliceSupplier,
		},
		{
			name:    "with error",
			s:       testUint32SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint32SliceSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestUint32SliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint32SliceSupplier, v)
			}
		})
	}
}

func TestSilentUint32SliceSupplier(t *testing.T) {
	var ss SilentUint32SliceSupplier = func() []uint32 {
		return resTestUint32SliceSupplier
	}
	v := ss()
	require.Equal(t, resTestUint32SliceSupplier, v)
}

func TestSilentUint32SliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint32SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint32SliceSupplier,
		},
		{
			name:    "with error",
			s:       testUint32SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentUint32SliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUint32SliceSupplier, v)
			}
		})
	}
}

func TestMustUint32SliceSupplier(t *testing.T) {
	var ms MustUint32SliceSupplier = func() []uint32 {
		return resTestUint32SliceSupplier
	}

	v := ms()
	require.Equal(t, resTestUint32SliceSupplier, v)
}

func TestMustUint32SliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint32SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint32SliceSupplier,
		},
		{
			name:    "with error",
			s:       testUint32SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustUint32SliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestUint32SliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUint32SliceSupplier, v)
			}
		})
	}
}

func TestMustUint32SliceSupplier_ToSilentUint32SliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint32SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint32SliceSupplier,
		},
		{
			name:    "with error",
			s:       testUint32SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint32SliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentUint32SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUint32SliceSupplier, v)
			}
		})
	}
}

func TestMustUint32SliceSupplier_ToUint32SliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Uint32SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUint32SliceSupplier,
		},
		{
			name:    "with error",
			s:       testUint32SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUint32SliceSupplier()
			r.NotNil(ms)

			s := ms.ToUint32SliceSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestUint32SliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUint32SliceSupplier, v)
			}
		})
	}
}
