package main

import (
	"io"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":4949")
	if err != nil {
		log.Println(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
		}

		written, err := io.CopyN(ioutil.Discard, conn, 10000000)
		if err != nil {
			log.Println(err)
		}
		if written != 10000000 {
			log.Println("Didn't write 10000000 bytes")
			continue
		}

		_, err = conn.Write([]byte("FIN!"))
		if err != nil {
			log.Println(err)
		}

		b := make([]byte, 20)
		for err == nil {
			_, err = conn.Read(b)
		}

		err = conn.Close()
		if err != nil {
			log.Println(err)
		}
	}
}
