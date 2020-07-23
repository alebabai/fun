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
	resTestBoolPtrSupplier *bool
	errTestBoolPtrSupplier = errors.New("error")
)

func testBoolPtrSupplier() (*bool, error) {
	return resTestBoolPtrSupplier, nil
}

func testBoolPtrSupplierWithError() (*bool, error) {
	return resTestBoolPtrSupplier, errTestBoolPtrSupplier
}

func TestBoolPtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BoolPtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBoolPtrSupplier,
		},
		{
			name:    "with error",
			s:       testBoolPtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestBoolPtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestBoolPtrSupplier, v)
			}
		})
	}
}

func TestBoolPtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BoolPtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBoolPtrSupplier,
		},
		{
			name:    "with error",
			s:       testBoolPtrSupplierWithError,
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
				r.EqualError(err, errTestBoolPtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestBoolPtrSupplier, v)
			}
		})
	}
}

func TestBoolPtrSupplier_ToSilentBoolPtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BoolPtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBoolPtrSupplier,
		},
		{
			name:    "with error",
			s:       testBoolPtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentBoolPtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestBoolPtrSupplier, v)
			}
		})
	}
}

func TestBoolPtrSupplier_ToMustBoolPtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BoolPtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBoolPtrSupplier,
		},
		{
			name:    "with error",
			s:       testBoolPtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustBoolPtrSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestBoolPtrSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestBoolPtrSupplier, v)
			}
		})
	}
}

func TestSilentBoolPtrSupplier(t *testing.T) {
	var ss SilentBoolPtrSupplier = func() *bool {
		return resTestBoolPtrSupplier
	}
	v := ss()
	require.Equal(t, resTestBoolPtrSupplier, v)
}

func TestSilentBoolPtrSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BoolPtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBoolPtrSupplier,
		},
		{
			name:    "with error",
			s:       testBoolPtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentBoolPtrSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestBoolPtrSupplier, v)
			}
		})
	}
}

func TestMustBoolPtrSupplier(t *testing.T) {
	var ms MustBoolPtrSupplier = func() *bool {
		return resTestBoolPtrSupplier
	}

	v := ms()
	require.Equal(t, resTestBoolPtrSupplier, v)
}

func TestMustBoolPtrSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BoolPtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBoolPtrSupplier,
		},
		{
			name:    "with error",
			s:       testBoolPtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustBoolPtrSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestBoolPtrSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestBoolPtrSupplier, v)
			}
		})
	}
}

func TestMustBoolPtrSupplier_ToSilentBoolPtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BoolPtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBoolPtrSupplier,
		},
		{
			name:    "with error",
			s:       testBoolPtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustBoolPtrSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentBoolPtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestBoolPtrSupplier, v)
			}
		})
	}
}

func TestMustBoolPtrSupplier_ToBoolPtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BoolPtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBoolPtrSupplier,
		},
		{
			name:    "with error",
			s:       testBoolPtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustBoolPtrSupplier()
			r.NotNil(ms)

			s := ms.ToBoolPtrSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestBoolPtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestBoolPtrSupplier, v)
			}
		})
	}
}
