package todo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewItem(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		text       string
		wantStatus Status
		wantErr    string
	}{
		{
			name:       "an item is created successfully",
			text:       "complete testing",
			wantStatus: "initialised",
		},
		{
			name:    "empty item is considered invalid",
			text:    "",
			wantErr: "missing values",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewItem(tt.text)
			if err != nil {
				assert.EqualError(t, err, tt.wantErr)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.wantStatus, got.Status)
			assert.Equal(t, tt.text, got.Text)
		})
	}
}
