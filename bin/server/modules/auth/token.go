package authentication

import (
	"crypto/rsa"
	"time"
)

type token struct {
	created       time.Time
	revision      time.Time
	clientPublic  rsa.PublicKey
	serverPublic  rsa.PublicKey
	serverPrivate rsa.PrivateKey
	symmetricKey  []byte
	highSecurity  bool
	tokenBase     string
}
