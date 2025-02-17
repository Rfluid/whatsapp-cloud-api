package business_model

import (
	"net/url"
	"strings"
)

// BusinessQueryParams holds the query parameters for a business request.
type BusinessQueryParams struct {
	Fields *[]string `json:"fields,omitempty"`
}

// BuildQuery constructs the query string from BusinessQueryParams.
// It returns an encoded query string that can be appended to a URL.
func (q BusinessQueryParams) BuildQuery() string {
	// Initialize a URL values object
	v := url.Values{}

	// If Fields are provided, join them with commas and add to the query
	if q.Fields != nil && len(*q.Fields) > 0 {
		v.Set("fields", strings.Join(*q.Fields, ","))
	}

	// Encode the query parameters
	return v.Encode()
}
