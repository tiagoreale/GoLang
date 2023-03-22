package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("cmd", "/C", "tree", "snmpwalk -v2c -c public 172.16.113.230 .1.3.6.1.4.1.3709.3.6.2.1.1.3").Output()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(string(out))
}

/*
snmpwalk -v2c -c public 172.16.113.230 .1.3.6.1.4.1.3709.3.6.2.1.1.22
snmpwalk -v2c -c public 172.16.113.230 .1.3.6.1.4.1.3709.3.6.2.1.1.3

package main

import (
    "os/exec"
    "fmt"
)

func main() {
    out, err := exec.Command("cmd", "/C", "tree", "").Output()
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }

    fmt.Println(string(out))
}
*/
