package godux

import (
	"reflect"
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

func TestCombineReducers(t *testing.T) {
	increment := Action{Type: "INCREMENT"}
	counter := func(state int, action Action) int {
		switch action.Type {
		case "INCREMENT":
			return state + 1
		default:
			return state
		}
	}

	toggle := Action{Type: "TOGGLE"}
	toggler := func(state bool, action Action) bool {
		switch action.Type {
		case "TOGGLE":
			return !state
		default:
			return state
		}
	}

	t.Run("CombineReducers returns combined reducer if one reducer provided", func(t *testing.T) {
		combinedReducer := CombineReducers(map[string]interface{}{
			"counter": counter,
		})

		got := combinedReducer(State{"counter": 1}, increment)
		want := State{"counter": 2}

		assertDeepEqual(t, got, want)
	})

	t.Run("CombineReducers returns combined reducer if several reducers provided", func(t *testing.T) {
		combinedReducer := CombineReducers(map[string]interface{}{
			"counter": counter,
			"toggler": toggler,
		})

		got := combinedReducer(State{"counter": 1, "toggler": false}, increment)
		want := State{"counter": 2, "toggler": false}

		assertDeepEqual(t, got, want)

		got = combinedReducer(State{"counter": 1, "toggler": false}, toggle)
		want = State{"counter": 1, "toggler": true}

		assertDeepEqual(t, got, want)
	})
}

func assertEqual(t *testing.T, actual interface{}, expected interface{}) {
	t.Helper()
	if actual != expected {
		t.Fatalf("Actual: '%v', expected: '%v'", actual, expected)
	}
}

func assertDeepEqual(t *testing.T, actual interface{}, expected interface{}) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Actual: '%v', expected: '%v'", actual, expected)
	}
}
