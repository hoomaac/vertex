package otp

import (
	"encoding/base32"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hoomaac/vertex/pkg/redis"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

const UnusedStrParam string = ""
const UnusedIntParam int = 0

func getTimeout() int {

	timeout := os.Getenv("OTP_TIMEOUT")

	if timeout == "" {
		log.Println("OTP_TIMEOUT is not set. Default timeout, which is 30s, will be used")
		timeout = "30"
	}

	otpTimeout, err := strconv.Atoi(timeout)

	if err != nil {
		log.Printf("converting otp timeout failed, %v\n", err)
		return 0
	}

	return otpTimeout
}

func getSecret() string {

	secret := os.Getenv("OTP_SECRET")

	if secret == "" {
		log.Fatalln("OTP_SECRET is empty")
	}

	return secret
}

func GenerateOtp(algorithm string) string {

	var algo otp.Algorithm
	var err error

	secret := getSecret()
	encodedSecret := base32.StdEncoding.EncodeToString([]byte(secret))

	if algorithm == UnusedStrParam {
		log.Println("OTP_ALGORITHM is not set. Default algorithm is used, SHA256")
		algorithm = "sha256"
	}

	switch strings.ToLower(algorithm) {
	case "sha256":
		algo = otp.AlgorithmSHA256
	case "sha512":
		algo = otp.AlgorithmSHA512
	}

	otpTimeout := getTimeout()

	if otpTimeout == 0 {
		return ""
	}

	passcode, err := totp.GenerateCodeCustom(encodedSecret, time.Now(),
		totp.ValidateOpts{Period: uint(otpTimeout), Skew: 1, Digits: otp.DigitsSix, Algorithm: algo})

	if err != nil {
		log.Printf("Generating OTP failed, %v", err)
		return ""
	}

	return passcode
}

func ValidateOtp(passcode string, secret string) bool {

	if secret == UnusedStrParam {
		secret = getSecret()
	}

	return totp.Validate(passcode, secret)
}

// Store value on RedisClient with this format:
// passcode@<username>, <username> should be replaced with any valid username
func StorePasscodeOnRedis(passcode string, username string, timeout int) bool {

	if timeout == UnusedIntParam {
		timeout = getTimeout()
		if timeout == 0 {
			return false
		}
	}

	return redis.SetValue(fmt.Sprintf("passcode@%s", username), passcode, time.Duration(timeout))
}

func GetPasscodeFromRedis(username string) string {
	return redis.GetValue(username)
}
