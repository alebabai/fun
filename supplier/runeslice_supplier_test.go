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
	resTestRuneSliceSupplier []rune
	errTestRuneSliceSupplier = errors.New("error")
)

func testRuneSliceSupplier() ([]rune, error) {
	return resTestRuneSliceSupplier, nil
}

func testRuneSliceSupplierWithError() ([]rune, error) {
	return resTestRuneSliceSupplier, errTestRuneSliceSupplier
}

func TestRuneSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       RuneSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testRuneSliceSupplier,
		},
		{
			name:    "with error",
			s:       testRuneSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			v, err := tt.s()
			if tt.wantErr {
				r.EqualError(err, errTestRuneSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestRuneSliceSupplier, v)
			}
		})
	}
}

func TestRuneSliceSupplier_ToSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       RuneSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testRuneSliceSupplier,
		},
		{
			name:    "with error",
			s:       testRuneSliceSupplierWithError,
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
				r.EqualError(err, errTestRuneSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestRuneSliceSupplier, v)
			}
		})
	}
}

func TestRuneSliceSupplier_ToSilentRuneSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       RuneSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testRuneSliceSupplier,
		},
		{
			name:    "with error",
			s:       testRuneSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ss := tt.s.ToSilentRuneSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestRuneSliceSupplier, v)
			}
		})
	}
}

func TestRuneSliceSupplier_ToMustRuneSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       RuneSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testRuneSliceSupplier,
		},
		{
			name:    "with error",
			s:       testRuneSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustRuneSliceSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestRuneSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestRuneSliceSupplier, v)
			}
		})
	}
}

func TestSilentRuneSliceSupplier(t *testing.T) {
	var ss SilentRuneSliceSupplier = func() []rune {
		return resTestRuneSliceSupplier
	}
	v := ss()
	require.Equal(t, resTestRuneSliceSupplier, v)
}

func TestSilentRuneSliceSupplier_ToSilentSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       RuneSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testRuneSliceSupplier,
		},
		{
			name:    "with error",
			s:       testRuneSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tss := tt.s.ToSilentRuneSliceSupplier()
			r.NotNil(tss)

			ss := tss.ToSilentSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestRuneSliceSupplier, v)
			}
		})
	}
}

func TestMustRuneSliceSupplier(t *testing.T) {
	var ms MustRuneSliceSupplier = func() []rune {
		return resTestRuneSliceSupplier
	}

	v := ms()
	require.Equal(t, resTestRuneSliceSupplier, v)
}

func TestMustRuneSliceSupplier_ToMustSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       RuneSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testRuneSliceSupplier,
		},
		{
			name:    "with error",
			s:       testRuneSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tms := tt.s.ToMustRuneSliceSupplier()
			r.NotNil(tms)

			ms := tms.ToMustSupplier()
			r.NotNil(ms)

			if tt.wantErr {
				r.PanicsWithError(errTestRuneSliceSupplier.Error(), func() {
					v := ms()
					r.Empty(v)
				})
			} else {
				v := ms()
				r.Equal(resTestRuneSliceSupplier, v)
			}
		})
	}
}

func TestMustRuneSliceSupplier_ToSilentRuneSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       RuneSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testRuneSliceSupplier,
		},
		{
			name:    "with error",
			s:       testRuneSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustRuneSliceSupplier()
			r.NotNil(ms)

			ss := ms.ToSilentRuneSliceSupplier()
			r.NotNil(ss)

			v := ss()
			if tt.wantErr {
				r.Empty(v)
			} else {
				r.Equal(resTestRuneSliceSupplier, v)
			}
		})
	}
}

func TestMustRuneSliceSupplier_ToRuneSliceSupplier(t *testing.T) {
	tests := []struct {
		name    string
		s       RuneSliceSupplier
		wantErr bool
	}{
		{
			name: "ok",
			s:    testRuneSliceSupplier,
		},
		{
			name:    "with error",
			s:       testRuneSliceSupplierWithError,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			ms := tt.s.ToMustRuneSliceSupplier()
			r.NotNil(ms)

			s := ms.ToRuneSliceSupplier()
			r.NotNil(s)

			v, err := s()
			if tt.wantErr {
				r.EqualError(err, errTestRuneSliceSupplier.Error())
				r.Empty(v)
			} else {
				r.Equal(resTestRuneSliceSupplier, v)
			}
		})
	}
}
