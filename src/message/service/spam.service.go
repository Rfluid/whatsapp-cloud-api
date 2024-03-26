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

// Spams message.
//
// Spam service retries to send message until WhatsApp accepts if the request gets ignored
// by WhatsApp. It can happen if you are trying to send 100 messages
// to the same user. If you don't want to retry on fail, use send service.
//
// This is not safe at all because of the lack of time limit and try limit.
func Spam(
	api bootstrap_model.WhatsAppAPI,
	data message_model.Message,
) (message_model.Response, error) {
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIdURL, common_enum.Messages),
		bytes.NewBuffer(jsonData),
	)
	req.Header = api.Headers

	resp, err := api.Client.Do(req)
	if err != nil {
		return message_model.Response{}, err
	}
	defer resp.Body.Close()

	var body message_model.Response

	json.NewDecoder(resp.Body).Decode(&body)

	if body.MessagingProduct.MessagingProduct == "" {
		return Spam(api, data)
	}

	return body, nil
}

// Spams message with media cache control.
//
// You may use this to control WhatsApp caching sending media via url.
//
// Spam service retries to send message until WhatsApp accepts if the request gets ignored
// by WhatsApp. It can happen if you are trying to send 100 messages
// to the same user. If you don't want to retry on fail, use send service.
//
// This is not safe at all because of the lack of time limit and try limit.
func SpamWithCacheControll(
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
	req.Header = api.Headers

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

	if body.MessagingProduct.MessagingProduct == "" {
		return SpamWithCacheControll(api, data, cacheControl)
	}

	return body, nil
}

// Spams many messages.
//
// All messages are sent using parallelism.
//
// Spam service retries to send message until WhatsApp accepts if the request gets ignored
// by WhatsApp. It can happen if you are trying to send 100 messages
// to the same user. If you don't want to retry on fail, use send service.
//
// This is not safe at all because of the lack of time limit and try limit.
func SpamMany(
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

			response, err := Spam(api, msg)

			if err == nil {
				respCh <- response
			} else {
				errCh <- err
			}
		}(msg)
	}
	wg.Wait()

	for response := range respCh {
		responses = append(responses, response)
	}
	for err := range errCh {
		errs = append(errs, err)
	}

	return responses, errs
}

// Spams many messages with media cache control.
//
// You may use this to control WhatsApp caching sending media via url.
//
// All messages are sent using parallelism.
//
// Spam service retries to send message until WhatsApp accepts if the request gets ignored
// by WhatsApp. It can happen if you are trying to send 100 messages
// to the same user. If you don't want to retry on fail, use send service.
//
// This is not safe at all because of the lack of time limit and try limit.
func SpamManyWithCacheControll(
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

			response, err := SpamWithCacheControll(api, msg, cacheControl)

			if err == nil {
				respCh <- response
			} else {
				errCh <- err
			}
		}(msg)
	}
	wg.Wait()

	for response := range respCh {
		responses = append(responses, response)
	}
	for err := range errCh {
		errs = append(errs, err)
	}

	return responses, errs
}

// Same as SpamMany but applies a callback for each result.
func SpamManyWithCallback(
	api bootstrap_model.WhatsAppAPI,
	data []message_model.Message,
	callback func(message_model.Message, message_model.Response, error),
) {
	var wg sync.WaitGroup

	for _, msg := range data {
		wg.Add(1)

		go func(msg message_model.Message) {
			defer wg.Done()

			response, err := Spam(api, msg)

			callback(msg, response, err)
		}(msg)
	}
	wg.Wait()
}

// Same as SpamWithCacheControll but applies a callback for each result.
func SpamManyWithCacheControllAndCallback(
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

			response, err := SpamWithCacheControll(api, msg, cacheControl)

			callback(msg, response, err)
		}(msg)
	}
	wg.Wait()
}
