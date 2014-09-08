package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net"
	"time"
	"strings"
)

func main() {
	host, port, err := MyServer("trim0055", "apollo.cselabs.umn.edu:9080")
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
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

	myIP := strings.Split(conn.LocalAddr().String(), ":")[0]
	elapsed := int(end.Sub(start) / time.Millisecond)
	err = ReportTime("trim0055", myIP, host, port, fmt.Sprintf("%d", elapsed), "apollo.cselabs.umn.edu:9090")
	if err != nil {
		log.Fatalln(err)
	}
}

func MyServer(x500, endpoint string) (string, string, error) {
	conn, err := net.Dial("tcp", endpoint)
	if err != nil {
		return "", "", err
	}

	message := fmt.Sprintf("serverlist\r\n")
	fmt.Printf("> %s", message)
	_, err = conn.Write([]byte(message))
	if err != nil {
		return "", "", err
	}

	received := make([]byte, 10000)
	for err == nil {
		_, err = conn.Read(received)
	}

	trimmed := bytes.TrimRight(received, "\n\x00")
	split := bytes.Split(trimmed, []byte("\r"))

	host := ""
	port := ""
	for _, server := range split[1:] {
		parts := bytes.Split(server, []byte(" "))
		if string(parts[0]) == "trim0055" {
			host = string(parts[1])
			port = string(parts[2])
		}
		fmt.Printf("< %s\n", server)
	}

	err = conn.Close()
	if err != nil {
		return "", "", err
	}

	return host, port, nil
}

func ReportTime(x500, clientIP, serverIP, serverPort, time, endpoint string) error {
	message := fmt.Sprintf("setrecord %s %s %s %s %s\r\n", x500, clientIP, serverIP, serverPort, time)
	fmt.Printf("> %s", message)
	conn, err := net.Dial("tcp", endpoint)
	if err != nil {
		return err
	}

	_, err = conn.Write([]byte(message))
	if err != nil {
		return err
	}

	received := make([]byte, 20)
	for err == nil {
		_, err = conn.Read(received)
	}

	trimmed := bytes.TrimRight(received, "\r\n\x00")
	fmt.Printf("< %s\n", trimmed)

	err = conn.Close()
	if err != nil {
		return err
	}

	// Get List

	conn, err = net.Dial("tcp", endpoint)
	if err != nil {
		return err
	}

	message = fmt.Sprintf("getrecord\r\n")
	fmt.Printf("> %s", message)
	_, err = conn.Write([]byte(message))
	if err != nil {
		return err
	}

	received = make([]byte, 10000)
	for err == nil {
		_, err = conn.Read(received)
	}

	trimmed = bytes.TrimRight(received, "\n\x00")
	split := bytes.Split(trimmed, []byte("\r"))

	for _, record := range split {
		fmt.Printf("< %s\n", record)
	}

	// Done

	err = conn.Close()
	if err != nil {
		return err
	}

	return nil
}
