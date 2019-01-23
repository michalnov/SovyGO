package main

import (
	"fmt"

	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/michalnov/SovyGo/bin/server"
)

func main() {
	fmt.Println("Hello Server")

	running := make(chan int)
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return
	}
	pemdata := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(&key.PublicKey),
		},
	)
	fmt.Println(string(pemdata))
	fmt.Println(key.PublicKey)
	y := x509.MarshalPKCS1PublicKey(&key.PublicKey)
	fmt.Println(y)
	go func() {
		server := server.Server{}
		server.SetupServer(running)
		server.StartServer()
	}()

	for i := 0; i < 1; i++ {
		_ = <-running
	}
}
