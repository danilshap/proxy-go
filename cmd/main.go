package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"proxy-go/config"
	"proxy-go/pkg/commands"
	"proxy-go/pkg/proxy"
)

func init() {
	config.InitConfig()
}

func main() {
	go func() {
		http.HandleFunc("/", proxy.HandleProxy)
		http.ListenAndServe(config.GetProxyDomain(), nil)
	}()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Go Tunnel Manager")
	commands.ShowInformation()
	for {
		fmt.Print("Enter command: ")
		scanner.Scan()
		input := scanner.Text()
		commands.ProcessCommand(input)
	}
}
