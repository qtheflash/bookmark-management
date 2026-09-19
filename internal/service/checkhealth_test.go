package service

import (
	"testing"

	"github.com/google/uuid"
	"github.com/qtheflash/bookmark-management/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestNewCheckHealth(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name       string
		cfg        *config.Config
		service    string
		instanceID string
		isUUID     bool
	}{
		{
			name: "TC01: Custom InstanceID in config",
			cfg: &config.Config{
				ServiceName: "bookmark-api",
				InstanceID:  "my-custom-id",
			},
			service:    "bookmark-api",
			instanceID: "my-custom-id",
			isUUID:     false,
		},
		{
			name: "TC02: Empty InstanceID generates UUID",
			cfg: &config.Config{
				ServiceName: "bookmark-api",
				InstanceID:  "",
			},
			service:    "bookmark-api",
			instanceID: "",
			isUUID:     true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			serv := NewCheckHealth(tc.cfg)
			assert.NotNil(t, serv)

			res := serv.CheckHealthImpl()

			assert.Equal(t, "OK", res.Message)
			assert.Equal(t, tc.service, res.ServiceName)
			if tc.isUUID {
				assert.NotEmpty(t, res.InstanceID)
				_, err := uuid.Parse(res.InstanceID)
				assert.NoError(t, err, "InstanceID phải là một UUID hợp lệ")
			} else {
				assert.Equal(t, tc.instanceID, res.InstanceID)
			}
		})
	}
}
