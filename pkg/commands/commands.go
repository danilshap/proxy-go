package commands

import (
	"fmt"
	"os"
	"proxy-go/pkg/proxy"
	"regexp"
	"strings"
)

func ProcessCommand(input string) {
	args := strings.Split(input, " ")

	switch args[0] {
	case "tunnel":
		handleTunnel(args)
	case "list":
		listTunnels()
	case "exit":
		fmt.Println("Exiting application...")
		os.Exit(0)
	case "help":
		ShowInformation()
	default:
		fmt.Println("Unknown command. Please try again.")
	}
}

func handleTunnel(args []string) {
	if len(args) < 2 {
		fmt.Println("Please provide a URL to crypt.")
		return
	}

	if !proxy.IsNewTunnelAvailable() {
		fmt.Println("Maximum number of tunnels reached. Cannot add more.")
		return
	}

	url := args[1]
	if isValidURL(url) {
		fmt.Println("Unfortunately, the string entered does not look like a link")
		return
	}

	encodedURL, err := proxy.AddNewProxyTunnel(url)
	if err != nil {
		fmt.Println("Invalid url")
		return;
	}
	
	fmt.Printf("Encrypted URL: %s\n", encodedURL)
}

func isValidURL(url string) bool {
	var re = regexp.MustCompile(`^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})(:[0-9]{1,5})?(\/[\w\.-]*)*\/?$`)
	return re.MatchString(url)
}

func listTunnels() {
	fmt.Println("Active tunnels:")
	tunnels := proxy.GetListOfTunnels()
	for encoded, original := range tunnels {
		fmt.Printf("Encoded: %s, Original: %s\n", encoded, original)
	}
}

func ShowInformation() {
	fmt.Println("\nAvailable commands:")
	fmt.Println("\t* tunnel [url] - Encrypts the URL and creates a tunnel")
	fmt.Println("\t* list - Lists all active tunnels")
	fmt.Println("\t* exit - Exits the application")
}
