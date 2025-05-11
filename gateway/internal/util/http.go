package util

import (
	"encoding/json"
	"maps"
	"net/http"

	"github.com/rs/zerolog/log"
)

type envelope map[string]any

func writeJSON(w http.ResponseWriter, status int, data map[string]any, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	maps.Copy(w.Header(), headers)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func ErrorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}

	err := writeJSON(w, status, env, nil)
	if err != nil {
		log.Error().Err(err).Msg("http request failed")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Error().Err(err)

	message := "the server encountered a problem and could not process your request"
	ErrorResponse(w, r, http.StatusInternalServerError, message)
}

func RateLimitExceededResponse(w http.ResponseWriter, r *http.Request) {
	log.Error().Msg("rate limit exceeded")

	message := "rate limit exceeded"
	ErrorResponse(w, r, http.StatusTooManyRequests, message)
}

func InvalidAuthenticationTokenResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("WWW-Authenticate", "Bearer")

	message := "invalid or missing authenticate token"
	ErrorResponse(w, r, http.StatusUnauthorized, message)
}

func UnauthorizedResponse(w http.ResponseWriter, r *http.Request) {
	log.Error().Msg("rate limit exceeded")

	message := "unauthorized request"
	ErrorResponse(w, r, http.StatusUnauthorized, message)
}
