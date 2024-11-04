package main

// Import necessary packages
import (
	"log"
	"os/exec"
)

func main() {
	// Run linux-exploit-suggester
	lesDir := "/usr/share/linux-exploit-suggester"
	out, err := exec.Command("sh", lesDir+"/les.sh", "--checksec").Output()

	// Check for runtime errors
	if err != nil {
		log.Fatal(err)
	} else {
		// Send output of linux-exploit-suggester to external listening server

	}
}
