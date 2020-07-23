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
	resTestBoolSupplier bool
	errTestBoolSupplier = errors.New("error")
)

func testBoolSupplier() (bool, error) {
	return resTestBoolSupplier, nil
}

func testBoolSupplierWithError() (bool, error) {
	return resTestBoolSupplier, errTestBoolSupplier
}

func TestBoolSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BoolSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBoolSupplier,
		},
		{
			name:    "with error",
			s:       testBoolSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestBoolSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestBoolSupplier, v)
			}
		})
	}
}

func TestBoolSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BoolSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBoolSupplier,
		},
		{
			name:    "with error",
			s:       testBoolSupplierWithError,
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
				r.EqualError(err, errTestBoolSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestBoolSupplier, v)
			}
		})
	}
}

func TestBoolSupplier_ToSilentBoolSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BoolSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBoolSupplier,
		},
		{
			name:    "with error",
			s:       testBoolSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentBoolSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestBoolSupplier, v)
			}
		})
	}
}

func TestBoolSupplier_ToMustBoolSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BoolSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBoolSupplier,
		},
		{
			name:    "with error",
			s:       testBoolSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustBoolSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestBoolSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestBoolSupplier, v)
			}
		})
	}
}

func TestSilentBoolSupplier(t *testing.T) {
	var ss SilentBoolSupplier = func() bool {
		return resTestBoolSupplier
	}
	v := ss()
	require.Equal(t, resTestBoolSupplier, v)
}

func TestSilentBoolSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BoolSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBoolSupplier,
		},
		{
			name:    "with error",
			s:       testBoolSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentBoolSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestBoolSupplier, v)
			}
		})
	}
}

func TestMustBoolSupplier(t *testing.T) {
	var ms MustBoolSupplier = func() bool {
		return resTestBoolSupplier
	}

	v := ms()
	require.Equal(t, resTestBoolSupplier, v)
}

func TestMustBoolSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BoolSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBoolSupplier,
		},
		{
			name:    "with error",
			s:       testBoolSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustBoolSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestBoolSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestBoolSupplier, v)
			}
		})
	}
}

func TestMustBoolSupplier_ToSilentBoolSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BoolSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBoolSupplier,
		},
		{
			name:    "with error",
			s:       testBoolSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustBoolSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentBoolSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestBoolSupplier, v)
			}
		})
	}
}

func TestMustBoolSupplier_ToBoolSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BoolSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBoolSupplier,
		},
		{
			name:    "with error",
			s:       testBoolSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustBoolSupplier()
			r.NotNil(ms)

			s := ms.ToBoolSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestBoolSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestBoolSupplier, v)
			}
		})
	}
}
