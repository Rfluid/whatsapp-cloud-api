package message_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	bootstrap_model "github.com/Rfluid/whatsapp-cloud-api/src/bootstrap/model"
	common_enum "github.com/Rfluid/whatsapp-cloud-api/src/common/enum"
	message_model "github.com/Rfluid/whatsapp-cloud-api/src/message/model"
)

// Sends message.
//
// Send service does nothing if the request gets ignored
// by WhatsApp. It can happen if you are trying to send 100 messages
// to the same user. If you want to retry on fail, use spam service.
func Send(
	api bootstrap_model.WhatsAppAPI,
	data message_model.Message,
) (message_model.Response, error) {
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIdURL, common_enum.Messages),
		bytes.NewBuffer(jsonData),
	)
	req.Header = api.JSONHeaders

	resp, err := api.Client.Do(req)
	if err != nil {
		return message_model.Response{}, err
	}
	defer resp.Body.Close()

	var body message_model.Response

	json.NewDecoder(resp.Body).Decode(&body)

	return body, nil
}

// Sends message with media cache control.
//
// You may use this to control WhatsApp caching sending media via url.
//
// Send service does nothing if the request gets ignored
// by WhatsApp. It can happen if you are trying to send 100 messages
// to the same user. If you want to retry on fail, use spam service.
func SendWithCacheControll(
	api bootstrap_model.WhatsAppAPI,
	data message_model.Message,
	cacheControl message_model.MediaCacheControl,
) (message_model.Response, error) {
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIdURL, common_enum.Messages),
		bytes.NewBuffer(jsonData),
	)
	req.Header = api.JSONHeaders

	for key, value := range cacheControl.ToMap() {
		req.Header.Add(key, value)
	}

	resp, err := api.Client.Do(req)
	if err != nil {
		return message_model.Response{}, err
	}
	defer resp.Body.Close()

	var body message_model.Response

	json.NewDecoder(resp.Body).Decode(&body)

	return body, nil
}

// Sends many messages.
//
// All messages are sent using parallelism.
//
// Send service does nothing if the request gets ignored
// by WhatsApp. It can happen if you are trying to send 100 messages
// to the same user. If you want to retry on fail, use spam service.
func SendMany(
	api bootstrap_model.WhatsAppAPI,
	data []message_model.Message,
) ([](message_model.Response), []error) {
	var mu sync.Mutex
	var errMu sync.Mutex
	responses := []message_model.Response{}
	errs := []error{}
	var wg sync.WaitGroup

	for _, msg := range data {
		wg.Add(1)

		go func(msg message_model.Message) {
			defer wg.Done()

			response, err := Send(api, msg)

			if err == nil {
				mu.Lock()
				responses = append(responses, response)
				mu.Unlock()
			} else {
				errMu.Lock()
				errs = append(errs, err)
				errMu.Unlock()
			}
		}(msg)
	}
	wg.Wait()

	return responses, errs
}

// Sends many messages with media cache control.
//
// You may use this to control WhatsApp caching sending media via url.
//
// All messages are sent using parallelism.
//
// Send service does nothing if the request gets ignored
// by WhatsApp. It can happen if you are trying to send 100 messages
// to the same user. If you want to retry on fail, use spam service.
func SendManyWithCacheControll(
	api bootstrap_model.WhatsAppAPI,
	data []message_model.Message,
	cacheControl message_model.MediaCacheControl,
) ([](message_model.Response), []error) {
	var mu sync.Mutex
	var errMu sync.Mutex
	responses := []message_model.Response{}
	errs := []error{}
	var wg sync.WaitGroup

	for _, msg := range data {
		wg.Add(1)

		go func(msg message_model.Message) {
			defer wg.Done()

			response, err := SendWithCacheControll(api, msg, cacheControl)

			if err == nil {
				mu.Lock()
				responses = append(responses, response)
				mu.Unlock()
			} else {
				errMu.Lock()
				errs = append(errs, err)
				errMu.Unlock()
			}
		}(msg)
	}
	wg.Wait()

	return responses, errs
}

// Same as SendMany but applies a callback for each result.
func SendManyWithCallback(
	api bootstrap_model.WhatsAppAPI,
	data []message_model.Message,
	callback func(message_model.Message, message_model.Response, error),
) {
	var wg sync.WaitGroup

	for _, msg := range data {
		wg.Add(1)

		go func(msg message_model.Message) {
			defer wg.Done()

			response, err := Send(api, msg)

			callback(msg, response, err)
		}(msg)
	}
	wg.Wait()
}

// Same as SendWithCacheControll but applies a callback for each result.
func SendManyWithCacheControllAndCallback(
	api bootstrap_model.WhatsAppAPI,
	data []message_model.Message,
	cacheControl message_model.MediaCacheControl,
	callback func(message_model.Message, message_model.Response, error),
) {
	var wg sync.WaitGroup

	for _, msg := range data {
		wg.Add(1)

		go func(msg message_model.Message) {
			defer wg.Done()

			response, err := SendWithCacheControll(api, msg, cacheControl)

			callback(msg, response, err)
		}(msg)
	}
	wg.Wait()
}
