package test

import "testing"

func assertEqual(t *testing.T, a interface{}, b interface{}) {

	if a != b {
		t.Errorf("%v and %v were expected to be equal", a, b)
	}
}

func assertNotEqual(t *testing.T, a interface{}, b interface{}) {

	if a == b {
		t.Errorf("%v and %v were expected to be equal", a, b)
	}
}
