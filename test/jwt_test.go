package test

import (
	"testing"

	"github.com/hoomaac/vertex/middleware/jwt"
)

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

func TestGenerateJwt(t *testing.T) {

	token := jwt.GenerateJwt([]byte("key"))

	assertNotEqual(t, len(token), 0)

}
