package business_service

import (
	"encoding/json"
	"fmt"
	"net/http"

	bootstrap_model "github.com/Rfluid/whatsapp-cloud-api/src/bootstrap/model"
	business_model "github.com/Rfluid/whatsapp-cloud-api/src/business/model"
	common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"
)

// Retrieves the business information using the provided API and query parameters.
//
// See https://developers.facebook.com/docs/marketing-api/reference/business
func Get(
	api bootstrap_model.WhatsAppAPI,
	query business_model.BusinessQueryParams,
) (business_model.Business, error) {
	// Build the query string from the query parameters
	queryString := query.BuildQuery()

	// Construct the full URL by appending the query string if it's not empty
	fullURL := api.WABAAccountIDURL
	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", api.WABAAccountIDURL, queryString)
	}

	// Create a new HTTP GET request with the full URL
	req, err := http.NewRequest(
		"GET",
		fullURL,
		nil,
	)
	if err != nil {
		// Return an empty Business and the encountered error if request creation fails
		return business_model.Business{}, err
	}

	// Execute the HTTP request using the provided HTTP client
	resp, err := api.Client.Do(req)
	if err != nil {
		// Return an empty Business and the encountered error if the request fails
		return business_model.Business{}, err
	}
	// Ensure the response body is closed after function execution
	defer resp.Body.Close()

	// Check if the response status code is not 200 OK
	if resp.StatusCode != http.StatusOK {
		var respErr common_model.ErrorResponse
		// Attempt to decode the error response into ErrorResponse struct
		if decodeErr := json.NewDecoder(resp.Body).Decode(&respErr); decodeErr != nil {
			// If decoding fails, assign the decoding error
			err = decodeErr
		} else {
			// If decoding succeeds, assign the API error message
			err = &respErr
		}
		// Return an empty Business and the encountered error
		return business_model.Business{}, err
	}

	// Initialize a Business struct to hold the successful response data
	var body business_model.Business

	// Decode the JSON response body into the Business struct
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		// Return an empty Business and the encountered error if decoding fails
		return business_model.Business{}, err
	}

	// Return the populated Business struct and nil error upon success
	return body, nil
}
