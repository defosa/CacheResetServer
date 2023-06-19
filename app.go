package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const (
	port         = ":8999"
	allowedCIDR  = "10.128.0.0/24"
	nginxCacheDir = "/path/to/nginx/cache/directory" // path to nginx cache directory
)

func main() {
	http.HandleFunc("/reset", handleCacheReset)
	fmt.Printf("Server listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handleCacheReset(w http.ResponseWriter, r *http.Request) {
	clientIP := getClientIP(r)
	if !isIPAllowed(clientIP) {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}

	resetCache()

	fmt.Fprint(w, "Cache reset successfully")
}

func getClientIP(r *http.Request) string {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return ""
	}
	return ip
}

func isIPAllowed(ip string) bool {
	_, ipNet, err := net.ParseCIDR(allowedCIDR)
	if err != nil {
		return false
	}
	clientIP := net.ParseIP(ip)
	return ipNet.Contains(clientIP)
}

func resetCache() {
	cmd := exec.Command("rm", "-r", nginxCacheDir, "*")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Printf("Error resetting cache: %v\n", err)
	}
}
