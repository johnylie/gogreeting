package gogreeting

import (
	"strings"
	"testing"
)

func TestGreeting(t *testing.T) {
	result := Greeting(
		"NAME",
		ID,
		nil,
	)

	if !strings.Contains(result, "NAME") {
		t.Error("Greeting should contain name")
	}
}

func TestCustomGreeting(t *testing.T) {
	custom := &Translation{
		Morning: "Halo pagi",
	}

	result := Greeting(
		"NAME",
		ID,
		custom,
	)

	if result == "" {
		t.Error("Custom greeting should not be empty")
	}
}
