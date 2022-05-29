package otp

import (
	"log"
	"strings"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func GenerateOtp(secret string, period uint, algorithm string) string {

	var algo otp.Algorithm

	switch strings.ToLower(algorithm) {
	case "sha256":
		algo = otp.AlgorithmSHA256
	case "sha512":
		algo = otp.AlgorithmSHA512
	}

	passcode, err := totp.GenerateCodeCustom(secret, time.Now(),
		totp.ValidateOpts{Period: period, Skew: 1, Digits: otp.DigitsSix, Algorithm: algo})

	if err != nil {
		log.Fatalf("Generating OTP failed, %v", err)
		return ""
	}

	return passcode
}

func ValidateOtp(passcode string, secret string) bool {
	return totp.Validate(passcode, secret)
}
