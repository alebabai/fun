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
	resTestUintSupplier uint
	errTestUintSupplier = errors.New("error")
)

func testUintSupplier() (uint, error) {
	return resTestUintSupplier, nil
}

func testUintSupplierWithError() (uint, error) {
	return resTestUintSupplier, errTestUintSupplier
}

func TestUintSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       UintSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUintSupplier,
		},
		{
			name:    "with error",
			s:       testUintSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestUintSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUintSupplier, v)
			}
		})
	}
}

func TestUintSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       UintSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUintSupplier,
		},
		{
			name:    "with error",
			s:       testUintSupplierWithError,
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
				r.EqualError(err, errTestUintSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUintSupplier, v)
			}
		})
	}
}

func TestUintSupplier_ToSilentUintSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       UintSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUintSupplier,
		},
		{
			name:    "with error",
			s:       testUintSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentUintSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUintSupplier, v)
			}
		})
	}
}

func TestUintSupplier_ToMustUintSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       UintSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUintSupplier,
		},
		{
			name:    "with error",
			s:       testUintSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUintSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestUintSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUintSupplier, v)
			}
		})
	}
}

func TestSilentUintSupplier(t *testing.T) {
	var ss SilentUintSupplier = func() uint {
		return resTestUintSupplier
	}
	v := ss()
	require.Equal(t, resTestUintSupplier, v)
}

func TestSilentUintSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       UintSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUintSupplier,
		},
		{
			name:    "with error",
			s:       testUintSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentUintSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUintSupplier, v)
			}
		})
	}
}

func TestMustUintSupplier(t *testing.T) {
	var ms MustUintSupplier = func() uint {
		return resTestUintSupplier
	}

	v := ms()
	require.Equal(t, resTestUintSupplier, v)
}

func TestMustUintSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       UintSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUintSupplier,
		},
		{
			name:    "with error",
			s:       testUintSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustUintSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestUintSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestUintSupplier, v)
			}
		})
	}
}

func TestMustUintSupplier_ToSilentUintSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       UintSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUintSupplier,
		},
		{
			name:    "with error",
			s:       testUintSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUintSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentUintSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestUintSupplier, v)
			}
		})
	}
}

func TestMustUintSupplier_ToUintSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       UintSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testUintSupplier,
		},
		{
			name:    "with error",
			s:       testUintSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustUintSupplier()
			r.NotNil(ms)

			s := ms.ToUintSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestUintSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestUintSupplier, v)
			}
		})
	}
}
