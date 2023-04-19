package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mockports "github.com/anuragaryan/ddd-with-hex-go/internal/ports/mocks"
)

func TestHandler_CreateList(t *testing.T) {
	tests := []struct {
		name                 string
		requestBody          createListRequest
		mocks                func(a *mockports.MockAPIPort)
		expectedResponseCode int
	}{
		{
			name: "a list is created successfully",
			requestBody: createListRequest{
				Name: "my most important list",
			},
			mocks: func(a *mockports.MockAPIPort) {
				a.EXPECT().CreateList("my most important list").Return(nil).Times(1)
			},
			expectedResponseCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			api := mockports.NewMockAPIPort(ctrl)
			tt.mocks(api)

			h, err := NewHandler(WithService(api))
			assert.NoError(t, err)

			r := chi.NewRouter()
			r.Post("/todo-list", h.CreateList)

			recorder := httptest.NewRecorder()

			bb, err := json.Marshal(tt.requestBody)
			assert.NoError(t, err)

			body := bytes.NewBuffer(bb)

			req := httptest.NewRequest("POST", "/todo-list", body)
			r.ServeHTTP(recorder, req)

			assert.Equal(t, tt.expectedResponseCode, recorder.Code)
		})
	}
}
