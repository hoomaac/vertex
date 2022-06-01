package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hoomaac/vertex/pkg/jwt"
	"github.com/hoomaac/vertex/pkg/otp"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	// Load the default env file
	err := godotenv.Load()

	if err != nil {
		os.Exit(1)
	}

	code := m.Run()

	os.Exit(code)
}

func TestGenerateJwt(t *testing.T) {

	token := jwt.GenerateJwt("username", "name")

	assertNotEqual(t, len(token), 0)
}

func TestGenerateOpt(t *testing.T) {

	// TODO: complete this test
	passcode := otp.GenerateOtp(os.Getenv("OTP_SECRET"), 30, "sha256")

	fmt.Printf("passcode: %s", passcode)
}
