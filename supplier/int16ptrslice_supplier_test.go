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
	resTestInt16PtrSliceSupplier []*int16
	errTestInt16PtrSliceSupplier = errors.New("error")
)

func testInt16PtrSliceSupplier() ([]*int16, error) {
	return resTestInt16PtrSliceSupplier, nil
}

func testInt16PtrSliceSupplierWithError() ([]*int16, error) {
	return resTestInt16PtrSliceSupplier, errTestInt16PtrSliceSupplier
}

func TestInt16PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int16PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt16PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt16PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestInt16PtrSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt16PtrSliceSupplier, v)
			}
		})
	}
}

func TestInt16PtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int16PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt16PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt16PtrSliceSupplierWithError,
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
				r.EqualError(err, errTestInt16PtrSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt16PtrSliceSupplier, v)
			}
		})
	}
}

func TestInt16PtrSliceSupplier_ToSilentInt16PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int16PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt16PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt16PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentInt16PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt16PtrSliceSupplier, v)
			}
		})
	}
}

func TestInt16PtrSliceSupplier_ToMustInt16PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int16PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt16PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt16PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt16PtrSliceSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestInt16PtrSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt16PtrSliceSupplier, v)
			}
		})
	}
}

func TestSilentInt16PtrSliceSupplier(t *testing.T) {
	var ss SilentInt16PtrSliceSupplier = func() []*int16 {
		return resTestInt16PtrSliceSupplier
	}
	v := ss()
	require.Equal(t, resTestInt16PtrSliceSupplier, v)
}

func TestSilentInt16PtrSliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int16PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt16PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt16PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentInt16PtrSliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt16PtrSliceSupplier, v)
			}
		})
	}
}

func TestMustInt16PtrSliceSupplier(t *testing.T) {
	var ms MustInt16PtrSliceSupplier = func() []*int16 {
		return resTestInt16PtrSliceSupplier
	}

	v := ms()
	require.Equal(t, resTestInt16PtrSliceSupplier, v)
}

func TestMustInt16PtrSliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int16PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt16PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt16PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustInt16PtrSliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestInt16PtrSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt16PtrSliceSupplier, v)
			}
		})
	}
}

func TestMustInt16PtrSliceSupplier_ToSilentInt16PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int16PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt16PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt16PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt16PtrSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentInt16PtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt16PtrSliceSupplier, v)
			}
		})
	}
}

func TestMustInt16PtrSliceSupplier_ToInt16PtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int16PtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt16PtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt16PtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt16PtrSliceSupplier()
			r.NotNil(ms)

			s := ms.ToInt16PtrSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestInt16PtrSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt16PtrSliceSupplier, v)
			}
		})
	}
}
