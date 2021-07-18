package easysdk

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"log"

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
