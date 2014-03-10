package main

import (
	"crypto/rand"
	"io"
	"log"
	"net"
	"time"
)

const bytes = 1000

func main() {
	conn, err := net.Dial("udp", "bulbasaur.tntapp.co:4949")
	defer conn.Close()
	if err != nil {
		log.Println(err)
		return
	}

	start := time.Now()
	for i := 0; i < 10000000; i += bytes {
		written, err := io.CopyN(conn, rand.Reader, bytes)
		if err != nil {
			log.Println(err)
			return
		}
		if written != bytes {
			log.Println("Did not send 100 bytes")
			return
		}
	}

	log.Println("Took", time.Since(start))
}
