package helper

import (
	"testing"

	"github.com/auctionee/core/internal/core/data"
)

func TestGetAUID(t *testing.T) {
	tests := []struct {
		name string
		want data.AUID
	}{
		{"", ""},
		{"", ""},
		{"", ""},
		{"", ""},
		{"", ""},
		{"", ""},
		{"", ""},
		{"", ""},
		{"", ""},
		{"", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetAUID()
			if len(got) != 6 {
				t.Errorf("GetAUID() = %v, want %v", got, tt.want)
			}
		})
	}
}
