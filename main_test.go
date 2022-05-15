package main

import (
	"testing"
)

func TestPasswordMatches(t *testing.T) {
	pw := password("k8V2z")

	if pw.matches([]rune("p86d9")) {
		t.Fatalf("Expected false for completely different guess")
	}

	if pw.matches([]rune("k8V2")) {
		t.Fatalf("Expected false for guess too short")
	}

	if pw.matches([]rune("k8V2zX")) {
		t.Fatalf("Expected false for guess too long")
	}

	if pw.matches([]rune("k8v2z")) {
		t.Fatalf("Expected false for guess with lowercase difference")
	}

	if !pw.matches([]rune("k8V2z")) {
		t.Fatalf("Expected true for correct guess")
	}
}
