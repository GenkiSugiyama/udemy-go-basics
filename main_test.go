package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "1 + 2 の結果として 3 が返ること",
			args: args{x: 1, y: 2},
			want: 3,
		},
		{
			name: "2 + 3 の結果として 5 が返ること",
			args: args{x: 2, y: 3},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		// TODO: Add test cases.
		{
			name: "1 / 4 で 0.25 が返ること",
			args: args{x: 1, y: 4},
			want: 0.25,
		},
		{
			name: "yが 0 の場合は 0 が返ること",
			args: args{x: 1, y: 0},
			want: 0.,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Divide(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
