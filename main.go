/**
   Copyright 2021 Suraj Muraleedharan

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
**/

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

