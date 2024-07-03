package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/mo0Oonnn/password-generator/internal/http-server/handlers"
	"github.com/mo0Oonnn/password-generator/internal/lib/api/response"
	"github.com/mo0Oonnn/password-generator/internal/lib/logger/slogdiscard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateGeneratePasswordHandler(t *testing.T) {
	cases := []struct {
		name    string
		size    string
		respErr string
	}{
		{
			name: "Success",
			size: "5",
		},
		{
			name:    "Invalid size",
			size:    "3",
			respErr: "invalid size",
		},
		{
			name:    "Not a number",
			size:    "a",
			respErr: "invalid size",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			handler := handlers.CreateGeneratePasswordHandler(slogdiscard.NewDiscardLogger())
			r := chi.NewMux()
			r.Get("/password/{size}", handler)

			req, err := http.NewRequest(http.MethodGet, "/password/"+tc.size, nil)
			require.NoError(t, err)
			rr := httptest.NewRecorder()

			r.ServeHTTP(rr, req)

			assert.Equal(t, rr.Code, http.StatusOK)

			body := rr.Body.String()

			var resp response.Response
			assert.NoError(t, json.Unmarshal([]byte(body), &resp))
			assert.Equal(t, tc.respErr, resp.Error)
		})
	}
}
