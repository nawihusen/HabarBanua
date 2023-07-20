package helper

import (
	"bytes"
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

// SignatureGetTokenMerchant generate signature to get token merchant
func SignatureGetTokenMerchant(merchantID string, timestamp string, privateKeyFile string) (signature string, err error) {
	// Read private key
	fmt.Println("Log for signature merchant token")
	pemString, err := os.ReadFile(privateKeyFile)
	if err != nil {
		return
	}

	// Hash message with SHA256
	h := sha256.New()
	h.Write([]byte(merchantID + "|" + timestamp))
	d := h.Sum(nil)

	// Sign with PKCS1v15
	block, _ := pem.Decode([]byte(pemString))
	privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	// fmt.Println(privateKey.N)
	sign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, d)
	if err != nil {
		return
	}
	signature = base64.StdEncoding.EncodeToString(sign)
	log.Debug("SignatureGetTokenMerchant: " + signature)

	return
}

// SignatureWithTokenMerchant generate signature with token merchant (symetric)
func SignatureWithTokenMerchant(timestamp string, token string, method string, endpoint string, requestBody string, merchantSecret string) (signature string, err error) {
	// Hash requestBody with SHA256
	fmt.Println("Log for signature merchant")
	var minified bytes.Buffer
	err = json.Compact(&minified, []byte(requestBody))
	if err != nil {
		return
	}
	h := sha256.New()
	h.Write(minified.Bytes())
	requestBodyHash := hex.EncodeToString(h.Sum(nil))
	// requestBodyHash := base64.URLEncoding.EncodeToString(h.Sum(nil))

	// Hash all with HMAC using algorithm SHA512
	mac := hmac.New(sha512.New, []byte(merchantSecret))
	message := method + ":" + endpoint + ":" + token + ":" + requestBodyHash + ":" + timestamp
	log.Debug("stringToSign: " + message)
	mac.Write([]byte(message))
	sign := mac.Sum(nil)
	signature = base64.StdEncoding.EncodeToString(sign)
	log.Debug("SignatureWithTokenMerchant: " + signature)

	return
}
