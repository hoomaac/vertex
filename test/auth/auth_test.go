package test

import (
	"log"
	"os"
	"testing"

	"github.com/hoomaac/vertex/pkg/jwt"
	"github.com/hoomaac/vertex/pkg/otp"
	"github.com/joho/godotenv"
	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// Load the default env file
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Printf("failed, %v\n", err)
		os.Exit(1)
	}

	code := m.Run()

	os.Exit(code)
}

func TestGenerateJwt(t *testing.T) {

	token := jwt.GenerateJwt("username", "name")

	require.NotEqual(t, len(token), 0)
}

func TestGenerateOtp(t *testing.T) {

	// TODO: complete this test
	passcode := otp.GenerateOtp("test@gmail.com")

	require.NotEqual(t, passcode, "")

	log.Printf("Passcode: %s\n", passcode)
}

func TestGenerateKey(t *testing.T) {

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "vertex",
		AccountName: "test@gmail.com",
	})

	if err != nil {
		t.Errorf("err: %v\n", err)
	}

	require.NotEqual(t, key.Secret(), "")
}
