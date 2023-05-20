package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

// json response
func respond(w http.ResponseWriter, data any) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// remote address
func getAddress(r *http.Request) string {
	address := r.Header.Get("x-forwarded-for")
	address = strings.TrimSpace(strings.Split(address, ",")[len(strings.Split(address, ","))-1])
	if address == "" {
		address = strings.Split(r.RemoteAddr, ":")[0]
	}
	return address
}