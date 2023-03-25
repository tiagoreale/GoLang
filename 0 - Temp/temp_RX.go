package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Create a Cmd struct that runs "echo hello; date; pwd" command
	cmd := exec.Command("bash", "-c", "echo hello; date; pwd; snmpwalk -v2c -c public 172.16.113.230 .1.3.6.1.4.1.3709.3.6.2.1.1.22")
	// Call Output method to get the result as a byte slice
	out, err := cmd.Output()
	// Check for errors
	if err != nil {
		fmt.Println(err)
	}
	// Convert the byte slice to a string and print it
	fmt.Println(string(out))
}
