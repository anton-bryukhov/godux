package core_test

import (
	"fmt"
	. "github.com/anton-bryukhov/godux/pkg/core"
	"testing"
)

func TestStore(t *testing.T) {
	reducer := CombineReducers(map[string]interface{}{
		"counter": CounterReducer,
		"toggler": TogglerReducer,
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

		store.Dispatch(Increment)

		got := store.GetState()
		want := map[string]interface{}{
			"counter": 1,
			"toggler": false,
		}

		AssertDeepEqual(t, got, want)

		store.Dispatch(Toggle)

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
		"counter": CounterReducer,
	})

	preloadedState := State{
		"counter": 0,
	}

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

		store.Dispatch(Increment)

		got := spy.calls
		want := []string{"first middleware: INCREMENT", "second middleware: INCREMENT"}

		AssertDeepEqual(t, got, want)

		store.Dispatch(Decrement)

		got = spy.calls
		want = []string{"first middleware: INCREMENT", "second middleware: INCREMENT", "first middleware: DECREMENT", "second middleware: DECREMENT"}

		AssertDeepEqual(t, got, want)
	})

	t.Run("Middleware can dispatch another action", func(t *testing.T) {
		store := CreateStore(reducer, preloadedState)

		middlewares := []Middleware{
			func(action Action, store *Store) {
				if action.Type == "INCREMENT" {
					store.Dispatch(Decrement)
				}
			},
		}

		store.ApplyMiddleware(middlewares...)

		store.Dispatch(Increment)

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

		store.Dispatch(Increment)

		got := states
		want := []State{State{"counter": 0}}

		AssertDeepEqual(t, got, want)

		store.Dispatch(Decrement)

		got = states
		want = []State{State{"counter": 0}, State{"counter": 1}}

		AssertDeepEqual(t, got, want)
	})
}
