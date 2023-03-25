package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Ajustes de integração bash mais asssociações de comandos
	cmd := exec.Command("bash", "-c", "echo hello; date; pwd; snmpwalk -v2c -c public 172.16.113.230 .1.3.6.1.4.1.3709.3.6.2.1.1.22")
	// Chamada do método Output para obter o resultado como uma fatia de byte
	out, err := cmd.Output()
	// Verificação da existência de erros
	if err != nil {
		fmt.Println(err)
	}
	// Converção da fatia de byte em uma string pronta para impressão
	fmt.Println(string(out))
}
