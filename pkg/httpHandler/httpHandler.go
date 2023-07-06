package httphandler

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marsalling json: %v", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(status)
	w.Write(dat)
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	RespondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func Readiness(w http.ResponseWriter, r *http.Request) {
	type readyResponse struct {
		Status string `json:"status"`
	}
	RespondWithJSON(w, 200, readyResponse{
		Status: "ok",
	})
}

func ErrHandler(w http.ResponseWriter, r *http.Request) {
	RespondWithError(w, 500, "Internal Server Error")
}
