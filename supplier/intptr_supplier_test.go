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
	resTestIntPtrSupplier *int
	errTestIntPtrSupplier = errors.New("error")
)

func testIntPtrSupplier() (*int, error) {
	return resTestIntPtrSupplier, nil
}

func testIntPtrSupplierWithError() (*int, error) {
	return resTestIntPtrSupplier, errTestIntPtrSupplier
}

func TestIntPtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       IntPtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testIntPtrSupplier,
		},
		{
			name:    "with error",
			s:       testIntPtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestIntPtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestIntPtrSupplier, v)
			}
		})
	}
}

func TestIntPtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       IntPtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testIntPtrSupplier,
		},
		{
			name:    "with error",
			s:       testIntPtrSupplierWithError,
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
				r.EqualError(err, errTestIntPtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestIntPtrSupplier, v)
			}
		})
	}
}

func TestIntPtrSupplier_ToSilentIntPtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       IntPtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testIntPtrSupplier,
		},
		{
			name:    "with error",
			s:       testIntPtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentIntPtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestIntPtrSupplier, v)
			}
		})
	}
}

func TestIntPtrSupplier_ToMustIntPtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       IntPtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testIntPtrSupplier,
		},
		{
			name:    "with error",
			s:       testIntPtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustIntPtrSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestIntPtrSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestIntPtrSupplier, v)
			}
		})
	}
}

func TestSilentIntPtrSupplier(t *testing.T) {
	var ss SilentIntPtrSupplier = func() *int {
		return resTestIntPtrSupplier
	}
	v := ss()
	require.Equal(t, resTestIntPtrSupplier, v)
}

func TestSilentIntPtrSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       IntPtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testIntPtrSupplier,
		},
		{
			name:    "with error",
			s:       testIntPtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentIntPtrSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestIntPtrSupplier, v)
			}
		})
	}
}

func TestMustIntPtrSupplier(t *testing.T) {
	var ms MustIntPtrSupplier = func() *int {
		return resTestIntPtrSupplier
	}

	v := ms()
	require.Equal(t, resTestIntPtrSupplier, v)
}

func TestMustIntPtrSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       IntPtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testIntPtrSupplier,
		},
		{
			name:    "with error",
			s:       testIntPtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustIntPtrSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestIntPtrSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestIntPtrSupplier, v)
			}
		})
	}
}

func TestMustIntPtrSupplier_ToSilentIntPtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       IntPtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testIntPtrSupplier,
		},
		{
			name:    "with error",
			s:       testIntPtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustIntPtrSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentIntPtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestIntPtrSupplier, v)
			}
		})
	}
}

func TestMustIntPtrSupplier_ToIntPtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       IntPtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testIntPtrSupplier,
		},
		{
			name:    "with error",
			s:       testIntPtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustIntPtrSupplier()
			r.NotNil(ms)

			s := ms.ToIntPtrSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestIntPtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestIntPtrSupplier, v)
			}
		})
	}
}
