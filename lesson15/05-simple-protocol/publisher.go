package main

import (
	"fmt"
	"log"

	nats "github.com/nats-io/nats.go"
)

const Subject = "test1"

func main() {
	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	fmt.Printf("Connected to %s\n", nats.DefaultURL)

	econn, err2 := nats.NewEncodedConn(conn, nats.DEFAULT_ENCODER)
	if err2 != nil {
		log.Fatal(err)
	}

	defer econn.Close()

	channel := make(chan string)
	econn.BindSendChan(Subject, channel)

	fmt.Println("Channel created")

	channel <- "Hello World #1"
	channel <- "Hello World #2"
	channel <- "Hello World #3"
	channel <- "EXIT"

	conn.Flush()
	fmt.Println("All messages sent")
}
