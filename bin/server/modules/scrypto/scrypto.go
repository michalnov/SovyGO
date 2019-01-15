package scrypto

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"hash"
	"os"
)

func newKeypair() (*rsa.PrivateKey, error) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func rsaEncrypt(data []byte, label []byte, hash hash.Hash, key *rsa.PublicKey) ([]byte, error) {
	rng := rand.Reader
	ciphertext, err := rsa.EncryptOAEP(hash, rng, key, data, label)
	if err != nil {
		return []byte(""), err
	}

	return ciphertext, nil
}

func rsaDecrypt(data []byte, label []byte, hash hash.Hash, key *rsa.PrivateKey) ([]byte, error) {
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(hash, rng, key, data, label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from decryption: %s\n", err)
		return []byte(""), err
	}
	return plaintext, nil
}

func aesEncrypt() ([]byte, error) {

}

func aesDecrypt() ([]byte, error) {

}
