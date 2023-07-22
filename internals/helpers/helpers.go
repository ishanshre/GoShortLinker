package helpers

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

func EnforceHttp(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}

func RemoveDomainError(url string) bool {
	if url == os.Getenv("DOMAIN") {
		return false
	}
	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0]
	return newURL != os.Getenv("DOMAIN")
}

func WriteJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func StatusBadRequest(w http.ResponseWriter, err string) {
	WriteJson(w, http.StatusBadRequest, err)
}

func StatusCreated(w http.ResponseWriter, data any) {
	WriteJson(w, http.StatusCreated, data)
}

func StatusAccepted(w http.ResponseWriter, data any) {
	WriteJson(w, http.StatusAccepted, data)
}

func StatusInternalServerError(w http.ResponseWriter, err string) {
	WriteJson(w, http.StatusInternalServerError, err)
}
