package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"proxy-go/pkg/commands"
	"proxy-go/pkg/proxy"
)

func main() {
	go func() {
		http.HandleFunc("/", proxy.HandleProxy)
		http.ListenAndServe(":8081", nil)
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
