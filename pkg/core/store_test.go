package core_test

import (
	"fmt"
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

func TestApplyMiddleware(t *testing.T) {
	reducer := CombineReducers(map[string]interface{}{
		"counter": func(state int, action Action) int {
			switch action.Type {
			case "INCREMENT":
				return state + 1
			case "DECREMENT":
				return state - 1
			default:
				return state
			}
		},
	})

	preloadedState := State{
		"counter": 0,
	}

	increment := Action{Type: "INCREMENT"}
	decrement := Action{Type: "DECREMENT"}

	t.Run("Middlewares get called in order and receives dispatched action and store", func(t *testing.T) {
		store := CreateStore(reducer, preloadedState)
		spy := struct{ calls []string }{[]string{}}

		middlewares := []Middleware{
			func(action Action, store *Store) {
				spy.calls = append(spy.calls, fmt.Sprintf("first middleware: %s", action.Type))
			},

			func(action Action, store *Store) {
				spy.calls = append(spy.calls, fmt.Sprintf("second middleware: %s", action.Type))
			},
		}

		store.ApplyMiddleware(middlewares...)

		store.Dispatch(increment)

		got := spy.calls
		want := []string{"first middleware: INCREMENT", "second middleware: INCREMENT"}

		AssertDeepEqual(t, got, want)

		store.Dispatch(decrement)

		got = spy.calls
		want = []string{"first middleware: INCREMENT", "second middleware: INCREMENT", "first middleware: DECREMENT", "second middleware: DECREMENT"}

		AssertDeepEqual(t, got, want)
	})

	t.Run("Middleware can dispatch another action", func(t *testing.T) {
		store := CreateStore(reducer, preloadedState)

		middlewares := []Middleware{
			func(action Action, store *Store) {
				if action.Type == "INCREMENT" {
					store.Dispatch(decrement)
				}
			},
		}

		store.ApplyMiddleware(middlewares...)

		store.Dispatch(increment)

		got := store.GetState()
		want := map[string]interface{}{"counter": 0}

		AssertDeepEqual(t, got, want)
	})

	t.Run("Middleware can access current state of store", func(t *testing.T) {
		store := CreateStore(reducer, preloadedState)
		states := []State{}

		middlewares := []Middleware{
			func(action Action, store *Store) {
				states = append(states, store.GetState())
			},
		}

		store.ApplyMiddleware(middlewares...)

		store.Dispatch(increment)

		got := states
		want := []State{State{"counter": 0}}

		AssertDeepEqual(t, got, want)

		store.Dispatch(decrement)

		got = states
		want = []State{State{"counter": 0}, State{"counter": 1}}

		AssertDeepEqual(t, got, want)
	})
}
