package core_test

import (
	. "github.com/anton-bryukhov/godux/pkg/core"
	"reflect"
	"testing"
)

const (
	increment = "INCREMENT"
	decrement = "DECREMENT"
	toggle    = "TOGGLE"
)

var Increment = Action{Type: increment}
var Decrement = Action{Type: decrement}
var Toggle = Action{Type: toggle}

func CounterReducer(state int, action Action) int {
	switch action.Type {
	case increment:
		return state + 1
	case decrement:
		return state - 1
	default:
		return state
	}
}

func TogglerReducer(state bool, action Action) bool {
	switch action.Type {
	case toggle:
		return !state
	default:
		return state
	}
}

func AssertEqual(t *testing.T, actual interface{}, expected interface{}) {
	t.Helper()
	if actual != expected {
		t.Fatalf("Actual: '%v', expected: '%v'", actual, expected)
	}
}

func AssertDeepEqual(t *testing.T, actual interface{}, expected interface{}) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Actual: '%v', expected: '%v'", actual, expected)
	}
}
