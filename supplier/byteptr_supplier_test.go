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
	resTestBytePtrSupplier *byte
	errTestBytePtrSupplier = errors.New("error")
)

func testBytePtrSupplier() (*byte, error) {
	return resTestBytePtrSupplier, nil
}

func testBytePtrSupplierWithError() (*byte, error) {
	return resTestBytePtrSupplier, errTestBytePtrSupplier
}

func TestBytePtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BytePtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBytePtrSupplier,
		},
		{
			name:    "with error",
			s:       testBytePtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestBytePtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestBytePtrSupplier, v)
			}
		})
	}
}

func TestBytePtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BytePtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBytePtrSupplier,
		},
		{
			name:    "with error",
			s:       testBytePtrSupplierWithError,
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
				r.EqualError(err, errTestBytePtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestBytePtrSupplier, v)
			}
		})
	}
}

func TestBytePtrSupplier_ToSilentBytePtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BytePtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBytePtrSupplier,
		},
		{
			name:    "with error",
			s:       testBytePtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentBytePtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestBytePtrSupplier, v)
			}
		})
	}
}

func TestBytePtrSupplier_ToMustBytePtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BytePtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBytePtrSupplier,
		},
		{
			name:    "with error",
			s:       testBytePtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustBytePtrSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestBytePtrSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestBytePtrSupplier, v)
			}
		})
	}
}

func TestSilentBytePtrSupplier(t *testing.T) {
	var ss SilentBytePtrSupplier = func() *byte {
		return resTestBytePtrSupplier
	}
	v := ss()
	require.Equal(t, resTestBytePtrSupplier, v)
}

func TestSilentBytePtrSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BytePtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBytePtrSupplier,
		},
		{
			name:    "with error",
			s:       testBytePtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentBytePtrSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestBytePtrSupplier, v)
			}
		})
	}
}

func TestMustBytePtrSupplier(t *testing.T) {
	var ms MustBytePtrSupplier = func() *byte {
		return resTestBytePtrSupplier
	}

	v := ms()
	require.Equal(t, resTestBytePtrSupplier, v)
}

func TestMustBytePtrSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BytePtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBytePtrSupplier,
		},
		{
			name:    "with error",
			s:       testBytePtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustBytePtrSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestBytePtrSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestBytePtrSupplier, v)
			}
		})
	}
}

func TestMustBytePtrSupplier_ToSilentBytePtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BytePtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBytePtrSupplier,
		},
		{
			name:    "with error",
			s:       testBytePtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustBytePtrSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentBytePtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestBytePtrSupplier, v)
			}
		})
	}
}

func TestMustBytePtrSupplier_ToBytePtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       BytePtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testBytePtrSupplier,
		},
		{
			name:    "with error",
			s:       testBytePtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustBytePtrSupplier()
			r.NotNil(ms)

			s := ms.ToBytePtrSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestBytePtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestBytePtrSupplier, v)
			}
		})
	}
}
