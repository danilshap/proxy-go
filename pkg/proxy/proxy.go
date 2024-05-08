package proxy

import (
	"net/http"
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

	proxyReq, err := http.NewRequest(r.Method, target+r.URL.RequestURI(), r.Body)
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
	w.Write([]byte("Proxy response: \n"))
	resp.Write(w)
}

func IsNewTunnelAvaliable() bool {
	return len(tunnels) >= maxTunnels
}

func AddNewProxyTunnel(url string) string {
	encodedURL := utils.GenerateUri()
	mu.Lock()
	tunnels[encodedURL] = url
	mu.Unlock()

	return encodedURL
}

func GetListOfTunels() map[string]string {
	return tunnels
}
