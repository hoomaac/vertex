package otp

import (
	"encoding/base32"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/hoomaac/vertex/pkg/redis"
	"github.com/pquerna/otp/totp"
)

// Redis OtpTimeout
const OtpTimeout = 30 * 1000 * 1000 * 1000

func getTimeout() int {

	strTimeout := os.Getenv("OTP_TIMEOUT")

	if strTimeout == "" {
		log.Println("OTP_TIMEOUT is not set. Default timeout, which is 30s, will be used")
		return OtpTimeout
	}

	timeout, err := strconv.Atoi(strTimeout)

	if err != nil {
		log.Printf("converting otp timeout failed, %v\n", err)
		return 0
	}

	return timeout
}

// Returns base32 encoded secret
func getSecret() string {

	secret := os.Getenv("OTP_SECRET")

	if secret == "" {
		log.Fatalln("OTP_SECRET is empty")
	}

	secret += "low"

	return base32.StdEncoding.EncodeToString([]byte(secret))
}

func generateKey(email string) string {

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "vertex",
		AccountName: email,
	})

	if err != nil {
		log.Printf("generating key for email: %s has been failed due to err: %v\n", email, err)
		return ""
	}

	return key.Secret()
}

func GenerateOtp(email string) string {

	var err error
	secret := generateKey(email)

	if secret == "" {
		log.Println("secret key is empty")
		return ""
	}

	passcode, err := totp.GenerateCode(secret, time.Now())

	if err != nil {
		log.Printf("Generating OTP failed, %v", err)
		return ""
	}

	log.Printf("secret key is: %s\n", secret)

	if !StoreSecretOnRedis(secret, email, getTimeout()) {
		log.Printf("Storing OTP secret on redis failed, %v\n", err)
		return ""
	}

	return passcode
}

func ValidateOtp(passcode string, email string) bool {

	secret := GetSecretFromRedis(email)

	log.Printf("value from redis: %s\n", secret)

	if secret == "" {
		return false
	}

	return totp.Validate(passcode, secret)
}

// Store value on RedisClient with this format:
// passcode@<username>, <username> should be replaced with any valid username
func StoreSecretOnRedis(secretKey string, email string, timeout int) bool {

	if timeout == 0 {
		timeout = 30 * 100
	}

	return redis.SetValue(fmt.Sprintf("otp:%s", email), secretKey, timeout)
}

func GetSecretFromRedis(email string) string {
	return redis.GetValue(fmt.Sprintf("otp:%s", email))
}
