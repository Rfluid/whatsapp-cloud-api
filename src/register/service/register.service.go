package register_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	bootstrap_model "github.com/Rfluid/whatsapp-cloud-api/src/bootstrap/model"
	common_enum "github.com/Rfluid/whatsapp-cloud-api/src/common/enum"
	common_model "github.com/Rfluid/whatsapp-cloud-api/src/common/model"
	register_model "github.com/Rfluid/whatsapp-cloud-api/src/register/model"
)

// Register WhatsApp business number.
func Register(
	api bootstrap_model.CloudApi,
	data register_model.Register,
) (common_model.SuccessResponse, error) {
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIdURL, common_enum.Register),
		bytes.NewBuffer(jsonData),
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

// Deregister WhatsApp business number.
func DeRegister(
	api bootstrap_model.CloudApi,
) (common_model.SuccessResponse, error) {
	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIdURL, common_enum.DeRegister),
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
