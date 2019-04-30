package main

import "testing"

func TestSayHello(t *testing.T) {
	expected := "Hello!"
	actual := sayHello()
	if actual != expected {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}
