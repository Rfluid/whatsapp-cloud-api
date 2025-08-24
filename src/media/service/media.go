// Handles media functions.
//
// To complete some of the following API calls, you need to have
// a media ID. There are two ways to get this ID:

// 1. From the API call: Once you have successfully uploaded
// media files to the API, the media ID is included in the
// response to your call.
//
// 2. From Webhooks: When a business account receives a media
// message, it downloads the media and uploads it to the
// Cloud API automatically. That event triggers the Webhooks
// and sends you a notification that includes the media ID.
package media_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	bootstrap_model "github.com/Rfluid/whatsapp-cloud-api/src/bootstrap/model"
	common_enum "github.com/Rfluid/whatsapp-cloud-api/src/common/enum"
	common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"
	media_model "github.com/Rfluid/whatsapp-cloud-api/src/media/model"
)

// Uploads media files. All media files sent through this endpoint are
// encrypted and persist for 30 days, unless they are deleted earlier.
// See https://developers.facebook.com/docs/whatsapp/cloud-api/reference/media/#upload-media
func Upload(
	api bootstrap_model.WhatsAppAPI,
	data media_model.Upload,
) (common_model.ID, error) {
	// Generate the body and content type for the request
	body, contentType, err := data.CreateForm()
	if err != nil {
		return common_model.ID{}, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIDURL, common_enum.Media),
		body,
	)
	if err != nil {
		return common_model.ID{}, err
	}
	req.Header = api.FormHeaders
	req.Header.Set("Content-Type", contentType)

	resp, err := api.Client.Do(req)
	if err != nil {
		return common_model.ID{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var respErr common_model.ErrorResponse
		if decodeErr := json.NewDecoder(resp.Body).Decode(&respErr); decodeErr != nil {
			err = decodeErr
		}
		return common_model.ID{}, err
	}

	var bodyResponse common_model.ID
	if err := json.NewDecoder(resp.Body).Decode(&bodyResponse); err != nil {
		return common_model.ID{}, err
	}

	return bodyResponse, err
}

// Retrievers media URL and info by media id.
// Use the returned URL to download the media file. Note that
// clicking this URL (i.e. performing a generic GET) will not
// return the media; you must include an access token.
// See https://developers.facebook.com/docs/whatsapp/cloud-api/reference/media?locale=pt_BR#download-media.
//
// @mediaID The id of media you want info about.
func RetrieveURL(
	api bootstrap_model.WhatsAppAPI,
	mediaID string,
	data media_model.RetrieveInfo,
) (media_model.MediaInfo, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return media_model.MediaInfo{}, err
	}
	req, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/%s", api.MainURL, mediaID),
		bytes.NewBuffer(jsonData),
	)
	req.Header = api.JSONHeaders

	resp, err := api.Client.Do(req)
	if err != nil {
		return media_model.MediaInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var respErr common_model.ErrorResponse
		if decodeErr := json.NewDecoder(resp.Body).Decode(&respErr); decodeErr != nil {
			err = decodeErr
		}
		return media_model.MediaInfo{}, err
	}

	var body media_model.MediaInfo

	err = json.NewDecoder(resp.Body).Decode(&body)

	return body, err
}

// Deletes media by media id.
func Delete(
	api bootstrap_model.WhatsAppAPI,
	mediaID string,
) (common_model.SuccessResponse, error) {
	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf("%s/%s", api.MainURL, mediaID),
		nil,
	)
	if err != nil {
		return common_model.SuccessResponse{}, err
	}
	req.Header = api.JSONHeaders

	resp, err := api.Client.Do(req)
	if err != nil {
		return common_model.SuccessResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var respErr common_model.ErrorResponse
		if decodeErr := json.NewDecoder(resp.Body).Decode(&respErr); decodeErr != nil {
			err = decodeErr
		}
		return common_model.SuccessResponse{}, err
	}

	var body common_model.SuccessResponse

	err = json.NewDecoder(resp.Body).Decode(&body)

	return body, err
}

// Downloads media from URL.
func Download(
	api bootstrap_model.WhatsAppAPI,
	url string,
) ([]byte, error) {
	req, err := http.NewRequest(
		"GET",
		url,
		nil,
	)
	if err != nil {
		return []byte{}, err
	}

	req.Header = api.JSONHeaders

	resp, err := api.Client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var respErr common_model.ErrorResponse
		if decodeErr := json.NewDecoder(resp.Body).Decode(&respErr); decodeErr != nil {
			err = decodeErr
		}
		return []byte{}, err
	}

	mediaBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return mediaBytes, nil
}
