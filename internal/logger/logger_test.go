// Singleton logger
package logger

import (
	"net/http"
	"reflect"
	"testing"
)

func TestInitialize(t *testing.T) {
	t.Logf("TODO: Add test cases.")

	type args struct {
		level string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Initialize(tt.args.level); (err != nil) != tt.wantErr {
				t.Errorf("Initialize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRequestLogger(t *testing.T) {
	t.Logf("TODO: Add test cases.")

	type args struct {
		h http.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RequestLogger(tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequestLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}
