package main

// Import necessary packages
import (
	"log"
	"net"
	"os/exec"
)

func main() {
	// Connect to vulnreport listening server
	conn, err := net.Dial("tcp", "babbage12.computing.clemson.edu:9090")
	// Error check connection
	if err != nil {
		log.Fatal(err)
		return
	}

	// Run linux-exploit-suggester and get output
	lesDir := "/usr/share/linux-exploit-suggester"
	out, err := exec.Command(lesDir+"/les.sh", "--checksec").Output()
	// Check for runtime errors
	if err != nil {
		log.Fatal(err)
	} else {
		// Send output of linux-exploit-suggester to external listening server
		_, err := conn.Write([]byte(string(out)))
		// Error check connection
		if err != nil {
			log.Fatal(err)
			return
		}

		// Close the connection
		conn.Close()
	}
}
