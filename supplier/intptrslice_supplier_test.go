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
	resTestIntPtrSliceSupplier []*int
	errTestIntPtrSliceSupplier = errors.New("error")
)

func testIntPtrSliceSupplier() ([]*int, error) {
	return resTestIntPtrSliceSupplier, nil
}

func testIntPtrSliceSupplierWithError() ([]*int, error) {
	return resTestIntPtrSliceSupplier, errTestIntPtrSliceSupplier
}

func TestIntPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       IntPtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testIntPtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testIntPtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestIntPtrSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestIntPtrSliceSupplier, v)
			}
		})
	}
}

func TestIntPtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       IntPtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testIntPtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testIntPtrSliceSupplierWithError,
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
				r.EqualError(err, errTestIntPtrSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestIntPtrSliceSupplier, v)
			}
		})
	}
}

func TestIntPtrSliceSupplier_ToSilentIntPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       IntPtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testIntPtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testIntPtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentIntPtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestIntPtrSliceSupplier, v)
			}
		})
	}
}

func TestIntPtrSliceSupplier_ToMustIntPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       IntPtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testIntPtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testIntPtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustIntPtrSliceSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestIntPtrSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestIntPtrSliceSupplier, v)
			}
		})
	}
}

func TestSilentIntPtrSliceSupplier(t *testing.T) {
	var ss SilentIntPtrSliceSupplier = func() []*int {
		return resTestIntPtrSliceSupplier
	}
	v := ss()
	require.Equal(t, resTestIntPtrSliceSupplier, v)
}

func TestSilentIntPtrSliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       IntPtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testIntPtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testIntPtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentIntPtrSliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestIntPtrSliceSupplier, v)
			}
		})
	}
}

func TestMustIntPtrSliceSupplier(t *testing.T) {
	var ms MustIntPtrSliceSupplier = func() []*int {
		return resTestIntPtrSliceSupplier
	}

	v := ms()
	require.Equal(t, resTestIntPtrSliceSupplier, v)
}

func TestMustIntPtrSliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       IntPtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testIntPtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testIntPtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustIntPtrSliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestIntPtrSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestIntPtrSliceSupplier, v)
			}
		})
	}
}

func TestMustIntPtrSliceSupplier_ToSilentIntPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       IntPtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testIntPtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testIntPtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustIntPtrSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentIntPtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestIntPtrSliceSupplier, v)
			}
		})
	}
}

func TestMustIntPtrSliceSupplier_ToIntPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       IntPtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testIntPtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testIntPtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustIntPtrSliceSupplier()
			r.NotNil(ms)

			s := ms.ToIntPtrSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestIntPtrSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestIntPtrSliceSupplier, v)
			}
		})
	}
}
