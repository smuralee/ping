package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	clientIP := GetClientIP(r)
	response, _ := json.Marshal(map[string]string{
		"IP":     clientIP,
		"Status": "Success",
	})
	fmt.Printf("Request received from %s\n", clientIP)

	logErr(w.Write(response))
}

func logErr(n int, err error) {
	if err != nil {
		log.Printf("Write failed: %v", err)
	}
}

// Return IP address by reading from the forwarded-for header for proxy and default fall back to use the remote address.
func GetClientIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

