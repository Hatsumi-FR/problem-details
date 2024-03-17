// Package problem_details provides utilities for creating and handling HTTP errors in a standardized way.
// It follows the problem details for HTTP APIs (RFC 9457) specification.
package problem_details

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ProblemDetails struct {
	Type     string            `json:"type"`
	Title    string            `json:"title"`
	Status   int               `json:"-"`
	Detail   string            `json:"detail"`
	Instance string            `json:"instance"`
	Members  map[string]string `json:"members,omitempty"`
}

func NewProblemDetails(typeURI, title string, status int, detail, instance string, members map[string]string) (*ProblemDetails, error) {
	if typeURI == "" {
		typeURI = "about:blank"
	}

	if status < 400 || status > 599 {
		return nil, errors.New("status code is not a valid HTTP error status code")
	}

	return &ProblemDetails{
		Type:     typeURI,
		Title:    title,
		Status:   status,
		Detail:   detail,
		Instance: instance,
		Members:  members,
	}, nil
}

func WriteProblemDetails(w http.ResponseWriter, pd *ProblemDetails) {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(pd.Status)
	err := json.NewEncoder(w).Encode(pd)
	if err != nil {
		return
	}
}
