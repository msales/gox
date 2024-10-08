package gdpr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/msales/gox/gdpr"
)

func Test_protector_ProtectRawIP(t *testing.T) {
	tests := []struct {
		name               string
		value              string
		wantProtectedValue string
	}{
		{
			name:               "Empty value",
			value:              "",
			wantProtectedValue: "",
		},
		{
			name:               "Value to protect",
			value:              "42.21.37.213",
			wantProtectedValue: "42.21.37.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			protectedValue := gdpr.ProtectRawIP(tt.value)

			assert.Equal(t, tt.wantProtectedValue, protectedValue)
		})
	}
}

func Test_protector_ProtectIPV4(t *testing.T) {
	tests := []struct {
		name               string
		value              uint32
		wantProtectedValue string
	}{
		{
			name:               "Empty value",
			value:              0,
			wantProtectedValue: "0.0.0.0",
		},
		{
			name:               "Value to protect",
			value:              706029013, // 42.21.37.213
			wantProtectedValue: "42.21.37.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			protectedValue := gdpr.ProtectIPV4(tt.value)

			assert.Equal(t, tt.wantProtectedValue, protectedValue)
		})
	}
}
