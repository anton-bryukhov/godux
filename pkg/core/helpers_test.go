package core_test

import (
	"reflect"
	"testing"
)

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
