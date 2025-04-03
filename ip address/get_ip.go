package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

type IPResponse struct {
	IP string `json:"ip"`
}

func getIP(w http.ResponseWriter, r *http.Request) {
	// Get the IP address from headers or RemoteAddr
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr) // removes port if present
	} else {
		// Handle multiple IPs in X-Forwarded-For, if present
		ip = strings.Split(ip, ",")[0]
		ip = strings.TrimSpace(ip) // clean any spaces around IP
	}

	// Print IP in the terminal in desired format
	fmt.Println("User IP:", ip)

	// Send the IP back as a JSON response
	response := IPResponse{IP: ip}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/get-ip", getIP)
	log.Println("Server running on http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
