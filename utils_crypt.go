package easysdk

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
)

// GetMD5Hash generate md5 hash from string
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// GetMD5Hash generate md5 hash from []byte
func GetMD5HashByte(text []byte) string {
	hasher := md5.New()
	hasher.Write(text)
	return hex.EncodeToString(hasher.Sum(nil))
}

func PgpEncrypt(publicKeyring string, secretString string) (string, error) {

	data, err := EncryptMailMessage([]byte(publicKeyring), []byte(secretString))
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// EncryptMailMessage encrypts a mail with given public key.
func EncryptMailMessage(key, body []byte) ([]byte, error) {
	rdr := bytes.NewBuffer(key)
	keyring, err := openpgp.ReadArmoredKeyRing(rdr)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	out := new(bytes.Buffer)
	ct, err := armor.Encode(out, "PGP MESSAGE", nil)
	if err != nil {
		log.Println("Can't create armorer: " + err.Error())
		return nil, err
	}
	wrt, err := openpgp.Encrypt(ct, []*openpgp.Entity{keyring[0]}, nil, &openpgp.FileHints{IsBinary: true}, nil)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	wrt.Write(body)
	wrt.Close()
	ct.Close()

	return out.Bytes(), nil
}

// PgpEncryptMessage encrypt message with public pgp key
func PgpEncryptMessage(key, body []byte) ([]byte, error) {
	return EncryptMailMessage(key, body)
}

// PgpDecryptMessage decrypt message with private key (password optional)
func PgpDecryptMessage(text string, key string, pass string) (string, error) {
	keyRing := strings.NewReader(key)

	entityList, err := openpgp.ReadArmoredKeyRing(keyRing)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	entity := entityList[0]
	passphraseByte := []byte(pass)
	err = entity.PrivateKey.Decrypt(passphraseByte)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	for _, subkey := range entity.Subkeys {
		subkey.PrivateKey.Decrypt(passphraseByte)
	}

	encryptedContent := strings.NewReader(text)

	md, err := openpgp.ReadMessage(encryptedContent, entityList, nil, nil)
	if err != nil {
		log.Println(err)
		return "", err
	}

	bytes, err := ioutil.ReadAll(md.UnverifiedBody)
	if err != nil {
		log.Println(err)
		return "", err
	}

	decStr := string(bytes)

	return decStr, nil
}

// RC4 Encryption

// StringBase64Rc4 encrypt string with key
func StringBase64Rc4(data string, keyString string) string {
	key, _ := NewCipherRC4([]byte(keyString)) //initialize our cipher with the given key

	buf := EncryptRC4([]byte(data), key)

	base64String := base64.StdEncoding.EncodeToString(buf)

	return string(base64String)
}

func EncryptRC4(data []byte, ciph *CipherRC4) []byte {
	buffer := make([]byte, len(data))
	for i, v := range data {
		buffer[i] = byte(v)
	}

	ciph.XorKeyStreamGeneric(buffer, buffer) //encrypt the data
	ciph.Reset()                             //reset since we cant rewind the rc4 state for working on the same dataset

	return buffer
}

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license

type CipherRC4 struct {
	s    [256]uint32
	i, j uint8
}

type KeySizeError int

func (k KeySizeError) Error() string {
	return "crypto/rc4: invalid key size " + strconv.Itoa(int(k))
}

func NewCipherRC4(key []byte) (*CipherRC4, error) {
	k := len(key)
	if k < 1 || k > 256 {
		return nil, KeySizeError(k)
	}
	var c CipherRC4
	for i := 0; i < 256; i++ {
		c.s[i] = uint32(i)
	}
	var j uint8 = 0
	for i := 0; i < 256; i++ {

		j += uint8(c.s[i]) + key[i%k]

		c.s[i], c.s[j] = c.s[j], c.s[i]
	}
	return &c, nil
}

func (c *CipherRC4) Reset() {
	for i := range c.s {
		c.s[i] = 0
	}
	c.i, c.j = 0, 0
}

func (c *CipherRC4) XorKeyStreamGeneric(dst, src []byte) {
	i, j := c.i, c.j
	for k, v := range src {
		i += 1
		j += uint8(c.s[i])
		c.s[i], c.s[j] = c.s[j], c.s[i]
		dst[k] = v ^ uint8(c.s[uint8(c.s[i]+c.s[j])])
	}
	c.i, c.j = i, j
}
