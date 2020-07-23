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
	resTestInt8SliceSupplier []int8
	errTestInt8SliceSupplier = errors.New("error")
)

func testInt8SliceSupplier() ([]int8, error) {
	return resTestInt8SliceSupplier, nil
}

func testInt8SliceSupplierWithError() ([]int8, error) {
	return resTestInt8SliceSupplier, errTestInt8SliceSupplier
}

func TestInt8SliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int8SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt8SliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt8SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestInt8SliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt8SliceSupplier, v)
			}
		})
	}
}

func TestInt8SliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int8SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt8SliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt8SliceSupplierWithError,
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
				r.EqualError(err, errTestInt8SliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt8SliceSupplier, v)
			}
		})
	}
}

func TestInt8SliceSupplier_ToSilentInt8SliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int8SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt8SliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt8SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentInt8SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt8SliceSupplier, v)
			}
		})
	}
}

func TestInt8SliceSupplier_ToMustInt8SliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int8SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt8SliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt8SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt8SliceSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestInt8SliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt8SliceSupplier, v)
			}
		})
	}
}

func TestSilentInt8SliceSupplier(t *testing.T) {
	var ss SilentInt8SliceSupplier = func() []int8 {
		return resTestInt8SliceSupplier
	}
	v := ss()
	require.Equal(t, resTestInt8SliceSupplier, v)
}

func TestSilentInt8SliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int8SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt8SliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt8SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentInt8SliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt8SliceSupplier, v)
			}
		})
	}
}

func TestMustInt8SliceSupplier(t *testing.T) {
	var ms MustInt8SliceSupplier = func() []int8 {
		return resTestInt8SliceSupplier
	}

	v := ms()
	require.Equal(t, resTestInt8SliceSupplier, v)
}

func TestMustInt8SliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int8SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt8SliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt8SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustInt8SliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestInt8SliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt8SliceSupplier, v)
			}
		})
	}
}

func TestMustInt8SliceSupplier_ToSilentInt8SliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int8SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt8SliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt8SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt8SliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentInt8SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt8SliceSupplier, v)
			}
		})
	}
}

func TestMustInt8SliceSupplier_ToInt8SliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int8SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt8SliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt8SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt8SliceSupplier()
			r.NotNil(ms)

			s := ms.ToInt8SliceSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestInt8SliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt8SliceSupplier, v)
			}
		})
	}
}
