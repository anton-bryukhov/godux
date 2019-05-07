package core_test

import (
	. "github.com/anton-bryukhov/godux/pkg/core"
	"testing"
)

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

		AssertDeepEqual(t, got, want)
	})

	t.Run("CombineReducers returns combined reducer if several reducers provided", func(t *testing.T) {
		combinedReducer := CombineReducers(map[string]interface{}{
			"counter": counter,
			"toggler": toggler,
		})

		got := combinedReducer(State{"counter": 1, "toggler": false}, increment)
		want := State{"counter": 2, "toggler": false}

		AssertDeepEqual(t, got, want)

		got = combinedReducer(State{"counter": 1, "toggler": false}, toggle)
		want = State{"counter": 1, "toggler": true}

		AssertDeepEqual(t, got, want)
	})
}
