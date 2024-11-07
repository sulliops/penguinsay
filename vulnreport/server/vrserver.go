package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"strconv"
	"time"
)

func handleConnection(conn net.Conn) {
	// Close the connection at the end of goroutine execution
	defer conn.Close()

	// Get client IP address
	addr := conn.RemoteAddr().String()
	fmt.Println("Receiving data from " + addr + "...")

	data := *new(string)
	// Read received data until EOF is reached
	read := bufio.NewReader(conn)
	for {
		// Read the next line of text data
		msg, err := read.ReadString('\n')
		if err == io.EOF {
			data += msg
			// Call helper goroutine to save received data to file
			go saveReport(addr, data)
			return
		} else if err != nil {
			log.Fatal(err)
			continue
		}

		// Add buffer contents to received data
		data += msg
	}
}

func saveReport(addr string, data string) {
	// Write received data to a file
	// Use the client's IP address and the current timestamp as the file name
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	name := "vulnreport-" + addr + "-" + timestamp + ".txt"
	err := ioutil.WriteFile(name, []byte(data), 0644)
	if err != nil {
		log.Fatal(err)
		return
	} else {
		fmt.Println("Saved vulnerability report as \"" + name + "\".")
	}
}

func main() {
	// Listen for incoming connections on port 9090
	port := fmt.Sprintf("%d", 9090)
	ln, err := net.Listen("tcp", ":"+port)
	// Error check listener
	if err != nil {
		log.Fatal(err)
		return
	} else {
		fmt.Println("Listening on port " + port + "...")
	}

	// Accept and handle incoming connections
	for {
		conn, err := ln.Accept()
		// Error check connection
		if err != nil {
			log.Fatal(err)
			continue
		}

		// Use helper goroutine to handle connection
		go handleConnection(conn)
	}
}
