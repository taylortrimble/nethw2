package main

import (
	"crypto/rand"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("udp", "bulbasaur.tntapp.co:4949")
	defer conn.Close()
	if err != nil {
		log.Println(err)
		return
	}

	durations := make([]time.Duration, 10)

	for i := 0; i < 10; i++ {
		start := time.Now()

		written, err := io.CopyN(conn, rand.Reader, 100)
		if err != nil {
			log.Println(err)
			return
		}
		if written != 100 {
			log.Println("Did not send 100 bytes")
			return
		}

		response := make([]byte, 100)
		received, err := conn.Read(response)
		if err != nil {
			log.Println(err)
			return
		}

		var end time.Time
		if received == 100 {
			end = time.Now()
		}

		durations[i] = end.Sub(start)

		log.Println("Took", durations[i])
	}

	var average int64 = 0
	for _, d := range durations {
		average += int64(d)
	}
	average /= int64(len(durations))

	log.Println("Average:", time.Duration(average))
}
