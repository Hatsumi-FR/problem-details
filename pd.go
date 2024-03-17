// Package problem_details provides utilities for creating and handling HTTP errors in a standardized way.
// It follows the problem details for HTTP APIs (RFC 9457) specification.
package problem_details

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ProblemDetails struct {
	Type     string         `json:"type"`
	Title    string         `json:"title"`
	Status   int            `json:"-"`
	Detail   string         `json:"detail"`
	Instance string         `json:"instance"`
	Members  map[string]any `json:"members,omitempty"`
}

func NewProblemDetails(typeURI, title string, status int, detail, instance string, members map[string]any) *ProblemDetails {
	if typeURI == "" {
		typeURI = "about:blank"
	}

	return &ProblemDetails{
		Type:     typeURI,
		Title:    title,
		Status:   status,
		Detail:   detail,
		Instance: instance,
		Members:  members,
	}
}

func WriteProblemDetails(w http.ResponseWriter, pd *ProblemDetails) {
	err := pd.Validate()
	if err != nil {
		pd = &ProblemDetails{
			Type:     "about:blank",
			Title:    "Problem details is not valid.",
			Detail:   "Problem details is not valid : " + err.Error(),
			Instance: pd.Instance,
			Status:   http.StatusInternalServerError,
		}
	}
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(pd.Status)
	err = json.NewEncoder(w).Encode(pd)
	if err != nil {
		return
	}
}

func (pb *ProblemDetails) Validate() error {
	if pb.Status < 400 || pb.Status > 599 {
		return errors.New("status code is not a valid HTTP error status code")
	}
	if pb.Type == "" {
		return errors.New("type is required")
	}
	if pb.Title == "" {
		return errors.New("title is required")
	}
	if pb.Detail == "" {
		return errors.New("detail is required")
	}
	if pb.Instance == "" {
		return errors.New("instance is required")
	}
	return nil
}
