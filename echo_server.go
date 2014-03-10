package main

import (
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":4949")
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	defer conn.Close()
	if err != nil {
		log.Println(err)
	}

	for {
		buf := make([]byte, 100)
		read, inAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println(err)
		}
		if read != 100 {
			log.Printf("Read %d bytes instead of 100.\n", read)
		}

		written, err := conn.WriteToUDP(buf, inAddr)
		if err != nil {
			log.Println(err)
		}
		if written != 100 {
			log.Println("Didn't write 100 bytes")
			continue
		}
	}
}
