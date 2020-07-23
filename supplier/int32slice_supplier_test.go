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
	resTestInt32SliceSupplier []int32
	errTestInt32SliceSupplier = errors.New("error")
)

func testInt32SliceSupplier() ([]int32, error) {
	return resTestInt32SliceSupplier, nil
}

func testInt32SliceSupplierWithError() ([]int32, error) {
	return resTestInt32SliceSupplier, errTestInt32SliceSupplier
}

func TestInt32SliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int32SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt32SliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt32SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestInt32SliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt32SliceSupplier, v)
			}
		})
	}
}

func TestInt32SliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int32SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt32SliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt32SliceSupplierWithError,
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
				r.EqualError(err, errTestInt32SliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt32SliceSupplier, v)
			}
		})
	}
}

func TestInt32SliceSupplier_ToSilentInt32SliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int32SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt32SliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt32SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentInt32SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt32SliceSupplier, v)
			}
		})
	}
}

func TestInt32SliceSupplier_ToMustInt32SliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int32SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt32SliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt32SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt32SliceSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestInt32SliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt32SliceSupplier, v)
			}
		})
	}
}

func TestSilentInt32SliceSupplier(t *testing.T) {
	var ss SilentInt32SliceSupplier = func() []int32 {
		return resTestInt32SliceSupplier
	}
	v := ss()
	require.Equal(t, resTestInt32SliceSupplier, v)
}

func TestSilentInt32SliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int32SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt32SliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt32SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentInt32SliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt32SliceSupplier, v)
			}
		})
	}
}

func TestMustInt32SliceSupplier(t *testing.T) {
	var ms MustInt32SliceSupplier = func() []int32 {
		return resTestInt32SliceSupplier
	}

	v := ms()
	require.Equal(t, resTestInt32SliceSupplier, v)
}

func TestMustInt32SliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int32SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt32SliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt32SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustInt32SliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestInt32SliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestInt32SliceSupplier, v)
			}
		})
	}
}

func TestMustInt32SliceSupplier_ToSilentInt32SliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int32SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt32SliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt32SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt32SliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentInt32SliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestInt32SliceSupplier, v)
			}
		})
	}
}

func TestMustInt32SliceSupplier_ToInt32SliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       Int32SliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testInt32SliceSupplier,
		},
		{
			name:    "with error",
			s:       testInt32SliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustInt32SliceSupplier()
			r.NotNil(ms)

			s := ms.ToInt32SliceSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestInt32SliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestInt32SliceSupplier, v)
			}
		})
	}
}
