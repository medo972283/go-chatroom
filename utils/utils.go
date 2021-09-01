package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Convert any type of str to the uuid.UUID type
func ConverseToUUID(value interface{}) uuid.UUID {
	res, _ := uuid.Parse(fmt.Sprintf("%v", value))
	return res
}

// Produces and returns a version 1 UUID.
func ProduceV1UUID() uuid.UUID {
	uuid1 := uuid.Must(uuid.NewUUID())
	return uuid1
}

// Salt and hash the password using the bcrypt algorithm
func BcryptString(str string) string {
	// The second paramerter is 'cost', the number of cost implies is the power-of-two number of rounds of hashing
	hashedkey, _ := bcrypt.GenerateFromPassword([]byte(str), 8)
	return string(hashedkey)
}

// Encrypts the string with MD5.
// Use bcrypt instead
func EncryptString(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	cipherStr := hasher.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
