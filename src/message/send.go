// Sends, spams and changes message status. You can also use this package
// to config user identity check.
//
// Use the SendMessageStatusCheck function to check the status of your API.
package message

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/Rfluid/whatsapp-cloud-api/src/bootstrap"
	"github.com/Rfluid/whatsapp-cloud-api/src/common"
)

// Sends 
//
// Send service does nothing if the request gets ignored
// by WhatsApp. It can happen if you are trying to send 100 messages
// to the same user. If you want to retry on fail, use spam service.
func Send(
	api bootstrap.WhatsAppAPI,
	data Message,
) (Response, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return Response{}, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIDURL, common.Messages),
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return Response{}, err
	}
	req.Header = api.JSONHeaders

	resp, err := api.Client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var respErr common.ErrorResponse
		if decodeErr := json.NewDecoder(resp.Body).Decode(&respErr); decodeErr != nil {
			err = decodeErr
		} else {
			err = &respErr
		}
		return Response{}, err
	}

	var body Response

	err = json.NewDecoder(resp.Body).Decode(&body)

	return body, err
}

// Sends message with media cache control.
//
// You may use this to control WhatsApp caching sending media via url.
//
// Send service does nothing if the request gets ignored
// by WhatsApp. It can happen if you are trying to send 100 messages
// to the same user. If you want to retry on fail, use spam service.
func SendWithCacheControll(
	api bootstrap.WhatsAppAPI,
	data Message,
	cacheControl MediaCacheControl,
) (Response, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return Response{}, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIDURL, common.Messages),
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return Response{}, err
	}
	req.Header = api.JSONHeaders

	for key, value := range cacheControl.ToMap() {
		req.Header.Add(key, value)
	}

	resp, err := api.Client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var respErr common.ErrorResponse
		if decodeErr := json.NewDecoder(resp.Body).Decode(&respErr); decodeErr != nil {
			err = decodeErr
		} else {
			err = &respErr
		}
		return Response{}, err
	}

	var body Response

	err = json.NewDecoder(resp.Body).Decode(&body)

	return body, err
}

// Sends many messages.
//
// All messages are sent using parallelism.
//
// Send service does nothing if the request gets ignored
// by WhatsApp. It can happen if you are trying to send 100 messages
// to the same user. If you want to retry on fail, use spam service.
func SendMany(
	api bootstrap.WhatsAppAPI,
	data []Message,
) ([](Response), []error) {
	respCh := make(chan Response)
	errCh := make(chan error)
	responses := []Response{}
	errs := []error{}
	var wg sync.WaitGroup

	for _, msg := range data {
		wg.Add(1)

		go func(msg Message) {
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
	api bootstrap.WhatsAppAPI,
	data []Message,
	cacheControl MediaCacheControl,
) ([](Response), []error) {
	respCh := make(chan Response)
	errCh := make(chan error)
	responses := []Response{}
	errs := []error{}
	var wg sync.WaitGroup

	for _, msg := range data {
		wg.Add(1)

		go func(msg Message) {
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
	api bootstrap.WhatsAppAPI,
	data []Message,
	callback func(Message, Response, error),
) {
	var wg sync.WaitGroup

	for _, msg := range data {
		wg.Add(1)

		go func(msg Message) {
			defer wg.Done()

			response, err := Send(api, msg)

			callback(msg, response, err)
		}(msg)
	}
	wg.Wait()
}

// Same as SendWithCacheControll but applies a callback for each result.
func SendManyWithCacheControllAndCallback(
	api bootstrap.WhatsAppAPI,
	data []Message,
	cacheControl MediaCacheControl,
	callback func(Message, Response, error),
) {
	var wg sync.WaitGroup

	for _, msg := range data {
		wg.Add(1)

		go func(msg Message) {
			defer wg.Done()

			response, err := SendWithCacheControll(api, msg, cacheControl)

			callback(msg, response, err)
		}(msg)
	}
	wg.Wait()
}
