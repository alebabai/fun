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
	resTestRunePtrSupplier *rune
	errTestRunePtrSupplier = errors.New("error")
)

func testRunePtrSupplier() (*rune, error) {
	return resTestRunePtrSupplier, nil
}

func testRunePtrSupplierWithError() (*rune, error) {
	return resTestRunePtrSupplier, errTestRunePtrSupplier
}

func TestRunePtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       RunePtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testRunePtrSupplier,
		},
		{
			name:    "with error",
			s:       testRunePtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestRunePtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestRunePtrSupplier, v)
			}
		})
	}
}

func TestRunePtrSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       RunePtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testRunePtrSupplier,
		},
		{
			name:    "with error",
			s:       testRunePtrSupplierWithError,
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
				r.EqualError(err, errTestRunePtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestRunePtrSupplier, v)
			}
		})
	}
}

func TestRunePtrSupplier_ToSilentRunePtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       RunePtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testRunePtrSupplier,
		},
		{
			name:    "with error",
			s:       testRunePtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentRunePtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestRunePtrSupplier, v)
			}
		})
	}
}

func TestRunePtrSupplier_ToMustRunePtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       RunePtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testRunePtrSupplier,
		},
		{
			name:    "with error",
			s:       testRunePtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustRunePtrSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestRunePtrSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestRunePtrSupplier, v)
			}
		})
	}
}

func TestSilentRunePtrSupplier(t *testing.T) {
	var ss SilentRunePtrSupplier = func() *rune {
		return resTestRunePtrSupplier
	}
	v := ss()
	require.Equal(t, resTestRunePtrSupplier, v)
}

func TestSilentRunePtrSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       RunePtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testRunePtrSupplier,
		},
		{
			name:    "with error",
			s:       testRunePtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentRunePtrSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestRunePtrSupplier, v)
			}
		})
	}
}

func TestMustRunePtrSupplier(t *testing.T) {
	var ms MustRunePtrSupplier = func() *rune {
		return resTestRunePtrSupplier
	}

	v := ms()
	require.Equal(t, resTestRunePtrSupplier, v)
}

func TestMustRunePtrSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       RunePtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testRunePtrSupplier,
		},
		{
			name:    "with error",
			s:       testRunePtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustRunePtrSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestRunePtrSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestRunePtrSupplier, v)
			}
		})
	}
}

func TestMustRunePtrSupplier_ToSilentRunePtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       RunePtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testRunePtrSupplier,
		},
		{
			name:    "with error",
			s:       testRunePtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustRunePtrSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentRunePtrSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestRunePtrSupplier, v)
			}
		})
	}
}

func TestMustRunePtrSupplier_ToRunePtrSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       RunePtrSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testRunePtrSupplier,
		},
		{
			name:    "with error",
			s:       testRunePtrSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustRunePtrSupplier()
			r.NotNil(ms)

			s := ms.ToRunePtrSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestRunePtrSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestRunePtrSupplier, v)
			}
		})
	}
}
