package main

import (
	"reflect"
	"testing"

	"../../pkg/mock"
)

func Test_run(t *testing.T) {
	var s mock.SiteScanner
	type args struct {
		s     Scanner
		urls  []string
		depth int
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "first",
			args: args{
				s:     s,
				urls:  []string{"fake"},
				depth: 2,
			},
			want: map[string]string{
				"https://go.dev/":     "go.dev",
				"https://golang.org/": "The Go Programming Language",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := run(tt.args.s, tt.args.urls, tt.args.depth)
			if (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
