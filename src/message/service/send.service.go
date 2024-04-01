// Sends, spams and changes message status. You can also use this package
// to config user identity check.
//
// Use the SendMessageStatusCheck function to check the status of your API.
package message_service

import (
	"bytes"
	"encoding/json"
	"errors"
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
	jsonData, err := json.Marshal(data)
	if err != nil {
		return message_model.Response{}, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIdURL, common_enum.Messages),
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return message_model.Response{}, err
	}
	req.Header = api.JSONHeaders

	resp, err := api.Client.Do(req)
	if err != nil {
		return message_model.Response{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errInt map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&errInt); err != nil {
			return message_model.Response{}, err
		}
		errMsgBytes, err := json.MarshalIndent(errInt, "", "    ")
		if err != nil {
			return message_model.Response{}, err
		}
		return message_model.Response{}, errors.New(string(errMsgBytes))
	}

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
	jsonData, err := json.Marshal(data)
	if err != nil {
		return message_model.Response{}, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIdURL, common_enum.Messages),
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return message_model.Response{}, err
	}
	req.Header = api.JSONHeaders

	for key, value := range cacheControl.ToMap() {
		req.Header.Add(key, value)
	}

	resp, err := api.Client.Do(req)
	if err != nil {
		return message_model.Response{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errInt map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&errInt); err != nil {
			return message_model.Response{}, err
		}
		errMsgBytes, err := json.MarshalIndent(errInt, "", "    ")
		if err != nil {
			return message_model.Response{}, err
		}
		return message_model.Response{}, errors.New(string(errMsgBytes))
	}

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
	respCh := make(chan message_model.Response)
	errCh := make(chan error)
	responses := []message_model.Response{}
	errs := []error{}
	var wg sync.WaitGroup

	for _, msg := range data {
		wg.Add(1)

		go func(msg message_model.Message) {
			defer wg.Done()

			response, err := Send(api, msg)

			if err == nil {
				respCh <- response
			} else {
				errCh <- err
			}
		}(msg)
	}

	go func() {
		wg.Wait()
		close(respCh)
		close(errCh)
	}()

	for response := range respCh {
		responses = append(responses, response)
	}
	for err := range errCh {
		errs = append(errs, err)
	}

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
	respCh := make(chan message_model.Response)
	errCh := make(chan error)
	responses := []message_model.Response{}
	errs := []error{}
	var wg sync.WaitGroup

	for _, msg := range data {
		wg.Add(1)

		go func(msg message_model.Message) {
			defer wg.Done()

			response, err := SendWithCacheControll(api, msg, cacheControl)

			if err == nil {
				respCh <- response
			} else {
				errCh <- err
			}
		}(msg)
	}

	go func() {
		wg.Wait()
		close(respCh)
		close(errCh)
	}()

	for response := range respCh {
		responses = append(responses, response)
	}
	for err := range errCh {
		errs = append(errs, err)
	}

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
