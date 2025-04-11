package main

import (
	"fmt"
	"log"
	"net/http"
)

func blockHostsMiddleware(next http.Handler) http.Handler {
	blockedHosts := []string{
		"adservice.google.com",
		"clients2.google.com",
		"clientservices.googleapis.com",
		"safebrowsingohttpgateway.googleapis.com",
		"ssl.google-analytics.com",
		"update.googleapis.com",
		"www.google-analytics.com",
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, host := range blockedHosts {
			if r.Host == host {
				log.Printf("💖 Blocked request to host: %s", r.Host)
				w.Header().Set("Content-Type", "text/plain; charset=utf-8")
				http.Error(w, fmt.Sprintf("🚫 Nope. '%s' is on the naughty list. Chrome doesn't need to know your secrets, and neither do these nosy hosts. Go bother someone else. ✋", r.Host), http.StatusForbidden)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
