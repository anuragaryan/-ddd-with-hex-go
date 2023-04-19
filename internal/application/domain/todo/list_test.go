package todo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		listName string
		wantErr  string
	}{
		{
			name:     "a list is created successfully",
			listName: "complete ddd-with-hex codebase",
		},
		{
			name:     "a list without name is invalid",
			listName: "",
			wantErr:  "list needs a name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewList(tt.listName)
			if err != nil {
				assert.EqualError(t, err, tt.wantErr)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.listName, got.Name)
		})
	}
}
