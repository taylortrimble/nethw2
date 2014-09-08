package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	ticker := time.NewTicker(35 * time.Second)
	go func() {
		for {
			err := Register("trim0055", 9949, "apollo.cselabs.umn.edu:9080")
			if err != nil {
				log.Fatalln(err)
			}

			<-ticker.C
		}
	}()

	listen, err := net.Listen("tcp", ":9949")
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

func Register(author string, port int, endpoint string) error {
	conn, err := net.Dial("tcp", endpoint)
	if err != nil {
		return err
	}

	remoteIP := strings.Split(conn.LocalAddr().String(), ":")[0]
	message := fmt.Sprintf("Register %s %s %d\r\n", author, remoteIP, port)
	fmt.Printf("> %s", message)
	_, err = conn.Write([]byte(message))
	if err != nil {
		return err
	}

	b := make([]byte, 20)
	for err == nil {
		_, err = conn.Read(b)
	}

	fmt.Printf("< %s", b)

	err = conn.Close()
	if err != nil {
		return err
	}

	return nil
}
