package main

import (
	"crypto/rand"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "bulbasaur.tntapp.co:4949")
	if err != nil {
		log.Println(err)
		return
	}

	start := time.Now()

	written, err := io.CopyN(conn, rand.Reader, 10000000)
	if err != nil {
		log.Println(err)
		return
	}
	if written != 10000000 {
		log.Println("Did not send 10000000 bytes")
		return
	}

	response := make([]byte, 4)
	_, err = conn.Read(response)
	if err != nil {
		log.Println(err)
		return
	}

	var end time.Time
	if string(response) == "FIN!" {
		end = time.Now()
		err = conn.Close()
		if err != nil {
			log.Println("Couldn't close connection properly")
		}
	}

	log.Println(string(response))
	log.Println("Took", end.Sub(start))
}
