package godux

import (
	"fmt"
	"testing"
)

func TestAction(t *testing.T) {
	t.Run("Action can be created", func(t *testing.T) {
		action := Action{Type: "INCREMENT"}

		got := action.Type
		want := "INCREMENT"

		assertEqual(t, got, want)
	})

	t.Run("Action has nil payload by default", func(t *testing.T) {
		action := Action{Type: "INCREMENT"}

		got := action.Payload

		assertEqual(t, got, nil)
	})

	t.Run("Action can have payload", func(t *testing.T) {
		action := Action{Type: "INCREMENT", Payload: 1}

		got, ok := action.Payload.(int)
		if !ok {
			t.Fatalf("wrong payload type")
		}
		want := 1

		assertEqual(t, got, want)
	})
}

func TestReducer(t *testing.T) {
	t.Run("Reducer receives state and action and returns state", func(t *testing.T) {
		action := Action{Type: "INCREMENT"}

		var reducer Reducer = func(state State, action Action) State {
			return state
		}

		_, got := reducer(42, action).(State)
		want := true

		assertEqual(t, got, want)
	})
}

func TestStore(t *testing.T) {
	t.Run("Store can be created with passed reducer and inital state", func(t *testing.T) {
		var reducer Reducer = func(state State, action Action) State {
			return state
		}
		store := CreateStore(0, reducer)

		got := fmt.Sprintf("%v", store.reducer)
		want := fmt.Sprintf("%v", reducer)

		assertEqual(t, got, want)
	})

	t.Run("Store returns initial state after creation", func(t *testing.T) {
		initailState := 0

		store := CreateStore(
			initailState,
			func(state State, action Action) State {
				return state
			},
		)

		got := store.GetState()
		want := initailState

		assertEqual(t, got, want)
	})
}

func assertEqual(t *testing.T, actual interface{}, expected interface{}) {
	t.Helper()
	if actual != expected {
		t.Fatalf("Actual: '%v', expected: '%v'", actual, expected)
	}
}
