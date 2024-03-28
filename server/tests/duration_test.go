package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/jasvtfvan/oms-admin/server/utils"
)

func TestParseDuration(t *testing.T) {
	type args struct {
		d string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Duration
		wantErr bool
	}{
		{
			name:    "5h20m",
			args:    args{"5h20m"},
			want:    time.Hour*5 + 20*time.Minute,
			wantErr: false,
		},
		{
			name:    "1d5h20m",
			args:    args{"1d5h20m"},
			want:    24*time.Hour + time.Hour*5 + 20*time.Minute,
			wantErr: false,
		},
		{
			name:    "1d",
			args:    args{"1d"},
			want:    24 * time.Hour,
			wantErr: false,
		},
		{
			name:    "1d4h1m3s",
			args:    args{"1d4h1m3s"},
			want:    24*time.Hour + time.Hour*4 + 1*time.Minute + 3*time.Second,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := utils.ParseDuration(tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDuration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseDuration() got = %v, want %v", got, tt.want)
			}
		})
	}
	buf, _ := utils.ParseDuration("1d")
	fmt.Println("**********************")
	fmt.Println(int64(buf / time.Second))
	fmt.Println(buf)
	fmt.Println("**********************")
}
