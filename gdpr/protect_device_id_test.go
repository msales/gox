package gdpr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/msales/gox/gdpr"
)

func Test_protector_ProtectDeviceID(t *testing.T) {
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
			value:              "some_value",
			wantProtectedValue: "some_val**",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			protectedValue := gdpr.ProtectDeviceID(tt.value)

			assert.Equal(t, tt.wantProtectedValue, protectedValue)
		})
	}
}
