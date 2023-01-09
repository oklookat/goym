package goym

import "testing"

func TestI2s(t *testing.T) {
	var i int = 1234
	var i64 int64 = 531135531
	var i32 int32 = 15135135
	var ui32 uint32 = 1551335113

	if res := i2s(i); res != "1234" {
		t.Fatalf("expected: %v, got: %v", i, res)
	}
	if res := i2s(i64); res != "531135531" {
		t.Fatalf("expected: %v, got: %v", i, res)
	}
	if res := i2s(i32); res != "15135135" {
		t.Fatalf("expected: %v, got: %v", i, res)
	}
	if res := i2s(ui32); res != "1551335113" {
		t.Fatalf("expected: %v, got: %v", i, res)
	}
}
