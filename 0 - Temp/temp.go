package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("terminal", "/C", "tree").Output()
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

aplicação com go que rode um comando msdos

Sim, é possível aplicar um comando MS-DOS em Go.

Para fazer isso, você precisa usar o pacote "os/exec" do Go. Esse pacote fornece APIs para executar comandos do sistema operacional, incluindo comandos MS-DOS.

Aqui está um exemplo de código que executa o comando MS-DOS "tree":
*/
