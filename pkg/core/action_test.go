package core_test

import (
	. "github.com/anton-bryukhov/godux/pkg/core"
	"testing"
)

func TestAction(t *testing.T) {
	t.Run("Action can be created", func(t *testing.T) {
		action := Action{Type: "INCREMENT"}

		got := action.Type
		want := "INCREMENT"

		AssertEqual(t, got, want)
	})

	t.Run("Action has nil payload by default", func(t *testing.T) {
		action := Action{Type: "INCREMENT"}

		got := action.Payload

		AssertEqual(t, got, nil)
	})

	t.Run("Action can have payload", func(t *testing.T) {
		action := Action{Type: "INCREMENT", Payload: 1}

		got, ok := action.Payload.(int)
		if !ok {
			t.Fatalf("wrong payload type")
		}
		want := 1

		AssertEqual(t, got, want)
	})
}
