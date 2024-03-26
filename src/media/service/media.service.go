// Package to handle media.
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
	"net/http"

	bootstrap_model "github.com/Rfluid/whatsapp-cloud-api/src/bootstrap/model"
	common_enum "github.com/Rfluid/whatsapp-cloud-api/src/common/enum"
	common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"
	media_model "github.com/Rfluid/whatsapp-cloud-api/src/media/model"
)

// Uploads media files. All media files sent through this endpoint are
// encrypted and persist for 30 days, unless they are deleted earlier.
func Upload(
	api bootstrap_model.WhatsAppAPI,
	data media_model.Upload,
) (common_model.Id, error) {
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIdURL, common_enum.Media),
		bytes.NewBuffer(jsonData),
	)
	req.Header = api.Headers

	resp, err := api.Client.Do(req)
	if err != nil {
		return common_model.Id{}, err
	}
	defer resp.Body.Close()

	var body common_model.Id

	json.NewDecoder(resp.Body).Decode(&body)

	return body, nil
}

// Retrievers media URL and info by media id.
// Use the returned URL to download the media file. Note that
// clicking this URL (i.e. performing a generic GET) will not
// return the media; you must include an access token.
// See https://developers.facebook.com/docs/whatsapp/cloud-api/reference/media?locale=pt_BR#download-media.
//
// @mediaId The id of media you want info about.
func RetrieveURL(
	api bootstrap_model.WhatsAppAPI,
	mediaId string,
	data media_model.RetrieveInfo,
) (media_model.MediaInfo, error) {
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/%s", api.MainURL, mediaId),
		bytes.NewBuffer(jsonData),
	)
	req.Header = api.Headers

	resp, err := api.Client.Do(req)
	if err != nil {
		return media_model.MediaInfo{}, err
	}
	defer resp.Body.Close()

	var body media_model.MediaInfo

	json.NewDecoder(resp.Body).Decode(&body)

	return body, nil
}

// Deletes media by media id.
func Delete(
	api bootstrap_model.WhatsAppAPI,
	mediaId string,
) (common_model.SuccessResponse, error) {
	req, _ := http.NewRequest(
		"DELETE",
		fmt.Sprintf("%s/%s", api.MainURL, mediaId),
		nil,
	)
	req.Header = api.Headers

	resp, err := api.Client.Do(req)
	if err != nil {
		return common_model.SuccessResponse{}, err
	}
	defer resp.Body.Close()

	var body common_model.SuccessResponse

	json.NewDecoder(resp.Body).Decode(&body)

	return body, nil
}

func Download(
	api bootstrap_model.WhatsAppAPI,
	url string,
) (common_model.SuccessResponse, error) {
	req, _ := http.NewRequest(
		"GET",
		url,
		nil,
	)
	req.Header = api.Headers

	resp, err := api.Client.Do(req)
	if err != nil {
		return common_model.SuccessResponse{}, err
	}
	defer resp.Body.Close()

	var body common_model.SuccessResponse

	json.NewDecoder(resp.Body).Decode(&body)

	return body, nil
}
