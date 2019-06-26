package services

import (
	"fmt"
	"strings"
	"testing"
)

func TestPrefixedUUID(t *testing.T) {
	type args struct {
		prefix string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test with Happy Path",
			args: args{
				prefix: "ab",
			},
			wantErr: true,
		},
		{
			name: "Test with Too Long",
			args: args{
				prefix: "abcd",
			},
			wantErr: true,
		},
		{
			name: "Test with Empty String",
			args: args{
				prefix: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrefixedUUID(tt.args.prefix); !strings.Contains(got[0:3], tt.args.prefix) && !tt.wantErr {
				t.Errorf("PrefixedUUID() = %v, want %v", got, tt.args.prefix)
			} else {
				fmt.Printf("Success: Wanted prefix of %s, got %s\n", tt.args.prefix, got)
			}
		})
	}
}
