package main

import (
	"crypto/tls"
	"log"
)

func main() {
	fooserver()
}

func fooserver() {
	// Simple static webserver:
	cer, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatal(err)
	}
	l, err := tls.Listen("tcp", ":8080", &tls.Config{
		Certificates: []tls.Certificate{cer},
		MinVersion:   tls.VersionSSL30,
	})
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		_, _ = c.Write([]byte("hello"))
		_ = c.Close()
	}
}
