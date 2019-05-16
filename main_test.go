package main

import (
	"testing"
)

func TestPilha(t *testing.T) {
	p := newPilha(10)
	if !p.isEmpty() {
		t.FailNow()
	}
	for i := 0; i < 5; i++ {
		p.push(i)
		if p.isFull() {
			t.FailNow()
		}
	}

	if p.isEmpty() {
		t.FailNow()
	}
}

func TestPilhaPop(t *testing.T) {
	p := newPilha(10)
	if !p.isEmpty() {
		t.FailNow()
	}
	for i := 0; i < 5; i++ {
		p.push(i)
		if p.isFull() {
			t.FailNow()
		}
	}

	for i := 0; i < 5; i++ {
		p.pop()
	}
	if p.isEmpty() {
		t.FailNow()
	}
}

func Test_convert(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		want    string
		wantErr bool
	}{
		{"penis case", 13, "1101", false},
		{"", 2, "10", false},
		{"not penis case", 123457, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convert(tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
