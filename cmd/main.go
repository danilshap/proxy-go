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
		proxyAddress := config.GetProxyDomain()
		if err := http.ListenAndServe(proxyAddress, nil); err != nil {
			fmt.Printf("Failed to start proxy server: %v\n", err)
			os.Exit(1)
		}
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
