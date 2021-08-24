package wsse

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	HMAC_SHA256 = "HmacSHA256"
	HMAC_SHA1   = "HmacSHA1"
	SHA256      = "SHA256"
	SHA1        = "SHA1"
)

// Transport implements http.Transport. It adds X-WSSE header
// to client requests.
type Transport struct {
	Username   string
	Password   string
	SignMethod string
	Transport  http.RoundTripper
}

func (t *Transport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}

	return http.DefaultTransport
}

const nonceSize uint = 16

// RoundTrip executes the HTTP request with an X-WSSE header.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	nonce, err := Nonce(nonceSize)
	if err != nil {
		return nil, err
	}

	created := time.Now().UTC().Format(time.RFC3339)

	passwordDigest := base64.StdEncoding.EncodeToString(createPasswordDigest(nonce, created, t.Password))
	wsseHeader := fmt.Sprintf(`UsernameToken Username="%s",PasswordDigest="%s",Nonce="%s",Created="%s"`, t.Username, passwordDigest, nonce, created)
	req.Header.Set("X-WSSE", wsseHeader)

	return t.transport().RoundTrip(req)
}

func createPasswordDigest(nonce, created, password string) []byte {
	//switch signMethod {
	//case HMAC_SHA1:
	//	hashed := hmac.New(sha1.New, []byte(password))
	//	hashed.Write([]byte(nonce))
	//	hashed.Write([]byte(created))
	//	return hashed.Sum(nil)
	//case HMAC_SHA256:
	//	hashed := hmac.New(sha256.New, []byte(password))
	//	hashed.Write([]byte(nonce))
	//	hashed.Write([]byte(created))
	//	return hashed.Sum(nil)
	//case SHA1:
	//	digest := sha1.New()
	//	digest.Write([]byte(nonce))
	//	digest.Write([]byte(created))
	//	digest.Write([]byte(password))
	//	return digest.Sum(nil)
	//default:
	//}
	digest := sha256.New()
	digest.Write([]byte(nonce))
	digest.Write([]byte(created))
	digest.Write([]byte(password))
	return digest.Sum(nil)
}

// Nonce returns nonce string, param `size` better for even number.
func Nonce(size uint) (string, error) {
	nonce := make([]byte, size/2)
	_, err := io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(nonce), nil
}
