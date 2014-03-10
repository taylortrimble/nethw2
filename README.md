nethw2
======

This project is written in Go, and can be run from any go environment. Go is installed on the CSElabs computers,
and you can learn how to set up a Go environment at golang.org.

To run the TCP version:

go run server.go
go run client.go

To run UDP 10x echo:

go run echo_server.go
go run echo_client.go

To run UDP 10,000,000 byte Transmit:

go run spam_server.go
go run spam_client.go
