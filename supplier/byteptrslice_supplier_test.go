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
	resTestBytePtrSliceSupplier []*byte
	errTestBytePtrSliceSupplier = errors.New("error")
)

func testBytePtrSliceSupplier() ([]*byte, error) {
	return resTestBytePtrSliceSupplier, nil
}

func testBytePtrSliceSupplierWithError() ([]*byte, error) {
	return resTestBytePtrSliceSupplier, errTestBytePtrSliceSupplier
}

func TestBytePtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BytePtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBytePtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testBytePtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestBytePtrSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestBytePtrSliceSupplier, v)
			}
		})
	}
}

func TestBytePtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BytePtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBytePtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testBytePtrSliceSupplierWithError,
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
				r.EqualError(err, errTestBytePtrSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestBytePtrSliceSupplier, v)
			}
		})
	}
}

func TestBytePtrSliceSupplier_ToSilentBytePtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BytePtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBytePtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testBytePtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentBytePtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestBytePtrSliceSupplier, v)
			}
		})
	}
}

func TestBytePtrSliceSupplier_ToMustBytePtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BytePtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBytePtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testBytePtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustBytePtrSliceSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestBytePtrSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestBytePtrSliceSupplier, v)
			}
		})
	}
}

func TestSilentBytePtrSliceSupplier(t *testing.T) {
	var ss SilentBytePtrSliceSupplier = func() []*byte {
		return resTestBytePtrSliceSupplier
	}
	v := ss()
	require.Equal(t, resTestBytePtrSliceSupplier, v)
}

func TestSilentBytePtrSliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BytePtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBytePtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testBytePtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentBytePtrSliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestBytePtrSliceSupplier, v)
			}
		})
	}
}

func TestMustBytePtrSliceSupplier(t *testing.T) {
	var ms MustBytePtrSliceSupplier = func() []*byte {
		return resTestBytePtrSliceSupplier
	}

	v := ms()
	require.Equal(t, resTestBytePtrSliceSupplier, v)
}

func TestMustBytePtrSliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BytePtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBytePtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testBytePtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustBytePtrSliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestBytePtrSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestBytePtrSliceSupplier, v)
			}
		})
	}
}

func TestMustBytePtrSliceSupplier_ToSilentBytePtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BytePtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBytePtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testBytePtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustBytePtrSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentBytePtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestBytePtrSliceSupplier, v)
			}
		})
	}
}

func TestMustBytePtrSliceSupplier_ToBytePtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BytePtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBytePtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testBytePtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustBytePtrSliceSupplier()
			r.NotNil(ms)

			s := ms.ToBytePtrSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestBytePtrSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestBytePtrSliceSupplier, v)
			}
		})
	}
}
