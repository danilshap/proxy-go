package proxy

import (
	"io"
	"net/http"
	"net/url"
	"proxy-go/utils"
	"sync"
)

const maxTunnels int = 5

var mu sync.Mutex
var tunnels = make(map[string]string)

func HandleProxy(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	target, ok := tunnels[r.URL.Path]
	mu.Unlock()

	if !ok {
		http.Error(w, "No tunnel for this URL", http.StatusNotFound)
		return
	}

	proxyReq, err := http.NewRequest(r.Method, target, r.Body)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	proxyReq.Header = r.Header

	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, "Failed to reach the destination", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(resp.StatusCode)
	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, "Failed to copy response body", http.StatusInternalServerError)
	}
}

func IsNewTunnelAvailable() bool {
	return len(tunnels) < maxTunnels
}

func AddNewProxyTunnel(url string) (string, error) {
	encodedURL := utils.GenerateUri()
	encodedPath, err := getPathFromEncodedURL(encodedURL)
	if err != nil {
		return "", err
	}

	mu.Lock()
	tunnels[encodedPath] = url
	mu.Unlock()

	return encodedURL, nil
}

func getPathFromEncodedURL(encodedURL string) (string, error) {
	parsedURL, err := url.Parse(encodedURL)
	if err != nil {
		return "", err
	}
	return parsedURL.Path, nil
}

func GetListOfTunnels() map[string]string {
	return tunnels
}
