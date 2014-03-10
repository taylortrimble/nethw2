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
		buf := make([]byte, 10000000)
		_, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println(err)
		}
	}
}
