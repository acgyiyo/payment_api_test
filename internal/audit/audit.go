package audit

import (
	"log"
	"net/http"
)

func LogRequest(r *http.Request, status string) {
	log.Printf("Request: %s %s | Status: %s", r.Method, r.URL.Path, status)
}
