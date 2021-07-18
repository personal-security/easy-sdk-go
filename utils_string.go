package easysdk

import (
	"bytes"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func RandStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func RandIntBytes(n int) string {
	const letterBytes = "123456789"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func InsertIntoString(s string, interval int, sep rune) string {
	var buffer bytes.Buffer
	before := interval - 1
	last := len(s) - 1
	for i, char := range s {
		buffer.WriteRune(char)
		if i%interval == before && i != last {
			buffer.WriteRune(sep)
		}
	}
	return buffer.String()
}

func GenerateUrlHash() string {
	unix := uint64(time.Now().Unix())
	md5Unix := GetMD5Hash(strconv.FormatUint(unix, 10))
	hash := md5Unix + "-" + uuid.New().String()
	return hash
}

func MakeSecureString(text string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9 -.?!)(,:]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(text, "")
	return processedString
}
