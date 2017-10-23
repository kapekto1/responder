package main

import (
	"testing"
)

func TestRespond(t *testing.T) {
	expected := Respond()
	if len(expected) == 0 {
		t.Error("No hostname!")
	}
}