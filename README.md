# Problem details

## Import package

```shell
go get github.com/hatsumi-fr/problem-details

```

## Basic usage

This package provides a set of utilities for creating and handling HTTP errors in a standardized way, following the problem details for HTTP APIs (RFC9457) specification.

RFC is available at https://datatracker.ietf.org/doc/rfc9457/

The Problem Details struct represents a problem that can be returned in an HTTP response. It includes the following fields:  
- Type: A URI that identifies the problem type.
- Title: A short, human-readable summary of the problem type.
- Status: The HTTP status code for this occurrence of the problem.
- Detail: A human-readable explanation specific to this occurrence of the problem.
- Instance: A URI that identifies the specific occurrence of the problem.
- Members: A map of additional members that may be attached to the problem details object.

## Basic usage example

```go
package main

import (
	"net/http"
	"github.com/hatsumi-fr/problem-details"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pd, err := problem_details.NewProblemDetails(
			"https://example.com/probs/out-of-credit",
			"You do not have enough credit.",
			http.StatusPaymentRequired,
			"Your current balance is 30, but that costs 50.",
			"https://example.com/account/12345/msgs/abc",
			map[string]string{
				"balance": "30",
				"accounts": "https://example.com/account/12345",
			},
		)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		problem_details.WriteProblemDetails(w, pd)
	})

	http.ListenAndServe(":8080", nil)
}
```

## Response example

```json
{
	"type": "about:blank",
	"title": "Request body is invalid.",
	"detail": "Field 'name' is required.",
	"instance": "/api/articles/5"
}
```

## Contribution

Feel free to contribute to this project by forking the repository and submitting pull requests. Every contribution is welcome! 