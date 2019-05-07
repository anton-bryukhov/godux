package core_test

import (
	. "github.com/anton-bryukhov/godux/pkg/core"
	"testing"
)

func TestStore(t *testing.T) {
	reducer := CombineReducers(map[string]interface{}{
		"counter": func(state int, action Action) int {
			switch action.Type {
			case "INCREMENT":
				return state + 1
			default:
				return state
			}
		},

		"toggler": func(state bool, action Action) bool {
			switch action.Type {
			case "TOGGLE":
				return !state
			default:
				return state
			}
		},
	})
	preloadedState := State{
		"counter": 0,
		"toggler": false,
	}

	t.Run("Store has an initial state", func(t *testing.T) {
		store := CreateStore(reducer, preloadedState)

		got := store.GetState()
		want := map[string]interface{}{
			"counter": 0,
			"toggler": false,
		}

		AssertDeepEqual(t, got, want)
	})

	t.Run("Store transforms state on actions dispatch", func(t *testing.T) {
		store := CreateStore(reducer, preloadedState)
		increment := Action{Type: "INCREMENT"}
		toggle := Action{Type: "TOGGLE"}

		store.Dispatch(increment)

		got := store.GetState()
		want := map[string]interface{}{
			"counter": 1,
			"toggler": false,
		}

		AssertDeepEqual(t, got, want)

		store.Dispatch(toggle)

		got = store.GetState()
		want = map[string]interface{}{
			"counter": 1,
			"toggler": true,
		}

		AssertDeepEqual(t, got, want)
	})
}
