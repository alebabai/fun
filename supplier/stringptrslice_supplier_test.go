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
	resTestStringPtrSliceSupplier []*string
	errTestStringPtrSliceSupplier = errors.New("error")
)

func testStringPtrSliceSupplier() ([]*string, error) {
	return resTestStringPtrSliceSupplier, nil
}

func testStringPtrSliceSupplierWithError() ([]*string, error) {
	return resTestStringPtrSliceSupplier, errTestStringPtrSliceSupplier
}

func TestStringPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       StringPtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testStringPtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testStringPtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestStringPtrSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestStringPtrSliceSupplier, v)
			}
		})
	}
}

func TestStringPtrSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       StringPtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testStringPtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testStringPtrSliceSupplierWithError,
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
				r.EqualError(err, errTestStringPtrSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestStringPtrSliceSupplier, v)
			}
		})
	}
}

func TestStringPtrSliceSupplier_ToSilentStringPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       StringPtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testStringPtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testStringPtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentStringPtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestStringPtrSliceSupplier, v)
			}
		})
	}
}

func TestStringPtrSliceSupplier_ToMustStringPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       StringPtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testStringPtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testStringPtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustStringPtrSliceSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestStringPtrSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestStringPtrSliceSupplier, v)
			}
		})
	}
}

func TestSilentStringPtrSliceSupplier(t *testing.T) {
	var ss SilentStringPtrSliceSupplier = func() []*string {
		return resTestStringPtrSliceSupplier
	}
	v := ss()
	require.Equal(t, resTestStringPtrSliceSupplier, v)
}

func TestSilentStringPtrSliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       StringPtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testStringPtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testStringPtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentStringPtrSliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestStringPtrSliceSupplier, v)
			}
		})
	}
}

func TestMustStringPtrSliceSupplier(t *testing.T) {
	var ms MustStringPtrSliceSupplier = func() []*string {
		return resTestStringPtrSliceSupplier
	}

	v := ms()
	require.Equal(t, resTestStringPtrSliceSupplier, v)
}

func TestMustStringPtrSliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       StringPtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testStringPtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testStringPtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustStringPtrSliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestStringPtrSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestStringPtrSliceSupplier, v)
			}
		})
	}
}

func TestMustStringPtrSliceSupplier_ToSilentStringPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       StringPtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testStringPtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testStringPtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustStringPtrSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentStringPtrSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestStringPtrSliceSupplier, v)
			}
		})
	}
}

func TestMustStringPtrSliceSupplier_ToStringPtrSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       StringPtrSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testStringPtrSliceSupplier,
		},
		{
			name:    "with error",
			s:       testStringPtrSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustStringPtrSliceSupplier()
			r.NotNil(ms)

			s := ms.ToStringPtrSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestStringPtrSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestStringPtrSliceSupplier, v)
			}
		})
	}
}
