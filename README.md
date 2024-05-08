
# Go Tunnel Manager

## Overview
The Go Tunnel Manager is a simple tunneling application written in Go. It allows users to create secure tunnels to local servers, making them accessible via a generated encrypted URL. This functionality mimics services like ngrok but is a simplified version for educational purposes.

## Features
- **Tunnel Creation**: Users can encrypt a URL to create a new tunnel.
- **Tunnel Listing**: Displays all active tunnels with their original and encrypted URLs.
- **Dynamic Tunnel Management**: Supports up to 5 active tunnels at any given time.
- **Simple HTTP Proxy**: Includes a basic HTTP proxy server to handle requests through encrypted URLs.

## How to Run the Project
1. **Clone the Repository**:
   ```bash
   git clone <repository-url>
   ```
2. **Navigate to the project directory**:
   ```bash
   cd proxy-go
   ```
3. **Build the project** (if necessary):
   ```bash
   go build
   ```
4. **Run the application**:
   ```bash
   go run main.go
   ```

## How to Use
After starting the application, use the command line interface to manage tunnels:
- **Create a tunnel**: Enter `tunnel [url]` to create a new tunnel.
- **List active tunnels**: Enter `list` to view all active tunnels.
- **Exit the application**: Enter `exit` to stop the application and clear all tunnels.

## Commands
- `tunnel [url]`: Encrypts the provided URL and creates a tunnel.
- `list`: Lists all active tunnels with their original and encrypted URLs.
- `exit`: Exits the application and clears all active tunnels.
- `help`: Displays information about available commands.

## System Requirements
- Go version 1.15 or higher

## License
This project is licensed under the MIT License - see the LICENSE file for details.
