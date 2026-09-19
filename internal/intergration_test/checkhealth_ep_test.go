package intergration_test

import (
	"net/http/httptest"
	"testing"

	"github.com/qtheflash/bookmark-management/internal/api"
	"github.com/qtheflash/bookmark-management/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestCheckHealthEndpoint(t *testing.T) {
	testcases := []struct {
		name                 string
		url                  func(api api.Engine) *httptest.ResponseRecorder
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "TC01: Check health endpoint",
			url: func(api api.Engine) *httptest.ResponseRecorder {
				request := httptest.NewRequest("GET", "/health-check", nil)
				resRecorder := httptest.NewRecorder()

				api.ServeHTTP(resRecorder, request)
				return resRecorder
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"message":"OK","service_name":"bookmark-api","instance_id":"my-instance-id"}`,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			cfg := &config.Config{
				ServiceName: "bookmark-api",
				InstanceID:  "my-instance-id",
			}
			app := api.NewEngine(cfg)

			resRecorder := tc.url(app)

			assert.Equal(t, tc.expectedStatusCode, resRecorder.Code)
			assert.Contains(t, resRecorder.Body.String(), tc.expectedResponseBody)

		})
	}
}
