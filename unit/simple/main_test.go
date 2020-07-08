package main

import "testing"

func Test_hello(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		expected := "Hello World"
		got := hello("World")
		if got != expected {
			t.Errorf("Expected: %v got: %v", expected, got)
		}
	})
}

func Test_helloMultiples(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name     string
		args     args
		want     string
		wantCtrl bool
	}{
		{
			name: "hello world",
			args: args{
				s: "world",
			},
			want:     "Hello world",
			wantCtrl: true,
		},
		{
			name: "control test",
			args: args{
				s: "you",
			},
			want:     "Hello world",
			wantCtrl: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hello(tt.args.s); (got != tt.want) == tt.wantCtrl {
				t.Errorf("hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
