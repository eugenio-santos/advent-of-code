package main

import (
	"os"
	"testing"
)

func TestD01P1(t *testing.T) {
	result := d01p1()
	if result != 1938424 {
		t.Error(result)
	}
}

func TestD01P2(t *testing.T) {
	result := d01p2()
	if result != 22014209 {
		t.Error(result)
	}
}

func TestD02P1(t *testing.T) {
	result := d02p1()
	if result != 510 {
		t.Error(result)
	}
}

func TestD02P2(t *testing.T) {
	result := d02p2()
	if result != 553 {
		t.Error(result)
	}
}

func TestD03P1(t *testing.T) {
	f, err := os.ReadFile("inputs/d03")
	if err != nil {
		panic(err)
	}

	tests := []struct {
		id       string
		input    string
		expected int
	}{
		{"example", "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))", 161},
		{"example", string(f), 188741603},
	}

	for _, tt := range tests {
		result := d03p1(tt.input)
		if tt.expected != result {
			t.Error(tt.id, " - expected ", tt.expected, " got ", result)
		}
	}
}

func TestD03P2(t *testing.T) {
	f, err := os.ReadFile("inputs/d03")
	if err != nil {
		panic(err)
	}

	tests := []struct {
		id       string
		input    string
		expected int
	}{
		{"example", "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))", 48},
		{"p2", string(f), 67269798},
	}

	for _, tt := range tests {
		result := d03p2(tt.input)
		if tt.expected != result {
			t.Error(tt.id, " - expected ", tt.expected, " got ", result)
		}
	}
}

func TestD04P1(t *testing.T) {
	tests := []struct {
		id       string
		file     string
		expected int
	}{
		{"example", "inputs/d04test", 18},
		{"p1", "inputs/d04", 2567},
	}
	for _, tt := range tests {
		lines := lines(tt.file)
		m := [][]rune{}
		for _, l := range lines {
			m = append(m, []rune(l))
		}

		if result := d04p1(m); result != tt.expected {
			t.Error("expected", tt.expected, "result", result)
		}
	}
}

func TestD04P2(t *testing.T) {
	tests := []struct {
		id       string
		file     string
		expected int
	}{
		{"example", "inputs/d04test", 9},
		{"p1", "inputs/d04", 2029},
	}
	for _, tt := range tests {
		lines := lines(tt.file)
		m := [][]rune{}
		for _, l := range lines {
			m = append(m, []rune(l))
		}

		if result := d04p2(m); result != tt.expected {
			t.Error("expected", tt.expected, "result", result)
		}
	}
}

func TestD05P1(t *testing.T) {
	tests := []struct {
		id       string
		file     string
		expected int
	}{
		{"example", "inputs/d05t", 143},
		{"p1", "inputs/d05", 4135},
	}
	for _, tt := range tests {
		if result, _, _ := d05p1(tt.file); result != tt.expected {
			t.Error("expected", tt.expected, "result", result)
		}
	}
}

func TestD05P2(t *testing.T) {
	tests := []struct {
		id       string
		file     string
		expected int
	}{
		{"example", "inputs/d05t", 123},
		{"p1", "inputs/d05", 5285},
	}
	for _, tt := range tests {
		if result := d05p2(tt.file); result != tt.expected {
			t.Error("expected", tt.expected, "result", result)
		}
	}
}
