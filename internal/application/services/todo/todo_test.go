package todo

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mockports "github.com/anuragaryan/ddd-with-hex-go/internal/ports/mocks"
)

func TestService_CreateList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		nameList string
		mocks    func(s *mockports.MockStoragePort, e *mockports.MockEventHandlerPort)
		wantErr  string
	}{
		{
			name:     "the list is created successfully",
			nameList: "my most important list",
			mocks: func(s *mockports.MockStoragePort, e *mockports.MockEventHandlerPort) {
				s.EXPECT().Add(gomock.Any()).Return(nil).Times(1)
				e.EXPECT().Notify(gomock.Any()).Return(nil).Times(1)
			},
			wantErr: "",
		},
		{
			name:     "storage returns error",
			nameList: "my most important list",
			mocks: func(s *mockports.MockStoragePort, e *mockports.MockEventHandlerPort) {
				s.EXPECT().Add(gomock.Any()).Return(errors.New("can't reach mysql")).Times(1)
			},
			wantErr: "can't reach mysql",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			//TALK: Dependency injection.
			storage := mockports.NewMockStoragePort(ctrl)
			eventHandler := mockports.NewMockEventHandlerPort(ctrl)
			tt.mocks(storage, eventHandler)

			s, err := NewService(
				withRepository(storage),
				WithEventsHandlers(eventHandler),
			)
			assert.NoError(t, err)

			err = s.CreateList(tt.nameList)
			if err != nil {
				assert.EqualError(t, err, tt.wantErr)
				return
			}

			assert.NoError(t, err)
		})
	}
}
