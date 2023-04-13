package main

import (
	"fmt"

	"github.com/AlekSi/go-zabbix"
)

func main() {
	// Cria a conexão com o servidor Zabbix
	zabbixServer := zabbix.New("http://localhost/zabbix/api_jsonrpc.php")
	_, err := zabbixServer.Login("admin", "password")
	if err != nil {
		panic(err)
	}

	// Realiza a consulta LLD para o host "MeuHost"
	params := zabbix.Params{"output": "extend", "host": "MeuHost", "application": "CPU"}
	items, err := zabbixServer.ItemGet(params)
	if err != nil {
		panic(err)
	}

	// Imprime os resultados
	for _, item := range items {
		fmt.Println(item.Name)
		fmt.Println(item.Key)
		fmt.Println(item.ValueType)
		fmt.Println(item.LastValue)
		fmt.Println(item.Status)
	}
}
