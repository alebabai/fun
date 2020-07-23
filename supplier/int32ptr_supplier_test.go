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
	resTestInt32PtrSupplier *int32
	errTestInt32PtrSupplier = errors.New("error")
)

func testInt32PtrSupplier() (*int32, error) {
	return resTestInt32PtrSupplier, nil
}

func testInt32PtrSupplierWithError() (*int32, error) {
	return resTestInt32PtrSupplier, errTestInt32PtrSupplier
}

func TestInt32PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int32PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt32PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt32PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestInt32PtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt32PtrSupplier, v)
			}
		})
	}
}

func TestInt32PtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int32PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt32PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt32PtrSupplierWithError,
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
				r.EqualError(err, errTestInt32PtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt32PtrSupplier, v)
			}
		})
	}
}

func TestInt32PtrSupplier_ToSilentInt32PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int32PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt32PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt32PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentInt32PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt32PtrSupplier, v)
			}
		})
	}
}

func TestInt32PtrSupplier_ToMustInt32PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int32PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt32PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt32PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt32PtrSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestInt32PtrSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt32PtrSupplier, v)
			}
		})
	}
}

func TestSilentInt32PtrSupplier(t *testing.T) {
	var ss SilentInt32PtrSupplier = func() *int32 {
		return resTestInt32PtrSupplier
	}
	v := ss()
	require.Equal(t, resTestInt32PtrSupplier, v)
}

func TestSilentInt32PtrSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int32PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt32PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt32PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentInt32PtrSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt32PtrSupplier, v)
			}
		})
	}
}

func TestMustInt32PtrSupplier(t *testing.T) {
	var ms MustInt32PtrSupplier = func() *int32 {
		return resTestInt32PtrSupplier
	}

	v := ms()
	require.Equal(t, resTestInt32PtrSupplier, v)
}

func TestMustInt32PtrSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int32PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt32PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt32PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustInt32PtrSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestInt32PtrSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt32PtrSupplier, v)
			}
		})
	}
}

func TestMustInt32PtrSupplier_ToSilentInt32PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int32PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt32PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt32PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt32PtrSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentInt32PtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt32PtrSupplier, v)
			}
		})
	}
}

func TestMustInt32PtrSupplier_ToInt32PtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int32PtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt32PtrSupplier,
		},
		{
			name:    "with error",
			s:       testInt32PtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt32PtrSupplier()
			r.NotNil(ms)

			s := ms.ToInt32PtrSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestInt32PtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt32PtrSupplier, v)
			}
		})
	}
}
