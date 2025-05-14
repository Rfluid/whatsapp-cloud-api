package template_service

import (
	"encoding/json"
	"fmt"
	"net/http"

	bootstrap_model "github.com/Rfluid/whatsapp-cloud-api/src/bootstrap/model"
	common_enum "github.com/Rfluid/whatsapp-cloud-api/src/common/enum"
	common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"
	template_model "github.com/Rfluid/whatsapp-cloud-api/src/template/model"
)

// Gets templates.
//
// See https://developers.facebook.com/docs/graph-api/reference/whats-app-business-account/message_templates/
//
// @fields Are the fields that must be returned by query.
// @query Is mainly used to paginate.
func Get(
	api bootstrap_model.WhatsAppAPI,
	query template_model.TemplateQueryParams,
) (template_model.GetTemplateResponse, error) {
	// Build the query string from the query parameters
	queryString := query.BuildQuery()

	// Construct the full URL by appending the query string if it's not empty
	fullURL := fmt.Sprintf("%s/%s", api.WABAAccountIdURL, common_enum.MessageTemplates)
	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}

	// Create a new HTTP GET request with the full URL
	req, err := http.NewRequest(
		"GET",
		fullURL,
		nil,
	)
	if err != nil {
		// Return an empty GetTemplateResponse and the encountered error if request creation fails
		return template_model.GetTemplateResponse{}, err
	}

	// Set the necessary headers
	req.Header = api.JSONHeaders

	// Execute the HTTP request using the provided HTTP client
	resp, err := api.Client.Do(req)
	if err != nil {
		// Return an empty GetTemplateResponse and the encountered error if the request fails
		return template_model.GetTemplateResponse{}, err
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
		// Return an empty GetTemplateResponse and the encountered error
		return template_model.GetTemplateResponse{}, err
	}

	// Initialize a GetTemplateResponse struct to hold the successful response data
	var body template_model.GetTemplateResponse

	// Decode the JSON response body into the GetTemplateResponse struct
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		// Return an empty GetTemplateResponse and the encountered error if decoding fails
		return template_model.GetTemplateResponse{}, err
	}

	// Return the populated GetTemplateResponse struct and nil error upon success
	return body, nil
}
