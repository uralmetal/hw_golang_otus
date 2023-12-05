package main

import (
	"bytes"
	"testing"
)

func TestSayHelloOtus(t *testing.T) {
	tests := []struct {
		name  string
		wantW string
	}{
		{"Simple test", "!SUTO ,olleH\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			SayHelloOtus(w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("SayHelloOtus() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
