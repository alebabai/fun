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
	resTestIntSliceSupplier []int
	errTestIntSliceSupplier = errors.New("error")
)

func testIntSliceSupplier() ([]int, error) {
	return resTestIntSliceSupplier, nil
}

func testIntSliceSupplierWithError() ([]int, error) {
	return resTestIntSliceSupplier, errTestIntSliceSupplier
}

func TestIntSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       IntSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testIntSliceSupplier,
		},
		{
			name:    "with error",
			s:       testIntSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestIntSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestIntSliceSupplier, v)
			}
		})
	}
}

func TestIntSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       IntSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testIntSliceSupplier,
		},
		{
			name:    "with error",
			s:       testIntSliceSupplierWithError,
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
				r.EqualError(err, errTestIntSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestIntSliceSupplier, v)
			}
		})
	}
}

func TestIntSliceSupplier_ToSilentIntSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       IntSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testIntSliceSupplier,
		},
		{
			name:    "with error",
			s:       testIntSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentIntSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestIntSliceSupplier, v)
			}
		})
	}
}

func TestIntSliceSupplier_ToMustIntSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       IntSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testIntSliceSupplier,
		},
		{
			name:    "with error",
			s:       testIntSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustIntSliceSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestIntSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestIntSliceSupplier, v)
			}
		})
	}
}

func TestSilentIntSliceSupplier(t *testing.T) {
	var ss SilentIntSliceSupplier = func() []int {
		return resTestIntSliceSupplier
	}
	v := ss()
	require.Equal(t, resTestIntSliceSupplier, v)
}

func TestSilentIntSliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       IntSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testIntSliceSupplier,
		},
		{
			name:    "with error",
			s:       testIntSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentIntSliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestIntSliceSupplier, v)
			}
		})
	}
}

func TestMustIntSliceSupplier(t *testing.T) {
	var ms MustIntSliceSupplier = func() []int {
		return resTestIntSliceSupplier
	}

	v := ms()
	require.Equal(t, resTestIntSliceSupplier, v)
}

func TestMustIntSliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       IntSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testIntSliceSupplier,
		},
		{
			name:    "with error",
			s:       testIntSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustIntSliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestIntSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestIntSliceSupplier, v)
			}
		})
	}
}

func TestMustIntSliceSupplier_ToSilentIntSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       IntSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testIntSliceSupplier,
		},
		{
			name:    "with error",
			s:       testIntSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustIntSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentIntSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestIntSliceSupplier, v)
			}
		})
	}
}

func TestMustIntSliceSupplier_ToIntSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       IntSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testIntSliceSupplier,
		},
		{
			name:    "with error",
			s:       testIntSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustIntSliceSupplier()
			r.NotNil(ms)

			s := ms.ToIntSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestIntSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestIntSliceSupplier, v)
			}
		})
	}
}
