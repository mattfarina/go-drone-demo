package main

import "testing"

func TestContent(t *testing.T) {
	c := content()
	if c != "Hello world!" {
		t.Error("Wrong content")
	}
}
