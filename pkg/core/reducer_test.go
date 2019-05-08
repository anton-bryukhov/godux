package core_test

import (
	. "github.com/anton-bryukhov/godux/pkg/core"
	"testing"
)

func TestCombineReducers(t *testing.T) {
	t.Run("CombineReducers returns combined reducer if one reducer provided", func(t *testing.T) {
		combinedReducer := CombineReducers(map[string]interface{}{
			"counter": CounterReducer,
		})

		got := combinedReducer(State{"counter": 1}, Increment)
		want := State{"counter": 2}

		AssertDeepEqual(t, got, want)
	})

	t.Run("CombineReducers returns combined reducer if several reducers provided", func(t *testing.T) {
		combinedReducer := CombineReducers(map[string]interface{}{
			"counter": CounterReducer,
			"toggler": TogglerReducer,
		})

		got := combinedReducer(State{"counter": 1, "toggler": false}, Increment)
		want := State{"counter": 2, "toggler": false}

		AssertDeepEqual(t, got, want)

		got = combinedReducer(State{"counter": 1, "toggler": false}, Toggle)
		want = State{"counter": 1, "toggler": true}

		AssertDeepEqual(t, got, want)
	})
}
